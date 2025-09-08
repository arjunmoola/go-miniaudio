//go:build dev
package gominiaudio
// #cgo LDFLAGS: ${SRCDIR}/build/miniaudio.o -ldl -lpthread -lm
// #include "thirdparty/miniaudio.h"
import "C"
