package node

import (
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

const (
	fileEvsPath = "../encoder.evs"
)

func TestDecoder(t *testing.T) {
	dec := NewEvsDecoder()
	dec.IsG192 = 0
	dec.SampleRate = 16000
	dec.BitRate = 24400
	dec.StartDecoder()

	file, err := os.Open(fileEvsPath)
	if err != nil {
		return
	}
	defer file.Close()

	buffer := make([]byte, 61)
	ticker := time.Tick(20 * time.Millisecond)

	for range ticker {
		n, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Read file error:", err)
			}
			break
		}

		dec.DecodeEvsToPcm(buffer[:n])
	}
	dec.StopDecoder()
}
