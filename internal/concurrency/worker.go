package concurrency

import (
    "math/rand"
    "time"
    "fmt"

    "github.com/marlkalone/tradesim/internal/domain"
    "github.com/marlkalone/tradesim/internal/utils"
)

// StartWorkers inicia N workers que enviam ordens de forma assíncrona.
func StartWorkers(engine *domain.MatchingEngine, numWorkers int, logger utils.Logger) {
    for i := 1; i <= numWorkers; i++ {
        go worker(i, engine, logger)
    }
}

// worker simula o envio de ordens continuamente
func worker(id int, engine *domain.MatchingEngine, logger utils.Logger) {
    rand.Seed(time.Now().UnixNano() + int64(id))

    for {
        // Sorteia se é BUY ou SELL
        orderType := domain.Buy
        if rand.Intn(2) == 0 {
            orderType = domain.Sell
        }

        amount := rand.Intn(10) + 1             // 1..10
        price := float64(rand.Intn(10) + 45)    // 45..54

        o := domain.Order{
            ID:     rand.Intn(999999),
            Type:   orderType,
            Amount: amount,
            Price:  price,
        }

        logger.Info("Worker", 
            "Worker #"+fmtInt(id)+" gerando ordem: "+fmtInt(o.ID))

        engine.ProcessOrder(o)

        // Pausa aleatória antes de gerar outra ordem
        time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
    }
}

func fmtInt(n int) string {
    return fmt.Sprintf("%d", n)
}