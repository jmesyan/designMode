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


type StudentDao interface {
	getAllStudents() map[int]*Student;
	getStudent(rollNo int)*Student;
        updateStudent(student *Student);
	deleteStudent(student *Student);
}

type StudentDaoImpl struct {
	students map[int]*Student
}

func(s *StudentDaoImpl)init(){
	s.students=make(map[int]*Student)
	s.students[0]=&Student{"Robert",0}
	s.students[1]=&Student{"Lucy",1}
}

func(s *StudentDaoImpl)getAllStudents() map[int]*Student{
	return s.students
}

func(s *StudentDaoImpl)getStudent(rollNo int)*Student{
	if _,ok:=s.students[rollNo];ok{
		return s.students[rollNo]
	}
	return nil;
}

func(s *StudentDaoImpl)updateStudent(student *Student){
	if _,ok:=s.students[student.getRollNo()];ok{
		s.students[student.getRollNo()].setName(student.getName())
		fmt.Println("Student: Roll No ",student.getRollNo(),", updated in the database")
	}
}

func(s *StudentDaoImpl)deleteStudent(student *Student){
	if _,ok:=s.students[student.getRollNo()];ok{
		fmt.Println("Student: Roll No ",student.getRollNo(),", delete in the database")
		delete(s.students,student.getRollNo())
	}
}

func main(){
	studentDao:=new(StudentDaoImpl)
	studentDao.init()
	for _,student:=range studentDao.getAllStudents(){
		fmt.Println("Student: [RollNo : ",student.getRollNo(),", Name : ",student.getName()," ]");
	}

	student:=studentDao.getAllStudents()[0];
	student.setName("Michael");
	studentDao.updateStudent(student);

	student=studentDao.getStudent(0)
	fmt.Println("Student: [RollNo : ",student.getRollNo(),", Name : ",student.getName()," ]");

}

