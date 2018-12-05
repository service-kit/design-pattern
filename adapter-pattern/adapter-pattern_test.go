package adapter_pattern

import (
	"fmt"
	"testing"
)

func Test_AdapterPattern(t *testing.T) {
	ma := new(AudioPlayer)
	err := ma.Play(MP4, "test.mp4")
	if nil != err {
		fmt.Println("play err ", err.Error())
	}
	err = ma.Play(VLC, "test.vlc")
	if nil != err {
		fmt.Println("play err ", err.Error())
	}
	err = ma.Play("pcm", "test.pcm")
	if nil != err {
		fmt.Println("play err ", err.Error())
	}
}
