package gominiaudio

// #include "thirdparty/miniaudio.h"
import "C"

import (
	"unsafe"
)

type EnumerateDevicesCallback func(context *Context, deviceType DeviceType, info *DeviceInfo, userData unsafe.Pointer) bool

//export goDevicesCallback
func goDevicesCallback(contextPtr unsafe.Pointer, deviceType C.ma_device_type, infoPtr unsafe.Pointer, userData unsafe.Pointer) C.ma_bool32 {
	c := contextFromPtr(contextPtr)
	info := deviceInfoFromPtr(infoPtr)
	dt := DeviceType(deviceType)
	res := enumerateDevicesCallback(c, dt, info, userData)
	return toMABool32(res) 
}
