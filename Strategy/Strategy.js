/**
 * Created by jamesyan on 2017/3/25.
 */

var OperationAdd=function(){}
OperationAdd.prototype.doOperation=function(num1,num2){
    return num1+num2;
}

var OperationSubstract=function(){}
OperationSubstract.prototype.doOperation=function(num1,num2){
    return num1-num2;
}

var OperationMultipy=function(){}
OperationMultipy.prototype.doOperation=function(num1,num2){
    return num1*num2;
}


var Context=function(strategy){
    this.strategy=strategy;
}

Context.prototype.executeStrategy=function(num1,num2){
    return this.strategy.doOperation(num1,num2);
}