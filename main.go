package main

import (
	"fmt"
	"math/bits"
	"math/rand/v2"
	"sort"
	"strings"
)

type card struct {
	rank int    // Ace 1/11 (depending on situation), Jack, Queen, King are all worth 10
	suit string // e.g. Spades
}

const (
	highcard      = iota // 0
	onepair              // 1
	twopair              // 2
	threekind            // 3
	straight             // 4
	flush                // 5
	fullhouse            // 6
	fourkind             // 7
	straightflush        // 8
)

var turn string
var i int
var soft bool = false
var q bool = false
var suits = []string{"♠", "♥", "♦", "♣"}
var deck, ph, dh []card
var phv, dhv []int // hand values
var b string = "\x1b[1m"
var n int = 0
var drawn bool = false
var wins, losses, ties int
var difficulty = medium

const (
	easy = iota
	medium
	hard
	fuckYouMode
)

func main() {
	difficulty = chooseDifficulty()
	for q == false {
		clear()
		if n == 0 {
			drawn = false
			deck = newd()
			ph, dh = nil, nil
			n++
			for range 5 {
				ph = append(ph, deck[0])
				deck = deck[1:]
				dh = append(dh, deck[0])
				deck = deck[1:]
			}
		}
		fmt.Printf("\n")
		phv = eval(ph)
		dhv = eval(dh)
		printPlayerHeading()
		//		fmt.Printf("%v\n\n\n", ph)
		prhand(ph, false)
		//		fmt.Printf("%d\n", phv)
		fmt.Printf("Dealer hand:\n\n")
		//		fmt.Printf("%v\n\n\n", dh)
		prhand(dh, true) //  prints the hand
		fmt.Printf("\n")
		//		fmt.Printf("%d\n", dhv)

		var inp string
		if drawn == true {
			fmt.Printf("(E)nd turn, (f)old, or (q)uit: ")
		} else {

			fmt.Printf("(R)edraw, (e)nd turn, (f)old, or (q)uit: ")
		}
		fmt.Scanln(&inp)
		fmt.Printf("\n")
		switch inp {
		case "r", "R":

			var shit string
			fmt.Printf(": ")
			fmt.Scanln(&shit)
			if drawn == true && shit != "" {
				continue
			}
			redraw(shit)
			drawn = true

		case "f", "F":
			n = 0
			losses++
			clear()
			fmt.Printf("\n")
			phv = eval(ph)
			dhv = eval(dh)
			printPlayerHeading()
			//		fmt.Printf("%v\n\n\n", ph)
			prhand(ph, false)
			//		fmt.Printf("%d\n", phv)
			fmt.Printf("Dealer hand:\n\n")
			//		fmt.Printf("%v\n\n\n", dh)
			prhand(dh, false) //  prints the hand
			fmt.Printf("\n")
			var shit byte
			fmt.Printf("You pussy; the dealer wins by fold.\n")
			fmt.Printf("Click enter to continue")
			fmt.Scanln(&shit)
			for shit == '0' {
			}
			n = 0
			continue
		case "q", "Q":
			q = true
		case "e", "E":
			dealer()
			n = 0
			clear()
			fmt.Printf("\n")
			phv = eval(ph)
			dhv = eval(dh)
			printPlayerHeading()
			//		fmt.Printf("%v\n\n\n", ph)
			prhand(ph, false)
			//		fmt.Printf("%d\n", phv)
			fmt.Printf("Dealer hand:\n\n")
			//		fmt.Printf("%v\n\n\n", dh)
			prhand(dh, false) //  prints the hand
			fmt.Printf("\n")

			result := compare(phv, dhv)
			var inp byte
			switch result {
			case 1:
				wins++
				fmt.Printf("You win with %s.\n", handName(phv[0]))
			case -1:
				losses++
				fmt.Printf("Dealer wins with %s.\n", handName(dhv[0]))
			default:
				ties++
				fmt.Printf("Tie: both hands are %s.\n", handName(phv[0]))
			}
			fmt.Printf("Click enter to continue")
			fmt.Scanln(&inp)
			if inp != 0 {
				continue
			}

		}
	}
}

func chooseDifficulty() int {
	for {
		fmt.Print("Choose difficulty: (e)asy, (m)edium, or (h)ard: ")
		var input string
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Printf("Could not read difficulty: %v\n", err)
			return medium
		}
		switch strings.ToLower(input) {
		case "e", "easy":
			return easy
		case "m", "medium":
			return medium
		case "h", "hard":
			return hard
		case "fuckyou":
			return fuckYouMode
		default:
			fmt.Println("Please enter e, m, or h.")
		}
	}
}

func printPlayerHeading() {
	played := wins + losses + ties
	winPercent := 0.0
	if played > 0 {
		winPercent = float64(wins) * 100 / float64(played)
	}

	color := "\x1b[32m"
	if winPercent < 40 {
		color = "\x1b[31m"
	} else if winPercent < 60 {
		color = "\x1b[33m"
	}
	fmt.Printf("%sWin %.1f%%\x1b[0m — Your hand:\n\n", color, winPercent)
}

func handName(category int) string {
	names := [...]string{
		"high card",
		"one pair",
		"two pair",
		"three of a kind",
		"a straight",
		"a flush",
		"a full house",
		"four of a kind",
		"a straight flush",
	}
	if category < 0 || category >= len(names) {
		return "an unknown hand"
	}
	return names[category]
}

func newd() []card {
	deck := make([]card, 0, 52)

	for _, suit := range suits {
		for rank := 1; rank <= 13; rank++ {
			deck = append(deck, card{rank: rank, suit: suit})
		}
	}
	for range 5 {
		shuf(deck)
	}
	return deck
}

func shuf(deck []card) {
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
}
func reval(cunts map[int]int, wanted int) []int {
	r := make([]int, 0)
	for rank, count := range cunts {
		if count == wanted {
			r = append(r, penis(rank))
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(r)))
	return r

}

func urfat(ranks []int) (int, bool) {
	if len(ranks) != 5 {
		return 0, false
	}

	cuntsecutive := true

	for i := 1; i < len(ranks); i++ {
		if ranks[i] != ranks[i-1]+1 {
			cuntsecutive = false
			break
		}
	}

	if cuntsecutive {
		return penis(ranks[4]), true // shitty ranks
	}

	// good ranks
	if ranks[0] == 1 && ranks[1] == 10 && ranks[2] == 11 && ranks[3] == 12 && ranks[4] == 13 {
		return 14, true
	}
	return 0, false
}

func eval(user []card) []int {
	cunts := make(map[int]int)
	suits := make(map[string]int)
	ranks := make([]int, 0, len(user))
	for _, i := range user {
		cunts[i.rank]++
		suits[i.suit]++
		ranks = append(ranks, i.rank)

	}
	sort.Ints(ranks)

	// flush
	hasFlush := false
	for _, v := range suits {
		if v == 5 {
			hasFlush = true
			break
		}
	}

	fuuuuuuuuuuck, hasStraight := urfat(ranks)
	if len(ranks) == 5 {
		hasStraight = true
		for i := 1; i < len(ranks); i++ {
			if ranks[i] != ranks[i-1]+1 {
				hasStraight = false
				break
			}
		}
	}

	pairs := 0
	three := false
	four := false

	for _, count := range cunts {
		switch count {
		case 2:
			pairs++
		case 3:
			three = true
		case 4:
			four = true
		}
	}
	switch {
	case hasFlush && hasStraight:
		return []int{straightflush, fuuuuuuuuuuck}
		return []int{straightflush}
	case four:
		return []int{fourkind, reval(cunts, 4)[0], reval(cunts, 1)[0]}
		return []int{fourkind}
	case three && pairs == 1:
		return []int{
			fullhouse,
			reval(cunts, 3)[0],
			reval(cunts, 2)[0],
		}
	case hasFlush:
		return append([]int{flush}, hr(ranks)...)
	case hasStraight:
		return []int{straight, fuuuuuuuuuuck}
	case three:
		tr := 0
		k := make([]int, 0, 2)

		for r, c := range cunts {
			if c == 3 {
				tr = penis(r)
				continue
			}
			for range c {
				k = append(k, penis(r))
			}
		}

		sort.Sort(sort.Reverse(sort.IntSlice(k)))
		return append([]int{threekind, tr}, k...)

	case pairs == 2:
		rranks := make([]int, 0, 2)
		wtf := 0

		for rank, cunt := range cunts {
			if cunt == 2 {
				rranks = append(rranks, penis(rank))
			} else if cunt == 1 {
				wtf = penis(rank)
			}
		}

		sort.Sort(sort.Reverse(sort.IntSlice(rranks)))
		return []int{twopair, rranks[0], rranks[1], wtf}

	case pairs == 1:
		rrank := 0
		kickers := make([]int, 0, 3)

		for rank, count := range cunts {
			if count == 2 {
				rrank = penis(rank)
				continue
			}

			for range count {
				kickers = append(kickers, penis(rank))
			}
		}

		sort.Sort(sort.Reverse(sort.IntSlice(kickers)))
		return append([]int{onepair, rrank}, kickers...)
	case true:
		return append([]int{highcard}, hr(ranks)...)
		hcs := make([]int, 0, len(ranks))
		for _, rank := range ranks {
			hcs = append(hcs, penis(rank))
		}
		sort.Sort(sort.Reverse(sort.IntSlice(hcs)))
		return append([]int{highcard}, hcs...)
	}
	return []int{highcard}

}

func (c card) String() string {
	rank := fmt.Sprint(c.rank)

	switch c.rank {
	case 1:
		rank = "A"
	case 11:
		rank = "J"
	case 12:
		rank = "Q"
	case 13:
		rank = "K"
	}
	return c.suit + rank
}

func clear() {
	// \x1b[H moves cursor to top left, \x1b[2J clears the screen
	fmt.Print("\x1b[H\x1b[2J")
}

func stats() {

}

func checkcard(targ []card) card {
	lcard := targ[len(targ)-1]

	return lcard
}

func (c card) lines() [3]string {
	rank := fmt.Sprint(c.rank)
	switch c.rank {
	case 1:
		rank = "A"
	case 11:
		rank = "J"
	case 12:
		rank = "Q"
	case 13:
		rank = "K"
	}

	width := 3

	leftp := (width - len(rank)) / 2
	center := strings.Repeat(" ", leftp) + rank + strings.Repeat(" ", width-len(rank)-leftp)

	return [3]string{
		c.suit + strings.Repeat(" ", width-1), center, strings.Repeat(" ", width-1) + c.suit,
	}
}

func prhand(hand []card, hidden bool) {
	back := [4]string{"?  ", " ? ", "  ?"}

	for row := range 3 {
		for _, c := range hand {
			color := "\x1b[47;30m" // black default
			if hidden {
				color := "\x1b[47;30m" // black on white for middle row
				if row == 0 || row == 2 {
					color = "\x1b[47;31m" // red on white background
				}
				fmt.Printf("%s%s%s\x1b[0m ", b, color, back[row])
				continue
			}
			if c.suit == "♥" || c.suit == "♦" {
				color = "\x1b[47;31m"
			}
			fmt.Printf("%s%s%s\x1b[0m ", b, color, c.lines()[row])
		}
		fmt.Printf("\n")
	}
}

func redraw(shit string) {
	for _, char := range shit {
		if char < '1' || char > '5' {
			fmt.Printf("Input integers 1-5\n")
			return
		} else {
			i := int(char - '1')
			ph[i] = deck[0]
			deck = deck[1:]
		}

	}
}

func compare(a, b []int) int {
	// a and b are from eval
	if len(a) == 0 || len(b) == 0 {
		return 0
	}
	//compare hand catagory

	if a[0] != b[0] {
		if a[0] > b[0] {
			return 1
		}
		return -1
	}

	// same catagory. ties
	for i := 1; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			if a[i] > b[i] {
				return 1
			}
			return -1
		}
	}

	return 0
}

func dealer() {
	for range dealerRedrawLimit(difficulty) {
		var discards []int
		switch difficulty {
		case easy:
			discards = easyDealerDiscards(dh)
		case hard:
			discards = hardDealerDiscards(dh, 300)
		case fuckYouMode:
			discards = fuckYouDiscards(dh, ph, deck, 300)
		default:
			discards = dealerDiscards(dh)
		}
		if len(discards) == 0 {
			return
		}
		for _, index := range discards {
			dh[index] = deck[0]
			deck = deck[1:]
		}
	}
}

func dealerRedrawLimit(difficulty int) int {
	if difficulty == fuckYouMode {
		return 5
	}
	return 1
}

func fuckYouDiscards(hand, opponent, available []card, trials int) []int {
	if len(hand) != 10 || len(opponent) != 10 || trials <= 0 {
		return nil
	}

	opponentScore := eval(opponent)
	bestMask, bestWins := 0, -1
	for mask := 0; mask < 1<<len(hand); mask++ {
		discardCount := bits.OnesCount(uint(mask))
		if discardCount > len(available) {
			continue
		}

		wins := 0
		for range trials {
			drawPile := append([]card(nil), available...)
			rand.Shuffle(len(drawPile), func(i, j int) {
				drawPile[i], drawPile[j] = drawPile[j], drawPile[i]
			})

			candidate := append([]card(nil), hand...)
			nextDraw := 0
			for i := range candidate {
				if mask&(1<<i) != 0 {
					candidate[i] = drawPile[nextDraw]
					nextDraw++
				}
			}
			if compare(eval(candidate), opponentScore) > 0 {
				wins++
			}
		}

		if wins > bestWins {
			bestMask, bestWins = mask, wins
		}
	}

	keep := make([]bool, len(hand))
	for i := range hand {
		keep[i] = bestMask&(1<<i) == 0
	}
	return discardedIndexes(keep)
}

func easyDealerDiscards(hand []card) []int {
	score := eval(hand)
	if score[0] != highcard {
		return nil
	}

	keep := make([]bool, len(hand))
	highCards := make([]int, 0, 2)
	for i, c := range hand {
		if penis(c.rank) >= 11 {
			highCards = append(highCards, i)
		}
	}
	sort.Slice(highCards, func(i, j int) bool {
		return penis(hand[highCards[i]].rank) > penis(hand[highCards[j]].rank)
	})
	for _, i := range highCards {
		keep[i] = true
	}
	return discardedIndexes(keep)
}

func hardDealerDiscards(hand []card, trials int) []int {
	if trials <= 0 || len(hand) != 5 {
		return dealerDiscards(hand)
	}

	deck := make([]card, 0, 52-len(hand))
	for _, suit := range suits {
		for rank := 1; rank <= 13; rank++ {
			candidate := card{rank: rank, suit: suit}
			if !containsCard(hand, candidate) {
				deck = append(deck, candidate)
			}
		}
	}

	bestMask, bestWins := 0, -1.0
	for mask := 0; mask < 1<<len(hand); mask++ {
		discards := bits.OnesCount(uint(mask))
		if discards+5 > len(deck) {
			continue
		}

		wins := 0.0
		for range trials {
			drawPile := append([]card(nil), deck...)
			rand.Shuffle(len(drawPile), func(i, j int) {
				drawPile[i], drawPile[j] = drawPile[j], drawPile[i]
			})

			dealerHand := append([]card(nil), hand...)
			for i := range dealerHand {
				if mask&(1<<i) != 0 {
					dealerHand[i] = drawPile[0]
					drawPile = drawPile[1:]
				}
			}
			opponentHand := append([]card(nil), drawPile[:5]...)
			switch compare(eval(dealerHand), eval(opponentHand)) {
			case 1:
				wins++
			case 0:
				wins += 0.5
			}
		}

		if wins > bestWins {
			bestMask, bestWins = mask, wins
		}
	}

	keep := make([]bool, len(hand))
	for i := range hand {
		keep[i] = bestMask&(1<<i) == 0
	}
	return discardedIndexes(keep)
}

func containsCard(hand []card, target card) bool {
	for _, c := range hand {
		if c == target {
			return true
		}
	}
	return false
}

func dealerDiscards(hand []card) []int {
	score := eval(hand)
	keep := make([]bool, len(hand))
	rankCounts := make(map[int]int)
	for _, c := range hand {
		rankCounts[c.rank]++
	}

	switch score[0] {
	case onepair, twopair:
		for i, c := range hand {
			keep[i] = rankCounts[c.rank] == 2
		}
	case threekind:
		for i, c := range hand {
			keep[i] = rankCounts[c.rank] == 3
		}
	case fourkind:
		for i, c := range hand {
			keep[i] = rankCounts[c.rank] == 4
		}
	case highcard:
		suitIndexes := make(map[string][]int)
		for i, c := range hand {
			suitIndexes[c.suit] = append(suitIndexes[c.suit], i)
		}
		for _, indexes := range suitIndexes {
			if len(indexes) == 4 {
				for _, i := range indexes {
					keep[i] = true
				}
				return discardedIndexes(keep)
			}
		}

		if indexes := straightDrawIndexes(hand); len(indexes) == 4 {
			for _, i := range indexes {
				keep[i] = true
			}
			return discardedIndexes(keep)
		}

		highCards := make([]int, 0, 2)
		for i, c := range hand {
			if penis(c.rank) >= 11 {
				highCards = append(highCards, i)
			}
		}
		sort.Slice(highCards, func(i, j int) bool {
			return penis(hand[highCards[i]].rank) > penis(hand[highCards[j]].rank)
		})
		if len(highCards) > 2 {
			highCards = highCards[:2]
		}
		if len(highCards) == 0 {
			for i, c := range hand {
				if c.rank == 1 {
					highCards = append(highCards, i)
					break
				}
			}
		}
		for _, i := range highCards {
			keep[i] = true
		}
	default:
		return nil
	}

	return discardedIndexes(keep)
}

func discardedIndexes(keep []bool) []int {
	discards := make([]int, 0, len(keep))
	for i, kept := range keep {
		if !kept {
			discards = append(discards, i)
		}
	}
	return discards
}

func straightDrawIndexes(hand []card) []int {
	best := []int(nil)
	bestOuts := 0
	for mask := 0; mask < 1<<len(hand); mask++ {
		if bits.OnesCount(uint(mask)) != 4 {
			continue
		}

		ranks := make(map[int]bool, 4)
		indexes := make([]int, 0, 4)
		for i, c := range hand {
			if mask&(1<<i) == 0 {
				continue
			}
			ranks[penis(c.rank)] = true
			indexes = append(indexes, i)
		}
		if len(ranks) != 4 {
			continue
		}

		outs := make(map[int]bool)
		for low := 1; low <= 10; low++ {
			present, missing := 0, 0
			for rank := low; rank < low+5; rank++ {
				held := ranks[rank]
				if rank == 1 {
					held = ranks[14]
				}
				if held {
					present++
				} else {
					missing = rank
				}
			}
			if present == 4 {
				outs[missing] = true
			}
		}
		if len(outs) > bestOuts {
			best = indexes
			bestOuts = len(outs)
		}
	}
	return best
}

func penis(rank int) int {
	if rank == 1 {
		return 14 // this motherfucker is an ace, so we cant have it be a 1
	} else {
		return rank // we already have the card, and we're just changing the aces
	}
}

func hr(ranks []int) []int { // we dont like HR
	hcs := make([]int, 0, len(ranks))
	for _, rank := range ranks {
		hcs = append(hcs, penis(rank))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(hcs)))
	return hcs
}
