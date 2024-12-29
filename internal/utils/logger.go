package utils

import (
    "fmt"
    "log"
    "os"
)

// Logger é a interface de log
type Logger interface {
    Info(component, message string)
    Error(component, message string)
}

type SimpleLogger struct {
    logger *log.Logger
}

// NewLogger cria uma instância de SimpleLogger
func NewLogger() Logger {
    return &SimpleLogger{
        logger: log.New(os.Stdout, "", log.LstdFlags),
    }
}

// Info registra informações gerais
func (l *SimpleLogger) Info(component, message string) {
    l.logger.Printf("[INFO] [%s] %s\n", component, message)
}

// Error registra mensagens de erro
func (l *SimpleLogger) Error(component, message string) {
    l.logger.Printf("[ERROR] [%s] %s\n", component, message)
}

// Apenas para referência, se quiser um helper:
func Logf(format string, args ...interface{}) {
    fmt.Printf(format, args...)
}
