// Package tribonacci contains functions to calculate tribonacci numbers.
package tribonacci

import (
	"errors"
	"math/big"
)

// ErrInvalidArg indicates that a argument value is out of valid range.
var ErrInvalidArg = errors.New("source argument 'n' is invalid")

// ErrCalcInterrupted indicates that function calculation was interrupted.
var ErrCalcInterrupted = errors.New("calculation was interrupted")

// Simple calculate tribonacci number with specified position (n)
// using Dynamic Programming.
//
// n int - natural integer contains tribonacci number position.
//
// Function returns *big.Int contains tribonacci number with specified position.
//
// Function use three variables to keep track of previous three numbers.
// Time complexity of this function is O(n).
func Simple(n int) (*big.Int, error) {

	if n <= 0 {
		return big.NewInt(0), ErrInvalidArg
	}

	nMinus3, nMinus2, nMinus1 := big.NewInt(0), big.NewInt(0), big.NewInt(1)

	firstThreeValues := [3]*big.Int{nMinus3, nMinus2, nMinus1}

	for i := 0; i < len(firstThreeValues); i++ {
		if n == i+1 {
			return firstThreeValues[i], nil
		}
	}

	nValue := calcNValue(nMinus1, nMinus2, nMinus3)

	for i := 4; i < n; i++ {

		nMinus3, nMinus2, nMinus1 = nMinus2, nMinus1, nValue

		nValue = calcNValue(nMinus1, nMinus2, nMinus3)
	}

	return nValue, nil
}

func calcNValue(nMinus1, nMinus2, nMinus3 *big.Int) *big.Int {

	nValue := new(big.Int).Set(nMinus1)
	nValue.Add(nValue, nMinus2)
	nValue.Add(nValue, nMinus3)

	return nValue
}

// MatrixManaged calculate tribonacci number with specified position (n)
// using matrix exponentiation.
//
// n int - natural integer contains tribonacci number position.
// quit chan bool - channel for interrupting the function
//
//
// If n is negative or zero function return 0 and invalid argument error.
// else if the function is complete, function return tribonnaci number with
// specified position as first argument, and nil as second.
//
// If the channel 'quit' was closed, then the function is interrupted. In this case, the
// function returns 0 as first argument and ErrCalcInterrupted as second.
//
// Function use matrix multiplication to calculate tribonacci number.
// Time complexity of this function is O(log n).
func MatrixManaged(n int, quit chan bool) (*big.Int, error) {

	if n <= 0 {
		return big.NewInt(0), ErrInvalidArg
	}

	if n == 1 || n == 2 {
		return big.NewInt(0), nil
	}

	matrixE := [3][3]*big.Int{
		{big.NewInt(1), big.NewInt(1), big.NewInt(1)},
		{big.NewInt(1), big.NewInt(0), big.NewInt(0)},
		{big.NewInt(0), big.NewInt(1), big.NewInt(0)},
	}

	var ok bool

	matrixE, ok = powerManaged(matrixE, n-3, quit)

	if !ok {
		return big.NewInt(0), ErrCalcInterrupted
	}

	// T[0][0] contains the tribonacci number
	// so return it
	return matrixE[0][0], nil
}

// powerManaged raises the matrix matrixA to the power n
// If second returned argument is true, first returned argument contains matrixA^n.
// else first returned argument contains current calculated matrix just as
// the channel 'quit' was closed.
func powerManaged(matrixA [3][3]*big.Int, n int, quit chan bool) ([3][3]*big.Int, bool) {

	select {
	case <-quit:
		return matrixA, false
	default:
		if n == 0 || n == 1 {
			return matrixA, true
		}

		var ok bool

		matrixA, ok = powerManaged(matrixA, n/2, quit)

		if !ok {
			return matrixA, ok
		}

		matrixA = multiply(matrixA, matrixA)

		if n%2 != 0 {

			matrixE := [3][3]*big.Int{
				{big.NewInt(1), big.NewInt(1), big.NewInt(1)},
				{big.NewInt(1), big.NewInt(0), big.NewInt(0)},
				{big.NewInt(0), big.NewInt(1), big.NewInt(0)},
			}

			matrixA = multiply(matrixA, matrixE)
		}

		return matrixA, true
	}
}

type matrixElement struct {
	i, j  int
	value *big.Int
}

func multiply(matrixA, matrixB [3][3]*big.Int) [3][3]*big.Int {

	var n = 3

	result := [3][3]*big.Int{
		{big.NewInt(0), big.NewInt(0), big.NewInt(0)},
		{big.NewInt(0), big.NewInt(0), big.NewInt(0)},
		{big.NewInt(0), big.NewInt(0), big.NewInt(0)},
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				mulResult := new(big.Int).Mul(matrixA[i][k], matrixB[k][j])
				result[i][j] = result[i][j].Add(result[i][j], mulResult)
			}
		}
	}

	return result
}
