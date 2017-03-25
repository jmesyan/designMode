/**
 * Created by jamesyan on 2017/3/25.
 */



var NameReposity=function(){
    var names=["Robert","John","Julie","Lora"];
    this.getIterator=function(){
        return new NameIterator();
    }
    var NameIterator =function(){
        this.index=0;
    }

    NameIterator.prototype.hasNext=function(){
        if(this.index<names.length){
            return true;
        }
        return false;
    }

    NameIterator.prototype.next=function(){
        if(this.hasNext()){
            return names[this.index++];
        }
        return null;
    }
}


