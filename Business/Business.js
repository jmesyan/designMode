/**
 * Created by jamesyan on 2017/3/25.
 */

var EJBService=function(){
    this.doProcessing=function(){
        console.log("Processing task by invoking EJB Service");
    }
};

var JMSService=function(){
    this.doProcessing=function(){
        console.log("Processing task by invoking JMS Service");
    }
};

var BusinessLookUp=function(){
    this.getBusinessService=function(serviceType){
        if(serviceType=="EJB"){
            return new EJBService();
        }else{
            return new JMSService();
        }
    }
}


var BusinessDelegate=function(){
    this.bussinessLookup=new BusinessLookUp();
    this.setSeviceType=function(serviceType){
        this.serviceType=serviceType
    }
    this.doTask=function(){
        var service=this.bussinessLookup.getBusinessService(this.serviceType);
        service.doProcessing();
    }
}

var Client=function(delegate){
      this.businessDelegate=delegate;
      this.doTask=function(){
          this.businessDelegate.doTask();
      }
}
