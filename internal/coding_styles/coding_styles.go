package coding_styles

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/codersgyan/expressify/internal/selector"
)

type CodingStyle string

const (
	Functional     CodingStyle = "Functional"
	ObjectOriented CodingStyle = "Object Oriented"
)

var codingStyles = []selector.Item{
	selector.Item(Functional),
	selector.Item(ObjectOriented),
}

func NewCodingStyleSelector() *selector.Selector {
	var items []list.Item
	for _, item := range codingStyles {
		items = append(items, list.Item(item))
	}
	return selector.NewSelector("\n😎 Choose a coding style", items)
}
