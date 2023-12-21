package coding_styles

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/codersgyan/expressify/internal/selector"
)

var codingStyles = []selector.Item{
	selector.Item("Functional"),
	selector.Item("Object Oriented"),
}

func NewCodingStyleSelector() *selector.Selector {
	var items []list.Item
	for _, item := range codingStyles {
		items = append(items, list.Item(item))
	}
	return selector.NewSelector("😎 Choose a coding style", items)
}
