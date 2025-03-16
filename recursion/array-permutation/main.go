// Given some array [a,b,c,d], how can we find all possible permutations?
// For any given index, every element has to occupy it at some point. For example for the 0th position.
// By swapping the 0th position with all others (including itself), we can ensure that we have cycled all elements through it

// [a,b,c,d]
// [b,a,c,d]
// [c,b,a,d]
// [d,b,c,a]

// Now for each of these arrays, we call our procedure again on the remaining elements, and we keep doing this

// 	[a,b,c,d]
// 		[a,b,c,d]
// 			[a,b,c,d]
// 			[a,b,d,c]
// 		[a,c,b,d]
// 			[a,c,b,d]
// 			[a,c,d,b]
// 		[a,d,c,b]
// 			[a,d,c,b]
// 			[a,d,b,x]

// And so on so forth...