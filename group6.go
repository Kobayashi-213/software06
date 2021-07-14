package main

import (
	"fmt"
	"os"
)

type Board struct {
	tokens []int // tokens[0] -> (0,0), tokens[1] -> (0,1), ...
}

func (b *Board) get(x, y int) int{
    return b.tokens[x + 9*y]    
}

func (b *Board) hyouji(){
    for i := 0; i < 9; i++ {
				for j := 0; j < 9; j++ {
					fmt.Printf(b.get(i, j))
				}
				fmt.Println("")
    }
}