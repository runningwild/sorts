Currently this repository only contains one sort, Smoothsort, that can be used as an alternative to Quicksort.  Details on the algorithm can be found at http://en.wikipedia.org/wiki/Smoothsort.  The smooth package can be goinstalled as follows:

    go get github.com/runningwild/sorts/smooth

The package uses the same interface as the standard Go sort package, namely it sorts any structure conforming to sort.interface.  The same convenient wrappers are included for sorting slices of ints, float64s, and strings as in the Go's standard sort package, so this package can be used as a drop-in replacement for the standard sort package.

Here is a comparison of the number of swaps and comparisons performed when using Go's Quicksort or this package's Smoothsort on arrays of 1k elements that are already sorted or reverse-sorted.

                            swaps comparisons
    Quicksort on sorted:      682       15896
    Smoothsort on sorted:       0        2481
    Quicksort on reversed:   1122       15587
    Smoothsort on reversed: 10708       26436


Here are some benchmarks on my laptop (weekly.2012-02-14, darwin/amd64).  MostlySorted is an array, A, that is first sorted, then every fifth element in A is swapped with a random element in A.

    Smoothsort on 1k elements
    Sorted        101431 ns/op
    MostlySorted  550455 ns/op
    Shuffled       95635 ns/op
    Reversed      730083 ns/op

    Quicksort on 1k elements
    Sorted        243270 ns/op
    MostlySorted  299113 ns/op
    Shuffled       39798 ns/op
    Reversed      245046 ns/op

    Smoothsort on 1M elements
    Sorted         101.388 ms/op
    MostlySorted  1002.774 ms/op
    Shuffled        93.598 ms/op
    Reversed      1290.697 ms/op

    Quicksort on 1M elements
    Sorted         519.893 ms/op
    MostlySorted   625.448 ms/op
    Shuffled        37.308 ms/op
    Reversed       523.922 ms/op

And to show that Smoothsort takes linear time when the input is sorted:

    Elements             Time

          10       1326 ns/op
         100      10682 ns/op
        1000     102360 ns/op
       10000    1018905 ns/op
      100000   10223760 ns/op
     1000000  103879200 ns/op
