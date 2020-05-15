package main

import "fmt"

func main (){
  var a int
  var b = 5
  var (
    c=6
    c1=8
  )
  c2 := 6
  fmt.Printf("a     变量的地址是: %x\n", &a  )
  fmt.Printf("b     变量的地址是: %x\n", &b  )
  fmt.Printf("c     变量的地址是: %x\n", &c  )
  fmt.Printf("c1    变量的地址是: %x\n", &c1  )
  fmt.Printf("c2    变量的地址是: %x\n", &c2  )
  var ip *int
  var ip0, ip1 *int


  fmt.Printf("ip    变量储存的指针地址: %x\n", ip )
  fmt.Printf("ip0   变量储存的指针地址: %x\n", ip0 )
  fmt.Printf("ip1   变量储存的指针地址: %x\n", ip1 )
  ip = &a
  ip0 = &b
  ip1 = &c
  var (
  ip2 = &c1
  ip3 = &c2
  )
  fmt.Printf(" ip   变量储存的指针地址: %x\n", ip )
  fmt.Printf("*ip   变量的值: %d\n", *ip )
  fmt.Printf(" ip0  变量储存的指针地址: %x\n", ip0 )
  fmt.Printf("*ip0  变量的值: %d\n", *ip0 )
  fmt.Printf(" ip1  变量储存的指针地址: %x\n", ip1 )
  fmt.Printf("*ip1  变量的值: %d\n", *ip1 )
  fmt.Printf(" ip2  变量储存的指针地址: %x\n", ip2 )
  fmt.Printf("*ip2  变量的值: %d\n", *ip2 )
  fmt.Printf(" ip3  变量储存的指针地址: %x\n", ip3 )
  fmt.Printf("*ip3  变量的值: %d\n", *ip3 )
}
