package package_managers

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/codersgyan/expressify/internal/selector"
)

var packageManagers = []selector.Item{
	selector.Item("NPM"),
	selector.Item("PNPM"),
}

func NewPManagerSelector() *selector.Selector {
	var items []list.Item
	for _, item := range packageManagers {
		items = append(items, list.Item(item))
	}
	return selector.NewSelector("\n😎 Choose a package manager", items)
}
