//mutualexclusion - обоюдное исключение/ Mutex - примитив синхронизации потоков.

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var online map[string]int
var mux sync.Mutex

func main() {
	online = make(map[string]int, 3)
	for {
		var str = "post/"
		n := rand.Intn(3) + 1
		url := str + strconv.Itoa(n)
		go work(url)
		c := rand.Intn(10) + 50
		time.Sleep(time.Millisecond * time.Duration(c))
		fmt.Println(online)
	}

}

func work(url string) {
	//mutex - blocked codesrc. позволяет блокировать участки кода. Когда заблокирован -
	// второй блокировщик станет в очередь
	mux.Lock()
	online[url]++
	mux.Unlock()

	//fake work
	time.Sleep(time.Millisecond * 100)
	mux.Lock()
	online[url]--
	mux.Unlock()

}
