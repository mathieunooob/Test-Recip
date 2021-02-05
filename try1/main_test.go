package main

import "testing"

// func TestAdditionOK(t *testing.T) {
//    var functReturn int
//    var param2 int

//    param1 := 3
//    param2 = 4

//    functReturn = monAddition(param1, param2)

//    if functReturn <= 0 {
// 	   t.Error("monAddition return is <= 0")
//    }
// }

// func TestAdditionFuncNOK(t *testing.T) {
// 	var functReturn int
// 	var param2 int
 
// 	param1 := 5
// 	param2 = 5
 
// 	functReturn = monAddition(param1, param2)
 
// 	if functReturn != 10 {
// 		t.Error("Fonctional ERROR : monAddition don't return 10")
// 	}
// }

// func TestSubtractionOK(t *testing.T) {
// 	var functReturn int
// 	param1 := 15
// 	param2 := 5

// 	functReturn = maSoustraction(param1, param2)

// 	if functReturn <= 0 {
// 		t.Error("maSoustraction's return is lower than 0")
// 	}
// }

// func TestSubtractionFuncNOK(t *testing.T) {
// 	var functReturn int
// 	param1 := -6
// 	param2 := 5

// 	functReturn = maSoustraction(param1, param2)

// 	if functReturn != param1-param2 {
// 		t.Error("maSoustraction's return bad result")
// 	}
// }

// func TestMultiplicationOK(t *testing.T) {
// 	var functReturn int
// 	param1 := 10
// 	param2 := -5

// 	functReturn = maMultiplication(param1, param2)

// 	if functReturn <= 0 {
// 		t.Error("maMultiplication's return is lower than 0")
// 	}
// }

// func TestMultiplicationFuncNOK(t *testing.T) {
// 	var functReturn int
// 	param1 := 10
// 	param2 := -5

// 	functReturn = maMultiplication(param1, param2)

// 	if functReturn != param1*param2 {
// 		t.Error("maMultiplication's result is false - The function has not been modified properly")
// 	}
// }

func TestDivisionOK(t *testing.T) {
	var functReturn int
	param1 := 10
	param2 := 5

	functReturn = maDivision(param1, param2)

	if functReturn <= 0 {
		t.Error("maDivision's return is not allowed")
	}
}


func TestDivisionFuncNOK(t *testing.T) {
	var functReturn int
	param1 := 10
	param2 := 3

	functReturn = maDivision(param1, param2)

	if functReturn != param1/param2 {
		t.Error("maDivision's result is false - The function has not been modified properly")
	}
}


