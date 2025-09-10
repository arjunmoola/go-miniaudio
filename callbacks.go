package gominiaudio

// #include "thirdparty/miniaudio.h"
import "C"

import (
	"unsafe"
)

type EnumerateDevicesCallback func(context *Context, deviceType DeviceType, info *DeviceInfo, userData unsafe.Pointer) bool
type DeviceDataProcCallback func(device *Device, output unsafe.Pointer, input unsafe.Pointer, frameCount uint32)
type DecoderReadProcCallback func(decoder *Decoder, bufferOut unsafe.Pointer, bytesToRead uint, bytesRead *uint) error
type DeviceDataProcCallbackF32[T SampleSize]func(device *Device, output []T, input []T, frameCount uint32)

//export goDevicesCallback
func goDevicesCallback(contextPtr unsafe.Pointer, deviceType C.ma_device_type, infoPtr unsafe.Pointer, userData unsafe.Pointer) C.ma_bool32 {
	c := contextFromPtr(contextPtr)
	info := deviceInfoFromPtr(infoPtr)
	dt := DeviceType(deviceType)
	res := enumerateDevicesCallback(c, dt, info, userData)
	return toMABool32(res) 
}

func deviceFromPtr(pDevice unsafe.Pointer) *Device {
	return &Device{
		device: (*C.struct_ma_device)(pDevice),
	}
}

//export goDataProcCallback
func goDataProcCallback(pDevice unsafe.Pointer, pOutput unsafe.Pointer, pInput unsafe.Pointer, frameCount C.ma_uint32) {
	device := deviceFromPtr(pDevice)
	dataCallback(device, pOutput,  pInput, uint32(frameCount))
}

//export goDataProcCallbackF32
func goDataProcCallbackF32(pDevice unsafe.Pointer, pOutput unsafe.Pointer, pInput unsafe.Pointer, frameCount C.ma_uint32) {
	device := deviceFromPtr(pDevice)
	chs := device.GetPlaybackChannels()
	outputBuf := newBuffer[Float32](pOutput, int(frameCount)*chs)
	//inputBuf := newBuffer[Float32](pInput, int(frameCount)*chs)
	dataCallbackF32(device, outputBuf,  nil, uint32(frameCount))
}

func decoderFromPtr(pDecoder unsafe.Pointer) *Decoder {
	return &Decoder{
		decoder: (*C.struct_ma_decoder)(pDecoder),
	}
}

//export goDecoderReadProcCallback
func goDecoderReadProcCallback(pDecoder unsafe.Pointer, pBufferOut unsafe.Pointer, bytesToRead C.size_t, pBytesRead *C.size_t) {
}
