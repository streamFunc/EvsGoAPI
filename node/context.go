package node

import "C"
import (
	"fmt"
	"unsafe"
)

//#cgo  CFLAGS:-I../source_code/fixed-point -I../source_code/fixed-point/basic_math -I../source_code/fixed-point/basic_op -I../source_code/fixed-point/lib_com -I../source_code/fixed-point/lib_enc -I../source_code/fixed-point/lib_dec
//#cgo LDFLAGS: -lm -L../source_code/fixed-point -levsEncoder -levsDecoder
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include "evs_encoder.h"
//#include "evs_decoder.h"
import "C"

type (
	EvsEncoderContext = C.struct_EvsEncoderContext
	EvsDecoderContext = C.struct_EvsDecoderContext
)

func newEvsEncoderContext() *EvsEncoderContext {
	return (*EvsEncoderContext)(C.NewEvsEncoder())
}

func (enc *EvsEncoderContext) initEvsEncoder(sample int, bitRate int, codec string, isG192Format int) int {
	p := C.CString(codec)
	defer C.free(unsafe.Pointer(p))
	res := int(C.InitEncoder(enc, (C.int)(sample), (C.int)(bitRate), p, (C.int)(isG192Format)))
	return res
}

func (enc *EvsEncoderContext) startEvsEncoder(payload []byte) []byte {
	length := len(payload)
	if length == 0 {
		return nil
	}

	if enc == nil {
		fmt.Printf("EvsEncoderContext enc nil error \n")
		return nil
	}

	C.EvsStartEncoder(enc, (*C.char)(unsafe.Pointer(&payload[0])), (C.int)(length))

	buffer := enc.buf
	if buffer.size > 0 {
		return C.GoBytes(unsafe.Pointer(&buffer.data[0]), buffer.size)
	} else {
		return nil
	}
}

func (enc *EvsEncoderContext) stopEvsEncoder() {
	C.StopEncoder(enc)
}

func UnitTestEncoder() {
	C.UnitTestEvsEncoder()
}

func newEvsDecoderContext() *EvsDecoderContext {
	return (*EvsDecoderContext)(C.NewEvsDecoder())
}

func (dec *EvsDecoderContext) initEvsDecoder(sample int, bitRate int, isG192Format int) int {
	res := int(C.InitDecoder(dec, (C.int)(sample), (C.int)(bitRate), (C.int)(isG192Format)))
	return res
}

func (dec *EvsDecoderContext) startEvsDecoder(payload []byte) []byte {
	length := len(payload)
	if length == 0 {
		return nil
	}

	C.EvsStartDecoder(dec, (*C.char)(unsafe.Pointer(&payload[0])))

	buffer := dec.buf
	if buffer.size > 0 {
		return C.GoBytes(unsafe.Pointer(&buffer.data[0]), C.int(buffer.size)*C.int(unsafe.Sizeof(buffer.data[0])))
	} else {
		return nil
	}
}

func (dec *EvsDecoderContext) stopEvsDecoder() {
	C.StopDecoder(dec)
}

func UnitTestDecoder() {
	C.UnitTestEvsDecoder()
}
