package gildedrose

const (
	maxQuality = 50
	minQuality = 0
)

type ItemInterface interface {
	updateQuality()
	updateDaysToSell()
}

type Item struct {
	Name           string
	DaysLeftToSell int
	Quality        int
}

func UpdateQuality(items []ItemInterface) {
	for _, item := range items {
		item.updateQuality()
	}
}

type BaseItem struct {
	DaysLeftToSell int
	Quality        int
	IsConjured     bool
}

// AgedBrie
type AgedBrie struct {
	Base BaseItem
}

func (brie *AgedBrie) updateQuality() {
	if brie.Base.DaysLeftToSell < 0 {
		brie.changeQuality(-2)
	} else {
		brie.changeQuality(1)
	}
	brie.updateDaysToSell()
}

func (brie *AgedBrie) updateDaysToSell() {
	brie.Base.DaysLeftToSell -= 1
}

func (brie *AgedBrie) changeQuality(change int) {
	if brie.Base.IsConjured && change < 0 { // conjured items degrade twice as fast
		change *= 2
	}
	newQuality := brie.Base.Quality + change
	brie.Base.Quality = min(maxQuality, max(minQuality, newQuality))
}

// BackStage
type Backstage struct {
	Base BaseItem
}

func (backstage *Backstage) updateQuality() {
	if backstage.Base.DaysLeftToSell < 0 {
		backstage.resetQuality()
	} else if backstage.Base.DaysLeftToSell <= 5 {
		backstage.changeQuality(3)
	} else if backstage.Base.DaysLeftToSell <= 10 {
		backstage.changeQuality(5)
	} else {
		backstage.changeQuality(1)
	}
	backstage.updateDaysToSell()
}

func (backstage *Backstage) resetQuality() {
	backstage.Base.Quality = 0
}

func (backstage *Backstage) changeQuality(change int) {
	if backstage.Base.IsConjured && change < 0 { // conjured items degrade twice as fast
		change *= 2
	}
	newQuality := backstage.Base.Quality + change
	backstage.Base.Quality = min(maxQuality, max(minQuality, newQuality))
}

func (backstage *Backstage) updateDaysToSell() {
	backstage.Base.DaysLeftToSell -= 1
}

// Sulfuras
type Sulfuras struct {
	Base BaseItem
}

func (sulfuras *Sulfuras) updateQuality() {
	sulfuras.Base.Quality = 80 // quality never changes
	sulfuras.updateDaysToSell()
}

func (sulfuras *Sulfuras) updateDaysToSell() {
	sulfuras.Base.DaysLeftToSell -= 1
}

// NormalItem
type NormalItem struct {
	Name string
	Base BaseItem
}

// NormalItem Names
const (
	NormalItemPlus5DexterityVest string = "+5 Dexterity Vest"
	NormalItemElixirMongoose     string = "Elixir of the Mongoose"
	NormalItemManaCake           string = "Mana Cake"
)

func (item *NormalItem) updateQuality() {
	if item.Base.DaysLeftToSell < 0 {
		item.changeQuality(-2)
	} else {
		item.changeQuality(-1)
	}
	item.updateDaysToSell()
}

func (item *NormalItem) updateDaysToSell() {
	item.Base.DaysLeftToSell -= 1
}

func (item *NormalItem) changeQuality(change int) {
	if item.Base.IsConjured && change < 0 { // conjured items degrade twice as fast
		change *= 2
	}
	newQuality := item.Base.Quality + change
	item.Base.Quality = min(maxQuality, max(minQuality, newQuality))
}
