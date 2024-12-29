package domain

import (
	"sync"

	"github.com/marlkalone/tradesim/internal/utils"
)

type OrderBook struct {
	BuyOrders 	[]Order
	SellOrders 	[]Order
	Mu 			sync.Mutex
	Logger 		utils.Logger
}

func NewOrderBook(logger utils.Logger) *OrderBook {
	return &OrderBook{
			BuyOrders: make([]Order, 0),
			SellOrders: make([]Order, 0),
			Logger: logger,
	}
}

func (ob *OrderBook) AddOrder(o Order) {
	ob.Mu.Lock()
	defer ob.Mu.Unlock()

	if o.Type == Buy {
		ob.BuyOrders = append(ob.BuyOrders, o)
		ob.Logger.Info("OrderBook", "Nova ordem de COMPRA adicionada")
	} else {
		ob.SellOrders = append(ob.SellOrders, o)
		ob.Logger.Info("OrderBook", "Nova ordem de VENDA adicionada")
	}
}