<?php
/**
 * Created by PhpStorm.
 * User: jamesyan
 * Date: 2017/3/25
 * Time: 19:09
 */

abstract class Game{
    abstract  function initialize();
    abstract  function startPlay();
    abstract  function endPlay();

    public final function play(){
        //初始化
        $this->initialize();
        //开始游戏
        $this->startPlay();
        //结束游戏
        $this->endPlay();
    }
}


class Cricket extends Game{
    function initialize()
    {
        // TODO: Implement initialize() method.
        var_dump("Cricket Game Initialized! Start playing.");
    }

    function startPlay()
    {
        // TODO: Implement startPlay() method.
        var_dump("Cricket Game Started. Enjoy the game!");
    }

    function endPlay()
    {
        // TODO: Implement endPlay() method.
        var_dump("Cricket Game Finished!");
    }
}



class Football extends Game{
    function initialize()
    {
        // TODO: Implement initialize() method.
        var_dump("Football Game Initialized! Start playing.");
    }

    function startPlay()
    {
        // TODO: Implement startPlay() method.
        var_dump("Football Game Started. Enjoy the game!");
    }

    function endPlay()
    {
        // TODO: Implement endPlay() method.
        var_dump("Football Game Finished!");
    }
}


//demo
$game=new Cricket();
$game->play();

$game=new Football();
$game->play();

