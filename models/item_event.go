package models

import (
	"time"
)

type ItemEvent struct {
	Id        int
	IsNeeded  bool
	Timestamp time.Time
	Item      Item
	ItemId    int
}

func NewItemEvent(item Item, isNeeded bool) *ItemEvent {
	return &ItemEvent{
		IsNeeded:  isNeeded,
		Timestamp: time.Now(),
		Item:      item,
	}
}
