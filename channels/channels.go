package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func chanIter() {

	// Without buffering, the receiving channel has to be along with the sending channel since it can only process one at the time
	ch := make(chan int)
	for j := 0; j < 5; j++ {
		wg.Add(2)
		// Receive only channel
		go func(ch <-chan int) {
			k := <-ch
			fmt.Println(k)
			wg.Done()
		}(ch)
		// Sending only channel
		go func(ch chan<- int) {
			ch <- 99
			wg.Done()
		}(ch)
	}
	wg.Wait()
}

// Select statements
const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // Signal only channel

func logger() {
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			break
			// default:
			//
		}
	}
}

func executeLog() {
	go logger()
	// Alternative way
	// defer func() {
	// 	close(logCh)
	// }()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Microsecond)
	doneCh <- struct{}{}
}

func bufferedChan() {
	ch := make(chan int, 50) // Can store now 50 ints
	wg.Add(2)
	// Receive only channel
	go func(ch <-chan int) {
		// for i := range ch {
		// 	fmt.Println(i)
		// }
		// Alternative way
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)
	// Sending only channel
	go func(ch chan<- int) {
		ch <- 287
		ch <- 18
		close(ch) // This is necessary to avoid deadlock, it is used to lock the length of the channel
		wg.Done()
	}(ch)
	wg.Wait()
}

func main() {
	ch := make(chan int)
	wg.Add(2)
	// Receiving
	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()
	// Sending
	go func() {
		i := 42
		ch <- i // This passes a copy of the data
		i = 27
		wg.Done()
	}()
	wg.Wait()

	// Using iterators
	chanIter()
	bufferedChan()

	// Select statemtent
	executeLog()

}
