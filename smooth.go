// Implementation of Edsger Dijkstra's smooth sort
// Worst case: O(n log n)
// Best case: O(n)
// Called 'smooth' because it has a smooth transition between linear time and
// O(n log n) time as the input set transitions from sorted to unsorted.
package smooth

import (
  "sort"
)

var leo []int

func init() {
  leo = []int{1,1}
  length := 2
  for leo[length-1] < 1000000000 {
    leo = append(leo, leo[length-1] + leo[length-2] + 1)
    length++
  }
}

type leap struct {
  root int  // index of the root of the heap in the input array
  size int  // the size of the heap is given by leo[heap.size]
}

// Stringify will reorder the root nodes to make sure that they are in
// increasing order.  This is called when a new heap is added at the end
// such that the only root node that is out of order is the new one.
func stringify(v sort.Interface, leaps []leap) int {
  k := len(leaps) - 1
  for j := k - 1; j >= 0; j-- {
    jr := leaps[j].root
    kr := leaps[k].root
    if v.Less(kr, jr) {
      size := leaps[k].size
      if size <= 1 {
        v.Swap(jr, kr)
        k = j
      } else {
        right := leaps[k].root - 1
        left := right - leo[leaps[k].size - 2]
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
func heapify(v sort.Interface, cleap leap) {
  for cleap.size > 1 {
    right := cleap.root - 1
    left := right - leo[cleap.size - 2]
    if v.Less(left, right) {
      if v.Less(cleap.root, right) {
        v.Swap(cleap.root, right)
        cleap.root = right
        cleap.size -= 2
      } else {
        break
      }
    } else {
      if v.Less(cleap.root, left) {
        v.Swap(cleap.root, left)
        cleap.root = left
        cleap.size -= 1
      } else {
        break
      }
    }
  }
}

func Sort(v sort.Interface) {
  if v.Len() <= 1 { return }
  leaps := make([]leap, 0, 5)
  leaps = append(leaps, leap{0,1})

  // Build
  for i := 1; i < v.Len(); i++ {
    // Add the next element to the string of heaps
    llen := len(leaps)
    if llen >= 2 && leaps[llen-2].size == leaps[llen-1].size + 1 {
      leaps = leaps[0 : len(leaps) - 1]
      leaps[len(leaps) - 1] = leap{ root : i, size : leaps[len(leaps)-1].size + 1 }
    } else {
      if leaps[len(leaps)-1].size == 1 {
        leaps = append(leaps, leap{ root : i, size : 0 })
      } else {
        leaps = append(leaps, leap{ root : i, size : 1 })
      }
    }

    // stringify - Despite what wikipedia says I think we only need to maintain
    // the string property when the heap that was just added has exactly one
    // element.  If we are combining heaps to make a new heap then those leaf nodes
    // already satisfy the string property and the larger of those will bubble up
    // when we heapify and will obviously still satisfy the string property.
    leapi := len(leaps) - 1
    if leaps[leapi].size <= 1 {
      leapi = stringify(v, leaps)
      if leapi != len(leaps) - 1 {
        heapify(v, leaps[leapi])
      }
    } else {
      heapify(v, leaps[leapi])
    }
  }

  // Shrink
  for len(leaps) > 0 {
    cleap := leaps[len(leaps) - 1]
    leaps = leaps[0 : len(leaps) - 1]
    if cleap.size > 1 {
      right := cleap.root - 1
      left := right - leo[cleap.size - 2]
      leaps = append(leaps, leap{ root : left, size : cleap.size-1 })
      leapi := stringify(v, leaps)
      if leapi < len(leaps) - 1 {
        heapify(v, leaps[leapi])
      }
      leaps = append(leaps, leap{ root : right, size : cleap.size-2 })
      leapi = stringify(v, leaps)
      if leapi < len(leaps) - 1 {
         heapify(v, leaps[leapi])
      }
    }
  }
}


func Ints(a []int) { Sort(sort.IntSlice(a)) }

