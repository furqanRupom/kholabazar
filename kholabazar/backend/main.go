package main

import (
	"fmt"
	"kholabazar/utils"
)

func main() {
	user := utils.Payload{
		Sub: 1,
		FirstName: "Furqan",
		Email: "furqanrupom978@gmail.com",
		LastName: "Ahmad",
		IsShopOwner: false,
	}
	jwt,err := utils.CreateJWT(user,"my-secret")
	if err != nil {
		fmt.Println("Something went wrong")
	}
	fmt.Println(jwt)

	// cmd.Serve()
}
