package test_frameworks

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/codersgyan/expressify/internal/selector"
)

var testFrameworks = []selector.Item{
	selector.Item("SuperTest with Jest"),
	selector.Item("Mocha with Chai HTTP"),
}

func NewTestFrameworkSelector() *selector.Selector {
	var items []list.Item
	for _, item := range testFrameworks {
		items = append(items, list.Item(item))
	}
	return selector.NewSelector("\n😎 Choose a test framework", items)
}
