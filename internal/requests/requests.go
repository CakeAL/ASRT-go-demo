package requests

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
)

type Wav struct {
	Channels   uint16
	SampleRate uint32
	ByteWidth  uint16
	WavBytes   []byte
}

type Resp struct {
	Result        string `json:"result"`
	StatusCode    int32  `json:"status_code"`
	StatusMessage string `json:"status_message"`
}

func (wav *Wav) SendPost(serviceUrl string) (Resp, error) {
	param := make(map[string]interface{})
	param["channels"] = wav.Channels
	param["sample_rate"] = wav.SampleRate
	param["byte_width"] = wav.ByteWidth
	
	samples := base64.URLEncoding.EncodeToString(wav.WavBytes)
	param["samples"] = samples
	bytesData, _ := json.Marshal(param)

	// 发送请求
	res, err := http.Post(serviceUrl, "application/json", bytes.NewBuffer([]byte(bytesData)))
	if err != nil {
		return Resp{}, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	// 读取响应 Body 
	content, err := io.ReadAll(res.Body)
	if err != nil {
		return Resp{}, err
	}

	// 反序列化
	var result Resp
	err = json.Unmarshal(content, &result)
	if err != nil {
		return Resp{}, err
	}

	return result, nil
}
