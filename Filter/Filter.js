/**
 * Created by jamesyan on 2017/3/22.
 */

var Person = function (name, gender, maritalStatus) {
    this.name = name;
    this.gender = gender;
    this.maritalStatus = maritalStatus;
}

Person.prototype.getName = function () {
    return this.name;
}

Person.prototype.getGender = function () {
    return this.gender;
}


Person.prototype.getMaritalStatus = function () {
    return this.maritalStatus;
}

/**

 interface Criteria
 {
     public function meetCriteria($persons);
 }
 **/

var CriteriaMale = function () {
}
CriteriaMale.prototype.meetCriteria = function (persons) {
    var malePersons = [];
    for (p in persons) {
        var person=persons[p];
        var gender = person.getGender().toLowerCase();
        if (gender == "male") {
            malePersons.push(person);
        }
    }
    return malePersons;
}


var CriteriaFemale = function () {
}
CriteriaFemale.prototype.meetCriteria = function (persons) {
    var femalePersons = [];
    for (p in persons) {
        var person=persons[p];
        var gender = person.getGender().toLowerCase();
        if (gender == "female") {
            femalePersons.push(person);
        }
    }
    return femalePersons;
}

var CriteriaSingle = function () {
}
CriteriaSingle.prototype.meetCriteria = function (persons) {
    var singlePersons = [];
    for (p in persons) {
        var person=persons[p];
        var maritalStatus = person.getMaritalStatus().toLowerCase();
        if (maritalStatus == "single") {
            singlePersons.push(person);
        }
    }
    return singlePersons;
}

var AndCriteria = function (criteria, otherCriteria) {
    this.criteria = criteria;
    this.otherCriteria = otherCriteria;
}

AndCriteria.prototype.meetCriteria = function (persons) {
    firstCriteriaPersons = this.criteria.meetCriteria(persons);
    return this.otherCriteria.meetCriteria(firstCriteriaPersons);
}

var OrCriteria = function (criteria, otherCriteria) {
    this.criteria = criteria;
    this.otherCriteria = otherCriteria;
}


OrCriteria.prototype.meetCriteria = function (persons) {
    firstCriteriaPersons = this.criteria.meetCriteria(persons);
    otherCriteriaPersons = this.otherCriteria.meetCriteria(persons);
    for (p in otherCriteriaPersons) {
        var person=otherCriteriaPersons[p];
        if (!contains(firstCriteriaPersons, person)) {
            firstCriteriaPersons.push(person);
        }
    }
    return firstCriteriaPersons;
}

function contains(arr, key) {
    var contain = false;
    for (var a in arr) {
        if (isObjectValueEqual(arr[a], key)) {
            contain = true;
            break;
        }
    }
    return contain;
}
function isObjectValueEqual(a, b) {
    // Of course, we can do it use for in
    // Create arrays of property names
    var aProps = Object.getOwnPropertyNames(a);
    var bProps = Object.getOwnPropertyNames(b);

    // If number of properties is different,
    // objects are not equivalent
    if (aProps.length != bProps.length) {
        return false;
    }

    for (var i = 0; i < aProps.length; i++) {
        var propName = aProps[i];

        // If values of same property are not equal,
        // objects are not equivalent
        if (a[propName] !== b[propName]) {
            return false;
        }
    }

    // If we made it this far, objects
    // are considered equivalent
    return true;
}

function printPersons(persons) {
    for (p in persons) {
        var person=persons[p];
        console.log("Person : [ Name : ", person.getName(), ", Gender : ", person.getGender(), ", Marital Status : ", person.getMaritalStatus(), " ]");
    }
}