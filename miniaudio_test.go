package gominiaudio

import (
	"testing"
	"unsafe"
)

func TestEngineConfigInitialization(t *testing.T) {
	config := NewEngineConfig()

	t.Log("ma_engine successfully initialized")

	config.SetListenerCount(1)
	config.SetChannels(1)
	config.SetPeriodSizeInFrames(1)
	config.SetPeriodSizeInMilliseconds(1)
	config.SetGainSmoothTimeInFrames(1)
	config.SetGainSmoothTimeInMilliseconds(1)
	config.SetDefaultVolumeSmoothTimeInPCMFrames(1)
	config.SetPreMixStackSizeInBytes(1)
	config.SetNoAutoStart(true)
	config.SetNoDevice(true)

	if value := config.GetListenerCount(); value != uint32(1) {
		t.Errorf("incorrect value. expected %d got %d", uint32(1), value)
	}

	if value := config.GetChannels(); value != uint32(1) {
		t.Errorf("incorrect value. expected %d got %d", uint32(1), value)
	}

	if value := config.GetPeriodSizeInFrames(); value != uint32(1) {
		t.Errorf("incorrect value. expected %d got %d", uint32(1), value)
	}

	if value := config.GetPeriodSizeInMilliseconds(); value != uint32(1) {
		t.Errorf("incorrect value. expected %d got %d", uint32(1), value)
	}

	if value := config.GetGainSmoothTimeInFrames(); value != uint32(1) {
		t.Errorf("incorrect value for gainSmoothTimeInFrames. expected %d got %d", uint32(1), value)
	}

	if value := config.GetGainSmoothTimeInMilliseconds(); value != uint32(1) {
		t.Errorf("incorrect value for gainSmoothTimeInMilliseconds. expected %d got %d", uint32(1), value)
	}

	if value := config.GetDefaultVolumeSmoothTimeInPCMFrames(); value != uint32(1) {
		t.Errorf("incorrect value for defaultVolumeSmoothTimeInPCMFrames. expected %d got %d", uint32(1), value)
	}

	if value := config.GetPreMixStackSizeInBytes(); value != uint32(1) {
		t.Errorf("incorrect value in preMixStackSizeInBytes. expected %d got %d", uint32(1), value)
	}

	if value := config.GetNoAutoStart(); value != true {
		t.Errorf("incorrect value for noAutoStart. expected %t got %t", true, value)
	}

	if value := config.GetNoDevice(); value != true {
		t.Errorf("incorrect value for noDevice. expected %t got %t", true, value)
	}
}

func TestEngineSoundConfigInitialization(t *testing.T) {
	engine := NewEngine()
	defer engine.Close()

	_ = NewSoundConfig(engine)

	t.Log("successfully initialized sound config")

}

func TestEngineInit(t *testing.T) {
	engine := NewEngine()
	defer engine.Close()

	err := engine.Init(nil)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestDecoderInitFromFile(t *testing.T) {
	decoder := NewDecoder()
	defer decoder.Close()

	testFile := "testdata/sine.wav"

	if err := decoder.InitFile(testFile, nil); err != nil {
		t.Error(err)
		t.FailNow()
	}

	defer decoder.Uninit()

	expectedFormat := decoder.GetOutputFormat()
	expectedSampleRate := decoder.GetOutputSampleRate()
	expectedChannels := decoder.GetOutputChannels()

	t.Logf("expectedFormat: %v", expectedFormat)
	t.Logf("expected sample rate: %v", expectedSampleRate)
	t.Logf("expected channels: %v", expectedChannels)

	deviceConfig := DeviceConfigInit(DeviceTypePlayback)

	deviceConfig.SetPlaybackFormat(expectedFormat)
	deviceConfig.SetSampleRate(expectedSampleRate)
	deviceConfig.SetPlaybackChannels(expectedChannels)
	deviceConfig.SetUserData(decoder)

	if value := deviceConfig.GetPlaybackFormat(); value != expectedFormat {
		t.Errorf("incorrect value for (*DeviceConfig).GetPlaybackFormat() got %v expected %v", value, expectedFormat)
	}

	if value := deviceConfig.GetPlaybackChannels(); value != expectedChannels {
		t.Errorf("incorrect value for (*DeviceConfig).GetPlaybackChannels() got %v expected %v", value, expectedChannels)
	}

	if value := deviceConfig.GetSampleRate(); value != expectedSampleRate {
		t.Errorf("incorrect value for (*DeviceConfig).GetSampleRate() got %v expected %v", value, expectedChannels)
	}

	dec := Decoder{}

	deviceConfig.GetUserData(&dec)

	if value := dec.cptr(); value != decoder.cptr() {
		t.Errorf("incorrect ptr for (*DeviceConfig).GetUserData()")
	}

	device := NewDevice()
	defer device.Close()

	if err := device.Init(nil, deviceConfig); err != nil {
		t.Errorf("DeviceInit(...) failed. unable to initialize device. got %v", err)
		t.FailNow()
	}

	defer device.Uninit()

}

func TestContextInitialization(t *testing.T) {
	contextConfig := ContextConfigInit()

	context := NewContext()
	defer context.Close()

	err := context.Init(contextConfig)

	if err != nil {
		t.Errorf("Context.Init(*ContextConfig) failed to initialize got: %v", err)
		t.FailNow()
	}

	var callback EnumerateDevicesCallback

	ch := make(chan string)
	done := make(chan struct{})

	callback = func(c *Context, deviceType DeviceType, pinfo *DeviceInfo, userData unsafe.Pointer) bool {
		if pinfo.IsDefault() {
			ch <- pinfo.Name()
		}
		return true
	}

	var devices []string

	go func() {
		defer close(ch)
		for {
			select {
			case <-done:
				t.Log("done is closed")
				return
			case device := <-ch:
				devices = append(devices, device)
			}
		}
	}()

	if err := context.EnumerateDevices(done, callback); err != nil {
		t.Errorf("Context.EnumerateDevices failed got %v", err)
		t.FailNow()
	}

	for _, device := range devices {
		t.Log(device)
	}

	if err := context.Uninit(); err != nil {
		t.Errorf("Context.Uninit() failed to uninitialize context got: %v", err)
		t.FailNow()
	}
}
