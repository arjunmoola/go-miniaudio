package gominiaudio

// #cgo LDFLAGS: -ldl -lpthread -lm
// #define MINIAUDIO_IMPLEMENTATION
// #include "miniaudio.h"
// #include <stdlib.h>
// #include <stdio.h>
import "C"

import (
	"unsafe"
)

type MADeviceState int

const (
	MADeviceStateUninitialized MADeviceState = C.ma_device_state_uninitialized
	MADeviceStateStopped MADeviceState = C.ma_device_state_stopped
	MADeviceStateStarted MADeviceState = C.ma_device_state_started
	MADeviceStateStarting MADeviceState = C.ma_device_state_starting
	MADeviceStateStopping MADeviceState = C.ma_device_state_stopping
)

type MADeviceType int

const (
	MADeviceTypePlayback MADeviceType = C.ma_device_type_playback
	MADeviceTypeCapture MADeviceType = C.ma_device_type_capture
	MADeviceTypeDuplex MADeviceType = C.ma_device_type_duplex
	MADeviceTypeLoopback MADeviceType = C.ma_device_type_loopback
)

type MAShareMode int

const (
	//MAShareModeShared MAShareMode = C.ma_share_mode
	MAShareModeExclusive MAShareMode = C.ma_share_mode_exclusive
)

const (
	MA_SUCCESS = C.MA_SUCCESS
	MA_ERROR = C.MA_ERROR
	MA_INVALID_ARGS = C.MA_INVALID_ARGS
	MA_INVALID_OPERATION = C.MA_INVALID_OPERATION
)

func toMABool32(b bool) C.ma_bool32 {
	if b {
		return C.ma_bool32(1)
	}
	return C.ma_bool32(0)
}

func maBoolToGoBool(b C.ma_bool32) bool {
	switch b {
	case C.ma_bool32(1):
		return true
	case C.ma_bool32(0):
		return false
	default:
		panic("unkown ma_bool32 value provided")
	}
}

type MAEngineConfig struct {
    config C.ma_engine_config
}

func NewMAEngineConfig() *MAEngineConfig {
	config := C.ma_engine_config_init()
	return &MAEngineConfig{
		config: config,
	}
}

func (c *MAEngineConfig) GetListenerCount() uint32 {
	return uint32(c.config.listenerCount)
}

func (c *MAEngineConfig) SetListenerCount(count uint32) {
	c.config.listenerCount = C.ma_uint32(count)
}

func (c *MAEngineConfig) GetChannels() uint32 {
	return uint32(c.config.channels)
}

func (c *MAEngineConfig) SetChannels(channels uint32) {
	c.config.channels = C.ma_uint32(channels)
}

func (c *MAEngineConfig) GetPeriodSizeInFrames() uint32 {
	return uint32(c.config.periodSizeInFrames)
}

func (c *MAEngineConfig) SetPeriodSizeInFrames(periodSize uint32) {
	c.config.periodSizeInFrames = C.ma_uint32(periodSize)
}

func (c *MAEngineConfig) GetPeriodSizeInMilliseconds() uint32 {
	return uint32(c.config.periodSizeInMilliseconds)
}

func (c *MAEngineConfig) SetPeriodSizeInMilliseconds(periodSize uint32) {
	c.config.periodSizeInMilliseconds = C.ma_uint32(periodSize)
}

func (c *MAEngineConfig) GetGainSmoothTimeInFrames() uint32 {
	return uint32(c.config.gainSmoothTimeInMilliseconds)
}

func (c *MAEngineConfig) SetGainSmoothTimeInFrames(gainSmoothTime uint32) {
	c.config.periodSizeInFrames = C.ma_uint32(gainSmoothTime)
}

func (c *MAEngineConfig) GetGainSmoothTimeInMilliseconds() uint32 {
	return uint32(c.config.gainSmoothTimeInMilliseconds)
}

func (c *MAEngineConfig) SetGainSmoothTimeInMilliseconds(gainSmoothTime uint32) {
	c.config.gainSmoothTimeInMilliseconds = C.ma_uint32(gainSmoothTime)
}

func (c *MAEngineConfig) GetDefaultVolumeSmoothTimeInPCMFrames() uint32 {
	return uint32(c.config.defaultVolumeSmoothTimeInPCMFrames)
}

func (c *MAEngineConfig) SetDefaultVolumeSmoothTimeInPCMFrames(defaultVolume uint32) {
	c.config.defaultVolumeSmoothTimeInPCMFrames = C.ma_uint32(defaultVolume)
}

func (c *MAEngineConfig) GetPreMixStackSizeInBytes() uint32 {
	return uint32(c.config.preMixStackSizeInBytes)
}

func (c *MAEngineConfig) SetPreMixStackSizeInBytes(stackSize uint32) {
	c.config.preMixStackSizeInBytes = C.ma_uint32(stackSize)
}

func (c *MAEngineConfig) GetNoAutoStart() bool {
	return maBoolToGoBool(c.config.noAutoStart)
}

func (c *MAEngineConfig) SetNoAutoStart(b bool) {
	c.config.noAutoStart = toMABool32(b)
}

func (c *MAEngineConfig) GetNoDevice() bool {
	return maBoolToGoBool(c.config.noDevice)
}

func (c *MAEngineConfig) SetNoDevice(b bool) {
	c.config.noDevice = toMABool32(b)
}



type MAEngine struct {
	engine *C.struct_ma_engine
}

func (e *MAEngine) Close() {
	C.free(unsafe.Pointer(e.engine))
}

func NewMAEngine() *MAEngine {
	engine := (*C.struct_ma_engine) (C.malloc(C.sizeof_struct_ma_engine))

	return &MAEngine{
		engine: engine,
	}
}

type MADeviceConfig struct {
	config *C.struct_ma_device_config
}

func NewMADeviceConfig() *MADeviceConfig {
	config := (*C.struct_ma_device_config) (C.malloc(C.sizeof_struct_ma_device_config))

	return &MADeviceConfig {
		config: config,
	}
}

func (c *MADeviceConfig) Close() {
	C.free(unsafe.Pointer(c.config))
}

func MaVersionString() string {
	return C.GoString(C.ma_version_string())
}

