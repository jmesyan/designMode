/**
 * Created by jamesyan on 2017/3/25.
 */

var RealCustomer=function(name){
    this.name=name;
}

RealCustomer.prototype.isNil=function(){
    return false;
}

RealCustomer.prototype.getName=function(){
    return this.name;
}

var NullCustomer=function(){

}
NullCustomer.prototype.isNil=function(){
    return true;
}

NullCustomer.prototype.getName=function(){
    return "no name in database";
}


var CustomerFactory=function(){
    this.names=["Rob", "Job", "Juile"];
}

CustomerFactory.prototype.getCustomer=function(name){
    for(var n in this.names){
        var na=this.names[n];
        if(na==name){
            return new RealCustomer(name);
        }
    }
    return new NullCustomer();
}
