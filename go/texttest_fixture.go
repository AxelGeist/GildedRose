package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func main() {
	fmt.Println("OMGHAI!")

	var items = []gildedrose.ItemInterface{
		// {"+5 Dexterity Vest", 10, 20},
		&gildedrose.AgedBrie{DaysLeftToSell: 2, Quality: 0},
		// {"Elixir of the Mongoose", 5, 7},
		&gildedrose.Sulfuras{DaysLeftToSell: 0, Quality: 80},
		&gildedrose.Sulfuras{DaysLeftToSell: -1, Quality: 80},
		&gildedrose.Backstage{DaysLeftToSell: 15, Quality: 20},
		&gildedrose.Backstage{DaysLeftToSell: 10, Quality: 49},
		&gildedrose.Backstage{DaysLeftToSell: 5, Quality: 49},
		// {"Conjured Mana Cake", 3, 6}, // <-- :O
	}

	days := 2
	var err error
	if len(os.Args) > 1 {
		days, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		days++
	}

	for day := 0; day < days; day++ {
		fmt.Printf("-------- day %d --------\n", day)
		fmt.Println("Name, SellIn, Quality")
		for i := 0; i < len(items); i++ {
			fmt.Println(items[i])
		}
		fmt.Println("")
		gildedrose.UpdateQuality(items)
	}
}
