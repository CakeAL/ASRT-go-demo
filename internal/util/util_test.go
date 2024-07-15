package util_test

import (
	"ASRT-go-demo/internal/util"
	"encoding/binary"
	"fmt"
	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
	"os"
	"testing"
)

func TestReadWAV(t *testing.T) {
	res, err := util.ReadWAV("../../test.wav")
	if err != nil {
		t.Fatal(err)
	}
	segmentIndex := 0
	var outFile *os.File
	var encoder *wav.Encoder
	for _, audioS := range res {
		// 创建新的输出文件
		if outFile != nil {
			encoder.Close()
			outFile.Close()
		}
		outFile, err = os.Create(fmt.Sprintf("test_segment_%d.wav", segmentIndex))
		if err != nil {
			t.Fatal(err)
		}
		encoder = wav.NewEncoder(outFile, int(audioS.SampleRate), int(audioS.ByteWidth*8), int(audioS.Channels), 1)

		audioBuf := &audio.IntBuffer{
			Data:   make([]int, len(audioS.WavBytes)/int(audioS.ByteWidth)),
			Format: &audio.Format{SampleRate: int(audioS.SampleRate), NumChannels: int(audioS.Channels)},
		}

		for i := 0; i < len(audioBuf.Data); i++ {
			start := i * int(audioS.ByteWidth)
			end := start + int(audioS.ByteWidth)
			if end > len(audioS.WavBytes) {
				break
			}
			audioBuf.Data[i] = int(binary.LittleEndian.Uint16(audioS.WavBytes[start:end]))
		}

		err = encoder.Write(audioBuf)
		if err != nil {
			t.Fatal(err)
		}

		err = encoder.Close()
		if err != nil {
			t.Fatal(err)
		}

		segmentIndex++
	}
}
