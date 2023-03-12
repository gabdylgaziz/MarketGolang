package packages

import (
	"fmt"
	"database/sql"
	"assignment/item"
	_ "github.com/lib/pq"
)

func SearchItems(name string) []item.Item{
	info := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, nickname, password, dbname)

	db, err := sql.Open("postgres", info)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var items []item.Item
	name = "%" + name + "%"
	res, err := db.Query(fmt.Sprintf("SELECT * FROM items WHERE Name LIKE '%s';", name))
	if err != nil {
		panic(err)
	}
	for res.Next(){
		var item item.Item
		err = res.Scan(&item.Name, &item.Price, &item.Rating)
		if err != nil {
			panic(err)
		}
		items = append(items, item)
	}
	return items
}
