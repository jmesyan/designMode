/**
 * Created by jamesyan on 2017/3/19.
 */


/**
 * interface MediaPlayer{
                public function play($audioType,$fileName);
            }

 interface AdvanceMediaPlayer{
                public function playVlc($fileName);
                public function playMp4($fileName);
            }
 */


var VlcPlayer={
    playVlc:function(fileName){
        console.log("Playing vlc file. Name: ", fileName);
    },

    playMp4:function(fileName){

    }
}

var Mp4Player={
    playVlc:function(fileName){

    },

    playMp4:function(fileName){
        console.log("Playing mp4 file. Name: ", fileName);
    }
}


var MediaAdapter=function(audioType){
    audioType=audioType.toLowerCase();
    if(audioType=="vlc"){
        var v=function(){}
        v.prototype=VlcPlayer;
        this.advanceMediaPlayer=new v();
    }else if(audioType=="mp4"){
        var v=function(){}
        v.prototype=Mp4Player;
        this.advanceMediaPlayer=new v();
    }
}

MediaAdapter.prototype.play=function(audioType,fileName)
{
    audioType = audioType.toLowerCase();
    if (audioType == "vlc") {
        this.advanceMediaPlayer.playVlc(fileName);
    } else if (audioType == "mp4") {
        this.advanceMediaPlayer.playMp4(fileName);
    }
}


var AudioPlayer=function(){};
AudioPlayer.prototype.play=function(audioType,fileName){
    audioType = audioType.toLowerCase();
    if(audioType=="mp3"){
        console.log("Playing mp3 file. Name: ", fileName);
    }else if(audioType=="vlc"||audioType=="mp4"){
        this.mediaAdapter=new MediaAdapter(audioType);
        this.mediaAdapter.play(audioType,fileName)
    }else{
        console.log("Invalid media. "+audioType+" format not supported");
    }
}


