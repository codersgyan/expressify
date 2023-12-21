package languages

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/codersgyan/expressify/internal/selector"
)

var languageList = []selector.Item{
	selector.Item("JavaScript"),
	selector.Item("TypeScript"),
}

func NewLanguageSelector() *selector.Selector {
	var items []list.Item
	for _, item := range languageList {
		items = append(items, list.Item(item))
	}
	return selector.NewSelector("😎 Choose a language", items)
}
