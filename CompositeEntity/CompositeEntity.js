/**
 * Created by jamesyan on 2017/3/25.
 */

var DependentObject1=function(){
    this.setData=function(data){
        this.data=data;
    }
    this.getData=function(){
        return this.data
    }
}

var DependentObject2=function(){
    this.setData=function(data){
        this.data=data;
    }
    this.getData=function(){
        return this.data
    }
}

var CoarseGrainedObject=function(){
    this.do1=new DependentObject1();
    this.do2=new DependentObject2();
    this.setData=function(data1,data2){
        this.do1.setData(data1);
        this.do2.setData(data2);
    }

    this.getData=function(){
        return [this.do1.getData(),this.do2.getData()];
    }
}

var CompositeEntity=function(){
    this.cgo=new CoarseGrainedObject();
    this.getData=function(){
        return this.cgo.getData();
    }
    this.setData=function(data1,data2){
        this.cgo.setData(data1,data2);
    }
}

var Client=function(){
    this.entity=new CompositeEntity();
    this.printData=function(){
        for(var d in this.entity.getData()){
            var data=this.entity.getData()[d];
            console.log("Data :"+data);
        }
    }
    this.setData=function(data1,data2){
        this.entity.setData(data1,data2);
    }

}


