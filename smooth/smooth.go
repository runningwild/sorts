// Implementation of Edsger Dijkstra's smooth sort
// Worst case time: O(n log n)
// Best case time: O(n)
// Called 'smooth' because it has a smooth transition between linear time and
// O(n log n) time as the input set transitions from sorted to unsorted.
// The space complexity can technically be constant, but only if a bit vector
// is used to store the sizes of the heaps.  In practice this is silly, so
// O(L^-1(n)) space is required, where L^-1(n) is the smallest Leonardo number
// greater than n.  This might as well be constant given how fast Leonardo
// numbers grow.
package smooth

import (
	"sort"
)

var leo []int

func init() {
	leo = []int{1, 1}
	length := 2
	for leo[length-1] < 1000000000 {
		leo = append(leo, leo[length-1]+leo[length-2]+1)
		length++
	}
}

// Stringify will reorder the root nodes to make sure that they are in
// increasing order.  This is called when a new heap is added at the end
// such that the only root node that is out of order is the new one.
func stringify(v sort.Interface, roots, sizes []int) int {
	k := len(roots) - 1
	for j := k - 1; j >= 0; j-- {
		jr := roots[j]
		kr := roots[k]
		if v.Less(kr, jr) {
			size := sizes[k]
			if size <= 1 {
				v.Swap(jr, kr)
				k = j
			} else {
				right := roots[k] - 1
				left := right - leo[sizes[k]-2]
				if size <= 1 || v.Less(right, jr) && v.Less(left, jr) {
					v.Swap(jr, kr)
					k = j
				}
			}
		} else {
			// Since the only node that is out of order is the one we start with,
			// once it is in order we can bail out.
			return k
		}
	}
	return k
}

// Heapify is called when two heaps are combined under a new root node.  Since
// the two sub-heaps are necessarily heaps it suffices to swap this node with
// its largest child repeatedly until it is larger than both of its children.
func heapify(v sort.Interface, root, size int) {
	for size > 1 {
		right := root - 1
		left := right - leo[size-2]
		if v.Less(left, right) {
			if v.Less(root, right) {
				v.Swap(root, right)
				root = right
				size -= 2
			} else {
				break
			}
		} else {
			if v.Less(root, left) {
				v.Swap(root, left)
				root = left
				size -= 1
			} else {
				break
			}
		}
	}
}

func Sort(v sort.Interface) {
	if v.Len() <= 1 {
		return
	}
	roots := make([]int, 0, 5)
	sizes := make([]int, 0, 5)
	roots = append(roots, 0)
	sizes = append(sizes, 1)

	// Build
	for i := 1; i < v.Len(); i++ {
		// Add the next element to the string of heaps
		llen := len(roots)
		if llen >= 2 && sizes[llen-2] == sizes[llen-1]+1 {
			roots = roots[0 : len(roots)-1]
			sizes = sizes[0 : len(sizes)-1]
			roots[len(roots)-1] = i
			sizes[len(sizes)-1]++
		} else {
			roots = append(roots, i)
			if sizes[len(sizes)-1] == 1 {
				sizes = append(sizes, 0)
			} else {
				sizes = append(sizes, 1)
			}
		}

		// stringify - Despite what wikipedia says I think we only need to maintain
		// the string property when the heap that was just added has exactly one
		// element.  If we are combining heaps to make a new heap then those leaf nodes
		// already satisfy the string property and the larger of those will bubble up
		// when we heapify and will obviously still satisfy the string property.
		rooti := len(roots) - 1
		if sizes[rooti] <= 1 {
			rooti = stringify(v, roots, sizes)
			if rooti != len(roots)-1 {
				heapify(v, roots[rooti], sizes[rooti])
			}
		} else {
			heapify(v, roots[rooti], sizes[rooti])
		}
	}

	// Shrink
	for len(roots) > 0 {
		root := roots[len(roots)-1]
		size := sizes[len(sizes)-1]
		roots = roots[0 : len(roots)-1]
		sizes = sizes[0 : len(sizes)-1]
		if size > 1 {
			right := root - 1
			left := right - leo[size-2]
			roots = append(roots, left)
			sizes = append(sizes, size-1)
			rooti := stringify(v, roots, sizes)
			if rooti < len(roots)-1 {
				heapify(v, roots[rooti], sizes[rooti])
			}
			roots = append(roots, right)
			sizes = append(sizes, size-2)
			rooti = stringify(v, roots, sizes)
			if rooti < len(roots)-1 {
				heapify(v, roots[rooti], sizes[rooti])
			}
		}
	}
}

func Ints(a []int)         { Sort(sort.IntSlice(a)) }
func Float64s(a []float64) { Sort(sort.Float64Slice(a)) }
func Strings(a []string)   { Sort(sort.StringSlice(a)) }
