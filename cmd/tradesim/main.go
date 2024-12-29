package main

import (
    "fmt"

    "github.com/marlkalone/tradesim/internal/domain"
    "github.com/marlkalone/tradesim/internal/concurrency"
    "github.com/marlkalone/tradesim/internal/utils"
)

func main() {
    logger := utils.NewLogger()

    // Criando o OrderBook
    orderBook := domain.NewOrderBook(logger)

    // Configurando o MatchingEngine
    matchingEngine := domain.NewMatchingEngine(orderBook, logger)

    // Iniciando algumas goroutines que simulam "traders"
    concurrency.StartWorkers(matchingEngine, 5, logger)

    // Exemplo de funcionamento simples (sincronamente):
    o1 := domain.Order{ID: 1, Type: domain.Buy, Amount: 10, Price: 50.0}
    o2 := domain.Order{ID: 2, Type: domain.Sell, Amount: 5, Price: 48.0}

    matchingEngine.ProcessOrder(o1)
    matchingEngine.ProcessOrder(o2)

    fmt.Println("TradeSim running... Pressione CTRL+C para sair.")
    select {} // mantém o programa rodando
}
