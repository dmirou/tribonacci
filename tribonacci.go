// Package tribonacci contains functions to calculate tribonacci numbers.
package tribonacci

import (
	"math/big"
)

// Simple calculate tribonacci number with specified position (n)
// using Dynamic Programming.
//
// n int - natural integer contains tribonacci number position.
//
// Function returns *big.Int contains tribonacci number with specified position.
//
// Function use three variables to keep track of previous three numbers.
// Time complexity of this function is O(n).
func Simple(n int) *big.Int {

	nMinus3, nMinus2, nMinus1 := big.NewInt(0), big.NewInt(0), big.NewInt(1)

	firstThreeValues := [3]*big.Int{nMinus3, nMinus2, nMinus1}

	for i := 0; i < len(firstThreeValues); i++ {
		if n == i+1 {
			return firstThreeValues[i]
		}
	}

	nValue := calcNValue(nMinus1, nMinus2, nMinus3)

	for i := 4; i < n; i++ {

		nMinus3, nMinus2, nMinus1 = nMinus2, nMinus1, nValue

		nValue = calcNValue(nMinus1, nMinus2, nMinus3)
	}

	return nValue
}

func calcNValue(nMinus1, nMinus2, nMinus3 *big.Int) *big.Int {

	nValue := new(big.Int).Set(nMinus1)
	nValue.Add(nValue, nMinus2)
	nValue.Add(nValue, nMinus3)

	return nValue
}

// Matrix calculate tribonacci number with specified position (n)
// using matrix exponentiation.
//
// n int - natural integer contains tribonacci number position.
//
// Function returns *big.Int contains tribonacci number with specified position.
//
// Function use three variables to keep track of previous three numbers.
// Time complexity of this function is O(log n).
func Matrix(n int) *big.Int {

	if n == 1 || n == 2 {
		return big.NewInt(0)
	}

	matrixE := [3][3]*big.Int{
		{big.NewInt(1), big.NewInt(1), big.NewInt(1)},
		{big.NewInt(1), big.NewInt(0), big.NewInt(0)},
		{big.NewInt(0), big.NewInt(1), big.NewInt(0)},
	}

	matrixE = power(matrixE, n-3)

	// T[0][0] contains the tribonacci number
	// so return it
	return matrixE[0][0]
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

// power raises the matrix matrixA to the power n
func power(matrixA [3][3]*big.Int, n int) [3][3]*big.Int {

	if n == 0 || n == 1 {
		return matrixA
	}

	matrixA = power(matrixA, n/2)

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
