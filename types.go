package main

import "github.com/dvl-numeez/jwt"



type Register struct{
	Username string `json:"username"`
	Email string `json:"email"`
}
type  RegisterResponse struct{
	Message string `json:"message"`
	Token string `json:"token"`
}

type CustomClaims struct{
	Username string
	jwt.RegisteredClaims
}


type ApiServer struct{
	listenAddr string
}