package main

import "fmt"

func IntSliceMax(l []int) int {
    if len(l) == 0 {
        panic("arg is an empty sequence")
    }

    max := l[0]
    for _, elem := range(l) {
        if elem > max {
            max = elem
        }
    }
    return max
}

func main() {
    fmt.Printf("3 = %v\n", IntSliceMax([]int{1,2,3}))
    fmt.Printf("-50 = %v\n", IntSliceMax([]int{-50}))
    fmt.Printf("%v = 54545\n", IntSliceMax([]int{-50, 30, 19, 54545, -56}))
    fmt.Printf("3 = %v\n", IntSliceMax([]int{}))
}
