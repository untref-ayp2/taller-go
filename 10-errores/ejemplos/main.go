package main

import "fmt"

type MyError struct {
	Code    int
	Message string
}

func (e MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, MyError{Code: 1, Message: "No se puede dividir por cero"}
	}
	return a / b, nil
}

func main() {
	res, err := Divide(10, 0)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("Resultado:", res)
}
