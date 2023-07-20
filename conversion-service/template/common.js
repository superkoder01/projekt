function toDate(object) {
    if (object === null || object === "" || object === undefined) {
        return ""
    }
    var date = new Date(object);
    return date.getDay() + "." + date.getMonth() + "." + date.getFullYear();
}

function toString(object) {
    if (object === null || object === undefined) {
        return ""
    }
    return object;
}

function addressToString(object) {
    if (object === null || object === undefined) {
        return ""
    }
    street = object.street;
    number = object.number;
    postCode = object.postCode;
    city = object.city;
    if (street === undefined && number === undefined && postCode === undefined && city === undefined) {
        return ""
    }
    return  street + " " + number + ", " + postCode + " " + city
}

function unitToPl(value, unit) {
    if (value === null || value === "" || value === undefined) {
        return ""
    }
    if (unit === null || unit === "" || unit === undefined) {
        return ""
    }
    if (unit == "month") {
        if (value == 1) {
            return value + " " + "miesiąc";
        } else if (value >= 2 && value <= 4) {
            return value + " " + "miesiące";
        } else if (value >= 5 && value <= 19) {
            return value + " " + "miesięcy";
        } else if (value % 1 == 0 || value % 2 == 0 || value % 3 == 0 || value % 4 == 0) {
            return value + " " + "miesiące";
        } else {
            return value + " " + "miesięcy";
        }
    } else {
        return value + " " + unit;
    }
}

function customerToString(object) {
    if (object === null || object === undefined) {
        return ""
    }
    firstName = object.firstName;
    lastName = object.lastName;
    if (firstName === undefined && lastName === undefined) {
        return ""
    }
    return  firstName + " " + lastName
}