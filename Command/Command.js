/**
 * Created by jamesyan on 2017/3/25.
 */

var Stock=function(name,quantity){
    this.name=name;
    this.quantity=quantity;
};

Stock.prototype.buy=function(){
    console.log("stock [name"+this.name+ ",quantity:"+this.quantity+"] bought\r\n");
}

Stock.prototype.sell=function(){
    console.log("stock [name"+this.name+ ",quantity:"+this.quantity+"] sold\r\n");
}


var BuyOrder=function(stock){
    this.stock=stock;
}

BuyOrder.prototype.execute=function(){
    this.stock.buy();
}


var SellOrder=function(stock){
    this.stock=stock;
}

SellOrder.prototype.execute=function(){
    this.stock.sell();
}


var Broker=function(){
    this.orderList=[];
}
Broker.prototype.takeOrder=function(order){
    this.orderList.push(order);
}

Broker.prototype.placeOrder=function(){
    for(i in this.orderList){
        var order=this.orderList[i];
        order.execute();
    }
    this.orderList=[];
}