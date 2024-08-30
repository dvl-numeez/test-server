package main

import (
	"encoding/json"
	"os"
)

func AddUser(user Register)error{
	var users []Register
	fileName:="db.json"
	data,err:=os.ReadFile(fileName)
	if err!=nil{
		return err
	}
	err=json.Unmarshal(data,&users)
	if err!=nil{
		return err
	}
	users = append(users, user)
	jsonData,err:=json.MarshalIndent(users,"","  ")
	if err!=nil{
		return err
	}
	err=os.WriteFile(fileName,jsonData,os.ModePerm)
	if err!=nil{
		return err
	}

	return nil

}
