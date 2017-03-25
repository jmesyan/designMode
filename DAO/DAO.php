<?php
/**
 * Created by PhpStorm.
 * User: jamesyan
 * Date: 2017/3/25
 * Time: 23:43
 */

class Student{
    public  $rollNo,$name;
    public function __construct($name,$rollNo)
    {
        $this->rollNo=$rollNo;
        $this->name=$name;
    }

    public function getRollNo(){
        return $this->rollNo;
    }

    public function setRollNo($rollNo){
        $this->rollNo=$rollNo;
    }

    public function getName(){
        return $this->name;
    }

    public function setName($name){
        $this->name=$name;
    }

}


interface StudentDao{
    public function getAllStudents();
    public function getStudent($rollNo);
    public function updateStudent($student);
    public function deleteStudent($student);
}

class StudentDaoImpl implements StudentDao{
    private $students;
    public function __construct()
    {
        $this->students=array();
        $this->students[]=new Student("Robert",0);
        $this->students[]=new Student("John",1);
    }

    public function getAllStudents()
    {
        // TODO: Implement getAllStudents() method.
        return $this->students;
    }

    public function getStudent($rollNo)
    {
        // TODO: Implement getStudent() method.
        foreach($this->students as $student){
            if($student->rollNo==$rollNo){
                return $student;
            }
        }

        return false;
    }

    public function updateStudent($student)
    {
        // TODO: Implement updateStudent() method.
        foreach($this->students as $key=>$stu){
            if($stu->rollNo==$student->rollNo){
                 $this->students[$key]->setName($student->getName());
                var_dump("Student: Roll No ".$student->getRollNo()
                    .", updated in the database");
                 break;
            }
        }

    }

    public function deleteStudent($student)
    {
        // TODO: Implement deleteStudent() method.
        foreach($this->students as $key=>$stu){
            if($stu->rollNo==$student->rollNo){
                unset($this->students[$key]);
                var_dump("Student: Roll No ".$student->getRollNo()
                    .", delete from the database");
                break;
            }
        }
    }


}

//demo
$studentDao=new StudentDaoImpl();
foreach($studentDao->getAllStudents() as $student){
    var_dump("Student: [RollNo : ".$student->getRollNo().", Name : ".$student->getName()." ]");
}

$student=$studentDao->getAllStudents()[0];
$student->setName("Michael");
$studentDao->updateStudent($student);

$student=$studentDao->getStudent(0);
var_dump("Student: [RollNo : ".$student->getRollNo().", Name : ".$student->getName()." ]");
