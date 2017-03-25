/**
 * Created by jamesyan on 2017/3/25.
 */

var ChatRoom=function(){
    this.sendMessage=function(user,message){
        var now=new Date().toLocaleString();
        console.log(now+":["+user.name+"]"+message+"\r\n");
    }
}

var chatRoom=new ChatRoom();

var User=function(name){
    this.name=name;
};

User.prototype.getName=function(){
    return this.name;
}

User.prototype.setName=function(name){
    this.name=name;
}

User.prototype.sendMessage=function(message){
    chatRoom.sendMessage(this,message);
}