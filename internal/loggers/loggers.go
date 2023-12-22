package loggers

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/codersgyan/expressify/internal/selector"
)

type Logger string

const (
	Winston Logger = "Winston"
	Bunyan  Logger = "Bunyan"
	Pino    Logger = "Pino"
)

var loggers = []selector.Item{
	selector.Item(Winston),
	selector.Item(Bunyan),
	selector.Item(Pino),
}

func NewLoggerSelector() *selector.Selector {
	var items []list.Item
	for _, item := range loggers {
		items = append(items, list.Item(item))
	}
	return selector.NewSelector("\n😎 Choose a logger", items)
}
