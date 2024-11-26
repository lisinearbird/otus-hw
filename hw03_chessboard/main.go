package main

import (
	"fmt"
)

func stringForChessboard(size int) (chessString string) {
	next := "#"
	prev := " "
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			chessString += next
			next, prev = prev, next // Чередую клетки
		}
		if chessString[len(chessString)-size] == next[0] { // Начало новой линии должно отличаться от начала предыдущей
			next, prev = prev, next
		}
		chessString += "\n"
	}
	return chessString
}

func main() {
	size := 8
	fmt.Printf("Enter size: ")
	fmt.Scanf("%d", &size)

	fmt.Println(stringForChessboard(size))
}
