package smooth

import (
  "fmt"
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

func (l leap) childRoots() (left,right int) {
  right = l.root - 1
  left = right - leo[l.size - 2]
  return
}

func makeLeaps(start,remaining int, l *[]leap) {
  if remaining == 0 { return }
  i := 0
  for leo[i] <= remaining {
    fmt.Printf("%d %d\n", leo[i], remaining)
    i++
  }
  i--
  root := start + leo[i]
  *l = append(*l, leap{ root : root, size : i })
  makeLeaps(root, remaining - leo[i], l)
}

func stringify(v []int, leaps []leap) int {
  k := len(leaps) - 1
  for j := k - 1; j >= 0; j-- {
    val := v[leaps[j].root]
    if val > v[leaps[k].root] {
      size := leaps[k].size
      if size > 1 {
        left,right := leaps[k].childRoots()
        if val > v[right] && val > v[left] {
          v[leaps[k].root],v[leaps[j].root] = v[leaps[j].root],v[leaps[k].root]
          k = j
        }
      } else {
        v[leaps[k].root],v[leaps[j].root] = v[leaps[j].root],v[leaps[k].root]
        k = j
      }
    } else {
      return k
    }
  }
  return k
}

func heapify(v []int, cleap leap) {
  for cleap.size > 1 {
    left,right := cleap.childRoots()
    if v[right] > v[left] {
      if v[right] > v[cleap.root] {
        v[cleap.root],v[right] = v[right],v[cleap.root]
        cleap.root = right
        cleap.size -= 2
      } else {
        break
      }
    } else {
      if v[left] > v[cleap.root] {
        v[cleap.root],v[left] = v[left],v[cleap.root]
        cleap.root = left
        cleap.size -= 1
      } else {
        break
      }
    }
  }
}

func fsckHeap(v []int, cleap leap) bool {
  if cleap.size <= 1 { return false }
  left,right := cleap.childRoots()
  if v[left] > v[cleap.root] || v[right] > v[cleap.root] {
    print("Failed on heap\n")
    return true
  }
  return fsckHeap(v, leap{left, cleap.size-1}) || fsckHeap(v, leap{right, cleap.size-2})
}
func fsck(v []int, leaps []leap) bool {
  for i := range leaps {
    if fsckHeap(v, leaps[i]) { return true }
  }
  for i := 1; i < len(leaps); i++ {
    if v[leaps[i-1].root] > v[leaps[i].root] {
      print("Failed on string\n")
      return true
    }
  }
  return false
}

func Sort(v []int) {
  if len(v) <= 1 { return }
  var leaps []leap
  leaps = append(leaps, leap{0,1})

  // Build
  for i := 1; i < len(v); i++ {
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
    }

    // Should be able to avoid doing this in some cases, not quite sure what though.
    heapify(v, leaps[leapi])
  }

  // Shrink
  for len(leaps) > 0 {
    cleap := leaps[len(leaps) - 1]
    leaps = leaps[0 : len(leaps) - 1]
    if cleap.size > 1 {
      left,right := cleap.childRoots()
      leaps = append(leaps, leap{ root : left, size : cleap.size-1 })
      leapi := stringify(v, leaps)
      heapify(v, leaps[leapi])
      leaps = append(leaps, leap{ root : right, size : cleap.size-2 })
      leapi = stringify(v, leaps)
      heapify(v, leaps[leapi])
    }
  }
}
