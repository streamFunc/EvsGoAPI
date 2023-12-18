package node

import (
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

const (
	filePcmPath = "../test16K.pcm"
)

func TestEncoder(t *testing.T) {
	enc := NewEvsEncoder()
	enc.IsG192 = 0
	enc.SampleRate = 16000
	enc.MaxBand = "WB"
	enc.BitRate = 24400
	enc.StartEncoder()

	file, err := os.Open(filePcmPath)
	if err != nil {
		return
	}
	defer file.Close()

	buffer := make([]byte, 160)
	ticker := time.Tick(20 * time.Millisecond)

	for range ticker {
		n, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Read file error:", err)
			}
			break
		}

		enc.EncodePcmToEvs(buffer[:n])
	}
	enc.StopEncoder()
}
