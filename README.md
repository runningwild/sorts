Currently this repository only contains one sort, Smoothsort, that can be used as an alternative to Quicksort.  Details on the algorithm can be found at http://en.wikipedia.org/wiki/Smoothsort.  The smooth package can be goinstalled as follows:

    goinstall github.com/runningwild/sorts/smooth

The package uses the same interface as the standard Go sort package, namely it sorts any structure conforming to sort.interface.  Convenient wrappers are included for sorting slices of ints, float64s, and strings.

Testing is done with gospec which can be found at git://github.com/orfjackal/gospec.git

Here is a comparison of the number of swaps and comparisons performed when using Go's Quicksort or this package's Smoothsort on arrays of 1k elements that are already sorted or reverse-sorted.

                            swaps comparisons
    Quicksort on sorted:      682       15896
    Smoothsort on sorted:       0        2481
    Quicksort on reversed:   1122       15587
    Smoothsort on reversed: 10708       26436


Here are some benchmarks on my laptop (r60.2, darwin/amd64).  Copy is a benchmark that just includes the copying required on Unsorted, Shuffled, and PartiallySorted since those slices cannot just be reused because they get sorted:

    Arrays of 1M elements:
    Copy                                500     3487014 ns/op
    Quicksort On Unsorted                 5   485252600 ns/op
    Smoothsort On Unsorted                1  1234045000 ns/op
    Quicksort On Sorted                   5   479280600 ns/op
    Smoothsort On Sorted                 20    96429150 ns/op
    Quicksort On Shuffled                 5   709480600 ns/op
    Smoothsort On Shuffled                1  1648161000 ns/op
    Quicksort On PartiallySorted          5   518328800 ns/op
    Smoothsort On PartiallySorted         5   525332400 ns/op

    Arrays of 1k elements:
    Copy                             500000        3379 ns/op
    Quicksort On Unsorted             10000      236227 ns/op
    Smoothsort On Unsorted             5000      696742 ns/op
    Quicksort On Sorted               10000      234545 ns/op
    Smoothsort On Sorted              20000       94916 ns/op
    Quicksort On Shuffled              5000      351280 ns/op
    Smoothsort On Shuffled             2000      755692 ns/op
    Quicksort On PartiallySorted      10000      266571 ns/op
    Smoothsort On PartiallySorted      5000      352206 ns/op

And to show that Smoothsort takes linear time when the input is sorted:

    Smoothsort On Sorted 1000         20000       94340 ns/op
    Smoothsort On Sorted 10000         2000      936181 ns/op
    Smoothsort On Sorted 100000         200     9415185 ns/op
    Smoothsort On Sorted 1000000         20    95552350 ns/op


