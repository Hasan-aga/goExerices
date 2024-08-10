package main

import (
	"fmt"
	"strings"
  "golang.org/x/tour/pic"
)

func main() {
	var a [5]string
	a[0] = "1"
	a[1] = "2"
	a[2] = "hello3"
	a[3] = "4"
	a[4] = "5"
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

//	slices()
//	sliceOfSlice()
//  fmt.Println(pic(100, 100))
    pic.Show(Pic)
}


func Pic(dx, dy int) [][]uint8{
  var slice [][]uint8
  for x := 0; x < dx; x++{
    var row []uint8
    for y := 0; y < dy; y++{
      row = append(row, uint8(x^y)) 
    }
    slice = append(slice, row)
  }
  return slice
}

func slices() {
	var a []string
	a = append(a, "111111111111")
	a = append(a, "2")
	a = append(a, "hello3")
	a = append(a, "4")
	a = append(a, "5")
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	b := []string{"11111", "2", "hello4", "4", "5"}
	fmt.Println(b)
}

func sliceOfSlice() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
