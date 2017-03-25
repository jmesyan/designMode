<?php
/**
 * Created by PhpStorm.
 * User: jamesyan
 * Date: 2017/3/25
 * Time: 20:39
 */

class Student{
    private $rollNo,$name;
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

class StudentView{
    public function printStudentDetails($studentName,$studentRollNo){
        var_dump("Student: ");
        var_dump("Name: ".$studentName);
        var_dump("Roll No: ".$studentRollNo);
    }
}

class StudentController{
    private $model,$view;
    public function __construct($model,$view)
    {
        $this->model=$model;
        $this->view=$view;
    }

    public function setStudentName($name){
        $this->model->setName($name);
    }

    public function getStudentName(){
        return $this->model->getName();
    }

    public function setStudentRollNo($rollNo){
        $this->model->setRollNo($rollNo);
    }

    public function getStudentRollNo(){
        return $this->model->getRollNo();
    }

    public function updateView(){
        $this->view->printStudentDetails($this->model->getName(),$this->model->getRollNo());
    }
}

//demo
$student = new Student();
$student->setName("Robert");
$student->setRollNo("10");

$view=new StudentView();

$controller =new StudentController($student,$view);
$controller->updateView();
$controller->setStudentName("John");
$controller->updateView();