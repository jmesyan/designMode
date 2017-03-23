/**
 * Created by Jmesyan on 2017/3/23.
 */

var Employee=function(name,dept,salary){
    this.name=name;
    this.dept=dept;
    this.salary=salary
    this.subordinates=[]
}

Employee.prototype.add=function(Employee){
    this.subordinates.push(Employee)
}

Employee.prototype.remove=function(Employee){
    for(p in this.subordinates){
       var per=this.subordinates[p];
        if(per==Employee){
             this.subordinates.slice(p,1)
             break;
        }
    }
}
Employee.prototype.getSubordinates=function(){
    return this.subordinates;
}

Employee.prototype.printEmployee=function(){
    console.log("Employee:[ Name : " ,this.name, ", dept : ",this.dept, ", salary :" ,this.salary , " ]");
}







