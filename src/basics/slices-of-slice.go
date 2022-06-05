/**
 * @Author: Henry
 * @Description:
 * @File:  slices-of-slice.go
 * @Version: 1.0.0
 * @Date: 2020/3/5 8:57 PM
 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	// creat a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"^", "_", "^"},
	}
	// The player takes turns.
	board[0][0] = "x"
	board[2][1] = "o"
	board[1][2] = "o"
	board[1][0] = "o"
	board[0][2] = "x"

	for i:=0; i < len(board); i++{
		fmt.Printf("%s\n", strings.Join(board[i]," "))
	}
}