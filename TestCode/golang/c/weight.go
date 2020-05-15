package main

import (
  "fmt"
)
//1 >现在有1克到127克之间127个整数重量，
//使用循环解决砝码问题的练习
func main()  {
  var weight, last int
  fmt.Printf("需要的砝码:\n")
  for weight = 1; weight <= 100; weight++ {
      if weight >= 2 * last {
        fmt.Printf("%d ", weight)
        last = weight
      }
  }
  fmt.Printf("\n")
}
