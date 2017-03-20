package main

import (
	"container/list"
	"strings"
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
	malePersons := list.New()
	for p := persons.Front(); p != nil; p = p.Next() {
		person := p.Value()
		gender := strings.ToLower(person.(Person).getGender())
		if (gender == "male") {
			malePersons.PushBack(person)
		}
	}
	return malePersons
}

type CriteriaFemale struct {

}

func (c *CriteriaFemale)meetCriteria(persons list.List) list.List {
	femalePersons := list.New()
	for p := persons.Front(); p != nil; p = p.Next() {
		person := p.Value()
		gender := strings.ToLower(person.(Person).getGender())
		if (gender == "female") {
			femalePersons.PushBack(person)
		}
	}
	return femalePersons
}

type CriteriaSingle struct {

}

func (c *CriteriaSingle)meetCriteria(persons list.List) list.List {
	singlePersons := list.New()
	for p := persons.Front(); p != nil; p = p.Next() {
		person := p.Value()
		maritalStatus := strings.ToLower(person.(Person).getMaritalStatus())
		if (maritalStatus == "single") {
			singlePersons.PushBack(person)
		}
	}
	return singlePersons
}

type AndCriteria struct {
	criteria Criteria
	otherCriteria Criteria
}

func(a *AndCriteria)meetCriteria(persons list.List) list.List{
	firstCriteriaPersons:=a.criteria.meetCriteria(persons)
	return a.otherCriteria.meetCriteria(firstCriteriaPersons)
}

type OrCriteria struct {
	criteria Criteria
	otherCriteria Criteria
}

func(o *OrCriteria)meetCriteria(persons list.List) list.List{
	firstCriteriaPersons:=o.criteria.meetCriteria(persons)
	otherCriteriaPersons:=o.otherCriteria.meetCriteria(persons)
	for p:=otherCriteriaPersons.Front();p!=nil;p=p.Next(){
		person:=p.Value()


	}
}


