package strUtil

import (
	"bytes"
	"testing"
)

func TestStrToBytes(t *testing.T) {
	expected := []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f}
	result := StrToBytes("hello")
	if !bytes.Equal(result, expected) {
		t.Errorf("StrToBytes test failed, expected %v but got %v", expected, result)
	}
}
