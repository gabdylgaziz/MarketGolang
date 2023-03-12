package packages

import (
	"assignment/item"
	"assignment/user"
	"fmt"
	"html/template"
	"net/http"
)

var temp user.User

type allinone struct{
	Userdata user.User
	Items []item.Item
}


func main_page(w http.ResponseWriter, r *http.Request){
	res, _ := template.ParseFiles("client/index.html")
	var data allinone
	data.Userdata = temp
	data.Items = GetItems()
	res.Execute(w, data)
}

func sec_page(w http.ResponseWriter, r *http.Request){
	fmt.Println("GOODBYE")
}

func Market(user user.User){
	temp = user
	http.HandleFunc("/", main_page)
	http.HandleFunc("/goodbye", sec_page)
	http.ListenAndServe(":2004", nil)
}

func MarketConsole(user user.User){
	fmt.Println(fmt.Sprintf("Welcome %s %s!", temp.Name, temp.Surname))
	for{
		fmt.Println("What do you want? 1 to view all available items, 2 to sort by price, 3 to sort by rating, 4 to search for items by name, 5 to exit")
		var a int
		fmt.Scan(&a)
		if a == 1{
			items := GetItems()
			for i := 0; i < len(items); i++ {
				fmt.Println(fmt.Sprintf("%s Price: %d Rating: %d", items[i].Name,  items[i].Price,  items[i].Rating))
			}
		}else if a == 2{
			items := FilterByPrice()
			for i := 0; i < len(items); i++ {
				fmt.Println(fmt.Sprintf("%s Price: %d Rating: %d", items[i].Name,  items[i].Price,  items[i].Rating))
			}
		}else if a == 3{
			items := FilterByRating()
			for i := 0; i < len(items); i++ {
				fmt.Println(fmt.Sprintf("%s Price: %d Rating: %d", items[i].Name,  items[i].Price,  items[i].Rating))
			}
		}else if a == 4{
			var name string
			fmt.Println("Please type name of item: ")
			fmt.Scan(&name)
			items := SearchItems(name)
			for i := 0; i < len(items); i++ {
				fmt.Println(fmt.Sprintf("%s Price: %d Rating: %d", items[i].Name,  items[i].Price,  items[i].Rating))
			}
		}else{
			break
		}
	}
}


