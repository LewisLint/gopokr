package main

import "fmt"

func main() {
	cards := newd()
	shuf(cards)
	//	p1, cards, e := deal(cards, 2)
	//	if e != nil {
	//		panic(e)
	//	}
	//	de, cards, e := deal(cards, 2)
	//	if e != nil {
	//		panic(e)
	//	}
	//	fmt.Printf("P1: %+v\nd: %+v\nREM: %+v\nDeck Length:  %+v\n", p1, de, cards, len(cards))
	//	fmt.Printf("%v \n", cards)
	//	fmt.Printf("%v\n", len(cards))

	players := []Player{
		{Name: "you"},
		{Name: "AI", Ai: true},
	}

	for r := 0; r < 2; r++ {
		for i := range players {
			dlt, rem, err := deal(cards, 1)
			if err != nil {
				panic(err)
			}
			players[i].Hand = append(players[i].Hand, dlt...)
			cards = rem
		}
	}

	for _, player := range players {
		fmt.Printf("%s: %+v\n")
	}
}
