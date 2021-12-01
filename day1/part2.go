package main

import (
    "fmt"
)

func sum_triplet(t []int) int {
    return t[0] + t[1] + t[2]
}

func main() {
    t := make([]int, 0) // Triplets
    s := make([]int, 0) // Sums of triplets
    for {
        d := 0 // Depth
        n, _ := fmt.Scanf("%d", &d)
        if n == 0 {
            break
        }

        // Append to a triplet, if all three values exist store them and remove first
        t = append(t, d)
        if len(t) == 3 {
           s = append(s, sum_triplet(t))
           t = t[1:]
        }
    }

    o := -1 // Old (negative means there was none)
    i := 0 // Number of increases in depth of triplets
    for _, v := range(s) {
        if o != -1 && v > o {
            i++
        }
        o = v
    }
    fmt.Println("Increases in value of triplets:", i)
}
