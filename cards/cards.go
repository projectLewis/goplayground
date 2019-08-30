package main

func main() {
	cards := newDeck()
	cards.saveToFile("deckOcards")
	// fmt.Println(deckFromFile("deckOcards"))
}
