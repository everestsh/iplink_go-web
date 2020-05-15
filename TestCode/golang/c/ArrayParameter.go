package main

import (
  "fmt"
)
func sum( ap []int, n int) int {
  var m int
  for i := 0; i < 10; i++ {
    m += ap[i]
  }
  return m
}
func change(a []int, n int)  {
  a[n] = 900
  fmt.Printf("change : &n=%p,&a=%p,a=%p\n", &n, &a, a)
  fmt.Printf("change a[%d] = %d\n", n, a[n])
}
func main()  {
//  var  a [10]int = [10]int {1,2,3,4,5,6,7,8,9,10}
  a := []int {1,2,3,4,5,6,7,8,9,10}
  fmt.Printf("sum = %d\n", sum(a, 10))

  n := 5
  fmt.Printf("a[%d] = %d\n", n, a[n])
  fmt.Printf("Before : &n=%p,&a=%p,a=%p\n", &n, &a, a)
  change(a, n)
  fmt.Printf("After  : &n=%p,&a=%p,a=%p\n", &n, &a, a)
  fmt.Printf("after a[%d] = %d\n", n, a[n])
}
