package test_frameworks

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/codersgyan/expressify/internal/selector"
)

type TestFramwork string

const (
	SuperTestWithJest TestFramwork = "SuperTest with Jest"
	MochaWithChaiHTTP TestFramwork = "Mocha with Chai HTTP"
)

var testFrameworks = []selector.Item{
	selector.Item(SuperTestWithJest),
	selector.Item(MochaWithChaiHTTP),
}

func NewTestFrameworkSelector() *selector.Selector {
	var items []list.Item
	for _, item := range testFrameworks {
		items = append(items, list.Item(item))
	}
	return selector.NewSelector("\n😎 Choose a test framework", items)
}
