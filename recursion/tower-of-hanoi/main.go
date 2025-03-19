package main

import "fmt"

func main() {
	TowerOfHanoi(5, "A", "B", "C")
}

func TowerOfHanoi(n int, src, dst, tmp string) {
	if n == 0 {
		return
	}
	TowerOfHanoi(n-1, src, tmp, dst)
	fmt.Printf("Move disk %d from peg %s to peg %s \n", n, src, dst)
	TowerOfHanoi(n-1, tmp, dst, src)
}

// Moving n rings from src to dst can be accomplished by first moving n-1 rings to tmp
// then moving the bottom ring of the n original rings to dst
// then moving the stack of rings at tmp to dst, we can utilise the previous src as our tmp since it will be empty at this point
// Time complexity is 2^n as detailed in https://proofwiki.org/wiki/Tower_of_Hanoi#Solution
