package main

//import "fmt"

func main(){
	//var card string ="Ace of Spades"
	// card := newCard()
	// fmt.Println(card)

		//fmt.Println((cards))

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	//cards := deck {newCard(),"Ace of Diamonds"}
	//cards = append(cards, "Six of Spades")
	
	//cards.print()

	//cards := newDeck()
	//cards.print()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	//cards := newDeck()
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")

	//cards := newDeckFromFile("my_cards")
	//cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
	

}

// func newCard() string{
// 	return "Five of  Diamonds"
// }