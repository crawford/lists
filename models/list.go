package models

import (
	"encoding/json"
)

type List struct {
	Id     string `sql:"type:uuid;"`
	//Id     int
	Name   string
	Items  []Item
	User   User
	UserId string `sql:"type:uuid;"`
	//UserId int
}

func (l List) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":    l.Id,
		"name":  l.Name,
		"items": l.Items,
	})
}
