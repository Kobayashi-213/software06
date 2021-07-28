package main

import (
	"fmt"
	//"math"
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
		s := ""
		for j := 0; j < 9; j++ {

			k := b.get(i, j)
			s += strconv.Itoa(k) + " "
		}
		fmt.Println(s)
	}
}

func (b *Board) JudgeTate() bool {
	n := 0
	m := 0
	k := true
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			n = b.get(j, i)
			if n != 0 {
				m = j
				break
			}
		}
		for j := 0; j < 9; j++ {
			if j == m {
				continue
			}
			if b.get(j, i) == 0 {
				break
			} else if n == b.get(j, i) {
				k = false
			}
		}

	}
	return k
}

func (b *Board) put(x, y, value int) bool {

	//すでに値が入っている場合 (初期状態は0)
	if b.tokens[x+9*y] != 0 {
		fmt.Println("Input Error !:すでに値が存在します")
		return false
	}

	b.tokens[x+9*y] = value

	if b.JudgeYoko() && b.JudgeTate() && b.Judge3x3() {
		return true
	} else {
		fmt.Println("Input Error !:　数字がかぶっています")
		b.tokens[x+9*y] = 0
		return false
	}

}

func (b *Board) JudgeYoko() bool {

	for i := 0; i < 9; i++ {

		alreadyExist := make([]bool, 9)
		for p := 0; p < 9; p++ {
			alreadyExist[p] = false
		}

		for k := 0; k < 9; k++ {
			//すでに存在する時
			if b.tokens[k+9*i] == 0 {
				continue
			} else if alreadyExist[b.tokens[k+9*i]-1] {
				return false
			} else {
				alreadyExist[b.tokens[k+9*i]-1] = true
			}
		}

	}
	return true
}

func (b *Board) Judge3x3() bool {

	//どのブロックを判定するか
	for i := 0; i < 9; i++ {
		alreadyExist := make([]bool, 9)

		//各ブロックの左上
		start := (i%3)*3 + int(i/3)*27

		//各ブロックについて判定
		for k := 0; k < 3; k++ {
			for l := 0; l < 3; l++ {
				//すでに存在する時
				if b.tokens[start+k+9*l] == 0 {
					continue
				} else if alreadyExist[b.tokens[start+k+9*l]-1] {
					//被ってる時
					return false
				} else {
					alreadyExist[b.tokens[start+k+9*l]-1] = true
				}
			}
		}
	}
	return true
}
func (b *Board) JudgeEmpty() bool {

	for i := 0; i < 9; i++ {
		for k := 0; k < 9; k++ {
			if b.get(k, i) == 0 {
				return false
			}
		}
	}
	return true
}

func (b *Board) input() bool {
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
	return b.put(x, y, value)

}

func main() {
	b := &Board{
		tokens: []int{1, 9, 5, 2, 6, 8, 4, 7, 3,
			2, 4, 8, 1, 3, 7, 5, 9, 6,
			6, 3, 7, 9, 4, 5, 2, 8, 1,
			9, 5, 6, 3, 2, 4, 7, 1, 8,
			8, 1, 2, 6, 7, 9, 3, 5, 4,
			3, 7, 4, 5, 8, 1, 6, 2, 9,
			5, 8, 3, 4, 9, 2, 1, 6, 7,
			7, 6, 1, 8, 5, 3, 9, 4, 0,
			4, 2, 9, 7, 1, 6, 8, 3, 0},
	}
	b.hyouji()

	for {
		if !b.input() {
			continue
		}

		b.hyouji()

		if b.JudgeEmpty() == true {
			fmt.Println("clear!!")
			break
		} else if b.JudgeEmpty() == false {
			fmt.Println("next...")
		}

	}

}
