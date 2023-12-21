package databases

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/codersgyan/expressify/internal/selector"
)

var databases = []selector.Item{
	selector.Item("MongoDB"),
	selector.Item("PostgreSQL"),
	selector.Item("MySQL"),
}

func NewDatabaseSelector() *selector.Selector {
	var items []list.Item
	for _, item := range databases {
		items = append(items, list.Item(item))
	}
	return selector.NewSelector("\n😎 Choose a database", items)
}
