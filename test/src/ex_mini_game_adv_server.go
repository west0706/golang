package main

import (
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"time"
)

func main() {
	myListen, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var channel_game []string

	for {
		connect, err := myListen.Accept()
		// =========== Listen ==========
		if err != nil {
			continue
		}
		go ConnectHandler(connect, channel_game)

		defer connect.Close()
	}

	defer func() {
		myListen.Close()
	}()
}

func ConnectHandler(connect net.Conn, channel_game []string) {
	recvBuf := make([]byte, 2048) //receive buffer: 4kb

	// channel_game := make(chan []string)

	for {
		n, err := connect.Read(recvBuf)
		if err != nil {
			if io.EOF == err {
				fmt.Println("connection is closed from client; %v", connect.RemoteAddr().String())
				return
			}
			fmt.Println(err)
			return
		}
		if 0 < n {
			data := recvBuf[:n]
			fmt.Println("-------------------")
			fmt.Println(string(data))
			fmt.Println("-------------------")

			channel_game = append(channel_game, string(data))
			channel_game = append(channel_game, string(data))

			fmt.Println("-------------------")
			fmt.Println(channel_game)
			fmt.Println("-------------------")

			fmt.Println(len(channel_game))

			if len(channel_game) > 0 { //TEMP
				play_game(len(channel_game))

				connect.Write([]byte("===================> OK"))
			}
		}
	}
}

func play_game(player_num int) {

	var cardDeck int
	var playerNumber int
	fmt.Println("######################")
	fmt.Println(player_num)
	fmt.Println("######################")
	// os.Exit(1)

	cardDeck = player_num
	playerNumber = player_num

	winRate := make([]int, playerNumber+1)

	for index := range winRate {
		winRate[index] = 0
	}

	for i := 0; i < 10; i++ {

		card := makeCard(cardDeck)
		fmt.Println(card)

		myCard := shuffleCard(card)
		fmt.Println(myCard)

		playerResult := make([][]int, playerNumber)
		for j := 0; j < playerNumber; j++ {
			playerResult[j] = make([]int, 2)
			playerResult[j][0] = myCard[j*2]
			playerResult[j][1] = myCard[j*2+1]
		}
		fmt.Println(playerResult)

		winner := compatition(playerResult)
		fmt.Println(winner)
		winRate[winner]++
	}

	fmt.Println("winRate check : ", winRate)
	for index, temp := range winRate {
		if index == 0 {
			fmt.Printf("Draw Rate : %0.2f %%\n", float32(temp))
		}
		fmt.Printf("Player %d Win Rate : %0.2f %%\n", index, float32(winRate[index]))
	}

}

func checkPair(playerResult [][]int) (int, int, bool) {
	myCheckMaxIndex := 0
	myCheckMax := 0
	drawFlag := false
	for index, buff := range playerResult {
		if buff[0] == buff[1] {
			if buff[0] > myCheckMax {
				myCheckMaxIndex = index
				myCheckMax = buff[0]
			} else if buff[0] == myCheckMax {
				drawFlag = true
			} else {
			}
		}
	}

	return myCheckMaxIndex, myCheckMax, drawFlag
}

func compareCard(playerResult [][]int) (int, int, bool) {
	myCheckMaxIndex := 0
	myCheckMax := -1
	drawFlag := false
	for index, buff := range playerResult {
		playerCard := (buff[0] + buff[1]) % 10
		if playerCard > myCheckMax {
			myCheckMaxIndex = index
			myCheckMax = playerCard
			drawFlag = false
		} else if playerCard == myCheckMax {
			drawFlag = true
		} else {
		}
	}
	return myCheckMaxIndex, myCheckMax, drawFlag
}

func compatition(playerResult [][]int) int {
	winner := 0
	pairindex, pairmax, drawflag := checkPair(playerResult)
	fmt.Println("checkPair", pairindex, pairmax, drawflag)
	if drawflag == true {
		winner = 0
	} else {
		if pairmax != 0 {
			winner = pairindex + 1
			fmt.Println("Winner is Pair : ", pairindex, " card num is : ", pairmax)
		} else {
			compareindex, comparemax, comparedrawflag := compareCard(playerResult)
			fmt.Println("checkCompare", compareindex, comparemax, comparedrawflag)
			if comparedrawflag == true {
				fmt.Println("compareDraw")
				winner = 0
			} else {
				winner = compareindex + 1
				fmt.Println("Winner is ", pairindex, " card num is ", comparemax)
			}
		}
	}
	return winner
}

func shuffleCard(card []int) []int {
	myCard := make([]int, len(card))

	var index = 0
	for {
		flag := checkCard(card)
		if flag == 0 || index == len(card) {
			break
		}

		s1 := rand.NewSource(time.Now().UnixNano())
		rand := rand.New(s1)
		randomNumber := rand.Intn(len(card))

		if card[randomNumber] != 0 {
			myCard[index] = card[randomNumber]
			card[randomNumber] = 0
			index++
		} else {
			continue
		}
		// fmt.Println(randomNumber)
	}

	return myCard
}

func checkCard(card []int) int {
	flag := 0
	for _, temp := range card {
		flag = flag + temp
	}
	return flag
}

func makeCard(cardDeck int) []int {
	result := make([]int, cardDeck*10)
	for i := 0; i < cardDeck*10; i++ {
		result[i] = (i + 1) % 10
		if result[i] == 0 {
			result[i] = 10
		}
	}
	return result
}
