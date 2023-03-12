package main

import (
	"assignment/packages"
	"fmt"
)

func main(){
	for{
		fmt.Println("What do you want? 1 for auth, 2 for registration")
		var a int 
		fmt.Scan(&a)
		if a == 1{
			var login, password string
			fmt.Println("Enter your username and password in a single line: ")
			fmt.Scan(&login, &password)
			if packages.CheckAuth(login, password) == true{
				user := packages.GetData(login)
				fmt.Println("Do you want to continue with the console version(1) or with the frontend(2)?")
				var b int 
				fmt.Scan(&b)
				if b == 1{
					packages.MarketConsole(user)
				}else{
					packages.Market(user)
				}
				break
			}else{
				fmt.Println("Try again!")
			}
		}else{
			var Name, Surname, login, password string
			fmt.Println("Enter your name, surname, username and password: ")
			fmt.Scan(&Name, &Surname, &login, &password)
			packages.NewAcc(Name, Surname, login, password)
		}
		
	}
	
}