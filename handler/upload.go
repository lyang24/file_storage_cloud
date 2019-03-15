package handler

import (
	"net/http"
	"ioutil"
	"io/ioutil"
	"io"
)

func UploadHandler(w http.ResopnseWriter, r*http.Request){
	if r.Method=="GET"{
		data, err :=ioutil.ReadFile("./static/view/index.html")
		if err!=nil{
			io.WriteString(w, "internal server error")
			return
		}
		io.WriteString(w, string(data))			
	} else if r.Method=="POST"{
	
	}	
}
