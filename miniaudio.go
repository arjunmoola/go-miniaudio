package gominiaudio

// #cgo LDFLAGS: ${SRCDIR}/build/miniaudio.o -ldl -lpthread -lm
// #include "thirdparty/miniaudio.h"
// #include <stdlib.h>
// #include <stdio.h>
import "C"

import (
	"unsafe"
	"fmt"
	//"errors"
)

type maContainer interface {
	cptr() unsafe.Pointer
}

func compare(a, b maContainer) bool {
	return a.cptr() == b.cptr()
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

func (d *DeviceConfig) SetUserData(data any) {
	var userDataPtr unsafe.Pointer
	switch userData := data.(type) {
	case *Decoder:
		userDataPtr = unsafe.Pointer(userData.cptr())
	}
	if userDataPtr != nil {
		d.config.pUserData = userDataPtr
	}
}

func (d *DeviceConfig) GetUserData(data any) {
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

func (d *Device) Uninit() {
	deviceUninit(d)
}

func DeviceUninit(device *Device) {
	deviceUninit(device)
}

func deviceUninit(device *Device) {
	C.ma_device_uninit(device.cptr())
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
	context C.struct_ma_context
}

func (c *Context) cptr() *C.struct_ma_context {
	if c == nil {
		return nil
	}
	return &c.context
}

func VersionString() string {
	return C.GoString(C.ma_version_string())
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
	return resourceManagerInit(config, m)
}

func ResourceManagerInit(config *ResourceManagerConfig, manager *ResourceManager) error {
	return resourceManagerInit(config, manager)
}

func resourceManagerInit(config *ResourceManagerConfig, manager *ResourceManager) error {
	configPtr := config.cptr()
	managerPtr := manager.cptr()
	res := C.ma_resource_manager_init(configPtr, managerPtr)
	return checkResult(res)
}

func (m *ResourceManager) Uninit() {
	resourceManagerUninit(m)
}

func ResourceManagerUninit(manager *ResourceManager) {
	resourceManagerUninit(manager)
}

func resourceManagerUninit(manager *ResourceManager) {
	C.ma_resource_manager_uninit(manager.cptr())
}

func (m *ResourceManager) RegisterFile(filePath string, flags uint32) error {
	return resourceManagerRegisterFile(m, filePath, flags)
}

func ResourceManagerRegisterFile(manager *ResourceManager, filePath string, flags uint32) error {
	return resourceManagerRegisterFile(manager, filePath, flags)
}

func resourceManagerRegisterFile(manager *ResourceManager, filePath string, flags uint32) error {
	ptr := manager.cptr()
	maFlags := C.ma_uint32(flags)
	path := C.CString(filePath)
	defer C.free(unsafe.Pointer(path))
	res := C.ma_resource_manager_register_file(ptr, path, maFlags)
	return checkResult(res)
}

func (m *ResourceManager) UnregisterFile(filePath string ) error {
	return resourceManagerUnregisterFile(m, filePath)
}

func ResourceManagerUnregisterFile(manager *ResourceManager, filePath string) error {
	return resourceManagerUnregisterFile(manager, filePath)
}

func resourceManagerUnregisterFile(manager *ResourceManager, filePath string) error {
	ptr := manager.cptr()
	path := C.CString(filePath)
	defer C.free(unsafe.Pointer(path))
	res := C.ma_resource_manager_unregister_file(ptr, path)
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
