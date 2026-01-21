package main

import "fmt"

const LoginToken string = "this could be anything."  // This is a public variable as the first letter of the variable is Capital.

func main() {
	var username string = "aryan"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var val int = 88888
	fmt.Println(val)
	fmt.Printf("Variable is of type: %T \n", val)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.987654345678756
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var float float64 = 255.987654345678756
	fmt.Println(float)
	fmt.Printf("Variable is of type: %T \n", float)

	//default value and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	//implicit type
	var website = "Hello everyone"
	fmt.Println(website)

	//no var style
    numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}