package gominiaudio

import (
	"testing"
)

func TestMAEngineConfigInitialization(t *testing.T) {
	config := NewMAEngineConfig()

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
