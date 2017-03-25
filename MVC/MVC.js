/**
 * Created by jamesyan on 2017/3/25.
 */

var Student=function(name,rollNo){
    this.name=name;
    this.rollNo=rollNo;
}

Student.prototype.getName=function(){
    return this.name;
}

Student.prototype.setName=function(name){
    this.name=name;
}

Student.prototype.getRollNo=function(){
    return this.rollNo;
}

Student.prototype.setRollNo=function(rollNo){
    this.rollNo=rollNo;
}

var View=function(){
    this.printStudentDetails=function(name,rollNo){
        console.log("Student :[name:"+name+",rollNo:"+rollNo+"]");
    }
}


var Controller=function(model,view){
    this.model=model;
    this.view=view;
}

Controller.prototype.getStudentName=function(){
    return this.model.getName();
}

Controller.prototype.setStudentName=function(name){
    this.model.setName(name);
}

Controller.prototype.getStudentRollNo=function(){
    return this.model.getRollNo();
}

Controller.prototype.setStudentRollNo=function(rollNo){
    this.model.setRollNo(rollNo);
}

Controller.prototype.updateView=function(){
    this.view.printStudentDetails(this.model.getName(),this.model.getRollNo())
}








