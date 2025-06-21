package main

import (
	"fmt"
	"image/color"

	"github.com/vikramshwetabh/podcast/auth"
)

func main() {
	auth.Login("admin", "password123")
	session := auth.Session()
	println("Session:", session)

	// user := user.UserId{
	// 	Email: "user@email.com",
	// 	// Name:  "John Doe",
	// }
	// fmt.Println("User Email:", user.Email)

	color.Red("sesion: %s", session)
	fmt.Println("This is main function")
}
