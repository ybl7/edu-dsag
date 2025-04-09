package main

import "fmt"

func main() {
	fmt.Println(QuickSelect([]int{2, 7, 3, 4, 11, 8, 1, 9, 6, 10, 5}, 6))
	fmt.Println(QuickSelect([]int{2, 7, 3, 4, 11, 8, 1, 9, 6, 10, 5}, 9))
	fmt.Println(QuickSelect([]int{2, 7, 3, 4, 11, 8, 1, 9, 6, 10, 5}, 2))
}

func QuickSelect(arr []int, k int) (int, error) {
	if k > len(arr)-1 {
		return 0, fmt.Errorf("index k greater than length of array")
	}
	low, high := 0, len(arr)-1
	m := 0

	for low < high && m < 100 {
		i, j := low, low
		p := arr[high]

		for i <= high {
			if p > arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
				j++
			}
			if i == high {
				// Swap but don't increment, we have already reached high, j can increase no further
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
		if j < k {
			low = j
		}
		if j > k {
			high = j - 1 // the current index j is the first element of the right subarray, we don't actually want it so -1 to eliminate it.
		}
		if j == k {
			return arr[j], nil
		}
		m++
	}
	return 0, nil
}

// Take an array
// [b,a,e,g,c,f,h,j,i,d]

// We don't care about the values, let's just suppose that if the array were to be ordered by value it would be [a,b,c,d,e,f,g,h,i,j]
// One way to get the kth smallest element is to sort the array then just select element k.
// Another would be to use an appropriately sized heap (more on this later, but ignore it for now).

// But we can accomplish this with swaps.
// The main idea is to choose one pivot element, and ensure that it lands in the correct place in the array.
// Then based on whether the index of pivot element is smaller/greater than index, we discard the left/right subarray w.r.t to the pivot element.

// Back to our example, suppose we want to find element at k=3.
// Instantiate l=0 and m=0. Select the final element (d in our case) as the pivot and compare its value to the value at index l.
// If the element at l is smaller than the final element, push it to the front of the array, do this for the entire array.
// Then at the end of the iteration over the array, swap the final element with the first element of the right subarray.
// In practice the left and right subarrays are just tracked with indexes, with our index m tracking the start of the right subarray.

// l = 0
// m = 0
// d > b so swap a[0] and a[0]
// l++ m++

// l = 1
// m = 1
// d > a so swap a[1] and a[1]
// l++ m++

// l = 2
// m = 2
// d < e so no swap required
// l++

// l = 3
// m = 2
// d < g so no swap required
// l++

// l = 4
// m = 2
// d > c so swap a[4] and a[1]
// l++ m++
// (at this point our array looks like [b,a,c,g,e,f,h,j,i,d])

// l = 5
// m = 3
// d < f so no swap required
// l++

// l = 6
// m = 3
// d < h so no swap required
// l++

// l = 7
// m = 3
// d < j so no swap required
// l++

// l = 8
// m = 3
// d < i so no swap required
// l++

// l = 9
// m = 3
// we've reached the end of the array so now need to place d in it's rightful place, which is at index m, so swap a[len(a)-1] and a[m]
// (at this point our array looks like [b,a,c,d,e,f,h,j,i,g])

// And we have now ensured that d is at the correct index, now we need to check if d < k or d > k, if d < k that means k lies in the right subarray so we must iterate on that.
// If d > k then it means k lies in the left subarray so we must iterate on that. The exit condition is when d = k.
// We will use two pointers to track the start and end of the array we want to operate on, instead of using a recursive call.
