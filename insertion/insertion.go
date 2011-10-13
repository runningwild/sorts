package insertion

import "sort"

func Sort(v sort.Interface) {
  for i := 1; i < v.Len(); i++ {
    var insert int
    for insert = i - 1; insert >= 0 && v.Less(i, insert); insert-- { }
    for target := i; target > insert + 1; target-- {
      v.Swap(target, target - 1)
    }
  }
}
