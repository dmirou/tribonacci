// Package tribonachi contains functions to calculate tribonachi numbers.
package tribonachi

import (
	"math/big"
)

// Simple calculate tribonachi number with specified index (nIdex)
// using Dynamic Programming.
//
// nIndex *big.Int - natural integer contains tribonachi number position.
//
// Function returns *big.Int contains tribonachi number with specified position.
//
// Function use three variables to keep track of previous three numbers.
// Time complexity of this function is linear.
func Simple(nIndex *big.Int) *big.Int {

	nMinus3, nMinus2, nMinus1 := big.NewInt(0), big.NewInt(0), big.NewInt(1)

	firstThreeValues := [3]*big.Int{nMinus3, nMinus2, nMinus1}

	for i := 0; i < len(firstThreeValues); i++ {
		if nIndex.Cmp(big.NewInt(int64(i+1))) == 0 {
			return firstThreeValues[i]
		}
	}

	nValue := calcNValue(nMinus1, nMinus2, nMinus3)

	bigOne := big.NewInt(1)

	for i := big.NewInt(4); i.Cmp(nIndex) < 0; i.Add(i, bigOne) {

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
