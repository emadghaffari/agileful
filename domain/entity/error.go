package entity

import "encoding/json"

type error struct{
	Message string
}

func Error(str string) []byte {
	bts,_:=json.Marshal(error{Message: str})
	return bts
}