package main

import (
"fmt"
"strings"
)

type MediaPlayer interface {
	play(audioType string, fileName string)
}

type AdvanceMediaPlay interface {
	playVlc(fileName string)
	playMp4(fileName string)
}

type VlcPlayer struct {

}

func (v *VlcPlayer)playVlc(fileName string) {
	fmt.Println("Playing vlc file. Name: ", fileName);
}

func (v *VlcPlayer)playMp4(fileName string) {

}

type Mp4Player struct {

}

func (v *Mp4Player)playVlc(fileName string) {

}

func (v *Mp4Player)playMp4(fileName string) {
	fmt.Println("Playing mp4 file. Name: ", fileName);
}

type MediaAdapter struct {
	advanceMediaPlay AdvanceMediaPlay
}

func (m *MediaAdapter)mediaAdapter(audioType string) {
	audioType = strings.ToLower(audioType)
	if (audioType == "vlc") {
		m.advanceMediaPlay = new(VlcPlayer)
	} else if (audioType == "mp4") {
		m.advanceMediaPlay = new(Mp4Player)
	}
}

func (m *MediaAdapter)play(audioType string, fileName string) {
	audioType = strings.ToLower(audioType)
	if (audioType == "vlc") {
		m.advanceMediaPlay.playVlc(fileName)
	} else if (audioType == "mp4") {
		m.advanceMediaPlay.playMp4(fileName)
	}
}


type AuidoPlayer struct {
	mediaAdapter *MediaAdapter
}

func(a *AuidoPlayer)play(audioType string, fileName string){
	audioType = strings.ToLower(audioType)
	if(audioType=="mp3"){
		fmt.Println("Playing mp3 file. Name: ", fileName);
	}else if(audioType=="vlc"||audioType=="mp4"){
		a.mediaAdapter=new(MediaAdapter)
		a.mediaAdapter.mediaAdapter(audioType)
		a.mediaAdapter.play(audioType,fileName)
	}else{
		fmt.Println("Invalid media. ",audioType," format not supported");
	}
}


func main(){
	//demo
	audioPlayer:=new(AuidoPlayer)
	audioPlayer.play("mp3", "beyond the horizon.mp3");
	audioPlayer.play("mp4", "alone.mp4");
	audioPlayer.play("vlc", "far far away.vlc");
	audioPlayer.play("avi", "mind me.avi");
}


