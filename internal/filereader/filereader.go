package filereader

import (
	"fmt"
	"math/rand"
	"time"
)

func ReadFile() string {
	seconds := time.Duration(rand.Intn(3) + 1)
	time.Sleep(seconds * time.Second)
	return fmt.Sprintf("Finished reading from Files. Duration: %d seconds", seconds)
}
