/*

Java ve C# gibi thread havuzu oluşturmak gibi pahalı bir işlem yerine Thread'leri soyutlaştırıp kullanma işlemi


func main() {
	var msg string = "Hello"
	go func() {
		fmt.Println(msg)
	}()
	time.Sleep(100 * time.Millisecond)
}


// Sync with WaitGroup

var wg = sync.WaitGroup{}

func main() {
	var msg string = "Hello"
	wg.Add(1)
	go func(m string) {
		fmt.Println(m)
		wg.Done()
	}(msg)
	msg = "Goodbye"
	wg.Wait()
}


// Mutex örnek

var wg = sync.WaitGroup{}
var counter = 1
var m = sync.RWMutex{}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}

// Thread sayısını görmek veya değiştirmek
fmt.Println("Threads: ", runtime.GOMAXPROCS(-1))
runtime.GOMAXPROCS(1)
*/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("Threads: ", runtime.GOMAXPROCS(-1))
}
