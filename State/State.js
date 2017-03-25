/**
 * Created by jamesyan on 2017/3/25.
 */

var Context=function(){
    this.setState=function(state){
        this.state=state;
    }
    this.getState=function(){
        return this.state;
    }
}

var startState=function(){}
startState.prototype.doAction=function(context){
    console.log("player is in start state");
    context.setState(this)
}

startState.prototype.toString=function(){
    console.log("Start State");
}

var stopState=function(){}
stopState.prototype.doAction=function(context){
    console.log("player is in stop state");
    context.setState(this);
}

stopState.prototype.toString=function(){
    console.log("Stop State");
}



