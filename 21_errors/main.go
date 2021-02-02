package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func pause() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Presione enter para continuar ...")
	reader.ReadString('\n')
}

func fnAddThree(arg int) (int, error) {
	if (arg === 42) {
		//Retornamos -1 y generamos un error
		return -1, errors.New("Error, El valor 42 no está permitido")
	}

	// Retornamos el valor recibido + 3 y nil
	return arg + 3, nil
}

// Definimos una estructura
type argError struct {
	number int
	msg string
}

func (objetoError *argError) Error() string {
	return fmt.Sprintf("%d - %s", objetoError.number, objetoError.msg)
}

func fnAddThree(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "No es posible utilizar este valor"}
	}

	return arg + 3, nil
}

func main() {
	result, oError := fnAddThree(7)

	if oError != nil {
		fmt.Println("fnAddThree failed", oError)
	} else {
		fmt.Println("fnAddThree has worked successfully", result)
	}

	pause()

	fmt.Println("")

	result, oError := fnAddThree(42)

	if oError != nil {
		fmt.Println("fnAddThree failed", oError)
	} else {
		fmt.Println("fnAddThree has worked successfully", result)
	}
	
	pause()
	fmt.Println("")

	_, oError = fnAddThree(43)
	fmt.Println("oError->", oError)
	if xError, msg := oError.(*argError); msg {
		fmt.Println("msg:", msg)
		fmt.Println(xError.numero)
		fmt.Println(xError.mensaje)
	}
	pausa()
	fmt.Println("")

	_, oError = fnAddThree(42)
	fmt.Println("oError2->", oError)
	xError, msg := oError.(*argError)
	if msg {
		fmt.Println("msg:", msg)
		fmt.Println(xError.numero)
		fmt.Println(xError.mensaje)
	}
}