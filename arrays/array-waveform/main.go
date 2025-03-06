package main

func main() {}

func waveForm() {}

// I can think of two approaches to solve this problemm

// -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
// Approach 1
// Given some array [1,4,1,9,2,,-7,6,8,0,-2]
// We can order it first in O(nlog(n)) time using the quicksort algorithm - this isn't a bad first step since sorting is a "for free" primitive operaion as Tim Roughgarden puts it in his book
// So the sorted array becomes [-7,-2,0,1,1,2,4,6,8,9]

// Now let's imagine this array is in fact two arrays, a left and a right part:
// leftArr = [-7,-2,0,1,1] and rightArr = [2,4,6,8,9]

// Why would we do this? Because all elements of leftArr < all elements in righArr
// This now allows us to swap elements between the two arrays and guarantee that we will always have a waveform, for instance

// leftArr[0] <> rightArr[0]
// leftArr =  [2,-2,0,1,1]
// rightArr = [-7,4,6,8,9]
// totalArr = [2,-2,0,1,1,-7,4,6,8,9]

// leftArr[2] <> rightArr[2]
// leftArr =  [2,-2,6,1,1]
// rightArr = [-7,4,0,8,9]
// totalArr = [2,-2,6,1,1,-7,4,0,8,9]

// leftArr[4] <> rightArr[4]
// leftArr =  [2,-2,6,1,9]
// rightArr = [-7,4,0,8,1]
// totalArr = [2,-2,6,1,9,-7,4,0,8,1]

// And we have our solution, in practice we won't create a leftArr or a rightArr, we'll just track the elements we want to swap

// -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
// Approach 2
// This follows a proof by induction type of approach. Suppose we have the simplest base case
// arr = [1,2,3]

// We then create a subroutine that looks at a window of 3 elements and organises them into a wave form
// waveform([1,2,3]) = [2,1,3]

// Now let's add an element to this array, call it a, so we have [2,1,3,a], let's look at what happens in a few cases
// if a<3 we don't need to do anything, our new array is still a waveform
// if a>3 we swap a and 3 to get [2,1,a,3] and we still have a waveform since a>3>1

// Now let's add another element b, so we have [2,1,3,a,b], let's examine what happens under both of the cases mentioned already
// [2,1,3,a]
// if b<=a then we need to swap a and b to maintain our waveform so we get [2,1,3,b,a]
// if b>a then we the waveform is maintained so we get [2,1,3,a,b]

// [2,1,a,3]
// if b<=3 then we need to swap a and b to maintain our waveform so we get [2,1,a,b,3]
// if b>3 then we the waveform is maintained so we get [2,1,a,3,b]

// So we see that via this method, we just need to keep track of the current index that we are "adding" an element at, and compare it with the previous element

// -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
// Of the two approaches, approach 2 is better since it requires only a traversal of the array so it is O(n), whereas the best sorting algorithms run in O(nlogn), so we'll go with approach 2
