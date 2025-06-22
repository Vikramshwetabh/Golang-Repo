// file of file handling 

package auth

import "fmt"

func Login(username, password string) { //capitalize the first letter of the function name to make it exported
	fmt.Println("Logging in with username:", username, password)
}
