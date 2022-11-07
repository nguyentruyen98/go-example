package main

func newCard() string {
	return "Five of Dinamonds"
}

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()

}
