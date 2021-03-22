package main

import (
	"bufio"
	"github.com/huagetai/painter/examples/share"
	"os"
)

func main() {
	a := share.Activity{
		Title: "有翼小助手",
		Images: []string{
			"./assets/test-1.jpeg",
			"./assets/test-2.jpeg",
			"./assets/test-3.jpeg",
			"./assets/test-4.jpeg",
			"./assets/test-5.jpeg",
			"./assets/test-6.jpeg",
		},
		BtnName:   "我要团购",
		AvatarUrl: "./assets/logo.png",
	}
	f, _ := os.Create("./assets/out.png")
	w := bufio.NewWriter(f)
	a.Gen(w)
	w.Flush()
}
