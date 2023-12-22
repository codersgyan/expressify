package configs

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/codersgyan/expressify/internal/selector"
)

type Config string

const (
	ENV  Config = "ENV based"
	JSON Config = "File based (JSON)"
	YAML Config = "File based (YAML)"
)

var configs = []selector.Item{
	selector.Item(ENV),
	selector.Item(JSON),
	selector.Item(YAML),
}

func NewConfigSelector() *selector.Selector {
	var items []list.Item
	for _, item := range configs {
		items = append(items, list.Item(item))
	}
	return selector.NewSelector("\n😎 Choose a config style", items)
}
