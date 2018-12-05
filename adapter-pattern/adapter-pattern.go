package adapter_pattern

import (
	"errors"
	"fmt"
)

const (
	VLC = "vlc"
	MP4 = "mp4"
)

type MediaPlayer interface {
	Play(audioType, fileName string)
}

type AdvanceMediaPlayer interface {
	PlayVLC(fileName string)
	PlayMp4(fileName string)
}

type VLCPlayer struct {
}

func (vp *VLCPlayer) PlayVLC(fileName string) {
	fmt.Println("play vlc file ", fileName)
}

func (vp *VLCPlayer) PlayMp4(fileName string) {

}

type Mp4Player struct {
}

func (vp *Mp4Player) PlayVLC(fileName string) {
}

func (vp *Mp4Player) PlayMp4(fileName string) {
	fmt.Println("play mp4 file ", fileName)
}

type MediaAdapter struct {
	Player AdvanceMediaPlayer
}

func (ma *MediaAdapter) Init(audioType string) {
	if VLC == audioType {
		ma.Player = &VLCPlayer{}
	} else if MP4 == audioType {
		ma.Player = &Mp4Player{}
	}
}

func (ma *MediaAdapter) Play(audioType, fileName string) {
	if VLC == audioType {
		ma.Player.PlayVLC(fileName)
	} else if MP4 == audioType {
		ma.Player.PlayMp4(fileName)
	}
}

type AudioPlayer struct {
}

func (ap *AudioPlayer) Play(audioType, fileName string) error {
	ma := &MediaAdapter{}
	if VLC == audioType {
		ma.Init(VLC)
	} else if MP4 == audioType {
		ma.Init(MP4)
	} else {
		return errors.New(audioType + " not supported")
	}
	ma.Play(audioType, fileName)
	return nil
}
