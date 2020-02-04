package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getCard(num int, except_num int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	var card_num int

	for {
		tmp_card_num := r1.Intn(num)
		if tmp_card_num != 0 && tmp_card_num != except_num {
			card_num = tmp_card_num
			break
		}
	}
	return card_num
}

func main() {

	//Card Set 1
	player1_card1 := getCard(10, 0)
	player2_card1 := getCard(10, player1_card1)

	//Card Set 2
	player1_card2 := getCard(10, 0)
	player2_card2 := getCard(10, player1_card2)

	player1_sum := player1_card1 + player1_card2
	player2_sum := player2_card1 + player2_card2

	fmt.Println(player1_card1, player1_card2, player1_sum)
	fmt.Println(player2_card1, player2_card2, player2_sum)

	// fmt.Println("Player1's CARD1, CARD2: ", player1_card1, player1_card2)
	// fmt.Println("Player2's CARD1, CARD2: ", player2_card1, player2_card2)

}
