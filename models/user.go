package models

import (
	"encoding/json"
)

type User struct {
	Id     string `sql:"type:uuid; unique; not null"`
	//Id    int
	Lists []List
}

func (u User) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":    u.Id,
		"lists": u.Lists,
	})
}
