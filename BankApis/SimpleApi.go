package bankApi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type SimpleApiReq struct {
	EmailId  string `json:"emailId"`
	Password string `json:"password"`
}

type SimpleApiResp struct {
	RespData SimpleApiReq `json:"respData"`
	Status   string       `json:"status"`
	Msg      string       `json:"msg"`
}

type Client struct {
	EmailId  string `json:"emailId"`
	Password string `json:"password"`
}

var ClientList = []Client{
	{EmailId: "user1@example.com", Password: "password1"},
	{EmailId: "user2@example.com", Password: "password2"},
	{EmailId: "user3@example.com", Password: "password3"},
	{EmailId: "user4@example.com", Password: "password4"},
	{EmailId: "user5@example.com", Password: "password5"},
}

func SimpleApi(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Methods","POST")
	w.Header().Set("Access-Control-Allow-Headers","Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	log.Println("SimpleApi(+)")

	var lSimpleApiReq SimpleApiReq
	var lSimpleApiResp SimpleApiResp

	//checking for incooming request is post or other
	if r.Method==http.MethodPost{
		//step1-->convert request body to byte
		lbody,lErr01:=io.ReadAll(r.Body);
		if lErr01!=nil{
			log.Println("ASA:Error01",lErr01)
			lSimpleApiResp.Status="E"
			lSimpleApiResp.Msg=lErr01.Error()
		}else{
			//step2-->unmarshall the request body(byte ) to respective struct
			lErr02:=json.Unmarshal(lbody,&lSimpleApiReq)
			if lErr02!=nil{
				log.Println("ASA:Error02",lErr02)
				lSimpleApiResp.Status="E"
				lSimpleApiResp.Msg=lErr02.Error()
			}else{
				//step3-->Validate the emailid
				for _,lClient:=range ClientList{
					if lSimpleApiReq.EmailId==lClient.EmailId{
						lSimpleApiResp.RespData.EmailId=lClient.EmailId
						lSimpleApiResp.RespData.Password=lClient.Password
						lSimpleApiResp.Status="S"
						lSimpleApiResp.Msg="Success"
						break;
					}else{
						lSimpleApiResp.Status="E"
						lSimpleApiResp.Msg="Not Found"
					}
				}
			}
		}
	}
	//step5-->Marshall the Structure Which is return as an response
	lResp,lErr04:=json.Marshal(lSimpleApiResp)
	if lErr04!=nil{
		log.Println("ASA:Error4",lErr04)
	}else{
		fmt.Fprintf(w,string(lResp))
	}
	log.Println("SimpleApi(-)")
}