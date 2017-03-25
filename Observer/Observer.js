/**
 * Created by jamesyan on 2017/3/25.
 */
var Subject = function () {
    this.state = 0;
    this.observers = [];
    this.getState = function () {
        return this.state;
    }
    this.setState = function (state) {
        this.state = state;
        this.notifyObservers();
    }

    this.attach = function (observer) {
        this.observers.push(observer);
    }

    this.notifyObservers = function () {
        for (var n in this.observers) {
            var observer = this.observers[n];
            observer.update();
        }
    }
}

var BinaryObserve=function(subject){
    this.subject=subject;
    this.subject.attach(this);
}

BinaryObserve.prototype.update=function(){
    console.log("this is binary string:"+this.subject.getState());
}

var OctalObserve=function(subject){
    this.subject=subject;
    this.subject.attach(this);
}

OctalObserve.prototype.update=function(){
    console.log("this is binary string:"+this.subject.getState());
}

var HexaObserve=function(subject){
    this.subject=subject;
    this.subject.attach(this);
}

HexaObserve.prototype.update=function(){
    console.log("this is binary string:"+this.subject.getState());
}


