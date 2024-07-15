package util

import (
	"ASRT-go-demo/internal/requests"
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/go-audio/audio"
	"io"
	"os"

	"github.com/go-audio/wav"
)

// / ReadWAV 输入文件路径，返回
func ReadWAV(path string) ([]*requests.Wav, error) {
	// 读取文件
	wavFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer wavFile.Close()

	// 解码文件
	wavDecoder := wav.NewDecoder(wavFile)
	if !wavDecoder.IsValidFile() {
		return nil, errors.New("invalid file")
	}

	sampleRate := wavDecoder.SampleRate
	channels := wavDecoder.NumChans
	byteWidth := wavDecoder.BitDepth / 8
	// 每一段字节数
	segmentSize := sampleRate * uint32(channels) * uint32(byteWidth) * 16

	var wavs []*requests.Wav

	for {
		buf := make([]byte, segmentSize)
		bytesRead, err := wavFile.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if bytesRead == 0 {
			break
		}

		audioBuf := &audio.IntBuffer{
			Data:   make([]int, bytesRead/int(byteWidth)),
			Format: wavDecoder.Format(),
		}
		// Read the byte slice into the IntBuffer's Data field
		for i := 0; i < len(audioBuf.Data); i++ {
			start := i * int(byteWidth)
			end := start + int(byteWidth)
			if end > len(buf) {
				break
			}
			audioBuf.Data[i] = int(binary.LittleEndian.Uint16(buf[start:end]))
		}

		bytesBuf := new(bytes.Buffer)
		for _, sample := range audioBuf.Data {
			if err := binary.Write(bytesBuf, binary.LittleEndian, int16(sample)); err != nil {
				return nil, err
			}
		}

		wavs = append(wavs, &requests.Wav{
			Channels:   wavDecoder.NumChans,
			SampleRate: wavDecoder.SampleRate,
			ByteWidth:  wavDecoder.BitDepth / 8,
			WavBytes:   bytesBuf.Bytes(),
		})
	}
	return wavs, nil
}
