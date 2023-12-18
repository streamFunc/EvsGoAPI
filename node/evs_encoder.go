package node

import (
	"errors"
	"fmt"
)

// EvsEncoder put received pcm data 16bit pcm_s16se
type EvsEncoder struct {
	MaxBand        string //NB, WB, SWB or FB
	SampleRate     int
	IsG192         int
	BitRate        int
	isEncoderStart bool
	debugPrint     bool
	ctx            *EvsEncoderContext
}

func NewEvsEncoder() *EvsEncoder {
	enc := EvsEncoder{}
	enc.ctx = newEvsEncoderContext()
	return &enc
}

func (n *EvsEncoder) StartEncoder() error {
	if n.SampleRate <= 0 {
		n.SampleRate = 8000
	}

	if n.BitRate <= 0 {
		n.BitRate = 9600
	}

	n.isEncoderStart = false
	n.debugPrint = true
	if n.MaxBand == "" {
		n.MaxBand = "WB"
	}

	fmt.Printf("EvsEncoder init stream isG192:%v maxBand:%v sample rate %v bitRate:%v\n", n.IsG192, n.MaxBand, n.SampleRate, n.BitRate)
	if n.ctx != nil {
		res := n.ctx.initEvsEncoder(n.SampleRate, n.BitRate, n.MaxBand, n.IsG192)
		if res != 0 {
			return errors.New(fmt.Sprintf("evsEncode init fail"))
		} else {
			n.isEncoderStart = true
			fmt.Printf("EvsEncoder StartEncoder success\n")
		}
	}
	return nil
}

// EncodePcmToEvs data then send result (if any) to receiver
// input pcm
// output evs
func (n *EvsEncoder) EncodePcmToEvs(data []byte) []byte {
	frame := n.ctx.startEvsEncoder(data)
	if frame == nil {
		fmt.Printf("node %v  encoderToReceiver frame is nil error\n", n)
		return nil
	}
	newFrame := n.delToCHeader(frame)
	if n.debugPrint {
		fmt.Printf("after EvsEncoder len newFrame:%v\n", len(newFrame))
		n.debugPrint = false
	}

	return newFrame
}

func (n *EvsEncoder) StopEncoder() {
	if n.ctx != nil {
		if n.isEncoderStart {
			n.isEncoderStart = false
			n.ctx.stopEvsEncoder()
			fmt.Printf("node %v StopEncoder success now\n", n)
		}
		n.ctx = nil
	}
}

// EVS Primary mode need to del Toc Header
func (n *EvsEncoder) delToCHeader(data []byte) []byte {
	if len(data) <= 0 {
		return nil
	}
	newData := make([]byte, len(data)-1)
	copy(newData, data[1:])

	return newData
}
