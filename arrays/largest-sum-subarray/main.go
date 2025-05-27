package main

import "fmt"

func main() {
	arr1 := []int{0, 1, 3, -2, 4, -5, -6, 2} // LLS = 6
	fmt.Printf("LSS %v: %v\n", arr1, LSS(arr1))

	arr2 := []int{10, 1, 3, -2, 4, -5, -6, 2} // LLS = 16
	fmt.Printf("LSS %v: %v\n", arr2, LSS(arr2))
}

func LSS(arr []int) int {
	var best int
	var val int

	for i, _ := range arr {
		if val < 0 {
			val = arr[i]
		} else {
			val += arr[i]
		}

		if val > best {
			best = val
		}
	}

	return best
}

// Explanation
// Consider an array R to which an element is being appended
// [A] -> [A, B] -> [A, B, C]

// We want to find continguous subsets of such an array that produce the greates sum,
// so for the above the array at each step these are:

// [A]: {A}
// [A, B]: {A}, {B}, {A, B}
// [A, B, C]: {A}, {B}, {C}, {A, B}, {B, C}, {A, B, C} But NOT {A, C}

// The calculations of sums overlap a lot, for example sum({A, B, C}) = sum({A, B}) + sum({C}) - how can we use this?
// Let's call the set containing the sums of possilbe contiguous subsets of R that end at and cinlucde R[i] S_i where S1, S2, S3 correspond to i = 1, 2, 3
// For example S1 = {{A}}
// 			   S2 = {{B}, {A+B}}
//             S3 = {{C}, {B+C}, {A+B+C}}

// S_i = {x + R[i] for x in S_i-1} U {R[i]}
// Which reads, the sums of all possible subests S_i is just the previous sums, elements of S_i-1, with the addition of the new element R[i],
// the union is required to include the sum of the new set containing just R[i] by itself

// LSS_i = max(S_i)
// The largest sum subarray of the array R is just the maximum value of the sums in S_i

// max(S_i) = max({x + R[i] for x in S_i-1} U {R[i]})  // The sum of {R[i]} is just R and max(A U B) = max(A, B)
// 		 = max({x + R[i] for x in S_i-1}, R[i])     // Since R[i] is just a constant that is added to all x
// 		 = max({x for x in S_i-1} + R[i], R[i])
// 		 = max(max(S_i-1) + R[i], R[i])
// 		 = max(LSS_i-1 + R[i], R[i])
// LSS_i    = max(LSS_i-1 + R[i], R[i])

// In English, what this says is that the new LSS after adding R[i] will always be the existing LSS in the case where R[i] decreases it,
// or it will be exising LSS plus R[i] in the case that R[i] increases it.

// So if we start in the base case with the single elemet array, and the LSS is just the first element, we can use this to calculate the lSS
// of the two element array after we add an element. In fact, every time we are adding an element, we calcuate the value of the LSS of the entire array
// up to that point. We can keep doing this until we reach the end of the input array. We find the best value for the current array, then we ask,
// can I improve it if I include more elements in the array - and we make use of the LSS_i to LSS_i-1 relationship derived above.

// For example for the array [0, 1, 3, -2, 4, -5, -6, 2]

// LSS_0 = 0
// LSS_1 = max(0+1, 1)   = 1
// LSS_2 = max(1+3, 3)   = 4
// LSS_3 = max(4-2, -2)  = 2
// LSS_4 = max(2+4, 4)   = 6
// LSS_5 = max(6-5, -5)  = 1
// LSS_6 = max(1+2, 2)   = 3

// So we see the best value was LSS_4 and we reutnr this.
// If we take a closer look at max(LSS_i-1 + R[i], R[i]), we can rewrite this in a conditional form

// If LSS_i-1 < 0 then R[i] will always be larger than it, so LSS_i = R[i] in these cases (this is what you often see on YouTube followed by some handwavy explanation)
// So we can rewrite max(LSS_i-1 + R[i], R[i]) into a conditional that checks if LSS_i-1 is less than 0, and if it is, we just set LSS_i to R[i]
