package main

import (
    "fmt"
    "math"
)

func main(){
    a := 2.4
    b := 1.6

    c := math.Pow(a,2)
    fmt.Printf("%.2f^%d = %.2f \n",a,2,c)

    c = math.Sin(a)
    fmt.Printf("Sin(%.2f) = %.2f \n",a,c)

    c = math.Cos(b)
    fmt.Printf("Cos(%.2f) = %.2f \n",b,c)

    c = math.Sqrt(a*b)
    fmt.Printf("Sqrt(%.2f*%.2f) = %.2f \n",a,b,c)
}
