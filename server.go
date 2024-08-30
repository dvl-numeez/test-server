package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dvl-numeez/jwt"
)


func GetNewServer(listenAddr string)*ApiServer{
	return &ApiServer{
		listenAddr: listenAddr,
	}
}

func(s *ApiServer)Run(){
	
	router:= http.NewServeMux()
	router.HandleFunc("/",HandleOrigin)
	router.HandleFunc("/register",HandleRegister)
	router.HandleFunc("/admin",HandleAdmin)
	err:=http.ListenAndServe(s.listenAddr,router)
	fmt.Println("Server is running on ",s.listenAddr)
	if err!=nil{
		log.Fatal(err)
	}
	
}
func HandleOrigin(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(struct{Message string}{Message:"This is blank route"})
}

func HandleRegister(w http.ResponseWriter, r *http.Request){
	registerRequest:= new(Register)
	err:=json.NewDecoder(r.Body).Decode(registerRequest)
	if err!=nil{
		w.Header().Add("content-type","json/application")
	    w.WriteHeader(404)
		json.NewEncoder(w).Encode(struct{Error string}{Error:err.Error()})
		return
	}
	customClaims:=CustomClaims{
		Username: registerRequest.Username,
	}
	token,err:=jwt.GetToken(jwt.HMAC256,customClaims)
	if err!=nil{
		w.Header().Add("content-type","json/application")
	    w.WriteHeader(404)
		json.NewEncoder(w).Encode(struct{Error string}{Error:err.Error()})
		return
	}
	err=AddUser(*registerRequest)
	if err!=nil{
		w.Header().Add("content-type","json/application")
	    w.WriteHeader(404)
		json.NewEncoder(w).Encode(struct{Error string}{Error:err.Error()})
		return
	}

	response:=RegisterResponse{
		Message: "User created successfully",
		Token: token,
	}
	w.Header().Add("content-type","json/application")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(&response)

}

func HandleAdmin(w http.ResponseWriter, r *http.Request) {
	claims:=CustomClaims{}
	token:=r.Header.Get("Authorization")[7:]
	tokenData,err:=jwt.ValidateToken(token)
	if err!=nil{
		w.Header().Add("content-type","json/application")
	    w.WriteHeader(404)
		json.NewEncoder(w).Encode(struct{Error string}{Error:err.Error()})
		return

	}
	payloadData:=tokenData.Payload
	err=json.Unmarshal(payloadData,&claims)
	if err!=nil{
		w.Header().Add("content-type","json/application")
	    w.WriteHeader(404)
		json.NewEncoder(w).Encode(struct{Error string}{Error:err.Error()})
		return

	}
	if claims.Username=="Numeez"{
		w.Header().Add("content-type","json/application")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct{Secret string}{Secret:"123"})
		return
	}

	w.Header().Add("content-type","json/application")
	w.WriteHeader(403)
	json.NewEncoder(w).Encode(struct{Error string}{Error:"Access not authorized"})

	
	
}