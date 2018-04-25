// Package tribonacci contains functions to calculate tribonacci numbers.
package tribonacci

import (
	"errors"
	"math/big"
)

// ErrInvalidArg indicates that a argument value is out of valid range.
var ErrInvalidArg = errors.New("source argument 'n' is invalid")

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
// Function returns *big.Int contains tribonacci number with specified position.
// Function is interrupted if the channel 'quit' was closed. In this case, the
// function returns the current computed value. You should not use this value
// in any way, because it is wrong.
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

	matrixE = powerManaged(matrixE, n-3, quit)

	// T[0][0] contains the tribonacci number
	// so return it
	return matrixE[0][0], nil
}

type matrixElement struct {
	i, j  int
	value *big.Int
}

func multiply(matrixA, matrixB [3][3]*big.Int) [3][3]*big.Int {

	var n = 3

	channel := make(chan *matrixElement)

	result := [3][3]*big.Int{
		{big.NewInt(0), big.NewInt(0), big.NewInt(0)},
		{big.NewInt(0), big.NewInt(0), big.NewInt(0)},
		{big.NewInt(0), big.NewInt(0), big.NewInt(0)},
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			go caclAij(matrixA, matrixB, i, j, channel)
		}
	}

	for i := 0; i < n*n; i++ {
		matrixEl := <-channel
		result[matrixEl.i][matrixEl.j] = matrixEl.value
	}

	return result
}

// calcAij calculates A[i,j] element of matrix multiplication (A*B)
// and send matrixElement result into channel
func caclAij(matrixA, matrixB [3][3]*big.Int, i, j int, channel chan *matrixElement) {

	var n = 3

	Aij := big.NewInt(0)

	for k := 0; k < n; k++ {

		mulResult := new(big.Int).Mul(matrixA[i][k], matrixB[k][j])

		Aij.Add(Aij, mulResult)
	}

	result := matrixElement{i: i, j: j, value: Aij}

	channel <- &result
}

// powerManaged raises the matrix matrixA to the power n
// Function is interrupted if the channel 'quit' was closed.
func powerManaged(matrixA [3][3]*big.Int, n int, quit chan bool) [3][3]*big.Int {

	select {
	case <-quit:
		return matrixA
	default:
		if n == 0 || n == 1 {
			return matrixA
		}

		matrixA = powerManaged(matrixA, n/2, quit)

		matrixA = multiply(matrixA, matrixA)

		if n%2 != 0 {

			matrixE := [3][3]*big.Int{
				{big.NewInt(1), big.NewInt(1), big.NewInt(1)},
				{big.NewInt(1), big.NewInt(0), big.NewInt(0)},
				{big.NewInt(0), big.NewInt(1), big.NewInt(0)},
			}

			matrixA = multiply(matrixA, matrixE)
		}

		return matrixA
	}
}
