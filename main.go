package main

import (
	"fmt"
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

func main() {

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
		fmt.Printf("Your hand:\n\n")
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
			clear()
			fmt.Printf("\n")
			phv = eval(ph)
			dhv = eval(dh)
			fmt.Printf("Your hand:\n\n")
			//		fmt.Printf("%v\n\n\n", ph)
			prhand(ph, false)
			//		fmt.Printf("%d\n", phv)
			fmt.Printf("Dealer hand:\n\n")
			//		fmt.Printf("%v\n\n\n", dh)
			prhand(dh, false) //  prints the hand
			fmt.Printf("\n")
			var shit byte
			fmt.Printf("You pussy\nClick enter to continue")
			fmt.Scanln(&shit)
			for shit == 0 {
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
			fmt.Printf("Your hand:\n\n")
			//		fmt.Printf("%v\n\n\n", ph)
			prhand(ph, false)
			//		fmt.Printf("%d\n", phv)
			fmt.Printf("Dealer hand:\n\n")
			//		fmt.Printf("%v\n\n\n", dh)
			prhand(dh, false) //  prints the hand
			fmt.Printf("\n")

			//			fmt.Printf("%v\n%v\n", phv[0], dhv[0])
			var inp byte
			// only pussies fold
			switch {
			case phv[0] > dhv[0]:
				fmt.Printf("You win\nClick enter to continue")
				fmt.Scanln(&inp)
				if inp != 0 {
					continue
				}
			case phv[0] < dhv[0]:
				fmt.Printf("Dealer wins\nClick enter to continue")
				fmt.Scanln(&inp)
				if inp != 0 {
					continue
				}
			default:
				win := compare(phv, dhv)
				switch win {
				case 1:
					fmt.Printf("You win\n")
				case -1:
					fmt.Printf("Dealer wins\n")
				default:
					fmt.Printf("Tie\n")
				}
				fmt.Printf("Click enter to continue")
				fmt.Scanln(&inp)
				if inp != 0 {
					continue
				}
			}

		}
	}
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
	for range 3 {
		score := eval(dh)

		if score[0] == highcard {
			// look for near flush or straight
			if hfd(dh) || hsd(dh) {
				ddraw(1)
				return
			}
			ddraw(3)
			return
		}
		switch score[0] {
		case highcard:
			ddraw(3)
			continue

		case onepair:
			pairr := score[1]
			if pairr <= 6 {
				ddraw(2)
				continue
			} else if pairr <= 10 {
				ddraw(1)
			} else {
				return
			}
			continue
		case twopair:
			return
		case threekind:
			return
		case straight:
			return
		case flush:
			return
		case fullhouse:
			return
		case fourkind:
			return
		case straightflush:
			return
		default:
			// two pair tripple, straight, flush, etc
			return
		}
	}
}

func ddraw(count int) {
	// pick the shittiest cards and redraw them
	for range count {
		weak := 0
		for i := 1; i < len(dh); i++ {
			if dh[i].rank < dh[weak].rank {
				weak = i
			}
		}

		dh[weak] = deck[0]
		deck = deck[1:]
	}
}

func hfd(hand []card) bool {
	scunts := map[string]int{}
	for _, aksfdhklasjddhksfkljsdhksffslahsakjldffgyusalhdkjffjskjdldfhsjsahsjdkfsaadfgjhahsdfhk := range hand {
		scunts[aksfdhklasjddhksfkljsdhksffslahsakjldffgyusalhdkjffjskjdldfhsjsahsjdkfsaadfgjhahsdfhk.suit]++
	}
	for _, v := range scunts {
		if v >= 4 {
			return true
		}
	}
	return false
}

func hsd(hand []card) bool {
	ranks := make([]int, len(hand))
	for i, c := range hand {
		ranks[i] = c.rank
	}
	sort.Ints(ranks)

	for i := 0; i <= len(ranks)-4; i++ {
		window := ranks[i : i+4]

		if window[3]-window[0] == 3 {
			return true
		}
	}
	return false
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
