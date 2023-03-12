package packages

import (
	"assignment/item"
	"sort"
	_ "github.com/lib/pq"
)

func FilterByPrice() []item.Item{
	items := GetItems()
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Price < items[j].Price
	})
	return items
}

func FilterByRating() []item.Item{
	items := GetItems()
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Rating < items[j].Rating
	})
	return items
}

