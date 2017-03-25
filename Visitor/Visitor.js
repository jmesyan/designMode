/**
 * Created by jamesyan on 2017/3/25.
 */

var Mouse=function(){
    this.accept=function(vistor){
        vistor.visit(this)
    }
}


var Keyboard=function(){
    this.accept=function(vistor){
        vistor.visit(this)
    }
}

var Monitor=function(){
    this.accept=function(vistor){
        vistor.visit(this)
    }
}


var Computer=function(){
    this.parts=[new Monitor(),new Keyboard(),new Mouse()];
    this.accept=function(vistor){
        for(var p in this.parts){
            var part=this.parts[p];
            part.accept(vistor)
        }
        vistor.visit(this)
    }
}

var ComputerVistor=function(){
    this.visit=function(part){
         var partName=part.constructor.name;
         switch (partName){
             case "Monitor":
                 console.log("Displaying Monitor");
                 break;
             case "Keyboard":
                 console.log("Displaying Keyboard");
                 break;
             case "Mouse":
                 console.log("Displaying Mouse");
                 break;
             case "Computer":
                 console.log("Displaying Computer");
                 break;
         }
    }
}