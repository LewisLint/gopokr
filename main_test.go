package main

import (
	"reflect"
	"testing"
)

func TestDealerDiscards(t *testing.T) {
	tests := []struct {
		name string
		hand []card
		want []int
	}{
		{
			name: "keeps the pair",
			hand: []card{
				{rank: 7, suit: "♠"}, {rank: 7, suit: "♥"},
				{rank: 2, suit: "♦"}, {rank: 5, suit: "♣"}, {rank: 9, suit: "♠"},
			},
			want: []int{2, 3, 4},
		},
		{
			name: "keeps a four-card flush draw",
			hand: []card{
				{rank: 2, suit: "♠"}, {rank: 5, suit: "♠"},
				{rank: 8, suit: "♠"}, {rank: 11, suit: "♠"}, {rank: 13, suit: "♥"},
			},
			want: []int{4},
		},
		{
			name: "keeps a four-card straight draw",
			hand: []card{
				{rank: 2, suit: "♠"}, {rank: 3, suit: "♥"},
				{rank: 4, suit: "♦"}, {rank: 5, suit: "♣"}, {rank: 9, suit: "♠"},
			},
			want: []int{4},
		},
		{
			name: "keeps a wheel draw",
			hand: []card{
				{rank: 1, suit: "♠"}, {rank: 2, suit: "♥"},
				{rank: 3, suit: "♦"}, {rank: 4, suit: "♣"}, {rank: 9, suit: "♠"},
			},
			want: []int{4},
		},
		{
			name: "keeps a made straight",
			hand: []card{
				{rank: 2, suit: "♠"}, {rank: 3, suit: "♥"},
				{rank: 4, suit: "♦"}, {rank: 5, suit: "♣"}, {rank: 6, suit: "♠"},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dealerDiscards(tt.hand); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("dealerDiscards() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandName(t *testing.T) {
	tests := []struct {
		category int
		want     string
	}{
		{highcard, "high card"},
		{onepair, "one pair"},
		{twopair, "two pair"},
		{threekind, "three of a kind"},
		{straight, "a straight"},
		{flush, "a flush"},
		{fullhouse, "a full house"},
		{fourkind, "four of a kind"},
		{straightflush, "a straight flush"},
		{-1, "an unknown hand"},
	}

	for _, tt := range tests {
		if got := handName(tt.category); got != tt.want {
			t.Errorf("handName(%d) = %q, want %q", tt.category, got, tt.want)
		}
	}
}

func TestEvalAceHighStraight(t *testing.T) {
	straightHand := []card{
		{rank: 1, suit: "♠"}, {rank: 10, suit: "♥"}, {rank: 11, suit: "♦"},
		{rank: 12, suit: "♣"}, {rank: 13, suit: "♠"},
	}
	if got, want := eval(straightHand), []int{straight, 14}; !reflect.DeepEqual(got, want) {
		t.Errorf("eval(ace-high straight) = %v, want %v", got, want)
	}

	royalFlush := []card{
		{rank: 1, suit: "♠"}, {rank: 10, suit: "♠"}, {rank: 11, suit: "♠"},
		{rank: 12, suit: "♠"}, {rank: 13, suit: "♠"},
	}
	if got, want := eval(royalFlush), []int{straightflush, 14}; !reflect.DeepEqual(got, want) {
		t.Errorf("eval(royal flush) = %v, want %v", got, want)
	}
}

func TestEasyDealerDiscardsOnlyHighCardHands(t *testing.T) {
	highCardHand := []card{
		{rank: 2, suit: "♠"}, {rank: 5, suit: "♥"},
		{rank: 8, suit: "♦"}, {rank: 11, suit: "♣"}, {rank: 13, suit: "♠"},
	}
	if got, want := easyDealerDiscards(highCardHand), []int{0, 1, 2}; !reflect.DeepEqual(got, want) {
		t.Errorf("easyDealerDiscards(high card) = %v, want %v", got, want)
	}

	pairHand := []card{
		{rank: 7, suit: "♠"}, {rank: 7, suit: "♥"},
		{rank: 2, suit: "♦"}, {rank: 5, suit: "♣"}, {rank: 9, suit: "♠"},
	}
	if got := easyDealerDiscards(pairHand); len(got) != 0 {
		t.Errorf("easyDealerDiscards(pair) = %v, want no discards", got)
	}
}

func TestHardDealerDiscardsAreValid(t *testing.T) {
	hand := []card{
		{rank: 2, suit: "♠"}, {rank: 5, suit: "♥"},
		{rank: 8, suit: "♦"}, {rank: 11, suit: "♣"}, {rank: 13, suit: "♠"},
	}

	discards := hardDealerDiscards(hand, 2)
	seen := make(map[int]bool)
	for _, index := range discards {
		if index < 0 || index >= len(hand) || seen[index] {
			t.Fatalf("hardDealerDiscards returned invalid/duplicate index %d", index)
		}
		seen[index] = true
	}
}

func TestDealerRedrawLimit(t *testing.T) {
	if got := dealerRedrawLimit(fuckYouMode); got != 5 {
		t.Errorf("dealerRedrawLimit(fuckYouMode) = %d, want 5", got)
	}
	for _, mode := range []int{easy, medium, hard} {
		if got := dealerRedrawLimit(mode); got != 1 {
			t.Errorf("dealerRedrawLimit(%d) = %d, want 1", mode, got)
		}
	}
}

func TestFuckYouDiscardsTargetKnownOpponent(t *testing.T) {
	hand := []card{
		{rank: 1, suit: "♠"}, {rank: 13, suit: "♥"}, {rank: 12, suit: "♣"},
		{rank: 11, suit: "♦"}, {rank: 9, suit: "♣"},
	}
	opponent := []card{
		{rank: 2, suit: "♠"}, {rank: 2, suit: "♥"}, {rank: 13, suit: "♣"},
		{rank: 12, suit: "♦"}, {rank: 11, suit: "♥"},
	}
	available := []card{{rank: 1, suit: "♦"}}

	if got, want := fuckYouDiscards(hand, opponent, available, 10), []int{1}; !reflect.DeepEqual(got, want) {
		t.Errorf("fuckYouDiscards() = %v, want %v", got, want)
	}
}

func TestFuckYouDiscardsKeepsHandThatAlreadyWins(t *testing.T) {
	hand := []card{
		{rank: 8, suit: "♠"}, {rank: 8, suit: "♥"}, {rank: 8, suit: "♦"},
		{rank: 13, suit: "♣"}, {rank: 13, suit: "♠"},
	}
	opponent := []card{
		{rank: 2, suit: "♠"}, {rank: 2, suit: "♥"}, {rank: 5, suit: "♦"},
		{rank: 7, suit: "♣"}, {rank: 9, suit: "♠"},
	}
	available := []card{{rank: 3, suit: "♠"}, {rank: 4, suit: "♥"}}

	if got := fuckYouDiscards(hand, opponent, available, 10); len(got) != 0 {
		t.Errorf("fuckYouDiscards() = %v, want no discards for a hand that already wins", got)
	}
}
