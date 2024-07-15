package main

import (
	"ASRT-go-demo/internal/util"
	"flag"
	"fmt"
	"os"
)

var f = flag.String("f", "", "wav/mp3 filename.")

func main() {
	flag.Parse()
	if *f == "" {
		fmt.Println("\tinput -help for help")
	}
	// fmt.Println(*f)
	// 读取音频
	audios, err := util.ReadWAV(*f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	serviceUrl := "http://127.0.0.1:20001/all"

	// 调用接口
	var result string
	for _, audio := range audios {
		res, err := audio.SendPost(serviceUrl)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		result = result + res.Result
	}

	fmt.Printf("转音频结果: %v", result)
}
