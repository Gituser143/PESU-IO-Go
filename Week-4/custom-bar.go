package main5

import (
	"time"

	"github.com/cheggaaa/pb/v3"
)

func main() {
	count := 100000

	// create and start new bar
	bar := pb.StartNew(count)

	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	bar.Finish()
}
