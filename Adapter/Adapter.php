<?php
/**
 * Created by PhpStorm.
 * User: jamesyan
 * Date: 2017/3/19
 * Time: 10:07
 */

interface MediaPlayer{
    public function play($audioType,$fileName);
}

interface AdvanceMediaPlayer{
    public function playVlc($fileName);
    public function playMp4($fileName);
}


class VlcPlayer implements AdvanceMediaPlayer{
    public function playVlc($fileName){
        echo "Playing vlc file. Name: ".$fileName."<br/>";
    }

    public function playMp4($fileName){

    }
}


class Mp4Player implements AdvanceMediaPlayer{
    public function playVlc($fileName){

    }

    public function playMp4($fileName){
        echo "Playing mp4 file. Name: ".$fileName."<br/>";
    }
}


class MediaAdapter implements MediaPlayer{
    static $advanceMuiscPlayer;
    public function __construct($audioType)
    {
        $audioType=strtolower($audioType);
        if($audioType=="vlc"){
            self::$advanceMuiscPlayer=new VlcPlayer();
        }else if($audioType=="mp4"){
            self::$advanceMuiscPlayer=new Mp4Player();
        }
    }
    public function play($audioType, $fileName)
    {
        // TODO: Implement play() method.
        $audioType=strtolower($audioType);
        if($audioType=="vlc"){
           self::$advanceMuiscPlayer->playVlc($fileName);
        }else if($audioType=="mp4"){
            self::$advanceMuiscPlayer->playMp4($fileName);
        }
    }
}


class AudioPlayer implements MediaPlayer{
    static $mediaAdapter;
    public function play($audioType, $fileName)
    {
        // TODO: Implement play() method.
        $audioType=strtolower($audioType);
        if($audioType=="mp3"){
            echo "Playing mp3 file. Name: $fileName"."<br/>";
        }else if($audioType=="vlc"||$audioType=="mp4"){
            self::$mediaAdapter=new MediaAdapter($audioType);
            self::$mediaAdapter->play($audioType,$fileName);
        }else{
            echo "Invalid media. ".$audioType." format not supported"."<br/>";
        }

    }
}

//demo
$audioPlayer=new AudioPlayer();
$audioPlayer->play("mp3", "beyond the horizon.mp3");
$audioPlayer->play("mp4", "alone.mp4");
$audioPlayer->play("vlc", "far far away.vlc");
$audioPlayer->play("avi", "mind me.avi");