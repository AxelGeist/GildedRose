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
		&gildedrose.NormalItem{Name: gildedrose.NormalItemPlus5DexterityVest, Base: gildedrose.BaseItem{DaysLeftToSell: 10, Quality: 20, IsConjured: false}},
		&gildedrose.AgedBrie{Base: gildedrose.BaseItem{DaysLeftToSell: 2, Quality: 0, IsConjured: false}},
		&gildedrose.NormalItem{Name: gildedrose.NormalItemElixirMongoose, Base: gildedrose.BaseItem{DaysLeftToSell: 5, Quality: 7, IsConjured: false}},
		&gildedrose.Sulfuras{Base: gildedrose.BaseItem{DaysLeftToSell: 0, Quality: 80, IsConjured: false}},
		&gildedrose.Sulfuras{Base: gildedrose.BaseItem{DaysLeftToSell: -1, Quality: 80, IsConjured: false}},
		&gildedrose.Backstage{Base: gildedrose.BaseItem{DaysLeftToSell: 15, Quality: 20, IsConjured: false}},
		&gildedrose.Backstage{Base: gildedrose.BaseItem{DaysLeftToSell: 10, Quality: 49, IsConjured: false}},
		&gildedrose.Backstage{Base: gildedrose.BaseItem{DaysLeftToSell: 5, Quality: 49, IsConjured: false}},
		&gildedrose.NormalItem{Name: gildedrose.NormalItemManaCake, Base: gildedrose.BaseItem{DaysLeftToSell: 3, Quality: 6, IsConjured: true}},
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
