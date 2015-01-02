package models

import (
	"encoding/json"
)

type Item struct {
	Id     string `sql:"type:uuid;"`
	Name   string `sql:"type:varchar(255);"`
	ListId string `sql:"type:uuid;"`
	List   List
}

/*func (i *Item) List() *List {
	if i.list != nil {
		return list
	}
	db = db.Model(i).Related(i.list)
	if db.RecordNotFound() {
		panic(fmt.Sprintf("orphaned item (%#v)", i))
	}
	db.Scan(i.list)
	return i.list
}*/

func (i Item) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":   i.Id,
		"name": i.Name,
	})
}
