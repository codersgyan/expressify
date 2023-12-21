package configs

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/codersgyan/expressify/internal/selector"
)

var configs = []selector.Item{
	selector.Item("ENV based"),
	selector.Item("File based (JSON)"),
	selector.Item("File based (YAML)"),
}

func NewConfigSelector() *selector.Selector {
	var items []list.Item
	for _, item := range configs {
		items = append(items, list.Item(item))
	}
	return selector.NewSelector("\n😎 Choose a config style", items)
}
