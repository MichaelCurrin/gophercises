package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const quantityConcurrent = 5

var ids = []string{"1", "abort", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}

func main() {
	var wg sync.WaitGroup
	fmt.Println("Hello, playground")
	var syncChan = make(chan bool, quantityConcurrent)
	var abortChan = make(chan bool)
	go func(_syncChan *chan bool, _abortChan *chan bool, _wg *sync.WaitGroup) {
		j := 1
		defer _wg.Done()
		<-*_abortChan
		_wg.Add(1)
		for i := 0; i <= quantityConcurrent-j; i++ {
			fmt.Println(i)
			syncChan <- true
		}
		fmt.Println("Finish channel all request")
		*_syncChan = make(chan bool, quantityConcurrent-j)
		*_abortChan = make(chan bool)

	}(&syncChan, &abortChan, &wg)
	for _, _id := range ids {
		wg.Wait()
		fmt.Println("go")
		syncChan <- true
		go func(id string) {
			err := makeRequest(id)
			if err != nil {
				fmt.Println(err)
				abortChan <- true
			}
			<-syncChan
		}(_id)
	}
	fmt.Println("flush")
	for i := 0; i < quantityConcurrent-1; i++ {
		syncChan <- true
	}

}

func makeRequest(str string) error {
	fmt.Println(str + " request")
	time.Sleep(time.Second * 2)
	if str == "abort" {
		return errors.New(str + " ABORT")
	}
	fmt.Println(str + " finish")
	return nil
}
