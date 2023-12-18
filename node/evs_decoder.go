package node

import (
	"errors"
	"fmt"
)

// EvsDecoder this node receives pcm data, keep tracking each ongoing telephone event code until end-bit in its last
// frame is seen, then send the finished event code to output in pcm packet 16bit pcm_s16se
type EvsDecoder struct {
	SampleRate     int
	BitRate        int
	IsG192         int //0 MIME 1 G192
	isDecoderStart bool
	realBitRate    int
	debugPrint     bool
	ctx            *EvsDecoderContext
}

const (
	Br7200   byte = 0x1
	Br8000   byte = 0x2
	Br9600   byte = 0x3
	Br13200  byte = 0x4
	Br16400  byte = 0x5
	Br24400  byte = 0x6
	Br32000  byte = 0x7
	Br48000  byte = 0x8
	Br64000  byte = 0x9
	Br96000  byte = 0xa
	Br128000 byte = 0xb

	Br6600  byte = 0x30
	Br8850  byte = 0x31
	Br12650 byte = 0x32
	Br14250 byte = 0x33
	Br15850 byte = 0x34
	Br18250 byte = 0x35
	Br19850 byte = 0x36
	Br23050 byte = 0x37
	Br23850 byte = 0x38
)

func NewEvsDecoder() *EvsDecoder {
	dec := EvsDecoder{}
	dec.ctx = newEvsDecoderContext()
	return &dec
}

func (n *EvsDecoder) StartDecoder() error {
	if n.SampleRate <= 0 {
		n.SampleRate = 8000
	}
	n.isDecoderStart = false
	if n.BitRate == 0 {
		n.BitRate = 9600
	}
	n.debugPrint = true

	fmt.Printf("EvsDecoder init stream isG192:%v sample rate %v bitRate:%v\n", n.IsG192, n.SampleRate, n.BitRate)

	if n.ctx != nil {
		res := n.ctx.initEvsDecoder(n.SampleRate, n.BitRate, n.IsG192)
		if res != 0 {
			return errors.New(fmt.Sprintf("evsDecoder init fail"))
		} else {
			n.isDecoderStart = true
			fmt.Printf("EvsDecoder StartDecoder success\n")
		}
	}
	return nil
}

// DecodeEvsToPcm input payload evs
// output pcm
func (n *EvsDecoder) DecodeEvsToPcm(payload []byte) []byte {
	if len(payload) <= 7 { //2.4 (EVS Primary SID) and Special case(see clause A.2.1.3)
		return nil
	}
	newPayload := n.addToCHeader(payload)
	evsDecoderData := n.ctx.startEvsDecoder(newPayload)
	if evsDecoderData == nil {
		fmt.Printf("evsDecode data nil error return\n")
		return nil
	}
	return evsDecoderData
}

func (n *EvsDecoder) StopDecoder() {
	if n.ctx != nil {
		if n.isDecoderStart {
			n.isDecoderStart = false
			n.ctx.stopEvsDecoder()
			fmt.Printf("node %v StopDecoder success now\n", n)
		}
		n.ctx = nil
	}
}

// EVS Primary mode need to set Toc Header
func (n *EvsDecoder) addToCHeader(data []byte) []byte {

	payLoadLen := len(data)
	var toc byte
	if payLoadLen == 17 { //EVS AMR-WB IO
		n.realBitRate = 6600
		toc = Br6600
	} else if payLoadLen == 18 {
		n.realBitRate = 7200
		toc = Br7200
	} else if payLoadLen == 20 {
		n.realBitRate = 8000
		toc = Br8000
	} else if payLoadLen == 23 { //EVS AMR-WB IO
		n.realBitRate = 8850
		toc = Br8850
	} else if payLoadLen == 24 {
		n.realBitRate = 9600
		toc = Br9600
	} else if payLoadLen == 32 { //EVS AMR-WB IO
		n.realBitRate = 12650
		toc = Br12650
	} else if payLoadLen == 33 {
		n.realBitRate = 13200
		toc = Br13200
	} else if payLoadLen == 36 { //EVS AMR-WB IO
		n.realBitRate = 14250
		toc = Br14250
	} else if payLoadLen == 40 { //EVS AMR-WB IO
		n.realBitRate = 15850
		toc = Br15850
	} else if payLoadLen == 41 {
		n.realBitRate = 16400
		toc = Br16400
	} else if payLoadLen == 46 { //EVS AMR-WB IO
		n.realBitRate = 18250
		toc = Br18250
	} else if payLoadLen == 50 { //EVS AMR-WB IO
		n.realBitRate = 19850
		toc = Br19850
	} else if payLoadLen == 58 { //EVS AMR-WB IO
		n.realBitRate = 23050
		toc = Br23050
	} else if payLoadLen == 60 { //EVS AMR-WB IO
		n.realBitRate = 23850
		toc = Br23850
	} else if payLoadLen == 61 {
		n.realBitRate = 24400
		toc = Br24400
	} else if payLoadLen == 80 {
		n.realBitRate = 32000
		toc = Br32000
	} else if payLoadLen == 120 {
		n.realBitRate = 48000
		toc = Br48000
	} else if payLoadLen == 160 {
		n.realBitRate = 64000
		toc = Br64000
	} else if payLoadLen == 240 {
		n.realBitRate = 96000
		toc = Br96000
	} else if payLoadLen == 320 {
		n.realBitRate = 128000
		toc = Br128000
	} else {
		toc = 0
	}

	if n.debugPrint {
		fmt.Printf("EvsDecoder data len:%v realBitRate:%v \n", payLoadLen, n.realBitRate)
		n.debugPrint = false
	}

	if toc == 0 {
		fmt.Printf("error,toc == 0,not support bitrate,set toc header\n")
		return data
	}

	newData := make([]byte, len(data)+1)
	newData[0] = toc
	copy(newData[1:], data)

	return newData
}
