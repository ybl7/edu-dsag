package main

import "fmt"

func main() {
	arr1 := []int{33, 9, 10, 3, 2, 60, 30, 32, 1}
	fmt.Println(arr1)
}

func IndexMaxDiff(arr []int) int {

	return 0
}

// We can think of this as a problem containing two sub problems.
// 1) Identify all pairs i, j that satify the conditions i < j AND arr[i] < arr[j]
// 2) Calculate the max(j - i) over these pairs
// We start at the value in the row and end at the colum, so the first row indicates the interval (i=0, j=5)

// 	   33  9   10  3   2   60  30  32  1
// 33  -   -   -   -   -   Y   -   -   -
// 9   -   -   Y   -   -   Y   Y   Y   -
// 10  -   -   -   -   -   Y   Y   Y   -
// 3   -   -   -   -   -   Y   Y   Y   -
// 2   -   -   -   -   -   Y   Y   Y   -
// 60  -   -   -   -   -   -   -   -   -
// 30  -   -   -   -   -   -   -   Y   -
// 32  -   -   -   -   -   -   -   -   -
// 1   -   -   -   -   -   -   -   -   -
// This will always be O(n^2)

// If we examine this table, one thing should become apparant when looking at the case for 9 and 10 when they are the lower bound of the interval.
// Both are good candidates, but, 9 is always the better candidate by virtue of being smaller than 10 and coming before.
// So in general for the start of an interval if arr[i] < arr[j] AND i < j, then the optimal pair will always include arr[i] over arr[j]

// A similar thing applies if we look at the case of 30 and 32 when they are the upper bound of the interval.
// Since 32 > 30 and comes later in the array, any interval that ends at 30 will be improved upon by 32.
// So in general for the end of an interval if arr[i] < arr[j] AND i < j, then the optimal pair will always include arr[j] over arr[i]

// So we cam apply these rules to narrow down the possible intervals we want to observe:
// [33, 9, -, 3, 2, -, -, -, 1]     <- we only need to consider these for the lower bound of any possible interval
// [ -, -, -, -, -, 60, -, 32, 1]   <- we only need to consider these for the upper bound of any possilbe interval

// So now, let's redraw the table of possible intervals

//     60  32  1
// 33  Y   -   -
// 9   Y   Y   -
// 3   Y   Y   -
// 2   Y   Y   -
// 1   -   -   -

// Of all of these candidates, thhe more to the right, the larger the value of j, and the higher up, the lower the value of i
// So our optimal solution minimises i and maximises j, this becomes a little clearer if we redraw the whole table (but keeping our improved intervals)

//          j=0 j=1 j=2 j=3 j=4 j=5 j=6 j=7 j=8
//          33  9   10  3   2   60  30  32  1
// i=0  33  -   -   -   -   -   5   -   -   -
// i=1  9   -   -   -   -   -   4   -   6   -
// i=2  10  -   -   -   -   -   -   -   -   -
// i=3  3   -   -   -   -   -   2   -   4   -
// i=4  2   -   -   -   -   -   1   -   3   -
// i=5  60  -   -   -   -   -   -   -   -   -
// i=6  30  -   -   -   -   -   -   -   -   -
// i=7  32  -   -   -   -   -   -   -   -   -
// i=8  1   -   -   -   -   -   -   -   -   -

// Now from this image we can see that values on the same column, the top is the best since j is constant and i increases as you go down
// and for values along the same row, the rightmost is the best since we keep i contant but j increases as you move to the right
// So this problem reduces to looking at just the top value of all of the rows/columns
// You don't actually need to calculate the difference for all the intervals as shown in the above diagram.
// You just need to pick rows/columns then calculate the value of the rightmost/top entry.
// There are O(n) rows/columns, so we turn a matrix calculation that was previously O(n^2), to O(n) using this fact.const

// Now how do you codify all of this in code?
