/**
 * Created by jamesyan on 2017/3/26.
 */

var AuthFilter={execute:function(request){console.log("Authing Request:"+request)}}
var DebugFilter={execute:function(request){console.log("log Request:"+request)}}
var Target={execute:function(request){console.log("executing Request:"+request)}}

var FilterChain=function(target){
    this.filterChains=[];
    this.target=target;
}
FilterChain.prototype.addFilter=function(filter){
    this.filterChains.push(filter);
}

FilterChain.prototype.setTarget=function(target){
    this.target=target;
}

FilterChain.prototype.execute=function(request){
    for(var f in this.filterChains){
        var filter=this.filterChains[f]
        filter.execute(request)
    }
    this.target.execute(request)
}

var FilterManager=function(){
    this.filterChain=new FilterChain(Target);
}
FilterManager.prototype.setFilter=function(filter){
    this.filterChain.addFilter(filter)
}

FilterManager.prototype.filterRequest=function(request){
    this.filterChain.execute(request)
}

var Client=function(filterManager){
    this.filterManager=filterManager;
}

Client.prototype.sendRequest=function(request){
    this.filterManager.filterRequest(request);
}