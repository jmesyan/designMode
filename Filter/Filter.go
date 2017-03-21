package main

import (
	"container/list"
	"strings"
	"fmt"
)

type Person struct {
	name          string
	gender        string
	maritalStatus string
}

func (p *Person)getName() string {
	return p.name
}

func (p *Person)getGender() string {
	return p.gender
}

func (p *Person)getMaritalStatus() string {
	return p.maritalStatus
}

type Criteria interface {
	meetCriteria(persons list.List) list.List
}

type CriteriaMale struct {

}

func (c *CriteriaMale)meetCriteria(persons list.List) list.List {
	var malePersons list.List
	for p := persons.Front(); p != nil; p = p.Next() {
		person := p.Value.(*Person)
		gender := strings.ToLower(person.getGender())
		if (gender == "male") {
			malePersons.PushBack(person)
		}
	}
	return malePersons
}

type CriteriaFemale struct {

}

func (c *CriteriaFemale)meetCriteria(persons list.List) list.List {
	var femalePersons list.List
	for p := persons.Front(); p != nil; p = p.Next() {
		person := p.Value.(*Person)
		gender := strings.ToLower(person.getGender())
		if (gender == "female") {
			femalePersons.PushBack(person)
		}
	}
	return femalePersons
}

type CriteriaSingle struct {

}

func (c *CriteriaSingle)meetCriteria(persons list.List) list.List {
	var singlePersons list.List
	for p := persons.Front(); p != nil; p = p.Next() {
		person := p.Value.(*Person)
		maritalStatus := strings.ToLower(person.getMaritalStatus())
		if (maritalStatus == "single") {
			singlePersons.PushBack(person)
		}
	}
	return singlePersons
}

type AndCriteria struct {
	criteria      Criteria
	otherCriteria Criteria
}

func (a *AndCriteria)meetCriteria(persons list.List) list.List {
	firstCriteriaPersons := a.criteria.meetCriteria(persons)
	return a.otherCriteria.meetCriteria(firstCriteriaPersons)
}

type OrCriteria struct {
	criteria      Criteria
	otherCriteria Criteria
}

func (o *OrCriteria)meetCriteria(persons list.List) list.List {
	firstCriteriaPersons := o.criteria.meetCriteria(persons)
	otherCriteriaPersons := o.otherCriteria.meetCriteria(persons)
	for p := otherCriteriaPersons.Front(); p != nil; p = p.Next() {
		person := p.Value.(*Person)
		contains, _ := Contains(firstCriteriaPersons, person)
		if (!contains) {
			firstCriteriaPersons.PushBack(person)
		}

	}
	return firstCriteriaPersons
}

func Contains(l list.List, value *Person) (bool, *list.Element) {
	for e := l.Front(); e != nil; e = e.Next() {
		if (e.Value == value) {
			return true, e
		}
	}
	return false, nil
}

func printPersons(p list.List) {
	for e := p.Front(); e != nil; e = e.Next() {
		person := e.Value.(*Person)
		fmt.Println("Person : [ Name : ", person.getName(), ", Gender : ", person.getGender(), ", Marital Status : ", person.getMaritalStatus(), " ]")
	}
}

func main() {
	//demo
	var persons list.List;
	persons.PushBack(&Person{"Robert", "Male", "Single"})
	persons.PushBack(&Person{"John", "Male", "Married"})
	persons.PushBack(&Person{"Laura", "Female", "Married"})
	persons.PushBack(&Person{"Diana", "Female", "Single"})
	persons.PushBack(&Person{"Mike", "Male", "Single"})
	persons.PushBack(&Person{"Bobby", "Male", "Single"})

	male := new(CriteriaMale)
	fmt.Println("Males: ")
	printPersons(male.meetCriteria(persons))

	female := new(CriteriaFemale)
	fmt.Println("Females: ")
	printPersons(female.meetCriteria(persons))

	single := new(CriteriaSingle)

	singleMale := &AndCriteria{single, male}
	fmt.Println("singleMales: ")
	printPersons(singleMale.meetCriteria(persons))

	singleOrFemale := &OrCriteria{single, female}
	fmt.Println("single or females: ")
	printPersons(singleOrFemale.meetCriteria(persons))
}


