package main

import (
	"go-concurrency-live-code/internal/service"
)

func main() {
	// service.ProcessData()
	// service.ProcessDataWithConcurrency()
	service.ProcessDataWithChannels()
}
