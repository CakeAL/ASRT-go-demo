# ASRT Go Demo 使用 Go 调用 ASRT 语言识别的客户端 Demo

## 简介

如题。

思路：

使用 go-audio 来转码 mp3/wav 为 16000Hz，并分割为 16s 的音频，然后分段调用 ASRT 的 POST 接口。 获取全部信息后返回。

> 暂未完成

## 使用方法

1. 下载[ASRT 语音识别服务端](https://wiki.ailemon.net/docs/asrt-doc/download)
2. 创建 `venv` 环境，使用 `Python 3.9` 装好所有依赖（不想费劲可以用 CPU 跑）。
3. 运行 `asrserver_http.py`。
4. 运行示例程序：

```bash
$ just r
# 或者
$ go mod tidy
$ go run main.go -f test.wav
```
