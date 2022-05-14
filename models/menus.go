package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	User            primitive.ObjectID `bson: "_id, omitempty"`
	Language        string             `bson: "language, omitempty"`
	Name            string             `bson: "name, omitempty"`
	Items           []string           `bson: "items, omitempyty"`
	ShowIngredients bool               `bson: "showIngredients, omitempty"`
	ShowAmounts     bool               `bson: "showAmounts, omitempty"`
}
