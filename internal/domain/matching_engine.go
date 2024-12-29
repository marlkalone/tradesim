package domain

import (
    "fmt"
    "sort"

    "github.com/marlkalone/tradesim/internal/utils"
)

// MatchingEngine é responsável por casar ordens de compra e venda
type MatchingEngine struct {
    OrderBook *OrderBook     // Uso maiúsculo, pois definimos o campo assim
    Logger    utils.Logger
}

// NewMatchingEngine cria um novo MatchingEngine
func NewMatchingEngine(ob *OrderBook, logger utils.Logger) *MatchingEngine {
    return &MatchingEngine{
        OrderBook: ob,
        Logger:    logger,
    }
}

// ProcessOrder recebe uma ordem, adiciona ao OrderBook e tenta casar
func (me *MatchingEngine) ProcessOrder(o Order) {
    me.OrderBook.AddOrder(o)
    me.Logger.Info("MatchingEngine", fmt.Sprintf("Processando ordem: %+v", o))

    me.matchOrders()
}

// matchOrders checa se há possibilidade de casar ordens
func (me *MatchingEngine) matchOrders() {
    // Trancar o mutex do OrderBook
    me.OrderBook.Mu.Lock()
    defer me.OrderBook.Mu.Unlock()

    // Ordenamos as listas
    sort.Slice(me.OrderBook.BuyOrders, func(i, j int) bool {
        return me.OrderBook.BuyOrders[i].Price > me.OrderBook.BuyOrders[j].Price
    })
    sort.Slice(me.OrderBook.SellOrders, func(i, j int) bool {
        return me.OrderBook.SellOrders[i].Price < me.OrderBook.SellOrders[j].Price
    })

    if len(me.OrderBook.BuyOrders) > 0 && len(me.OrderBook.SellOrders) > 0 {
        highestBuy := &me.OrderBook.BuyOrders[0]
        lowestSell := &me.OrderBook.SellOrders[0]

        // Se a maior oferta de compra >= menor oferta de venda, temos match
        if highestBuy.Price >= lowestSell.Price {
            tradedAmount := min(highestBuy.Amount, lowestSell.Amount)

            me.Logger.Info("MatchingEngine",
                fmt.Sprintf("Casando %d ações ao preço de %.2f", tradedAmount, lowestSell.Price))

            // Ajusta quantidades
            highestBuy.Amount -= tradedAmount
            lowestSell.Amount -= tradedAmount

            // Remove do slice se ficou zerada
            if highestBuy.Amount == 0 {
                me.OrderBook.BuyOrders = me.OrderBook.BuyOrders[1:]
            }
            if lowestSell.Amount == 0 {
                me.OrderBook.SellOrders = me.OrderBook.SellOrders[1:]
            }
        }
    }
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
