package main

import "testing"

func Test_InputNum(t *testing.T) {
	expected := 8
	//構造体
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	//入れたいマス (2,2)
	//入れたい数値 8
	b.put(2, 2, 8)
	if b.get(2, 2) != expected {
		t.Errorf("Test_InputNum Error")
	}

}
