package database

import (
	"fmt"
	"math/rand"
	"time"
)

func ReadFromDatabase() string {
	seconds := time.Duration(rand.Intn(3) + 1)
	time.Sleep(seconds * time.Second)
	return fmt.Sprintf("Finished reading from Database. Duration: %d seconds", seconds)
}
