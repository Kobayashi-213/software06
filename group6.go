package main

import (
	"fmt"
	"strconv"
)

type Board struct {
	tokens []int // tokens[0] -> (0,0), tokens[1] -> (0,1), ...
}

func (b *Board) get(x, y int) int {
	return b.tokens[x+9*y]
}

func (b *Board) hyouji() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			s := ""
			k = b.get(i, j)
			s += strconv.Itoa(i) + " "
		}
		fmt.Println(s)
	}
}

func (b *Board) judgeTate() bool {
	n := 0
	k := true
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			n = b.get(j, i)
			if n != 0 {
				break
			}
		}
		for j := 0; j < 9; j++ {
			if b.get(j, i) == 0 {
				break
			} else if n == b.get(j, i) {
				k = false
			}
		}

	}
	return k
}

func (b *Board) put(x, y, value int) {

	//すでに値が入っている場合 (初期状態は0)
	if b.tokens[x+9*y] != 0 {
		fmt.Println("Input Error !:すでに値が存在します")
		return
	}

	b.tokens[x+9*y] = value
	return
}

func (b *Board) input() {
	//Player1 がxとする
	var x, y, value int

	//座標を入力する
	fmt.Printf("Input (x,y):")
	//test用
	x = 1
	y = 2
	//標準入力から受け取る
	fmt.Scan(&x)
	fmt.Scan(&y)

	//0~8に変換
	x--
	y--

	//xとyの値が範囲外になっていないか判定
	if x < 0 || x >= 9 || y < 0 || y >= 9 {
		fmt.Println("Input Error :invalied value !")
	}

	//値を入力する
	fmt.Printf("value : ")
	fmt.Scan(&value)

	//入力した値を反映する
	b.put(x, y, value)
	return
}

func main() {
	b := &Board{
		tokens: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.input()
}
