package main

import "fmt"

type Student struct {
	name string
	rollNo int
}

func(s *Student)getName()string{
	return s.name
}

func(s *Student)setName(name string){
	s.name=name
}

func(s *Student)getRollNo()int{
	return s.rollNo
}

func(s *Student)setRollNo(rollNo int){
	s.rollNo=rollNo
}


type View struct {

}

func(v *View)printStudentDetails(name string,rollNo int){
	fmt.Println("Student:[name:",name,",rollNo:",rollNo,"]")
}


type Controller struct {
	model *Student
	view *View
}


func(s *Controller)getStudentName()string{
	return s.model.getName()
}

func(s *Controller)setStudentName(name string){
	s.model.setName(name)
}

func(s *Controller)getStudentRollNo()int{
	return s.model.getRollNo()
}

func(s *Controller)setStudentRollNo(rollNo int){
	s.model.setRollNo(rollNo)
}

func(s *Controller)updateView(){
	s.view.printStudentDetails(s.model.getName(),s.model.getRollNo())
}

func main(){
	studentModel:=new(Student)
	studentModel.setRollNo(10)
	studentModel.setName("Lucy")

	studentView:=new(View)

	studentController:=&Controller{studentModel,studentView}
	studentController.updateView()
	studentController.setStudentName("Joh")
	studentController.updateView()
}