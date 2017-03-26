/**
 * Created by jamesyan on 2017/3/26.
 */
var Service1=function(){
    return {
        getName:function(){
            return "Service1";
        },
        execute:function(){
            console.log("execute service1")
        }
    }
}

var Service2=function(){
    return {
        getName:function(){
            return "Service2";
        },
        execute:function(){
            console.log("execute service2")
        }
    }
}

var InitialContext=function(){
    return{
        lookup:function(serviceName){
            serviceName=serviceName.toLowerCase();
            if(serviceName=="service1"){
                console.log("Looking up and creating a new Service1 object");
                return new Service1();
            }else if(serviceName=="service2"){
                console.log("Looking up and creating a new Service2 object");
                return new Service2();
            }
            return null;
        }
    }
}

var Cache=function(){
    return {
        services:[],
        getService:function(serviceName){
            for(var s in this.services){
                var service=this.services[s]
                if(service.getName().toLowerCase()==serviceName){
                    return service;
                }
            }
            return null;
        },
        addService:function(newService){
            var exists=false;
            for(var s in this.services){
                var service=this.services[s];
                if(service.getName()==newService.getName()){
                    exists=true;
                    break;
                }
            }
            if(!exists){
                this.services.push(newService);
            }
        }
    }
}



var ServiceLocator=function(){
    var context=new InitialContext();
    var cache=new Cache();
    return {
        getService:function(serviceName){
            var service=cache.getService(serviceName)
            if(service!=null){
                return service;
            }
            var service1=context.lookup(serviceName);
            cache.addService(service1);
            return service1;
        }
    }
}