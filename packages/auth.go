package packages

import (
	"database/sql"
	"fmt"
	"assignment/user"
	_ "github.com/lib/pq"
)

func CheckAuth(login string, password string) bool{
	accounts := GetAccounts()
	for i := 0; i < len(accounts); i++ {
		if accounts[i].Login == login{
			if accounts[i].Password == password{
				return true
			}
		}
	}
	fmt.Println("Incorrect login or password, please try again")
	return false
}

func GetData(login string) user.User{
	accounts := GetAccounts()
	var user user.User
	for i := 0; i < len(accounts); i++ {
		if accounts[i].Login == login{
			user = accounts[i]
		}
	}
	return user
}

func NewAcc(name string, surname string, login string, passw string){
	info := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, nickname, password, dbname)

	db, err := sql.Open("postgres", info)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	res, err := db.Query(fmt.Sprintf("INSERT INTO acc VALUES('%s','%s', 0,'%s','%s')", name, surname, login, passw))
	if err != nil {
		panic(err)
	}
	defer res.Close()
}

