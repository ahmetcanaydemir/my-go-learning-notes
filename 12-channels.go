/*
Channels, race problemi olmadan güvenli bir şekilde goroutine kullanmayı amaçlar.


var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)
	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		ch <- 42
		wg.Done()
	}()
	wg.Wait()
}



func main() {
	ch := make(chan int)

	for j := 0; j < 5; j++ {
		wg.Add(2)
		go func() {
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()
		go func() {
			ch <- 42
			wg.Done()
		}()
	}
	wg.Wait()
}


// 2 goroutine de hem okuyup hem yazıyor :


func main() {
	ch := make(chan int)
	wg.Add(2)
	go func() {
		i := <-ch
		fmt.Println(i)
		ch <- 25
		wg.Done()
	}()
	go func() {
		ch <- 42
		fmt.Println(<-ch)
		wg.Done()
	}()
	wg.Wait()
}


// read-only ve send-only channel

func main() {
	ch := make(chan int)
	wg.Add(2)

	// read only channel
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)

	// send only channel
	go func(ch chan<- int) {
		ch <- 42
		wg.Done()
	}(ch)
	wg.Wait()
}


//Channel'ı kapatma (closed channel'a mesaj gönderilemez ve panic durumu olmadan kapalı mı anlaşılamaz)
func main() {
	ch := make(chan int, 50)
	wg.Add(2)

	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 25
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}

// Aynı işlemin for range olmadan yapma hali
func main() {
	ch := make(chan int, 50)
	wg.Add(2)

	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 25
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}

*/

package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50)
	wg.Add(2)

	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 25
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}
