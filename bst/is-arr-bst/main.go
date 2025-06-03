package bst

func IsArrBST(arr []int) bool {
	// Base case where we have no subtree or a single element subtree
	if len(arr) == 0 || len(arr) == 1 {
		return true
	}

	r := arr[0]
	var j int
	var ge bool

	for i := 0; i < len(arr); {
		// So long as we haven't hit an element that is greater than e, we are still in the left subtree
		if arr[i] <= r && !ge {
			i++
		} else if arr[i] > r && !ge { // The first time we hit an element that is greater than e, we set a boolean variable to indicate that we may have entered the right subtree
			ge = true
			j = i
			i++
		} else if ge && arr[i] <= r { // If we are in the right subtree, but we regress and find a value smaller than e, then this can't be a binary tree, so we immediately return false
			return false
		} else if ge && arr[i] > r {
			i++
		}
	}

	// If we get this far, then it means at least for element r, the binary tree structure is respected, but now we have to look at the left and right subtrees
	left := IsArrBST(arr[1:j])
	right := IsArrBST(arr[j:])

	return left && right
}

// Given an array of integers, find if it represents a pre-order traversal of a BST.
// The key here is the pre-order traversal, this means that element 0 of the arr will always be the root node
// Then the whole of the left subtree will be traversed, then the whole of the right subtree
// How does this look in an array? Something like this:
// [root l1 l2 ... ln r1 r2 ... rn], therefore a sufficient test is to iterate down the array and find if given a root element e:
// if at there is a continuous subset of elements all smaller than e, followed by a subset all greater than e
// We can do this recursively, where we shrink the array with each iteration
