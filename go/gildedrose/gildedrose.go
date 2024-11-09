package gildedrose

const maxQuality = 50
const minQuality = 0

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
	for i := 0; i < len(items); i++ {
		item := items[i]
		item.updateQuality()
		item.updateDaysToSell()
	}
}

// AgedBrie
type AgedBrie struct {
	DaysLeftToSell int
	Quality        int
}

func (brie *AgedBrie) updateQuality() {
	if brie.DaysLeftToSell < 0 {
		brie.decreaseQualityIfMinNotReached(2)
	} else {
		brie.increaseQualityIfMaxNotReached(1)
	}
}

func (brie *AgedBrie) updateDaysToSell() {
	brie.DaysLeftToSell -= 1
}

func (brie *AgedBrie) increaseQualityIfMaxNotReached(increase int) {
	brie.Quality = min(maxQuality, brie.Quality+increase)
}

func (brie *AgedBrie) decreaseQualityIfMinNotReached(decrease int) {
	brie.Quality = max(minQuality, brie.Quality-decrease)
}

// BackStage
type Backstage struct {
	DaysLeftToSell int
	Quality        int
}

func (backstage *Backstage) updateQuality() {
	if backstage.DaysLeftToSell < 0 {
		backstage.resetQuality()
	} else if backstage.DaysLeftToSell <= 5 {
		backstage.increaseQualityIfMaxNotReached(3)
	} else if backstage.DaysLeftToSell <= 10 {
		backstage.increaseQualityIfMaxNotReached(5)
	} else {
		backstage.increaseQualityIfMaxNotReached(1)
	}
}

func (backstage *Backstage) resetQuality() {
	backstage.Quality = 0
}

func (backstage *Backstage) increaseQualityIfMaxNotReached(increase int) {
	backstage.Quality = min(maxQuality, backstage.Quality+increase)
}

func (backstage *Backstage) updateDaysToSell() {
	backstage.DaysLeftToSell -= 1
}

// Sulfuras
type Sulfuras struct {
	DaysLeftToSell int
	Quality        int
}

func (sulfuras *Sulfuras) updateQuality() {
	sulfuras.Quality = 80 // quality never changes
}

func (sulfuras *Sulfuras) updateDaysToSell() {
	sulfuras.DaysLeftToSell -= 1
}
