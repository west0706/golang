package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var winRate [3]int
	winRate[0] = 0
	winRate[1] = 0
	winRate[2] = 0

	// for i := 0; i < 100; i++ {
	card := makeCard()
	// fmt.Println(card)

	mycard := shuffleCard(card)
	fmt.Println(mycard)

	// }
}

func makeCard() [20]int {
	var result [20]int
	for i := 0; i < 20; i++ {
		result[i] = (i + 1) % 10
		if result[i] == 0 {
			result[i] = 10
		}
	}
	return result
}

func shuffleCard(card [20]int) [20]int {
	var myCard [20]int
	fmt.Println(card)

	var index = 0
	for {
		flag := checkCard(card)
		if flag == 0 || index == 20 {
			break
		}

		s1 := rand.NewSource(time.Now().UnixNano())
		rand := rand.New(s1)
		randomNumber := rand.Intn(20)

		if card[randomNumber] != 0 {
			myCard[index] = card[randomNumber]
			card[randomNumber] = 0

			fmt.Println(myCard)
			fmt.Println(card)
			index++
		} else {
			continue
		}

		fmt.Println(randomNumber)

	}

	return myCard

}

func checkCard(card [20]int) int {
	flag := 0
	for _, temp := range card {
		flag = flag + temp
	}
	return flag
}
