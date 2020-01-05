package main

import "fmt"

func main() {

    selected := 2

    switch selected {
        case 0:
            fmt.Println("selected = 0")
        case 1:
            fmt.Println("selected = 1")
        case 2:
            fmt.Println("selected = 2")
        case 3:
            fmt.Println("selected = 3")
        default:
            fmt.Println("other..")

    }

}
