//go:build !dev
package gominiaudio
// #cgo LDFLAGS: -ldl -lpthread -lm
// #define MINIAUDIO_IMPLEMENTATION
// #include "thirdparty/miniaudio.h"
import "C"
