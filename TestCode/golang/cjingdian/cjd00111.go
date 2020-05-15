// 一维数组的倒置
package main

import "fmt"

const (
	M = 20
)

//*参数及功能说明如下：
//*x：要倒置的一维数组，其类型为整形指针。
//*n:数组元件的个数，其类型为整数
//*功能：实现一维数组的倒置
func fun(x []int, n int) int {
	var i, j, m int
	j = n - 1
	m = n / 2
	for i = 0; i < m; i++ {
		t := x[i]
		x[i] = x[j]
		x[j] = t
		j--
	}
	return m
}

func main() {
	var i, n int
	//  var a []int
	var m [M]int
	//var a[]int
	fmt.Printf("\nEnter n:\n")
	fmt.Scanf("%d", &n)
	fmt.Printf("\nn=%d:\n", n)

	fmt.Printf("The original array:\n")
	for i = 0; i < n; i++ {
		fmt.Scanf("%d", &m[i])
	}
	fmt.Println(m)
	//a = m[0 : n+1]
	//fun(a,i)
	fun(m[:], i)
	fmt.Printf("The array inverted:\n")
	for i = 0; i < n; i++ {
		fmt.Printf("%d\n", m[i])
	}
	fmt.Println(m)
}
