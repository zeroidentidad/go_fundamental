package main

import (
    "fmt"
    "math/rand"
)

type Node struct {
    value int
    next *Node
}

func main(){

    var head *Node
    head = nil

    // add 5 data
    fmt.Println("=== insert 5 data=== ")
    n := 0
    for n < 5 {
        fmt.Printf("data %d\n",n)

        head = add(head,rand.Intn(100))
        n++
    }
    fmt.Println("=== display ===")
    display(head)

}

func add(list *Node,data int) *Node {
    if list==nil {
        list := new(Node)
        list.value = data
        list.next = nil

        return list
    }else{
        temp := new(Node)
        temp.value = data
        temp.next = list

        list = temp

        return list
    }
}

func display(list *Node){
    var temp *Node
    temp = list

    for temp!=nil {

        fmt.Println(temp.value)
        temp = temp.next
    }
}

