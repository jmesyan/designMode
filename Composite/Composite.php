<?php

/**
 * Created by PhpStorm.
 * User: jamesyan
 * Date: 2017/3/22
 * Time: 21:46
 */
class Employee
{
    private $name, $dept, $salary;
    private $subordinates;

    public function __construct($name, $dept, $sal)
    {
        $this->name = $name;
        $this->dept = $dept;
        $this->salary = $sal;
        $this->subordinates = array();
    }

    public function add($e)
    {
        $this->subordinates[] = $e;
    }

    public function remove($e)
    {
        foreach ($this->subordinates as $key => $ordinate) {
            if ($ordinate == $e) {
                unset($this->subordinates[$key]);
                break;
            }
        }
    }

    public function getSubordinates()
    {
        return $this->subordinates;
    }


    public function printEmployees()
    {
        echo "Employee:[ Name : " . $this->name . ", dept : " . $this->dept . ", salary :" . $this->salary . " ]<br/>";
    }

}

//demo
$CEO = new Employee("John", "CEO", 30000);

$headSales = new Employee("Robert", "Head Sales", 20000);

$headMarketing = new Employee("Michel", "Head Marketing", 20000);

$clerk1 = new Employee("Laura", "Marketing", 10000);
$clerk2 = new Employee("Bob", "Marketing", 10000);

$salesExecutive1 = new Employee("Richard", "Sales", 10000);
$salesExecutive2 = new Employee("Rob", "Sales", 10000);

$CEO->add($headSales);
$CEO->add($headMarketing);

$headSales->add($salesExecutive1);
$headSales->add($salesExecutive2);

$headMarketing->add($clerk1);
$headMarketing->add($clerk2);

//打印该组织的所有员工
$CEO->printEmployees();
foreach ($CEO->getSubordinates() as $headEmployee) {
    $headEmployee->printEmployees();
    foreach ($headEmployee->getSubordinates() as $employee) {
        $employee->printEmployees();
    }
}