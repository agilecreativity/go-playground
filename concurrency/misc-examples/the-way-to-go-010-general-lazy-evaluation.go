// Chapter 14: page 418
package main

import "fmt"

type Any interface{}
type EvalFunc func(Any) (Any, Any)

func BuildLazyIntEvaluator(evalFunc EvalFunc, initState Any) func() Any {
	retValChan := make(chan Any)
	loopFunc := func() {
		var actState Any = initState
		var retVal Any
		for {
			retVal, actState = evalFunc(actState)
			retValChan <- retVal
		}
	}

	retFunc := func() Any {
		return <-retValChan
	}

	go loopFunc()
	return retFunc
}
func main() {
	evenFunc := func(state Any) (Any, Any) {
		os := state.(int)
		ns := os + 2
		return os, ns
	}

	even := BuildLazyIntEvaluator(evenFunc, 0)

	for i := 0; i < 5; i++ {
		fmt.Printf("%vth even: %v\n", i, even())
	}
}

/* Output:

0th even: 0
1th even: 2
2th even: 4
3th even: 6
4th even: 8

*/
