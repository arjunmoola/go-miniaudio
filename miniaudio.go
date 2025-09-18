package gominiaudio

// #include <stdlib.h>
// #include <stdio.h>
// #include "thirdparty/miniaudio.h"
//
//extern ma_bool32 goDevicesCallback(void *pContext, ma_device_type deviceType, void *pInfo, void *pUserdata);
//extern void goDataProcCallback(void *pDevice, void *pOutput, const void *pInput, ma_uint32 frameCount);
//extern void goDataProcCallbackF32(void *pDevice, void *pOutput, const void *pInput, ma_uint32 frameCount);
//
//ma_bool32 deviceEnumerationCallback(ma_context *pContext, ma_device_type deviceType, const ma_device_info* pInfo, void* pUserData) {
//  return goDevicesCallback((void*)pContext, deviceType, (void*)pInfo, pUserData);
//}
//
//ma_result go_ma_context_enumerate_devices(ma_context *pContext, void* pUserData) {
//	return ma_context_enumerate_devices(pContext, deviceEnumerationCallback, pUserData);
//}
//
//void maDataCallback(ma_device *pDevice, void *pOutput, const void *pInput, ma_uint32 frameCount) {
//	goDataProcCallback((void*)pDevice, pOutput, pInput, frameCount);
//}
//
//void maDataCallbackF32(ma_device *pDevice, void *pOutput, const void *pInput, ma_uint32 frameCount) {
//	goDataProcCallbackF32((void*)pDevice, pOutput, pInput, frameCount);
//}
//
//ma_result go_ma_device_config_set_data_callback(ma_device_config *pDeviceConfig) {
//	pDeviceConfig->dataCallback = maDataCallback;
//}
//
//ma_result go_ma_device_config_set_data_callbackF32(ma_device_config *pDeviceConfig) {
//	pDeviceConfig->dataCallback = maDataCallbackF32;
//}
import "C"

import (
	"unsafe"
	"fmt"
	//"errors"
)

type Float32 C.float
type Float64 C.double
type UInt8 C.ma_uint8
type UInt16 C.ma_uint16
type UInt32 C.ma_uint32
type UInt64 C.ma_uint64

type SampleSize interface {
	Float64|Float32|UInt8|UInt16|UInt32|UInt64
}

type Buffer struct {
	buffer unsafe.Pointer
}

func newBuffer(ptr unsafe.Pointer) *Buffer {
	return &Buffer{
		buffer: ptr,
	}
}

func (b *Buffer) cptr() unsafe.Pointer {
	if b == nil {
		return nil
	}
	return b.buffer
}

func Slice[T SampleSize](b *Buffer, n int) []T {
	ptr := b.cptr()
	return unsafe.Slice((*T)(ptr), n)
}

func bufferPointer[T SampleSize](buf []T) unsafe.Pointer {
	return unsafe.Pointer(&buf[0])
}

var enumerateDevicesCallback EnumerateDevicesCallback
var dataCallback DeviceDataProcCallback
var decaoderReadProcCallback DecoderReadProcCallback

type maContainer[T any] interface {
	cptr() T
}

func compare[T comparable](a, b maContainer[T]) bool {
	return a.cptr() == b.cptr()
}


type PCMFrameReader interface {
	ReadPCMFrames(*Buffer, int) (int, error)
}

type MAFormat uint32

type DeviceState int

const (
	DeviceStateUninitialized DeviceState = C.ma_device_state_uninitialized
	DeviceStateStopped DeviceState = C.ma_device_state_stopped
	DeviceStateStarted DeviceState = C.ma_device_state_started
	DeviceStateStarting DeviceState = C.ma_device_state_starting
	DeviceStateStopping DeviceState = C.ma_device_state_stopping
)

type DeviceType int

const (
	DeviceTypePlayback DeviceType = iota+1
	DeviceTypeCapture
	DeviceTypeDuplex
	DeviceTypeLoopback
)

func getMaDeviceType(d DeviceType) C.ma_device_type {
	var e C.ma_device_type
	switch d {
	case DeviceTypePlayback:
		e = C.ma_device_type_playback
	case DeviceTypeCapture:
		e = C.ma_device_type_capture
	case DeviceTypeDuplex:
		e = C.ma_device_type_duplex
	case DeviceTypeLoopback:
		e = C.ma_device_type_loopback 
	default:
		panic("invalid device type given")
	}	
	return e
}

type Format int

const (
	FormatUnknown Format = iota
	Format_u8
	Format_s16
	Format_s24
	Format_s32
	Format_f32
	Format_count
)

func getMaFormat(f Format) C.ma_format {
	var e C.ma_format
	switch f {
	case FormatUnknown:
		e = C.ma_format_unknown
	case Format_u8:
		e = C.ma_format_u8
	case Format_s16:
		e = C.ma_format_s16
	case Format_s24:
		e = C.ma_format_s24
	case Format_s32:
		e = C.ma_format_s32
	case Format_f32:
		e = C.ma_format_f32
	case Format_count:
		e = C.ma_format_count
	default:
		panic("invalid miniaudio format provided")
	}
	return e
}

type ShareMode int

const (
	//ShareModeShared ShareMode = C.ma_share_mode
	ShareModeExclusive ShareMode = C.ma_share_mode_exclusive
)

const (
	ma_SUCCESS =C.MA_SUCCESS
	ma_ERROR =C.MA_ERROR
	ma_INVALID_ARGS =C.MA_INVALID_ARGS
	ma_INVALID_OPERATION =C.MA_INVALID_OPERATION
	ma_OUT_OF_MEMORY =C.MA_OUT_OF_MEMORY
    ma_OUT_OF_RANGE =C.MA_OUT_OF_RANGE
    ma_ACCESS_DENIED =C.MA_ACCESS_DENIED
    ma_DOES_NOT_EXIST =C.MA_DOES_NOT_EXIST
    ma_ALREADY_EXISTS =C.MA_ALREADY_EXISTS
    ma_TOO_MANY_OPEN_FILES =C.MA_TOO_MANY_OPEN_FILES
    ma_INVALID_FILE =C.MA_INVALID_FILE
    ma_TOO_BIG =C.MA_TOO_BIG
    ma_PATH_TOO_LONG =C.MA_PATH_TOO_LONG
    ma_NAME_TOO_LONG =C.MA_NAME_TOO_LONG
    ma_NOT_DIRECTORY =C.MA_NOT_DIRECTORY
    ma_IS_DIRECTORY =C.MA_IS_DIRECTORY
    ma_DIRECTORY_NOT_EMPTY =C.MA_DIRECTORY_NOT_EMPTY
    ma_AT_END =C.MA_AT_END
    ma_NO_SPACE =C.MA_NO_SPACE
    ma_BUSY =C.MA_BUSY
    ma_IO_ERROR =C.MA_IO_ERROR
    ma_INTERRUPT =C.MA_INTERRUPT
    ma_UNAVAILABLE =C.MA_UNAVAILABLE
    ma_ALREADY_IN_USE =C.MA_ALREADY_IN_USE
    ma_BAD_ADDRESS =C.MA_BAD_ADDRESS
    ma_BAD_SEEK =C.MA_BAD_SEEK
    ma_BAD_PIPE =C.MA_BAD_PIPE
    ma_DEADLOCK =C.MA_DEADLOCK
    ma_TOO_MANY_LINKS =C.MA_TOO_MANY_LINKS
    ma_NOT_IMPLEMENTED =C.MA_NOT_IMPLEMENTED
    ma_NO_MESSAGE =C.MA_NO_MESSAGE
    ma_BAD_MESSAGE =C.MA_BAD_MESSAGE
    ma_NO_DATA_AVAILABLE =C.MA_NO_DATA_AVAILABLE
    ma_INVALID_DATA =C.MA_INVALID_DATA
    ma_TIMEOUT =C.MA_TIMEOUT
    ma_NO_NETWORK =C.MA_NO_NETWORK
    ma_NOT_UNIQUE =C.MA_NOT_UNIQUE
    ma_NOT_SOCKET =C.MA_NOT_SOCKET
    ma_NO_ADDRESS =C.MA_NO_ADDRESS
    ma_BAD_PROTOCOL =C.MA_BAD_PROTOCOL
    ma_PROTOCOL_UNAVAILABLE =C.MA_PROTOCOL_UNAVAILABLE
    ma_PROTOCOL_NOT_SUPPORTED =C.MA_PROTOCOL_NOT_SUPPORTED
    ma_PROTOCOL_FAMILY_NOT_SUPPORTED =C.MA_PROTOCOL_FAMILY_NOT_SUPPORTED
    ma_ADDRESS_FAMILY_NOT_SUPPORTED =C.MA_ADDRESS_FAMILY_NOT_SUPPORTED
    ma_SOCKET_NOT_SUPPORTED =C.MA_SOCKET_NOT_SUPPORTED
    ma_CONNECTION_RESET =C.MA_CONNECTION_RESET
    ma_ALREADY_CONNECTED =C.MA_ALREADY_CONNECTED
    ma_NOT_CONNECTED =C.MA_NOT_CONNECTED
    ma_CONNECTION_REFUSED =C.MA_CONNECTION_REFUSED
    ma_NO_HOST =C.MA_NO_HOST
    ma_IN_PROGRESS =C.MA_IN_PROGRESS
    ma_CANCELLED =C.MA_CANCELLED
    ma_MEMORY_ALREADY_MAPPED =C.MA_MEMORY_ALREADY_MAPPED

    //General non-standard errors. 
    ma_CRC_MISMATCH =C.MA_CRC_MISMATCH

    //General miniaudio-specific errors. 
    ma_FORMAT_NOT_SUPPORTED =C.MA_FORMAT_NOT_SUPPORTED
    ma_DEVICE_TYPE_NOT_SUPPORTED =C.MA_DEVICE_TYPE_NOT_SUPPORTED
    ma_SHARE_MODE_NOT_SUPPORTED =C.MA_SHARE_MODE_NOT_SUPPORTED
    ma_NO_BACKEND =C.MA_NO_BACKEND
    ma_NO_DEVICE =C.MA_NO_DEVICE
    ma_API_NOT_FOUND =C.MA_API_NOT_FOUND
    ma_INVALID_DEVICE_CONFIG =C.MA_INVALID_DEVICE_CONFIG
    ma_LOOP =C.MA_LOOP
    ma_BACKEND_NOT_ENABLED =C.MA_BACKEND_NOT_ENABLED

    //State errors. 
    ma_DEVICE_NOT_INITIALIZED =C.MA_DEVICE_NOT_INITIALIZED
    ma_DEVICE_ALREADY_INITIALIZED =C.MA_DEVICE_ALREADY_INITIALIZED
    ma_DEVICE_NOT_STARTED =C.MA_DEVICE_NOT_STARTED
    ma_DEVICE_NOT_STOPPED =C.MA_DEVICE_NOT_STOPPED

    //Operation errors. 
    ma_FAILED_TO_INIT_BACKEND =C.MA_FAILED_TO_INIT_BACKEND
    ma_FAILED_TO_OPEN_BACKEND_DEVICE =C.MA_FAILED_TO_OPEN_BACKEND_DEVICE
    ma_FAILED_TO_START_BACKEND_DEVICE =C.MA_FAILED_TO_START_BACKEND_DEVICE
    ma_FAILED_TO_STOP_BACKEND_DEVICE =C.MA_FAILED_TO_STOP_BACKEND_DEVICE
)

type StreamFormat int

const (
	StreamFormatPCM StreamFormat = iota
)

func getMaStreamFormat(f StreamFormat) C.ma_stream_format {
	var e C.ma_stream_format
	switch f {
	case StreamFormatPCM:
		e = C.ma_stream_format_pcm
	}
	return e
}

type StreamLayout int

const (
	StreamLayoutInterleaved StreamLayout = iota
	StreamLayoutDeinterleaved
)

func getMaStreamLayout(l StreamLayout) C.ma_stream_layout {
	return C.ma_stream_layout(l)
}

type DitherMode int

const (
	DitherModeNon DitherMode = iota
	DitherModeRectangle
	DitherModeTriange
)

func getMaDitherMode(d DitherMode) C.ma_dither_mode {
	return C.ma_dither_mode(d)
}

type StandardSampleRate int

const (
	StandardSampleRate48000 StandardSampleRate = 48000
	StandardSampleRate44100 StandardSampleRate = 44100
	StandardSampleRate32000 StandardSampleRate = 32000
	StandardSampleRate24000 StandardSampleRate = 24000
	StandardSampleRate22050 StandardSampleRate = 22050
	StandardSampleRate88200 StandardSampleRate = 88200
	StandardSampleRate96000 StandardSampleRate = 96000
	StandardSampleRate176400 StandardSampleRate = 176400
	StandardSampleRate192000 StandardSampleRate = 192000
	StandardSampleRate16000 StandardSampleRate = 16000
	StandardSampleRate11025 StandardSampleRate = 11025
	StandardSampleRate8000 StandardSampleRate = 8000
	StandardSampleRate352800 StandardSampleRate = 352800
	StandardSampleRate384000 StandardSampleRate = 384000
	StandardSampleRateMin StandardSampleRate = StandardSampleRate8000
	StandardSampleRateMax StandardSampleRate = StandardSampleRate384000
	StandardSampleRateCount StandardSampleRate = 14
)

func getMaStandardSapleRate(r StandardSampleRate) C.ma_standard_sample_rate {
	return C.ma_standard_sample_rate(r)
}

type ChannelMixMode int

const (
	ChannelMixModeRectangular ChannelMixMode = iota
	ChannelMixModeSimple
	ChannelMixModeCustomHeights
	ChannelMixModeDefault
)

func getMaChannelMixMode(m ChannelMixMode) C.ma_channel_mix_mode {
	return C.ma_channel_mix_mode(m)
}

type StandardChannelMap int

const (
	MICROSOFT StandardChannelMap = iota
	ALSA
	RFC3351
	FLAC
	VORBIS
	SOUND4
	SNDIO
	WEBAUDIO
	DEFAULT
)

func getMaStandardChannelMap(m StandardChannelMap) C.ma_standard_channel_map {
	return C.ma_standard_channel_map(m)
}

type PerformanceProfile int

const (
	LOW_LATENCY PerformanceProfile = iota
	PROFILE_CONSERVATIVE
)

func getMaPerformanceProfile(pp PerformanceProfile) C.ma_performance_profile {
	return C.ma_performance_profile(pp)
}

type EncodingFormat int

const (
	UNKNOWN_ENC EncodingFormat = iota
	WAV_ENC
	FLAC_ENC
	MP3_ENC
	VORBIS_ENC
)

type MABackend int

const (
	WASAPI MABackend = C.ma_backend_wasapi
	DSOUND MABackend = C.ma_backend_dsound
	WINMM MABackend = C.ma_backend_winmm
	COREAUDIO MABackend = C.ma_backend_coreaudio
	SNDIO_B MABackend = C.ma_backend_sndio
	AUDIO4 MABackend = C.ma_backend_audio4
	OSS MABackend = C.ma_backend_oss
	PULSEAUDIO MABackend = C.ma_backend_pulseaudio
	JACK MABackend = C.ma_backend_jack
	AAUDIO MABackend = C.ma_backend_aaudio
	OPENSL MABackend = C.ma_backend_opensl
	WEBAUDIO_B MABackend = C.ma_backend_webaudio
	Custom_B MABackend = C.ma_backend_custom
	Backend_NULL MABackend = C.ma_backend_custom
)

func (b MABackend) String() string {
	return GetBackendName(b)
}

func (b MABackend) IsEnabled() bool {
	return IsBackendEnabled(b)
}

func (b MABackend) IsLoopbackSupported() bool {
	return IsLoopbackSupported(b)
}

func GetBackendFromName(backendName string) (MABackend, error) {
	cstr := C.CString(backendName)
	defer C.free(unsafe.Pointer(cstr))
	var backend C.ma_backend
	
	res := C.ma_get_backend_from_name(cstr, &backend)

	if err := checkResult(res); err != nil {
		return MABackend(backend), err
	}
	return MABackend(backend), nil
}

func GetEnabledBackends() ([]MABackend, error) {
	var backends []MABackend
	buf := make([]C.ma_backend, C.MA_BACKEND_COUNT)
	var enabledBackends C.size_t

	res := C.ma_get_enabled_backends(&buf[0], C.MA_BACKEND_COUNT, &enabledBackends)

	if err := checkResult(res); err != nil {
		return backends, err
	}

	backends = make([]MABackend, 0, int(enabledBackends))

	for i := range int(enabledBackends) {
		backends = append(backends, MABackend(buf[i]))
	}


	return backends, nil
}

func IsBackendEnabled(backend MABackend) bool {
	res := C.ma_is_backend_enabled(C.ma_backend(backend))
	return maBoolToGoBool(res)
}

func IsLoopbackSupported(backend MABackend) bool {
	res := C.ma_is_loopback_supported(C.ma_backend(backend))
	return maBoolToGoBool(res)
}

func GetBackendName(backEnd MABackend) string {
	charPtr := C.ma_get_backend_name(C.ma_backend(backEnd))
	return C.GoString(charPtr)
}

type Channel uint8

type Event struct {
	event *C.ma_event
}

func NewEvent() *Event {
	event := (*C.ma_event)(C.malloc(C.sizeof_ma_event))
	return &Event{
		event: event,
	}
}

func (e *Event) cptr() *C.ma_event {
	if e == nil {
		return nil
	}
	return e.event
}

func (e *Event) Close() {
	C.free(unsafe.Pointer(e.cptr()))
}

func (e *Event) Init() error {
	res := C.ma_event_init(e.cptr())
	return checkResult(res)
}

func (e *Event) Uninit() {
	C.ma_event_uninit(e.cptr())
}

func (e *Event) Wait() error {
	res := C.ma_event_wait(e.cptr())
	return checkResult(res)
}

func (e *Event) Signal() error {
	res := C.ma_event_signal(e.cptr())
	return checkResult(res)
}

//Utilities

func CalculateBufferSizeInMillisecondsFromFrames(bufferSizeInFrames uint32, sampleRate uint32) uint32 {
	res := C.ma_calculate_buffer_size_in_milliseconds_from_frames(C.ma_uint32(bufferSizeInFrames), C.ma_uint32(sampleRate))
	return uint32(res)
}

func CalculateBufferSizeInFramesFromMilliseconds(bufferSize uint32, sampleRate uint32) uint32 {
	res := C.ma_calculate_buffer_size_in_frames_from_milliseconds(C.ma_uint32(bufferSize), C.ma_uint32(sampleRate))
	return uint32(res)
}

func CopyPCMFrames(dst, src *Buffer, frameCount int, format Format, channels int) {
	C.ma_copy_pcm_frames(dst.cptr(), src.cptr(), C.ma_uint64(frameCount), C.ma_format(format), C.ma_uint32(channels))
}

func SilencePCMFrames(dst *Buffer, frameCount int, format Format, channels int) {
	C.ma_silence_pcm_frames(dst.cptr(), C.ma_uint64(frameCount), C.ma_format(format), C.ma_uint32(channels))
}

func ClipSamplesF32(dst []Float32, src []Float32, count int) {
	pDst := (*C.float)(bufferPointer(dst))
	pSrc := (*C.float)(bufferPointer(src))
	C.ma_clip_samples_f32(pDst, pSrc, C.ma_uint64(count))
}

func ClipPCMFrames(dst *Buffer, src *Buffer, frameCount int, format Format, channels int) {
	pDst := dst.cptr()
	pSrc := src.cptr()
	C.ma_clip_pcm_frames(pDst, pSrc, C.ma_uint64(frameCount), C.ma_format(format), C.ma_uint32(channels))
}

func ApplyVolumeFactorF32(dst []Float32, sampleCount int, factor float32) {
	pDst := (*C.float)(bufferPointer(dst))
	C.ma_apply_volume_factor_f32(pDst, C.ma_uint64(sampleCount), C.float(factor))
}

func ApplyVolumeFactorPCMFrames(frames *Buffer, frameCount int, format Format,  channels int, factor float32) {
	pFrames := frames.cptr()
	C.ma_apply_volume_factor_pcm_frames(pFrames, C.ma_uint64(frameCount), C.ma_format(format), C.ma_uint32(channels), C.float(factor))
}

func MixPCMFramesF32(dst []Float32, src []Float32, frameCount int, channels int, volume float32) error {
	pDst := (*C.float)(bufferPointer(dst))
	pSrc := (*C.float)(bufferPointer(src))
	res := C.ma_mix_pcm_frames_f32(pDst, pSrc, C.ma_uint64(frameCount), C.ma_uint32(channels), C.float(volume))
	return checkResult(res)
}

func CopyApplyVolumeClipPCMFrames(dst *Buffer, src *Buffer, frameCount int, format Format, channels int, volume float32) {
	C.ma_copy_and_apply_volume_and_clip_pcm_frames(dst.cptr(), src.cptr(), C.ma_uint64(frameCount), C.ma_format(format), C.ma_uint32(channels), C.float(volume))
}

func VolumeLinearToDb(factor float32) float32 {
	res := C.ma_volume_linear_to_db(C.float(factor))
	return float32(res)
}

func VolumeDbToLinear(gain float32) float32 {
	res := C.ma_volume_db_to_linear(C.float(gain))
	return float32(res)
}

type MiniAudioErr struct {
	res C.ma_result
}

func (e *MiniAudioErr) Code() int {
	return int(e.res)
}

func (e *MiniAudioErr) Error() string {
	s := maResultString(e.res)
	d := int(e.res)
	return fmt.Sprintf("miniaudio: %s code: %d", s, d)
}

func newMiniAudioErr(res C.ma_result) *MiniAudioErr {
	return &MiniAudioErr{
		res: res,
	}
}

func maResultString(res C.ma_result) string {
	return C.GoString(C.ma_result_description(res))
}

func checkResult(res C.ma_result) error {
	switch res {
	case ma_SUCCESS:
		return nil
	default:
		return newMiniAudioErr(res)
	}
}

func toMABool32(b bool) C.ma_bool32 {
	if b {
		return C.ma_bool32(1)
	}
	return C.ma_bool32(0)
}

func toMABool8(b bool) C.ma_bool8 {
	if b {
		return C.ma_bool8(1)
	}
	return C.ma_bool8(0)
}

func maBool8ToGoBool(b C.ma_bool8) bool {
	switch b {
	case C.ma_bool8(1):
		return true
	case C.ma_bool8(0):
		return false
	default:
		panic("unknown value for ma_bool8")
	}
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

type BiQuadConfig struct {
	config C.ma_biquad_config
}

func BiQuadConfigInit(format Format, channels int, b0, b1, b2, a0, a1, a2 float64) *BiQuadConfig {
	f := getMaFormat(format)
	config := C.ma_biquad_config_init(f, C.ma_uint32(channels), C.double(b0), C.double(b1), C.double(b2), C.double(a0), C.double(a1), C.double(a2))
	return &BiQuadConfig{
		config: config,
	}
}

type BiQuad struct {
	biquad *C.ma_biquad
}

type EngineConfig struct {
    config C.ma_engine_config
}

func (c *EngineConfig) cptr() *C.ma_engine_config {
	if c == nil {
		return nil
	}
	return &c.config
}

func NewEngineConfig() *EngineConfig {
	config := C.ma_engine_config_init()
	return &EngineConfig{
		config: config,
	}
}

func (c *EngineConfig) GetListenerCount() uint32 {
	return uint32(c.config.listenerCount)
}

func (c *EngineConfig) SetListenerCount(count uint32) {
	c.config.listenerCount = C.ma_uint32(count)
}

func (c *EngineConfig) GetChannels() uint32 {
	return uint32(c.config.channels)
}

func (c *EngineConfig) SetChannels(channels uint32) {
	c.config.channels = C.ma_uint32(channels)
}

func (c *EngineConfig) SetSampleRate(sampleRate uint32) {
	c.config.sampleRate = C.ma_uint32(sampleRate)
}

func (c *EngineConfig) GetPeriodSizeInFrames() uint32 {
	return uint32(c.config.periodSizeInFrames)
}

func (c *EngineConfig) SetPeriodSizeInFrames(periodSize uint32) {
	c.config.periodSizeInFrames = C.ma_uint32(periodSize)
}

func (c *EngineConfig) GetPeriodSizeInMilliseconds() uint32 {
	return uint32(c.config.periodSizeInMilliseconds)
}

func (c *EngineConfig) SetPeriodSizeInMilliseconds(periodSize uint32) {
	c.config.periodSizeInMilliseconds = C.ma_uint32(periodSize)
}

func (c *EngineConfig) GetGainSmoothTimeInFrames() uint32 {
	return uint32(c.config.gainSmoothTimeInMilliseconds)
}

func (c *EngineConfig) SetGainSmoothTimeInFrames(gainSmoothTime uint32) {
	c.config.periodSizeInFrames = C.ma_uint32(gainSmoothTime)
}

func (c *EngineConfig) GetGainSmoothTimeInMilliseconds() uint32 {
	return uint32(c.config.gainSmoothTimeInMilliseconds)
}

func (c *EngineConfig) SetGainSmoothTimeInMilliseconds(gainSmoothTime uint32) {
	c.config.gainSmoothTimeInMilliseconds = C.ma_uint32(gainSmoothTime)
}

func (c *EngineConfig) GetDefaultVolumeSmoothTimeInPCMFrames() uint32 {
	return uint32(c.config.defaultVolumeSmoothTimeInPCMFrames)
}

func (c *EngineConfig) SetDefaultVolumeSmoothTimeInPCMFrames(defaultVolume uint32) {
	c.config.defaultVolumeSmoothTimeInPCMFrames = C.ma_uint32(defaultVolume)
}

func (c *EngineConfig) GetPreMixStackSizeInBytes() uint32 {
	return uint32(c.config.preMixStackSizeInBytes)
}

func (c *EngineConfig) SetPreMixStackSizeInBytes(stackSize uint32) {
	c.config.preMixStackSizeInBytes = C.ma_uint32(stackSize)
}

func (c *EngineConfig) GetNoAutoStart() bool {
	return maBoolToGoBool(c.config.noAutoStart)
}

//Setting noAutoStart to true will make sure that the engine does not automatically start
//after initialization
func (c *EngineConfig) SetNoAutoStart(b bool) {
	c.config.noAutoStart = toMABool32(b)
}

func (c *EngineConfig) GetNoDevice() bool {
	return maBoolToGoBool(c.config.noDevice)
}

func (c *EngineConfig) SetNoDevice(b bool) {
	c.config.noDevice = toMABool32(b)
}

type SoundConfig struct {
	config C.ma_sound_config
}

func (c *SoundConfig) cptr() *C.ma_sound_config {
	if c == nil {
		return nil
	}
	return &c.config
}

func NewSoundConfig(engine *Engine) *SoundConfig {
	config := C.ma_sound_config_init_2(engine.engine)
	return &SoundConfig{
		config: config,
	}
}

type Sound struct {
	sound C.struct_ma_sound
}

func (s *Sound) cptr() *C.struct_ma_sound {
	if s == nil {
		return nil
	}
	return &s.sound
}

func (s *Sound) InitFromFile(engine *Engine, filePath string, flags uint32, soundGroup *SoundGroup, fence *Fence) error {
	return soundInitFromFile(engine, filePath, flags, soundGroup, fence, s)
}

func (s *Sound) InitCopy(engine *Engine, existingSound *Sound, flags uint32, soundGroup *SoundGroup) error {
	return soundInitCopy(engine, existingSound, flags, soundGroup, s)
}

func SoundInitFromFile(engine *Engine, filePath string, flags uint32, soundGroup *SoundGroup, fence *Fence, sound *Sound) error {
	return soundInitFromFile(engine, filePath, flags, soundGroup, fence, sound)
}

func soundInitFromFile(engine *Engine, filePath string, flags uint32, soundGroup *SoundGroup, fence *Fence, sound *Sound) error {
	enginePtr := engine.cptr()
	sGroupPtr := soundGroup.cptr()
	fencePtr := fence.cptr()
	soundPtr := sound.cptr()
	path := C.CString(filePath)
	defer C.free(unsafe.Pointer(path))
	maFlags := C.ma_uint32(flags)
	res := C.ma_sound_init_from_file(enginePtr, path, maFlags, sGroupPtr, fencePtr, soundPtr)
	return newMiniAudioErr(res)
}

func SoundInitCopy(engine *Engine, existingSound *Sound, flags uint32, soundGroup *SoundGroup, sound *Sound) error {
	return soundInitCopy(engine, existingSound, flags, soundGroup, sound)
}

func soundInitCopy(engine *Engine, existingSound *Sound, flags uint32, soundGroup *SoundGroup, sound *Sound) error {
	enginePtr := engine.cptr()
	eSoundPtr := existingSound.cptr()
	groupPtr := soundGroup.cptr()
	soundPtr := sound.cptr()
	maFlags := C.ma_uint32(flags)
	res := C.ma_sound_init_copy(enginePtr, eSoundPtr, maFlags, groupPtr, soundPtr)
	return newMiniAudioErr(res)
}

type SoundGroup struct {
	soundGroup C.ma_sound_group
}

func (m *SoundGroup) cptr() *C.ma_sound_group {
	if m == nil {
		return nil
	}
	return &m.soundGroup
}

type Engine struct {
	engine *C.struct_ma_engine
}

func (m *Engine) cptr() *C.struct_ma_engine {
	if m == nil {
		return nil
	}
	return m.engine
}

func NewEngine() *Engine {
	engine := (*C.struct_ma_engine)(C.malloc(C.sizeof_struct_ma_engine))
	return &Engine{
		engine: engine,
	}
}

func (e *Engine) Close() {
	C.free(unsafe.Pointer(e.cptr()))
}

func (e *Engine) Init(config *EngineConfig) error {
	return engineInit(config, e)
}

func EngineInit(config *EngineConfig, engine *Engine) error {
	return engineInit(config, engine)
}

func engineInit(config *EngineConfig, engine *Engine) error {
	configPtr := config.cptr()
	enginePtr := engine.cptr()
	res := C.ma_engine_init(configPtr, enginePtr)
	return checkResult(res)
}

func (e *Engine) Uninit() {
	engineUninit(e)
}

func EngineUninit(engine *Engine) {
	engineUninit(engine)
}

func engineUninit(engine *Engine) {
	C.ma_engine_uninit(engine.cptr())
}

func (e *Engine) PlaySound(filePath string, soundGroup *SoundGroup) error {
	return enginePlaySound(e, filePath, soundGroup)
}

func EnginePlaySound(engine *Engine, filePath string, soundGroup *SoundGroup) error {
	return enginePlaySound(engine, filePath, soundGroup)
}

func enginePlaySound(engine *Engine, filePath string, soundGroup *SoundGroup) error {
	enginePtr := engine.cptr()
	soundGroupPtr := soundGroup.cptr()
	path := C.CString(filePath)
	defer C.free(unsafe.Pointer(path))
	res := C.ma_engine_play_sound(enginePtr, path, soundGroupPtr)
	return checkResult(res)
}

func (e *Engine) GetResourceManager() *ResourceManager {
	return engineGetResourceManager(e)
}

func EngineGetResourceManager(engine *Engine) *ResourceManager {
	return engineGetResourceManager(engine)
}

func engineGetResourceManager(engine *Engine) *ResourceManager {
	if engine == nil {
		panic("expected engine to be non nil")
	}
	enginePtr := engine.cptr()
	managerPtr := C.ma_engine_get_resource_manager(enginePtr)
	return &ResourceManager{
		manager: managerPtr,
	}
}

func EngineGetDevice(engine *Engine) {}

type DataConverterConfig struct {
	config C.ma_data_converter_config
}

func DataConverterConfigInitDefault() *DataConverterConfig {
	config := C.ma_data_converter_config_init_default()
	return &DataConverterConfig{
		config: config,
	}
}

func DataConverterConfigInit(formatIn Format, formatOut Format, channelsIn int, channelsOut int, sampleRateIn int, sampleRateOut int) *DataConverterConfig {
	config := C.ma_data_converter_config_init(C.ma_format(formatIn), C.ma_format(formatOut), C.ma_uint32(channelsIn), C.ma_uint32(channelsOut), C.ma_uint32(sampleRateIn), C.ma_uint32(sampleRateOut))
	return &DataConverterConfig{
		config: config,
	}
}

type DataConverter struct {
	c *C.ma_data_converter
}

type ConvertFramesArg struct {
	Ptr unsafe.Pointer
	FramesCount int
	Format Format
	Channels int
	SampleRate int
}

func ConvertFrames(out ConvertFramesArg, in ConvertFramesArg) int {
	res := C.ma_convert_frames(out.Ptr, C.ma_uint64(out.FramesCount), C.ma_format(out.Format), C.ma_uint32(out.Channels), C.ma_uint32(out.SampleRate), in.Ptr, C.ma_uint64(in.FramesCount), C.ma_format(in.Format), C.ma_uint32(in.Channels), C.ma_uint32(in.SampleRate))

	return int(res)
}

type DataSource struct {
	src unsafe.Pointer
}

func (d *DataSource) cptr() unsafe.Pointer {
	if d == nil {
		return nil
	}
	return d.src
}

func (d *DataSource) ReadPCMFrames(framesOut *Buffer, frameCount int) (int, error) {
	var framesRead C.ma_uint64

	res := C.ma_data_source_read_pcm_frames(d.cptr(), framesOut.cptr(), C.ma_uint64(frameCount), &framesRead)

	if err := checkResult(res); err != nil {
		return 0, err
	}

	return int(framesRead), nil
}

func DataSourceReadPCMFrames(d any, framesOut *Buffer, frameCount int) (int, error) {
	var framesRead C.ma_uint64
	ptr := getDataSourcePtr(d)

	res := C.ma_data_source_read_pcm_frames(ptr, framesOut.cptr(), C.ma_uint64(frameCount), &framesRead)

	if err := checkResult(res); err != nil {
		return 0, err
	}

	return int(framesRead), nil
}

func (d *DataSource) ReadPCMFramesF32(framesOut []Float32, frameCount int) (int, error) {
	var framesRead C.ma_uint64

	res := C.ma_data_source_read_pcm_frames(d.cptr(), bufferPointer(framesOut), C.ma_uint64(frameCount), &framesRead)

	if err := checkResult(res); err != nil {
		return 0, err
	}

	return int(framesRead), nil
}

func (d *DataSource) SeekPCMFrames(frameCount int) (int, error) {
	var framesSeeked C.ma_uint64

	res := C.ma_data_source_seek_pcm_frames(d.cptr(), C.ma_uint64(frameCount), &framesSeeked)

	if err := checkResult(res); err != nil {
		return 0, err
	}

	return int(framesSeeked), nil
}

func (d *DataSource) SeekToPCMFrame(frameIndex int) error {
	res := C.ma_data_source_seek_to_pcm_frame(d.cptr(), C.ma_uint64(frameIndex))
	return checkResult(res)
}

func (d *DataSource) SeekSeconds(secondCount float32) (float32, error) {
	var secondsSeeked C.float

	res := C.ma_data_source_seek_seconds(d.cptr(), C.float(secondCount), &secondsSeeked)

	if err := checkResult(res); err != nil {
		return float32(secondsSeeked), err
	}

	return float32(secondsSeeked), nil
}

func (d *DataSource) SeekToSecond(seekPointInSeconds float32) error {
	res := C.ma_data_source_seek_to_second(d.cptr(), C.float(seekPointInSeconds))
	return checkResult(res)
}

func (d *DataSource) GetDataFormat(src *DataSource, channelMapCap int) (DataFormat, error) {
	var dataFormat DataFormat

	var format C.ma_format
	var channels C.ma_uint32
	var sampleRate C.ma_uint32
	var channelMap C.ma_channel

	res := C.ma_data_source_get_data_format(d.cptr(), &format, &channels, &sampleRate, &channelMap, C.size_t(channelMapCap))

	if err := checkResult(res); err != nil {
		return dataFormat, err
	}

	dataFormat.Format = Format(format)
	dataFormat.Channels = int(channels)
	dataFormat.SampleRate = int(sampleRate)
	dataFormat.ChannelMapCap = channelMapCap


	return dataFormat, nil
}

func (d *DataSource) GetCursorInSeconds() (float32, error) {
	var pCursor C.float

	res := C.ma_data_source_get_cursor_in_seconds(d.cptr(), &pCursor)

	if err := checkResult(res); err != nil {
		return float32(pCursor), err
	}

	return float32(pCursor), nil
}

func (d *DataSource) GetLengthInSeconds() (float32, error) {
	var pLength C.float

	res := C.ma_data_source_get_length_in_seconds(d.cptr(), &pLength)

	if err := checkResult(res); err != nil {
		return float32(pLength), err
	}

	return float32(pLength), nil
}

func (d *DataSource) GetCursorInPCMFrames() (int, error) {
	var pCursor C.ma_uint64

	res := C.ma_data_source_get_cursor_in_pcm_frames(d.cptr(), &pCursor)
	
	if err := checkResult(res); err != nil {
		return int(pCursor), nil
	}
	return int(pCursor), nil
}

func (d *DataSource) GetLengthInPCMFrames() (int, error) {
	var pLength C.ma_uint64

	res := C.ma_data_source_get_length_in_pcm_frames(d.cptr(), &pLength)

	if err := checkResult(res); err != nil {
		return int(pLength), err
	}

	return int(pLength), nil
}

func (d *DataSource) SetLooping(isLooping bool) error {
	res := C.ma_data_source_set_looping(d.cptr(), toMABool32(isLooping))
	return checkResult(res)
}

func (d *DataSource) IsLooping() bool {
	res := C.ma_data_source_is_looping(d.cptr())
	return maBoolToGoBool(res)
}

func (d *DataSource) SetRangeInPCMFrames(begInFrames, endInFrames int) error {
	res := C.ma_data_source_set_range_in_pcm_frames(d.cptr(), C.ma_uint64(begInFrames), C.ma_uint64(endInFrames))
	return checkResult(res)
}

type PCMFrameRange struct {
	Start, End int
}

func (d *DataSource) GetRangeInPCMFrameRange() PCMFrameRange {
	var rng PCMFrameRange

	var start, end C.ma_uint64

	C.ma_data_source_get_range_in_pcm_frames(d.cptr(), &start, &end)

	rng.Start = int(start)
	rng.End = int(end)

	return rng
}

func (d *DataSource) SetLoopPointInPCMFrames(start, end int) error {
	res := C.ma_data_source_set_loop_point_in_pcm_frames(d.cptr(), C.ma_uint64(start), C.ma_uint64(end))
	return checkResult(res)
}

type PCMFrameLoopPoint struct {
	Start, End int
}

func (d *DataSource) GetLoopPointInPCMFrames() PCMFrameLoopPoint {
	var lp PCMFrameLoopPoint
	var start, end C.ma_uint64
	C.ma_data_source_get_loop_point_in_pcm_frames(d.cptr(), &start, &end)

	lp.Start = int(start)
	lp.End = int(end)

	return lp
}

func (d *DataSource) SetCurrent(current *DataSource) error {
	res := C.ma_data_source_set_current(d.cptr(), current.cptr())
	return checkResult(res)
}

func (d *DataSource) GetCurrent() *DataSource {
	ptr := C.ma_data_source_get_current(d.cptr())
	return &DataSource{
		src: ptr,
	}
}

func (d *DataSource) SetNext(nextSource *DataSource) error {
	res := C.ma_data_source_set_next(d.cptr(), nextSource.cptr())
	return checkResult(res)
}

func DataSourceSetNext(current, next any) error {
	res := C.ma_data_source_set_next(getDataSourcePtr(current), getDataSourcePtr(next))
	return checkResult(res)
}

func getDataSourcePtr(d any) unsafe.Pointer {
	var ptr unsafe.Pointer
	switch ds :=  d.(type) {
	case *Decoder:
		ptr = unsafe.Pointer(ds.cptr())
	case *Waveform:
		ptr = unsafe.Pointer(ds.cptr())
	}
	return ptr
}

func (d *DataSource) GetNext() *DataSource {
	ptr := C.ma_data_source_get_next(d.cptr())
	return &DataSource{
		src: ptr,
	}
}

type DeviceInfo struct {
	info *C.ma_device_info
}

func (d *DeviceInfo) cptr() *C.ma_device_info {
	if d == nil {
		return nil
	}
	return d.info
}

func (d *DeviceInfo) GetDeviceId() string {
	return ""
}

func (d *DeviceInfo) Name() string {
	return C.GoString(&d.info.name[0])
}

func (d *DeviceInfo) IsDefault() bool {
	return maBoolToGoBool(d.info.isDefault)
}

type DeviceDescriptor struct {
	des *C.ma_device_descriptor
}

func (d *DeviceDescriptor) cptr() *C.ma_device_descriptor {
	if d == nil {
		return nil
	}
	return d.des
}

type DeviceConfig struct {
	config C.struct_ma_device_config
}


func (c *DeviceConfig) cptr() *C.struct_ma_device_config {
	if c == nil {
		return nil
	}
	return &c.config
}

func DeviceConfigInit(deviceType DeviceType) *DeviceConfig {
	maDeviceType := getMaDeviceType(deviceType)
	config := C.ma_device_config_init(maDeviceType)
	return &DeviceConfig{
		config: config,
	}
}

func (d *DeviceConfig) SetPlaybackFormat(f Format) {
	d.config.playback.format = getMaFormat(f)
}

func (d *DeviceConfig) GetPlaybackFormat() Format {
	return Format(d.config.playback.format)
}

func (d *DeviceConfig) SetPlaybackChannels(channels int) {
	d.config.playback.channels = C.ma_uint32(channels)
}

func (d *DeviceConfig) GetPlaybackChannels() int {
	return int(d.config.playback.channels)
}

func (d *DeviceConfig) SetSampleRate(sampleRate int) {
	d.config.sampleRate = C.ma_uint32(sampleRate)
}

func (d *DeviceConfig) GetSampleRate() int {
	return int(d.config.sampleRate)
}

func (d *DeviceConfig) GetPeriods() int {
	return int(d.config.periods)
}

func (d *DeviceConfig) SetPeriods(periods int) {
	d.config.periods = C.ma_uint32(periods)
}

func (d *DeviceConfig) GetPerformanceProfile() PerformanceProfile {
	return PerformanceProfile(d.config.performanceProfile)
}

func (d *DeviceConfig) SetPerformanceProfile(pp PerformanceProfile) {
	d.config.performanceProfile = getMaPerformanceProfile(pp)
}

func (d *DeviceConfig) SetNoPreSilencedOutputBuffer(b bool) {
	d.config.noPreSilencedOutputBuffer = toMABool8(b)
}

func (d *DeviceConfig) GetNoPreSilencedOutputBuffer() bool {
	return maBool8ToGoBool(d.config.noPreSilencedOutputBuffer)
}

func (d *DeviceConfig) SetNoClip(b bool) {
	d.config.noClip = toMABool8(b)
}

func (d *DeviceConfig) GetNoClip() bool {
	return maBool8ToGoBool(d.config.noClip)
}

func (d *DeviceConfig) GetNoDisabledDenormals() bool {
	return maBool8ToGoBool(d.config.noClip)
}

func (d *DeviceConfig) SetNoDisabledDenormals(b bool) {
	d.config.noDisableDenormals = toMABool8(b)
}

func (d *DeviceConfig) GetNoFixedSizedCallback() bool {
	return maBool8ToGoBool(d.config.noFixedSizedCallback)
}

func (d *DeviceConfig) SetNoFixedSizedCallback(b bool) {
	d.config.noFixedSizedCallback = toMABool8(b)
}

func (d *DeviceConfig) SetPeriodSizeInMilliseconds(m int) {
	d.config.periodSizeInMilliseconds = C.ma_uint32(m)
}

func (d *DeviceConfig) GetPeriodSizeInMilliseconds() int {
	return int(d.config.periodSizeInMilliseconds)
}



//TODO: DeviceConfig DataCallback and notification callback
func (d *DeviceConfig) SetDataCallback(callback DeviceDataProcCallback) {
	dataCallback = callback
	C.go_ma_device_config_set_data_callback(d.cptr())
}

func (d *DeviceConfig) SetUserData(data PCMFrameReader) {
	var userDataPtr unsafe.Pointer
	switch userData := data.(type) {
	case *Decoder:
		userDataPtr = unsafe.Pointer(userData.cptr())
	case *Waveform:
		userDataPtr = unsafe.Pointer(userData.cptr())
	}
	if userDataPtr != nil {
		d.config.pUserData = userDataPtr
	}
}

func (d *DeviceConfig) GetUserData(data PCMFrameReader) {
	switch t := data.(type) {
	case *Decoder:
		ptr := (*C.struct_ma_decoder)(d.config.pUserData)
		t.decoder = ptr
	}
}

func GetUserDataFromDeviceConfig(config *DeviceConfig, data any) {
	switch t := data.(type) {
	case *Decoder:
		ptr := (*C.struct_ma_decoder)(config.config.pUserData)
		t.decoder = ptr
	}
}

type Device struct {
	device *C.struct_ma_device
}

func NewDevice() *Device {
	ptr := (*C.struct_ma_device)(C.malloc(C.sizeof_struct_ma_device))
	return &Device{
		device: ptr,
	}
}

func (d *Device) Close() {
	C.free(unsafe.Pointer(d.cptr()))
}

func (c *Device) cptr() *C.struct_ma_device {
	if c == nil {
		return nil
	}
	return c.device
}

func (d *Device) Init(context *Context, config *DeviceConfig) error {
	return deviceInit(context, config, d)
}

func DeviceInit(context *Context, config *DeviceConfig, device *Device) error {
	return deviceInit(context, config, device)
}

func deviceInit(context *Context, config *DeviceConfig, device *Device) error {
	contextPtr := context.cptr()
	configPtr := config.cptr()
	devicePtr := device.cptr()
	res := C.ma_device_init(contextPtr, configPtr, devicePtr)
	return checkResult(res)
}

func toMaBackendArray(backends []MABackend) []C.ma_backend {
	maBackends := make([]C.ma_backend, 0, len(backends))
	for _, backend := range backends {
		maBackends = append(maBackends, C.ma_backend(backend))
	}
	return maBackends
}

func (d *Device) InitEx(backends []MABackend, contextConfig *ContextConfig, deviceConfig *DeviceConfig) error {
	return deviceInitEx(backends, contextConfig, deviceConfig, d)
}

func deviceInitEx(backEnds []MABackend, contextConfig *ContextConfig, deviceConfig *DeviceConfig, device *Device) error {
	contextConfigPtr := contextConfig.cptr()
	deviceConfigPtr := deviceConfig.cptr()
	devicePtr := device.cptr()
	maBackends := toMaBackendArray(backEnds)
	backendcount := len(maBackends)
	var res C.ma_result

	if backendcount == 0 {
		res =  C.ma_device_init_ex(nil, 0, contextConfigPtr, deviceConfigPtr, devicePtr)
	} else {
		res = C.ma_device_init_ex(&maBackends[0], C.ma_uint32(backendcount), contextConfigPtr, deviceConfigPtr, devicePtr)
	}

	return checkResult(res)
}

func (d *Device) Uninit() {
	deviceUninit(d)
}

func DeviceUninit(device *Device) {
	deviceUninit(device)
}

func deviceUninit(device *Device) {
	C.ma_device_uninit(device.cptr())
}

func (d *Device) GetContext() *Context {
	ptr := C.ma_device_get_context(d.cptr())
	return &Context{
		context: ptr,
	}
}

func (d *Device) GetDeviceInfo(deviceType DeviceType, deviceInfo *DeviceInfo) error {
	res := C.ma_device_get_info(d.cptr(), C.ma_device_type(deviceType), deviceInfo.cptr())
	return checkResult(res)
}

func (d *Device) GetName(deviceType DeviceType) (string, error) {
	nameCap := C.MA_MAX_DEVICE_NAME_LENGTH+1
	nameBuf := make([]byte, nameCap)
	charPtr := (*C.char)(C.CBytes(nameBuf))
	defer C.free(unsafe.Pointer(charPtr))
	res := C.ma_device_get_name(d.cptr(), C.ma_device_type(deviceType), charPtr, C.size_t(nameCap), nil)
	if err := checkResult(res); err != nil {
		return "", err
	}
	return string(nameBuf), nil
}

func (d *Device) Start() error {
	res := C.ma_device_start(d.cptr())
	return checkResult(res)
}

func (d *Device) Stop() error {
	res := C.ma_device_stop(d.cptr())
	return checkResult(res)
}

func (d *Device) IsStarted() bool {
	res := C.ma_device_is_started(d.cptr())
	return maBoolToGoBool(res)
}

func (d *Device) GetState() DeviceState {
	state := C.ma_device_get_state(d.cptr())
	return DeviceState(state)
}

func (d *Device) SetMasterVolume(volume float32) error {
	res := C.ma_device_set_master_volume(d.cptr(), C.float(volume))
	return checkResult(res)
}

func (d *Device) GetMasterVolume() (float32, error) {
	var volume C.float
	res := C.ma_device_get_master_volume(d.cptr(), &volume)
	if err := checkResult(res); err != nil {
		return float32(volume), err
	}
	return float32(volume), nil
}

func (d *Device) SetMasterVolumeDB(gainDB float32) error {
	res := C.ma_device_set_master_volume_db(d.cptr(), C.float(gainDB))
	return checkResult(res)
}

func (d *Device) GetMasterVolumeDB() (float32, error) {
	var volume C.float
	res := C.ma_device_get_master_volume_db(d.cptr(), &volume)
	if err := checkResult(res); err != nil {
		return float32(volume), err
	}
	return float32(volume), nil
}

func (d *Device) HandleBackendDataCallback(pOutput unsafe.Pointer, pInput unsafe.Pointer, frameCount uint32) error {
	res := C.ma_device_handle_backend_data_callback(d.cptr(), pOutput, pInput, C.ma_uint32(frameCount))
	return checkResult(res)
}

func (d *Device) CalculateBufferSizeInFramesFromDescriptor(descriptor *DeviceDescriptor, nativeSampleRate uint32, performanceProfile PerformanceProfile) int {
	bufferSize := C.ma_calculate_buffer_size_in_frames_from_descriptor(descriptor.cptr(), C.ma_uint32(nativeSampleRate), C.ma_performance_profile(performanceProfile))
	return int(bufferSize)
}

func (d *Device) GetDeviceType() DeviceType {
	return DeviceType(d.device._type)
}

func (d *Device) SetDeviceType(deviceType DeviceType) {
	d.device._type = C.ma_device_type(deviceType)
}

func (d *Device) GetSampleRate() int {
	return int(d.device.sampleRate)
}

func (d *Device) SetSampleRate(sampleRate int) {
	d.device.sampleRate = C.ma_uint32(sampleRate)
}

func (d *Device) GetPlaybackFormat() Format {
	return Format(d.device.playback.format)
}

func (d *Device) GetPlaybackChannels() int {
	return int(d.device.playback.channels)
}

func (d *Device) GetPlaybackPeriods() int {
	return int(d.device.playback.internalPeriods)
}

func (d *Device) GetPlaybackPeriodSizeInFrames() int {
	return int(d.device.playback.internalPeriodSizeInFrames)
}

func (d *Device) GetPlaybackInputCache() *Buffer {
	ptr := d.device.playback.pInputCache
	return &Buffer{
		buffer: ptr,
	}
}

func (d *Device) GetPlaybackInputCacheCap() int {
	return int(d.device.playback.inputCacheCap)
}

func (d *Device) GetPlaybackInputCacheConsumed() int {
	return int(d.device.playback.inputCacheConsumed)
}

func (d *Device) GetPlaybackInputCacheRemaining() int {
	return int (d.device.playback.inputCacheRemaining)
}

func (d *Device) GetPlaybackIntermediaryBufferCap() int {
	return int(d.device.playback.intermediaryBufferCap)
}

func (d *Device) GetPlaybackIntermediaryBufferLen() int {
	return int(d.device.playback.intermediaryBufferLen)
}

func (d *Device) GetPlaybackIntermediaryBuffer() *Buffer {
	ptr := d.device.playback.pIntermediaryBuffer
	return &Buffer{
		buffer: ptr,
	}
}

func (d *Device) GetCaptureIntermediaryBufferCap() int {
	return int(d.device.capture.intermediaryBufferCap)
}

func (d *Device) GetCaptureIntermediaryBufferLen() int {
	return int(d.device.capture.intermediaryBufferLen)
}

func (d *Device) GetCaptureIntermediaryBuffer() *Buffer {
	return &Buffer{
		buffer: d.device.capture.pIntermediaryBuffer,
	}
}

func (d *Device) GetCaptureSampleRate() int {
	return int(d.device.capture.internalSampleRate)
}

func (d *Device) GetCapturePeriodSizeInFrames() int {
	return int(d.device.capture.internalPeriodSizeInFrames)
}

func (d *Device) GetCaptureChannels() int {
	return int(d.device.capture.internalChannels)
}

func (d *Device) GetCapturePeriods() int {
	return int(d.device.capture.internalPeriods)
}

func (d *Device) GetCapturceFormat() Format {
	return Format(d.device.capture.format)
}

func (d *Device) GetUserData(userData PCMFrameReader) {
	switch t := userData.(type) {
	case *Decoder:
		t.decoder = (*C.struct_ma_decoder)(d.device.pUserData)
	case *Waveform:
		t.wf = (*C.ma_waveform)(d.device.pUserData)
	}
}

func (d *Device) GetDataSource() *DataSource {
	var dataSource DataSource
	ptr := (*C.ma_data_source)(d.device.pUserData)
	dataSource.src = unsafe.Pointer(ptr)
	return &dataSource
}

type ContextConfig struct {
	config C.struct_ma_context_config
}

func (c *ContextConfig) cptr() *C.struct_ma_context_config {
	if c == nil {
		return nil
	}
	return &c.config
}

func ContextConfigInit() *ContextConfig {
	config := C.ma_context_config_init()
	return &ContextConfig{
		config: config,
	}
}

type Context struct {
	context *C.struct_ma_context
}

func contextFromPtr(ptr unsafe.Pointer) *Context {
	return &Context{
		context: (*C.struct_ma_context)(ptr),
	}
}

func deviceInfoFromPtr(ptr unsafe.Pointer) *DeviceInfo {
	return &DeviceInfo{
		info: (*C.ma_device_info)(ptr),
	}
}

func (c *Context) cptr() *C.struct_ma_context {
	if c == nil {
		return nil
	}
	return c.context
}

func NewContext() *Context {
	context := (*C.struct_ma_context)(C.malloc(C.sizeof_struct_ma_context))
	return &Context{
		context: context,
	}
}

func (c *Context) Close() {
	C.free(unsafe.Pointer(c.cptr()))
}

func (c *Context) Init(config *ContextConfig) error {
	res := C.ma_context_init(nil, C.ma_uint32(0), config.cptr(), c.cptr())
	return checkResult(res)
}

func (c *Context) Uninit() error {
	res := C.ma_context_uninit(c.cptr())
	return checkResult(res)
}

func (c *Context) EnumerateDevices(done chan struct{}, callback EnumerateDevicesCallback) error {
	defer close(done)
	enumerateDevicesCallback = callback
	return checkResult(C.go_ma_context_enumerate_devices(c.cptr(), nil))
}

func VersionString() string {
	return C.GoString(C.ma_version_string())
}

type ResourceManagerPipelineNotifications struct {
	n C.ma_resource_manager_pipeline_notifications
}
 
func ResourceManagerPipelineNotificationsInit() *ResourceManagerPipelineNotifications {
	n := C.ma_resource_manager_pipeline_notifications_init()
	return &ResourceManagerPipelineNotifications{
		n: n,
	}
}

func (n *ResourceManagerPipelineNotifications) cptr() *C.ma_resource_manager_pipeline_notifications {
	if n == nil {
		return nil
	}
	return &n.n
}

type ResourceManagerDataBuffer struct {
	buffer *C.struct_ma_resource_manager_data_buffer
}

func NewResourceManagerDataBuffer() *ResourceManagerDataBuffer {
	buffer := (*C.struct_ma_resource_manager_data_buffer)(C.malloc(C.sizeof_struct_ma_resource_manager_data_buffer))
	return &ResourceManagerDataBuffer{
		buffer: buffer,
	}
}

func (b *ResourceManagerDataBuffer) cptr() *C.struct_ma_resource_manager_data_buffer {
	if b == nil {
		return nil
	}
	return b.buffer
}

func (b *ResourceManagerDataBuffer) Close() {
	C.free(unsafe.Pointer(b.cptr()))
}

func (b *ResourceManagerDataBuffer) Init(m *ResourceManager, filePath string, flags uint, n *ResourceManagerPipelineNotifications) error {
	pFilePath := C.CString(filePath)
	defer C.free(unsafe.Pointer(pFilePath))
	res := C.ma_resource_manager_data_buffer_init(m.cptr(), pFilePath, C.ma_uint32(flags), n.cptr(), b.cptr())
	return checkResult(res)
}

func (b *ResourceManagerDataBuffer) InitCopy(m *ResourceManager, existingBuffer *ResourceManagerDataBuffer) error {
	res := C.ma_resource_manager_data_buffer_init_copy(m.cptr(), existingBuffer.cptr(), b.cptr())
	return checkResult(res)
}

func (b *ResourceManagerDataBuffer) Uninit() error {
	res := C.ma_resource_manager_data_buffer_uninit(b.cptr())
	return checkResult(res)
}

func (b *ResourceManagerDataBuffer) ReadPCMFrames(framesOut *Buffer, frameCount int) (int, error) {
	var framesRead C.ma_uint64

	res := C.ma_resource_manager_data_buffer_read_pcm_frames(b.cptr(), framesOut.cptr(), C.ma_uint64(frameCount), &framesRead)

	if err := checkResult(res); err != nil {
		return int(framesRead), err
	}

	return int(framesRead), nil
}

func (b *ResourceManagerDataBuffer) SeekToPCMFrame(frameIndex int) error {
	res := C.ma_resource_manager_data_buffer_seek_to_pcm_frame(b.cptr(), C.ma_uint64(frameIndex))
	return checkResult(res)
}

type DataFormat struct {
	Format Format
	Channels int
	SampleRate int
	ChannelMap Channel
	ChannelMapCap int
}

func (b *ResourceManagerDataBuffer) GetDataFormat(channelMapCap int) (DataFormat, error) {
	var dataFormat DataFormat

	var pFormat C.ma_format
	var pChannels C.ma_uint32
	var pSampleRate C.ma_uint32
	var pChannelMap C.ma_channel

	res := C.ma_resource_manager_data_buffer_get_data_format(b.cptr(), &pFormat, &pChannels, &pSampleRate, &pChannelMap, C.size_t(channelMapCap))

	if err := checkResult(res); err != nil {
		return dataFormat, err
	}

	dataFormat.Format = Format(pFormat)
	dataFormat.Channels = int(pChannels)
	dataFormat.SampleRate = int(pSampleRate)
	dataFormat.ChannelMap = Channel(pChannelMap)

	return dataFormat, nil
}

func (b *ResourceManagerDataBuffer) GetCursorInPCMFrames() (int, error) {
	var cursor C.ma_uint64

	res := C.ma_resource_manager_data_buffer_get_cursor_in_pcm_frames(b.cptr(), &cursor)

	if err := checkResult(res); err != nil {
		return int(cursor), err
	}

	return int(cursor), nil
}

func (b *ResourceManagerDataBuffer) GetLengthInPCMFrames() (int, error) {
	var length C.ma_uint64

	res := C.ma_resource_manager_data_buffer_get_length_in_pcm_frames(b.cptr(), &length)

	if err := checkResult(res); err != nil {
		return int(length), err
	}

	return int(length), nil
}

func (b *ResourceManagerDataBuffer) Result() error {
	res := C.ma_resource_manager_data_buffer_result(b.cptr())
	return checkResult(res)
}

func (b *ResourceManagerDataBuffer) SetLooping(isLooping bool) error {
	res := C.ma_resource_manager_data_buffer_set_looping(b.cptr(), toMABool32(isLooping))
	return checkResult(res)
}

func (b *ResourceManagerDataBuffer) IsLooping() bool {
	res := C.ma_resource_manager_data_buffer_is_looping(b.cptr())
	return maBoolToGoBool(res)
}

func (b *ResourceManagerDataBuffer) GetAvailableFrames() (int, error) {
	var frames C.ma_uint64
	res := C.ma_resource_manager_data_buffer_get_available_frames(b.cptr(), &frames)
	if err := checkResult(res); err != nil {
		return int(frames), err
	}
	return int(frames), nil
}

type ResourceManagerConfig struct {
	config C.ma_resource_manager_config
}

func (c *ResourceManagerConfig) cptr() *C.ma_resource_manager_config {
	if c == nil {
		return nil
	}
	return &c.config
}

func NewResourceManagerConfig() *ResourceManagerConfig {
	config := C.ma_resource_manager_config_init()
	return &ResourceManagerConfig{
		config: config,
	}
}

type ResourceManager struct {
	manager *C.struct_ma_resource_manager
}

func (m *ResourceManager) cptr() *C.ma_resource_manager {
	if m == nil {
		return nil
	}
	return m.manager
}

func NewResourceManager() *ResourceManager {
	manager := (*C.struct_ma_resource_manager)(C.malloc(C.sizeof_struct_ma_resource_manager))
	return &ResourceManager{
		manager: manager,
	}
}

func (m *ResourceManager) Close() {
	ptr := m.cptr()
	C.free(unsafe.Pointer(ptr))
}

func (m *ResourceManager) Init(config *ResourceManagerConfig) error {
	configPtr := config.cptr()
	managerPtr := m.cptr()
	res := C.ma_resource_manager_init(configPtr, managerPtr)
	return checkResult(res)
}

func (m *ResourceManager) Uninit() {
	C.ma_resource_manager_uninit(m.cptr())
}

func (m *ResourceManager) RegisterFile(filePath string, flags uint32) error {
	ptr := m.cptr()
	maFlags := C.ma_uint32(flags)
	path := C.CString(filePath)
	defer C.free(unsafe.Pointer(path))
	res := C.ma_resource_manager_register_file(ptr, path, maFlags)
	return checkResult(res)
}

func (m *ResourceManager) UnregisterFile(filePath string ) error {
	ptr := m.cptr()
	path := C.CString(filePath)
	defer C.free(unsafe.Pointer(path))
	res := C.ma_resource_manager_unregister_file(ptr, path)
	return checkResult(res)
}

func (m *ResourceManager) RegisterDecodedData(name string, pData unsafe.Pointer, frameCount int, format Format, channels int, sampleRate int) error {
	pName := C.CString(name)
	defer C.free(unsafe.Pointer(pName))
	res := C.ma_resource_manager_register_decoded_data(m.cptr(), pName, pData, C.ma_uint64(frameCount), C.ma_format(format), C.ma_uint32(channels), C.ma_uint32(sampleRate))
	return checkResult(res)
}

func (m *ResourceManager) RegisterEncodedData(name string, pData unsafe.Pointer, sizeInBytes int) error {
	pName := C.CString(name)
	defer C.free(unsafe.Pointer(pName))
	res := C.ma_resource_manager_register_encoded_data(m.cptr(), pName, pData, C.size_t(sizeInBytes))
	return checkResult(res)
}

func (m *ResourceManager) UnregisterData(name string) error {
	pName := C.CString(name)
	defer C.free(unsafe.Pointer(pName))
	res := C.ma_resource_manager_unregister_data(m.cptr(), pName)
	return checkResult(res)
}

type DecoderConfig struct {
	config C.ma_decoder_config
}

func (c *DecoderConfig) cptr() *C.ma_decoder_config {
	if c == nil {
		return nil
	}
	return &c.config
}

func DecoderConfigInit(outputFormat MAFormat, outputChannels uint32, outputSampleRate uint32) *DecoderConfig {
	format := C.ma_format(outputFormat)
	chs := C.ma_uint32(outputChannels)
	sampleRate := C.ma_uint32(outputSampleRate)
	config := C.ma_decoder_config_init(format, chs, sampleRate)
	return &DecoderConfig{
		config: config,
	}
}

func DecoderConfigInitDefault() *DecoderConfig {
	config := C.ma_decoder_config_init_default()
	return &DecoderConfig{
		config: config,
	}
}

type Decoder struct {
	decoder *C.struct_ma_decoder
}

func (d *Decoder) cptr() *C.struct_ma_decoder {
	if d == nil {
		return nil
	}
	return d.decoder
}

func NewDecoder() *Decoder {
	ptr := (*C.struct_ma_decoder)(C.malloc(C.sizeof_struct_ma_decoder))
	return &Decoder{
		decoder: ptr,
	}
}

func (d *Decoder) Close() {
	ptr := d.cptr()
	C.free(unsafe.Pointer(ptr))
}

func (d *Decoder) InitFile(filePath string, config *DecoderConfig) error {
	return decoderInitFile(filePath, config, d)
}

func DecoderInitFile(filePath string, config *DecoderConfig, decoder *Decoder) error {
	return decoderInitFile(filePath, config, decoder)
}

func decoderInitFile(filePath string, config *DecoderConfig, decoder *Decoder) error {
	configPtr := config.cptr()
	decoderPtr := decoder.cptr()
	path := C.CString(filePath)
	defer C.free(unsafe.Pointer(path))
	res := C.ma_decoder_init_file(path, configPtr, decoderPtr)
	return checkResult(res)
}

func (d *Decoder) Uninit() error {
	return decoderUninit(d)
}

func DecoderUninit(decoder *Decoder) error {
	return decoderUninit(decoder)
}

func decoderUninit(decoder *Decoder) error {
	ptr := decoder.cptr()
	res := C.ma_decoder_uninit(ptr)
	return checkResult(res)
}

func (d *Decoder) GetOutputFormat() Format {
	return Format(d.decoder.outputFormat)
}

func (d *Decoder) GetOutputChannels() int {
	return int(d.decoder.outputChannels)
}

func (d *Decoder) GetOutputSampleRate() int {
	return int(d.decoder.outputSampleRate)
}

func (d *Decoder) ReadPCMFrames(pFramesOut *Buffer, frameCount int) (int, error) {
	var framesRead C.ma_uint64

	res := C.ma_decoder_read_pcm_frames(d.cptr(), pFramesOut.cptr(), C.ma_uint64(frameCount), &framesRead)

	if err := checkResult(res); err != nil {
		return int(framesRead), err
	}
	return int(framesRead), nil
}

func (d *Decoder) SeekToPCMFrame(frameIndex int) error {
	res := C.ma_decoder_seek_to_pcm_frame(d.cptr(), C.ma_uint64(frameIndex))
	return checkResult(res)
}

func (d *Decoder) GetDataFormat(channelMapCap uint) (DataFormat, error) {
	var dataFormat DataFormat

	var pFormat C.ma_format
	var pChannels C.ma_uint32
	var pSampleRate C.ma_uint32
	var pChannelMap C.ma_channel

	res := C.ma_decoder_get_data_format(d.cptr(), &pFormat, &pChannels, &pSampleRate, &pChannelMap, C.size_t(channelMapCap))

	if err := checkResult(res); err != nil {
		return dataFormat, err
	}

	dataFormat.Format = Format(pFormat)
	dataFormat.Channels = int(pChannels)
	dataFormat.SampleRate = int(pSampleRate)
	dataFormat.ChannelMap = Channel(pChannelMap)

	return dataFormat, nil
}

func (d *Decoder) GetCursorInPCMFrames() (int, error) {
	var pCursor C.ma_uint64

	res := C.ma_decoder_get_cursor_in_pcm_frames(d.cptr(), &pCursor)

	if err := checkResult(res); err != nil {
		return 0, err
	}

	return int(pCursor), nil
}

func (d *Decoder) GetLengthInPCMFrames() (int, error) {
	var length C.ma_uint64

	res := C.ma_decoder_get_length_in_pcm_frames(d.cptr(), &length)

	if err := checkResult(res); err != nil {
		return 0, err
	}

	return int(length), nil
}

func (d *Decoder) GetAvailableFrames() (int, error) {
	var availableFrames C.ma_uint64

	res := C.ma_decoder_get_available_frames(d.cptr(), &availableFrames)

	if err := checkResult(res); err != nil {
		return 0, err
	}

	return int(availableFrames), nil
}

type EncoderConfig struct {
	config C.ma_encoder_config
}

func (e *EncoderConfig) cptr() *C.ma_encoder_config {
	if e == nil {
		return nil
	}
	return &e.config
}

func EncoderConfigInit(encodingFormat EncodingFormat, format Format, channels int, sampleRate int) *EncoderConfig {
	config := C.ma_encoder_config_init(C.ma_encoding_format(encodingFormat), C.ma_format(format), C.ma_uint32(channels), C.ma_uint32(sampleRate))
	return &EncoderConfig{
		config: config,
	}
}

type Encoder struct {
	encoder *C.struct_ma_encoder
}

func (e *Encoder) cptr() *C.struct_ma_encoder {
	if e == nil {
		return nil
	}
	return e.encoder
}

func NewEncoder() *Encoder {
	encoder := (*C.struct_ma_encoder)(C.malloc(C.sizeof_struct_ma_encoder))
	return &Encoder{
		encoder: encoder,
	}
}

func (e *Encoder) Close() {
	C.free(unsafe.Pointer(e.cptr()))
}

func (e *Encoder) InitFile(filePath string, config *EncoderConfig) error {
	cPath := C.CString(filePath)
	defer C.free(unsafe.Pointer(cPath))
	res := C.ma_encoder_init_file(cPath, config.cptr(), e.cptr())
	return checkResult(res)
}

func (e *Encoder) Uninit() {
	C.ma_encoder_uninit(e.cptr())
}

func (e *Encoder) WritePCMFrames(pFramesIn unsafe.Pointer, frameCount int) (int, error) {
	var framesWritten C.ma_uint64

	res := C.ma_encoder_write_pcm_frames(e.cptr(), pFramesIn, C.ma_uint64(frameCount), &framesWritten)

	if err := checkResult(res); err != nil {
		return 0, err
	}

	return int(framesWritten), nil
}

type AllocationCallbacks struct {
	ac *C.ma_allocation_callbacks
}

func (a *AllocationCallbacks) cptr() *C.ma_allocation_callbacks {
	if a == nil {
		return nil
	}
	return a.ac
}

type AudioBufferConfig struct {
	config C.ma_audio_buffer_config
}

func AudioBufferConfigInit(format Format, channels int, sizeInFrames int, data *Buffer, allocationCallbacks *AllocationCallbacks) *AudioBufferConfig {
	config := C.ma_audio_buffer_config_init(C.ma_format(format), C.ma_uint32(channels), C.ma_uint64(sizeInFrames), data.cptr(), allocationCallbacks.cptr())
	return &AudioBufferConfig{
		config: config,
	}
}

func (a *AudioBufferConfig) cptr() *C.ma_audio_buffer_config {
	if a == nil {
		return nil
	}
	return &a.config
}

type AudioBuffer struct {
	buffer *C.ma_audio_buffer
}

func NewAudioBuffer() *AudioBuffer {
	buffer := (*C.ma_audio_buffer)(C.malloc(C.sizeof_ma_audio_buffer))
	return &AudioBuffer{
		buffer: buffer,
	}
}

func (a *AudioBuffer) cptr() *C.ma_audio_buffer {
	if a == nil {
		return nil
	}
	return a.buffer
}

func (a *AudioBuffer) Close() {
	C.free(unsafe.Pointer(a.cptr()))
}

func (a *AudioBuffer) Init(config *AudioBufferConfig) error {
	res := C.ma_audio_buffer_init(config.cptr(), a.cptr())
	return checkResult(res)
}

func (a *AudioBuffer) InitCopy(config *AudioBufferConfig) error {
	res := C.ma_audio_buffer_init_copy(config.cptr(), a.cptr())
	return checkResult(res)
}

func (a *AudioBuffer) Uninit() {
	C.ma_audio_buffer_uninit(a.cptr())
}

func (a *AudioBuffer) ReadPCMFrames(framesOut *Buffer, frameCount int, loop bool) int {
	res := C.ma_audio_buffer_read_pcm_frames(a.cptr(), framesOut.cptr(), C.ma_uint64(frameCount), toMABool32(loop))
	return int(res)
}

func (a *AudioBuffer) SeekToPCMFrame(frameIndex int) error {
	res := C.ma_audio_buffer_seek_to_pcm_frame(a.cptr(), C.ma_uint64(frameIndex))
	return checkResult(res)
}

type WaveformType int

const (
	SINE WaveformType = iota
	SQUARE
	TRIANGLE
	SAWTOOTH
)

type WaveformConfig struct {
	config C.ma_waveform_config
}

func (c *WaveformConfig) cptr() *C.ma_waveform_config {
	if c == nil {
		return nil
	}
	return &c.config
}

func WaveformConfigInit(format Format, channels uint, sampleRate uint, wfType WaveformType, amplitude float64, frequency float64) *WaveformConfig {
	config := C.ma_waveform_config_init(C.ma_format(format), C.ma_uint32(channels), C.ma_uint32(sampleRate), C.ma_waveform_type(wfType), C.double(amplitude), C.double(frequency))
	return &WaveformConfig{
		config: config,
	}
}

type Waveform struct {
	wf *C.ma_waveform
}

func NewWaveform() *Waveform {
	wf := (*C.ma_waveform)(C.malloc(C.sizeof_ma_waveform))
	return &Waveform{
		wf: wf,
	}
}

func (wf *Waveform) cptr() *C.ma_waveform {
	if wf == nil {
		return nil
	}
	return wf.wf
}

func (wf *Waveform) Close() {
	C.free(unsafe.Pointer(wf.cptr()))
}

func (wf *Waveform) Init(config *WaveformConfig) error {
	res := C.ma_waveform_init(config.cptr(), wf.cptr())
	return checkResult(res)
}

func (wf *Waveform) Uninit() {
	C.ma_waveform_uninit(wf.cptr())
}

func (wf *Waveform) ReadPCMFrames(pFramesOut *Buffer, frameCount int) (int, error) {
	var framesRead C.ma_uint64

	res := C.ma_waveform_read_pcm_frames(wf.cptr(), pFramesOut.cptr(), C.ma_uint64(frameCount), &framesRead)

	if err := checkResult(res); err != nil {
		return 0, err
	}

	return int(framesRead), nil
}

func (wf *Waveform) SeekToPCMFrame(frameIndex int) error {
	res := C.ma_waveform_seek_to_pcm_frame(wf.cptr(), C.ma_uint64(frameIndex))
	return checkResult(res)
}

func (wf *Waveform) SetAmplitude(amplitude float64) error {
	res := C.ma_waveform_set_amplitude(wf.cptr(), C.double(amplitude))
	return checkResult(res)
}

func (wf *Waveform) SetFrequency(frequency float64) error {
	res := C.ma_waveform_set_frequency(wf.cptr(), C.double(frequency))
	return checkResult(res)
}

func (wf *Waveform) SetType(wfType WaveformType) error {
	res := C.ma_waveform_set_type(wf.cptr(), C.ma_waveform_type(wfType))
	return checkResult(res)
}

func (wf *Waveform) SetSampleRate(sampleRate int) error {
	res := C.ma_waveform_set_sample_rate(wf.cptr(), C.ma_uint32(sampleRate))
	return checkResult(res)
}

type PulsewaveConfig struct {
	config C.ma_pulsewave_config
}

func PulsewaveConfigInit(format Format, channels int, sampleRate int, dutyCycle float64, amplitude float64, frequency float64) *PulsewaveConfig {
	config := C.ma_pulsewave_config_init(C.ma_format(format), C.ma_uint32(channels), C.ma_uint32(sampleRate), C.double(dutyCycle), C.double(amplitude), C.double(frequency))
	return &PulsewaveConfig{
		config: config,
	}
}

func (c *PulsewaveConfig) cptr() *C.ma_pulsewave_config {
	if c == nil {
		return nil
	}
	return &c.config
}

type Pulsewave struct {
	pw *C.ma_pulsewave
}

func NewPulsewave() *Pulsewave {
	pw := (*C.ma_pulsewave)(C.malloc(C.sizeof_ma_pulsewave))
	return &Pulsewave{
		pw: pw,
	}
}

func (pw *Pulsewave) cptr() *C.ma_pulsewave {
	if pw == nil {
		return nil
	}
	return pw.pw
}

func (pw *Pulsewave) Close() {
	C.free(unsafe.Pointer(pw.cptr()))
}

func (pw *Pulsewave) Init(config *PulsewaveConfig) error {
	res := C.ma_pulsewave_init(config.cptr(), pw.cptr())
	return checkResult(res)
}

func (pw *Pulsewave) Uninit() {
	C.ma_pulsewave_uninit(pw.cptr())
}

func (pw *Pulsewave) ReadPCMFrames(pFramesOut *Buffer, frameCount int) (int, error) {
	var framesRead C.ma_uint64

	res := C.ma_pulsewave_read_pcm_frames(pw.cptr(), pFramesOut.cptr(), C.ma_uint64(frameCount), &framesRead)

	if err := checkResult(res); err != nil {
		return 0, err
	}

	return int(framesRead), nil
}

func (pw *Pulsewave) SeekToPCMFrame(frameIndex int) error {
	res := C.ma_pulsewave_seek_to_pcm_frame(pw.cptr(), C.ma_uint64(frameIndex))
	return checkResult(res)
}

func (pw *Pulsewave) SetAmplitude(amplitude float64) error {
	res := C.ma_pulsewave_set_amplitude(pw.cptr(), C.double(amplitude))
	return checkResult(res)
}

func (pw *Pulsewave) SetFrequency(frequency float64) error {
	res := C.ma_pulsewave_set_frequency(pw.cptr(), C.double(frequency))
	return checkResult(res)
}

func (pw *Pulsewave) SetSampleRate(sampleRate int) error {
	res := C.ma_pulsewave_set_sample_rate(pw.cptr(), C.ma_uint32(sampleRate))
	return checkResult(res)
}

func (pw *Pulsewave) SetDutyCycle(dutyCycle float64) error {
	res := C.ma_pulsewave_set_duty_cycle(pw.cptr(), C.double(dutyCycle))
	return checkResult(res)
}

type NoiseType int

const (
	WHITE NoiseType = iota
	PINK
	BROWNIAN
)

type NoiseConfig struct {
	config C.ma_noise_config
}

func NoiseConfigInit(format Format, channels int, noiseType NoiseType, seed int, amplitude float64) *NoiseConfig {
	config := C.ma_noise_config_init(C.ma_format(format), C.ma_uint32(channels), C.ma_noise_type(noiseType), C.ma_int32(seed), C.double(amplitude))
	return &NoiseConfig {
		config: config,
	}
}

func (c *NoiseConfig) cptr() *C.ma_noise_config {
	if c == nil {
		return nil
	}
	return &c.config
}

func (c *NoiseConfig) GetHeapSize() (int, error) {
	var size C.size_t
	res := C.ma_noise_get_heap_size(c.cptr(), &size)

	if err := checkResult(res); err != nil {
		return 0, err
	}

	return int(size), nil
}

type Noise struct {
	noise *C.ma_noise
}

type Fence struct {
	fence C.ma_fence
}

func (f *Fence) cptr() *C.ma_fence {
	if f == nil {
		return nil
	}
	return &f.fence
}

func (f *Fence) Init() error {
	return fenceInit(f)
}

func FenceInit(fence *Fence) error {
	return fenceInit(fence)
}

func fenceInit(fence *Fence) error {
	res := C.ma_fence_init(fence.cptr())
	return checkResult(res)
}

func (f *Fence) Uninit() {
	fenceUninit(f)
}

func FenceUninit(fence *Fence) {
	fenceUninit(fence)
}

func fenceUninit(fence *Fence) {
	C.ma_fence_uninit(fence.cptr())
}

func (f *Fence) Acquire() error {
	return fenceAcquire(f)
}

func FenceAcquire(fence *Fence) error {
	return fenceAcquire(fence)
}

func fenceAcquire(fence *Fence) error {
	res := C.ma_fence_acquire(fence.cptr())
	return checkResult(res)
}

func (f *Fence) Release() error {
	return fenceRelease(f)
}

func FenceRelease(fence *Fence) error {
	return fenceRelease(fence)
}

func fenceRelease(fence *Fence) error {
	res := C.ma_fence_release(fence.cptr())
	return checkResult(res)
}

func (f *Fence) Wait() error {
	return fenceWait(f)
}

func FenceWait(fence *Fence) error {
	return fenceWait(fence)
}

func fenceWait(fence *Fence) error {
	res := C.ma_fence_wait(fence.cptr())
	return checkResult(res)
}
