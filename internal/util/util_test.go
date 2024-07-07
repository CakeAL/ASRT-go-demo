package util_test

import (
	"ASRT-go-demo/internal/util"
	"testing"
)

func TestReadWAV(t *testing.T) {
	res, err := util.ReadWAV("test.wav")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res.ByteWidth, res.Channels, res.SampleRate)
}