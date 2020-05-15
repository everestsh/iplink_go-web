package main

import "fmt"
import "time"

type User struct {
	username string
}

func (this *User) close() {
	fmt.Println(this.username, "Closed !")
}
func f(u *User) {
	defer u.close()
	//u1 do something
}

func main() {
	u1 := &User{"jack"}
	f(u1) //这样的方式，u1才不会不依赖main函数的执行

	//这样的方式，u2也不会依赖main函数的执行
	u2 := &User{"lily"}
	m := func() {
		defer u2.close()
		//u2 do something
	}
	//  m()<pre name="code" class="plain">
	{
		defer u2.close()
		//u2 do something
		//  }() </pre>
		time.Sleep(10 * time.Second)
		fmt.Println("Dome !")
	}
}
