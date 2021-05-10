package data

import "encoding/json"

const PROJECT_ID = "auctionee"

type Credentials struct{
	Login		string `json:"login"`
	Password 	string `json:"password"`
}

type AuthResponse struct {
	Success 	string `json:"success"`
	Error	 	string `json:"error"`
}

func (a *AuthResponse)ToJSON(s string, err string)([]byte){
	a.Success = s
	a.Error = err
	res, _ :=  json.Marshal(a)
	return res
}