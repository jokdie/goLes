package main

import "jun3_1/service"

func main() {
	userS := service.UserService{}
	userS.CreateUser("Tomas")

	admS := service.AdminService{}
	admS.CreateAdmin("Yarik")
	admS.Logger.Log("first log")
	admS.Log("second log")
}
