package util

import (
	"ASRT-go-demo/internal/requests"
	"bytes"
	"encoding/binary"
	"os"

	"github.com/go-audio/wav"
)

func ReadWAV(path string) (requests.Wav, error) {
    // 读取文件
	wavFile, err := os.Open(path)
	if err != nil {
		return requests.Wav{}, err
	}

	// 解码文件
	wavDecoder := wav.NewDecoder(wavFile)
	buf, err := wavDecoder.FullPCMBuffer()
	if err != nil {
		return requests.Wav{}, err
	}
	byteBuffer := new(bytes.Buffer)
	for _, sample := range buf.Data { // 小端
		err = binary.Write(byteBuffer, binary.LittleEndian, int16(sample))
		if err != nil {
			return requests.Wav{}, err
		}
	}
	
	myWav := requests.Wav{
		Channels:   wavDecoder.NumChans,
		SampleRate: wavDecoder.SampleRate,
		ByteWidth:  wavDecoder.BitDepth / 8,
		WavBytes:   byteBuffer.Bytes(),
	}
	return myWav, nil
}
