// Package tribonachi contains functions to calculate tribonachi numbers.
package tribonachi

import (
	"math/big"
)

// Simple calculate tribonachi number with specified index (nIdex)
// using Dynamic Programming.
//
// nIndex int - natural integer contains tribonachi number position.
//
// Function returns *big.Int contains tribonachi number with specified position.
//
// Function use three variables to keep track of previous three numbers.
// Time complexity of this function is O(n).
func Simple(nIndex int) *big.Int {

	nMinus3, nMinus2, nMinus1 := big.NewInt(0), big.NewInt(0), big.NewInt(1)

	firstThreeValues := [3]*big.Int{nMinus3, nMinus2, nMinus1}

	for i := 0; i < len(firstThreeValues); i++ {
		if nIndex == i+1 {
			return firstThreeValues[i]
		}
	}

	nValue := calcNValue(nMinus1, nMinus2, nMinus3)

	for i := 4; i < nIndex; i++ {

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

// Matrix calculate tribonachi number with specified index (nIdex)
// using matrix exponentiation.
//
// nIndex int - natural integer contains tribonachi number position.
//
// Function returns *big.Int contains tribonachi number with specified position.
//
// Function use three variables to keep track of previous three numbers.
// Time complexity of this function is O(log n).
func Matrix(nIndex int) *big.Int {

	if nIndex == 1 || nIndex == 2 {
		return big.NewInt(0)
	}

	matrixE := [3][3]*big.Int{
		{big.NewInt(1), big.NewInt(1), big.NewInt(1)},
		{big.NewInt(1), big.NewInt(0), big.NewInt(0)},
		{big.NewInt(0), big.NewInt(1), big.NewInt(0)},
	}

	matrixE = power(matrixE, nIndex-3)

	// T[0][0] contains the tribonacci number
	// so return it
	return matrixE[0][0]
}

func multiply(matrixA, matrixB [3][3]*big.Int) [3][3]*big.Int {

	result := [3][3]*big.Int{
		{big.NewInt(0), big.NewInt(0), big.NewInt(0)},
		{big.NewInt(0), big.NewInt(0), big.NewInt(0)},
		{big.NewInt(0), big.NewInt(0), big.NewInt(0)},
	}

	var n = 3

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {

				mulResult := new(big.Int).Mul(matrixA[i][k], matrixB[k][j])

				result[i][j].Add(result[i][j], mulResult)
			}
		}
	}

	return result
}

// Recursive function to raise the matrix
// matrixA to the power n
func power(matrixA [3][3]*big.Int, n int) [3][3]*big.Int {

	if n == 0 || n == 1 {
		return matrixA
	}

	matrixA = power(matrixA, n/2)

	matrixA = multiply(matrixA, matrixA)

	if n%2 != 0 {

		M := [3][3]*big.Int{
			{big.NewInt(1), big.NewInt(1), big.NewInt(1)},
			{big.NewInt(1), big.NewInt(0), big.NewInt(0)},
			{big.NewInt(0), big.NewInt(1), big.NewInt(0)},
		}

		matrixA = multiply(matrixA, M)
	}

	return matrixA
}
