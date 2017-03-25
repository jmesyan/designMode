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

var StudentDao=function(){
    this.students=[];
    this.students[0]=new Student("Robert",0);
    this.students[1]=new Student("Lucy",1);
}
StudentDao.prototype.getAllStudents=function(){
    return this.students;
}

StudentDao.prototype.getStudent=function(rollNo){
    return this.students[rollNo];
}

StudentDao.prototype.updateStudent=function(student){
    if(this.students[student.getRollNo()]){
        this.students[student.getRollNo()].setName(student.getName())
        console.log("Student: Roll No "+student.getRollNo()+", updated in the database")
    }
}

StudentDao.prototype.deleteStudent=function(student){
   if(this.students[student.getRollNo()]){
       this.students.slice(student.getRollNo(),1);
   }
}



