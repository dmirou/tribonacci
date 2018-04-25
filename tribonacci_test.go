package tribonacci

import (
	"testing"
)

var nToTribonacciValidMap = map[int]string{
	1:    "0",
	2:    "0",
	3:    "1",
	4:    "1",
	5:    "2",
	6:    "4",
	7:    "7",
	8:    "13",
	9:    "24",
	10:   "44",
	20:   "19513",
	50:   "1697490356184",
	100:  "28992087708416717612934417",
	1000: "443382579490226307661986241584270009256355236429858450381499235934108943134478901646797270328593836893366107162717822510963842586116043942479088674053663996392411782672993524690287662511197858910187264664163782145563472265666010074477859199789932765503984125240893",
}

func TestSimpleValidDataSource(t *testing.T) {

	for n, expectedTribonacci := range nToTribonacciValidMap {

		actualTribonacci, actualErr := Simple(n)

		checkTribonnaciErr(t, nil, actualErr)

		checkTribonnaci(t, expectedTribonacci, actualTribonacci.String())
	}
}

func TestMatrixManagedValidDataSource(t *testing.T) {

	for n, expectedTribonacci := range nToTribonacciValidMap {

		quit := make(chan bool)

		actualTribonacci, actualErr := MatrixManaged(n, quit)

		checkTribonnaciErr(t, nil, actualErr)

		checkTribonnaci(t, expectedTribonacci, actualTribonacci.String())
	}
}

func checkTribonnaciErr(t *testing.T, expected, actual error) {
	if expected != actual {
		t.Errorf("Expected tribonacci error \"%s\". Got \"%s\"\n", expected, actual)
	}
}

func checkTribonnaci(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Expected tribonacci %s. Got %s\n", expected, actual)
	}
}
