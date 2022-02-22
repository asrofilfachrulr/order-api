package model

import (
	"fmt"
	"testing"
)

type test struct {
	N int
}

func Add() int {
	fmt.Println("Additon operated")
	return 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1
}

func something(t test) {

}

func TestSomething(t *testing.T) {
	ts := &test{
		N: Add(),
	}
	something(*ts)
	// fmt.Println(ts.N)
	// fmt.Println(ts.N)
	// fmt.Println(ts.N)
	// fmt.Println(ts.N)
	// fmt.Println(ts.N)

}
