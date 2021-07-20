package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var board []string
var mapWinGroup = map[int][]int{1: {1, 2, 3}, 2: {4, 5, 6}, 3: {7, 8, 9}, 4: {1, 4, 7}, 5: {2, 5, 8}, 6: {3, 6, 9}, 7: {1, 5, 9}, 8: {3, 5, 7}}
var mapInGroup = map[int][]int{1: {1, 4, 7}, 2: {1, 5}, 3: {1, 6, 8}, 4: {2, 4}, 5: {2, 5, 7, 8}, 6: {2, 6}, 7: {3, 4, 8}, 8: {3, 5}, 9: {3, 6, 7}}

func main() {
	pa := 0
	bo := reset(board)
	for {
		if pa == 1 {
			bo = reset(bo)

			pa = 0
		} else if pa == -1 {
			break
		}
		printBoard(bo)
		bo = playerTurn(bo)
		fmt.Println("Player has placed \"x\"..")
		if isWinning(bo, "x") {
			fmt.Println("Well Done!! You are the winner!!")
			pa = playAgain()
			continue
		}
		index, bo := comTurn(bo)
		bo[index] = "o"
		// printBoard(bo)
		if isWinning(bo, "o") {
			printBoard(bo)
			fmt.Println("Player \"o\" is the winner!!")
			pa = playAgain()
			continue
		} else if isFullBoard(bo) {
			pa = playAgain()
			continue
		} else {
			fmt.Println("Com has placed \"o\"..")
		}
	}
}

func playAgain() int {
	for {
		fmt.Println()
		fmt.Println("Do you want replay? (Y/N):")
		input := userInput()
		if strings.ToUpper(input) == "Y" {
			return 1
		} else if strings.ToUpper(input) == "N" {
			return -1
		} else {
			fmt.Println("Invalid input")
			continue
		}
	}
}

func reset(bo []string) []string {
	return []string{"", " ", " ", " ", " ", " ", " ", " ", " ", " "}
}

func randNumber(no int) int {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(no)
	return i
}

func printBoard(bo []string) {
	fmt.Println("    |     |")
	fmt.Println("", bo[1], " | ", bo[2], " | ", bo[3])
	fmt.Println("    |     |")
	fmt.Println("---------------")
	fmt.Println("    |     |")
	fmt.Println("", bo[4], " | ", bo[5], " | ", bo[6])
	fmt.Println("    |     |")
	fmt.Println("---------------")
	fmt.Println("    |     |")
	fmt.Println("", bo[7], " | ", bo[8], " | ", bo[9])
	fmt.Println("    |     |")
}

func playerTurn(bo []string) []string {
	var in int
	for {
		fmt.Println("Please place \"x\" in the board, (1-9):")
		input := userInput()
		pos := checkInput(bo, input)
		fmt.Println(pos)
		if pos == 0 {
			printBoard(bo)
			continue
		} else {
			in = pos
			break
		}
	}
	bo[in] = "x"
	return bo

}

func checkInput(bo []string, s string) int {
	var pos int
	in, err := strconv.Atoi(s)
	pos = in

	if err != nil {
		fmt.Println("Invalid input")
		pos = 0
	} else if in < 1 || in > 9 {
		fmt.Println("Invalid slot number")
		pos = 0
	} else {
		if !checkEmpty(bo, pos) {
			fmt.Println("The slot is unavailable!")
			pos = 0
		}
	}

	fmt.Println()
	return pos
}

func checkEmpty(bo []string, pos int) bool {
	if bo[pos] == " " {
		return true
	}
	return false
}

func isFullBoard(bo []string) bool {
	var count int
	for _, v := range bo {
		if v == " " {
			count++
		}
	}
	if count <= 0 {
		fmt.Println("Tie Game!!!")
		return true
	}
	return false
}

func isWinning(bo []string, le string) bool {
	return (bo[1] == le && bo[2] == le && bo[3] == le) || (bo[4] == le && bo[5] == le && bo[6] == le) || (bo[7] == le && bo[8] == le && bo[9] == le) || (bo[1] == le && bo[4] == le && bo[7] == le) || (bo[2] == le && bo[5] == le && bo[8] == le) || (bo[3] == le && bo[6] == le && bo[9] == le) || (bo[1] == le && bo[5] == le && bo[9] == le) || (bo[3] == le && bo[5] == le && bo[7] == le)
}

func possibleMoves(bo []string) []int {
	var pm []int
	for i, v := range bo {
		if v == " " {
			pm = append(pm, i)
		}
	}
	return pm
}

func comTurn(bo []string) (int, []string) {
	pm := possibleMoves(bo)
	for _, v := range []string{"o", "x"} {
		for _, i := range pm {
			boardCopy := append([]string{}, bo...)
			boardCopy[i] = v
			if isWinning(boardCopy, v) {
				move := i
				return move, bo
			}
		}
	}

	for _, v := range []string{"o", "x"} {
		if v == "o" {

			for _, z := range []int{2, 1} {

				i := aiMove(bo, v, pm, z)
				if i != nil {
					if z == 1 {
						afteri := map[int]bool{}
						for _, j := range i {
							boardCopy := append([]string{}, bo...)
							boardCopy[j] = "o"
							k := aiMove(boardCopy, "x", possibleMoves(boardCopy), 2)
							if len(k) > 0 && k != nil {
								afteri[j] = true
							}
						}
						l := []int{}
						for _, w := range i {
							if _, ok := afteri[w]; !ok {
								l = append(l, w)
							}
						}
						i = l[:]

					}
					if len(i) > 0 {
						move := i[randNumber(len(i))]
						return move, bo
					}
				}
			}
		} else {
			i := aiMove(bo, v, pm, 2)
			if i != nil {
				afteri := map[int]bool{}
				for _, j := range i {
					boardCopy := append([]string{}, bo...)
					boardCopy[j] = "o"
					k := aiMove(boardCopy, "x", possibleMoves(boardCopy), 2)
					if len(k) > 0 && k != nil {
						afteri[j] = true
					}
				}
				l := []int{}
				for _, w := range i {
					if _, ok := afteri[w]; !ok {
						l = append(l, w)
					}
				}
				i = l[:]
				if len(i) > 0 {
					move := i[randNumber(len(i))]
					return move, bo
				}
			}
		}
	}

	if len(pm) > 0 {
		move := pm[randNumber(len(pm))]
		return move, bo
	}

	return 0, bo
}

func aiMove(bo []string, let string, po []int, noInterset int) []int {
	var sliAiStep []int
	for _, i := range po {
		boCopy := append([]string{}, bo...)
		boCopy[i] = let
		sliSet := 0
		sliIG := mapInGroup[i]
		for _, j := range sliIG {
			sliWG := mapWinGroup[j]

			if (boCopy[sliWG[0]] == let || boCopy[sliWG[0]] == " ") && (boCopy[sliWG[1]] == let || boCopy[sliWG[1]] == " ") && (boCopy[sliWG[2]] == let || boCopy[sliWG[2]] == " ") {
				letNo := 0
				for _, v := range []int{0, 1, 2} {
					if boCopy[sliWG[v]] == let {
						letNo++
					}
				}
				if letNo >= 2 {
					sliSet++
				}
			}
		}
		if sliSet >= noInterset {
			sliAiStep = append(sliAiStep, i)
		}
	}
	if len(sliAiStep) > 0 {
		return sliAiStep
	}
	return nil
}

func userInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
