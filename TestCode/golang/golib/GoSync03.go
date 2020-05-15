//?????

package main

import (
	"fmt"
	"sync"
)

func main() {
	var l *sync.RWMutex
	l = new(sync.RWMutex)
	l.Unlock() //1个RUnlock
	fmt.Println("1")
	l.RLock()
}
