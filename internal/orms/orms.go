package orms

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/codersgyan/expressify/internal/selector"
)

type ORM string

const (
	Mongoose  ORM = "Mongoose"
	Sequelize ORM = "Sequelize"
	TypeORM   ORM = "TypeORM"
	Prisma    ORM = "Prisma"
	None      ORM = "None"
)

var orms = []selector.Item{
	selector.Item(Mongoose),
	selector.Item(Sequelize),
	selector.Item(TypeORM),
	selector.Item(Prisma),
	selector.Item(None),
}

func NewORMSelector() *selector.Selector {
	var items []list.Item
	for _, item := range orms {
		items = append(items, list.Item(item))
	}
	return selector.NewSelector("\n😎 Choose an ORM", items)
}
