package main

import "fmt"

func main() {
    var (
        a = 5
        b = 8
    )

    fmt.Println(a>b && a!=b)
    fmt.Println(!(a>=b))
    fmt.Println(a==b || a>b)

}
