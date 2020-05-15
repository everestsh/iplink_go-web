package main

import "fmt"
import "time"

type User struct {
	username string
}

func (this *User) Close() {
	fmt.Println(this.username, "Closed !!!")
}

func main() {
	{
		//
		u1 := &User{"jack"}
		defer u1.Close()
	}
	u2 := &User{"lily"}
	defer u2.Close()
	time.Sleep(10 * time.Second)

	fmt.Println("Done !")
}
