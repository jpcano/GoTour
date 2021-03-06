package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkIn(t, ch)
	close(ch)
}

func WalkIn(t *tree.Tree, ch chan int) {
	if t != nil {
		WalkIn(t.Left, ch)
		ch <- t.Value
		WalkIn(t.Right, ch)
	}
}
// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i:= range(ch1) {
		if i != <- ch2 {
			return false
		}
	}
	return true
}	

func main() {
	ch := make(chan int)
	go Walk(tree.New(2), ch)
	for n := range ch {
		fmt.Printf("%v ", n)
	}
	fmt.Println()
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
