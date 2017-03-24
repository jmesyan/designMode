/**
 * Created by jamesyan on 2017/3/25.
 */


/**
 * interface display(){}
 */

var RealImage=function(fileName){
    this.fileName=fileName;

    var loadFromDisk=function(){
        console.log("load disk file:"+fileName+"\r\n");
    };

    loadFromDisk()
};


RealImage.prototype.display=function(){
    console.log("displaying file:"+this.fileName);
}

var ProxyImage=function(fileName){
    this.fileName=fileName
}
ProxyImage.prototype.display=function (){
    if(this.realImage==null){
        this.realImage=new RealImage(this.fileName)
    }
    this.realImage.display()
}

