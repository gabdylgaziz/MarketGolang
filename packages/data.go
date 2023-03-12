package packages

import (
	"assignment/item"
	"assignment/user"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	nickname     = "postgres"
	password = "your_password"
	dbname   = "your_db_name"
)

func GetAccounts() []user.User{
	info := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, nickname, password, dbname)

	db, err := sql.Open("postgres", info)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var accounts []user.User
	res, err := db.Query(`SELECT * FROM acc;`)
	if err != nil {
		panic(err)
	}
	for res.Next(){
		var user user.User
		err = res.Scan(&user.Name, &user.Surname, &user.Balance, &user.Login, &user.Password)
		if err != nil {
			panic(err)
		}
		accounts = append(accounts, user)
	}
	return accounts
}



func GetItems() []item.Item{
	info := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, nickname, password, dbname)

	db, err := sql.Open("postgres", info)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var items []item.Item
	res, err := db.Query(`SELECT * FROM items;`)
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
