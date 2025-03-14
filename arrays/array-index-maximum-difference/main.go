package main

import "fmt"

func main() {
	arr1 := []int{33, 9, 10, 3, 2, 60, 30, 32, 1}
	arr2 := []int{8, 5, 6, 9, 3, 2, 7, 4, 10}
	arr3 := []int{1, 3, 4, 2, 6}
	arr4 := []int{3, 7, 5, 6, 1, 4, 2}
	fmt.Println(IndexMaxDiff(arr1))
	fmt.Println(IndexMaxDiff(arr2))
	fmt.Println(IndexMaxDiff(arr3))
	fmt.Println(IndexMaxDiff(arr4))
}

func IndexMaxDiff(arr []int) int {
	// We want to store the index and value, we map the value to the index i.e. {33: 0} or {9: 1}
	// In the explanation we could determine the index because I wrote out the whole array, but in practice go slices have to be of a certain type
	// And since the inputs are integers, we can't just set all the indexes we aren't interesed in to -1 or 0 to fill in the indexes that we don't care about, we might actually have inputs of -1 or 0 in the original array
	// We don't want to lose the crucial information about the index of the elements in the original array when determining which pairs of elements we want to calculate the difference across
	var lower []map[string]int
	var upper []map[string]int

	for i := 0; i < len(arr); {
		if i == 0 || arr[i] < lower[len(lower)-1]["value"] {
			lower = append(lower, map[string]int{
				"value": arr[i],
				"index": i,
			})
		}
		i++
	}
	// Remember that upper is reversed
	for j := len(arr) - 1; j >= 0; {
		if j == len(arr)-1 || arr[j] > upper[len(upper)-1]["value"] {
			upper = append(upper, map[string]int{
				"value": arr[j],
				"index": j,
			})
		}
		j--
	}

	var best int
	var curr int
	for j := len(upper) - 1; j >= 0; j-- {
		for i, _ := range lower {
			if upper[j]["value"] > lower[i]["value"] {
				curr = upper[j]["index"] - lower[i]["index"]
				if curr > best {
					best = curr
				}
				// Escape this inner for loop since we already found the maximum for the endpoint
				// The very first value of i for the current value of j to satisfy arr[i] < arr[j] will also be the best (since I will onbly get bigger with subsequent iterations), so there is no need to keep iterating over i
				break
			}
		}
	}
	return best
}

// We can think of this as a problem containing two sub problems.
// 1) Identify all pairs i, j that satify the conditions i < j AND arr[i] < arr[j]
// 2) Calculate the max(j - i) over these pairs
// The x, y axis correspond to j, i respectively (i=0, j=5)

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

// Now from this image we can see that for values on the same column, the top is the best since j is constant and i increases as you go down.
// And for values along the same row, the rightmost is the best since we keep i contant but j increases as you move to the right

// You don't actually need to calculate the difference j - i for all the intervals as shown in the above diagram.
// You just need to pick either rows/columns then calculate the value of the rightmost/top entry respectively.
// There are O(n) rows/columns, so we turn a matrix calculation that was previously O(n^2), to O(n) using this fact.

// Now how do you translate all of this in code?
// First let's create two slices, one for the lower interval and one for the upper
// Then we can iterate down just ONE of either the upper or lower array
// Suppose we choose to iterate over the upper array, so we walk down j
// For each value of j, we walk up i until we find the first element where arr[i] < arr[j] and calculate j - i and check if this is the best value yet
// If so we store it, if not, we move to the next element and continue

// This is a better solution that I found after the fact

// func ArrayIndexMaxDiff2(arr []int, size int) int {
//     leftMin := make([]int, size)
//     rightMax := make([]int, size)
//     leftMin[0] = arr[0]
//     i, j := 0, 0
//     var maxDiff = 0
//     for i = 1; i < size; i++ {
//         if leftMin[i-1] < arr[i] {
//             leftMin[i] = leftMin[i-1]
//         } else {
//             leftMin[i] = arr[i]
//         }
//     }
//     rightMax[size-1] = arr[size-1]
//     for i = size - 2; i >= 0; i-- {
//         if rightMax[i+1] > arr[i] {
//             rightMax[i] = rightMax[i+1]
//         } else {
//             rightMax[i] = arr[i]
//           }
//       }
//     fmt.Println("leftMin:", leftMin)
//     fmt.Println("rightMax: ", rightMax)
//     i=0
//     j=0
//     maxDiff = -1
//     for j < size && i < size {
//         if leftMin[i] < rightMax[j] {
//             if maxDiff < j-i {
//                 maxDiff = j - i
//             }
//             j=j+1
//           } else {
//             i=i+1
//           }
//       }
//     return maxDiff
// }

// Instead of outputing only the endpoints that are optimal and using a map to store their index
// It just replaces suboptimal indexes with the optimal one like so:
// leftMin: [33 9 9 3 2 2 2 2 1]
// rightMax:  [60 60 60 60 60 60 33 33 1]
// Then walks down both arrays
