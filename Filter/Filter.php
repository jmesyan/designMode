<?php

/**
 * Created by PhpStorm.
 * User: jamesyan
 * Date: 2017/3/19
 * Time: 20:15
 */
class Person
{
    private $name, $gender, $maritalStatus;

    public function __construct($name, $gender, $maritalStatus)
    {
        $this->name = $name;
        $this->gender = $gender;
        $this->maritalStatus = $maritalStatus;
    }

    public function getName()
    {
        return $this->name;
    }

    public function getGender()
    {
        return $this->gender;
    }

    public function getMaritalStatus()
    {
        return $this->maritalStatus;
    }
}


interface Criteria
{
    public function meetCriteria($persons);
}

class CriteriaMale implements Criteria
{
    public function meetCriteria($persons)
    {
        $malePersons = array();
        // TODO: Implement meetCriteria() method.
        foreach ($persons as $person) {
            $gender = $person->getGender();
            $gender = strtolower($gender);
            if ($gender == "male") {
                array_push($malePersons, $person);
            }
        }
        return $malePersons;
    }

}


class CriteriaFemale implements Criteria
{
    public function meetCriteria($persons)
    {
        // TODO: Implement meetCriteria() method.
        $femalePersons = array();
        foreach ($persons as $person) {
            $gender = $person->getGender();
            $gender = strtolower($gender);
            if ($gender == "female") {
                array_push($femalePersons, $person);
            }
        }

        return $femalePersons;
    }
}


class CriteriaSingle implements Criteria
{
    public function meetCriteria($persons)
    {
        // TODO: Implement meetCriteria() method.
        $singlePersons = array();
        foreach ($persons as $person) {
            $maritalStatus = strtolower($person->getMaritalStatus());
            if ($maritalStatus == "single") {
                array_push($singlePersons, $person);
            }
        }
        return $singlePersons;
    }
}

class AndCriteria implements Criteria
{
    private $criteria;
    private $otherCriteria;

    public function __construct($criteria, $otherCriteria)
    {
        $this->criteria = $criteria;
        $this->otherCriteria = $otherCriteria;
    }

    public function meetCriteria($persons)
    {
        // TODO: Implement meetCriteria() method.
        $firstCriteriaPersons = $this->criteria->meetCriteria($persons);
        return $this->otherCriteria->meetCriteria($firstCriteriaPersons);
    }
}

class OrCriteria implements Criteria
{
    private $criteria;
    private $otherCriteria;

    public function __construct($criteria, $otherCriteria)
    {
        $this->criteria = $criteria;
        $this->otherCriteria = $otherCriteria;
    }

    public function meetCriteria($persons)
    {
        // TODO: Implement meetCriteria() method.
        $firstCriteriaItems = $this->criteria->meetCriteria($persons);
        $otherCriteriaItem = $this->criteria->meetCriteria($persons);
        foreach ($otherCriteriaItem as $person) {
            if (!in_array($person, $firstCriteriaItems)) {
                array_push($firstCriteriaItems, $person);
            }
        }
        return $firstCriteriaItems;
    }
}

function printPersons($persons)
{
    foreach ($persons as $person) {
        echo "Person : [ Name : " . $person->getName()
            . ", Gender : " . $person->getGender()
            . ", Marital Status : " . $person->getMaritalStatus()
            . " ]<br/>";
    }
}

//demo
$persons = array();
array_push($persons, new Person("Robert", "Male", "Single"));
array_push($persons, new Person("John", "Male", "Married"));
array_push($persons, new Person("Laura", "Female", "Married"));
array_push($persons, new Person("Diana", "Female", "Single"));
array_push($persons, new Person("Mike", "Male", "Single"));
array_push($persons, new Person("Bobby", "Male", "Single"));

$male = new CriteriaMale();
$female = new CriteriaFemale();
$single = new CriteriaSingle();
$singleMale = new AndCriteria($single, $male);
$singleOrFemale = new OrCriteria($single, $female);

var_dump("Males: ");
printPersons($male->meetCriteria($persons));

var_dump("Females: ");
printPersons($female->meetCriteria($persons));

var_dump("Single Males: ");
printPersons($singleMale->meetCriteria($persons));

var_dump("Single Or Females: ");
printPersons($singleOrFemale->meetCriteria($persons));


