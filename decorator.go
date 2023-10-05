package main

import (
	"fmt"
	"time"
)

func DecoratorFactory(f func(...interface{}) (int, error), functionName string) func(...interface{}) (int, error) {
	return func(args ...interface{}) (int, error) {
		startTime := time.Now()
		fmt.Printf("Calling function %s with arguments %v\n", functionName, args)

		result, err := f(args...)

		endTime := time.Now()
		fmt.Printf("Function %s took %v to execute\n", functionName, endTime.Sub(startTime))
		fmt.Printf("Function %s returned result: %v, error: %v\n", functionName, result, err)

		return result, err
	}
}

func ExampleFunction(args ...interface{}) (int, error) {
	time.Sleep(1 * time.Second)
	return 42, nil
}

func main() {
	decoratedFunction := DecoratorFactory(ExampleFunction, "ExampleFunction")

	result, err := decoratedFunction("arg1", "arg2")

	fmt.Printf("Result: %d, Error: %v\n", result, err)
}
