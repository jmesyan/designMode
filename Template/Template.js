/**
 * Created by jamesyan on 2017/3/25.
 */

var Cricket=function(){
    this.initlize=function(){
        console.log("Cricket initlize\r\n");
    }

    this.startPlay=function(){
        console.log("Cricket startPlay\r\n");
    }

    this.endPlay=function(){
        console.log("Cricket endPlay\r\n");
    }
}

var FootbBall=function(){
    this.initlize=function(){
        console.log("FootbBall initlize\r\n");
    }

    this.startPlay=function(){
        console.log("FootbBall startPlay\r\n");
    }

    this.endPlay=function(){
        console.log("FootbBall endPlay\r\n");
    }
}

function Play(game){
    game.initlize();
    game.startPlay();
    game.endPlay();
}

