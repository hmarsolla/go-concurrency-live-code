package service

import (
	"fmt"
	"go-concurrency-live-code/internal/database"
	"go-concurrency-live-code/internal/filereader"
	"math/rand"
	"sync"
	"time"
)

func ProcessData() {
	start := time.Now()

	filereader.ReadFile()
	database.ReadFromDatabase()

	elapsed := time.Since(start)
	fmt.Printf("Duration: %.2f seconds", elapsed.Seconds())
}

func ProcessDataWithConcurrency() {
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		filereader.ReadFile()
	}()

	go func() {
		defer wg.Done()
		database.ReadFromDatabase()
	}()

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Duration: %.2f seconds", elapsed.Seconds())
}

func ProcessDataWithChannels() {
	start := time.Now()

	dataChannel := make(chan string)
	done := make(chan bool)

	go gatherAllData(dataChannel)
	go readReceivedData(dataChannel, done)

	<-done

	elapsed := time.Since(start)
	fmt.Printf("Duration: %.2f seconds", elapsed.Seconds())
}

func gatherAllData(dataChannel chan string) {
	var wg sync.WaitGroup
	numRoutinesFile := rand.Intn(3) + 1
	numRoutinesDatabase := rand.Intn(3) + 1
	wg.Add(numRoutinesFile + numRoutinesDatabase)

	for i := 0; i < numRoutinesFile; i++ {
		go func() {
			defer wg.Done()
			dataChannel <- filereader.ReadFile()
		}()
	}

	for i := 0; i < numRoutinesDatabase; i++ {
		go func() {
			defer wg.Done()
			dataChannel <- database.ReadFromDatabase()
		}()
	}

	wg.Wait()
	close(dataChannel)
}

func readReceivedData(dataChannel <-chan string, done chan<- bool) {
	for data := range dataChannel {
		fmt.Println(data)
	}
	done <- true
}
