/**
 * Created by jamesyan on 2017/3/14.
 */

var Singleton=(function () {
    var instance;
    function init(){
        var privateRandomNumber=Math.random()
        return {
            getRandomNumber:function(){
                return privateRandomNumber;
            }
        }
    }
    return{
        getInstance:function(){
            if(!instance){ //instance==null||instance=="undefined"等
                console.log(11111);
                instance=init();
            }
            console.log(2222);
            return instance;
        }
    }
})()