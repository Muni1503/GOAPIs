package bankApi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)


type UpdateBankApiResp struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

func UpdateBankApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-control-Allow-origin","*")
	w.Header().Set("Access-Control-Allow-Methods","POST")
	w.Header().Set("Access-Control-Allow-Headers","Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	log.Println("UpdateBankApi(+)")

	var lUpdateBankApiReq Banks
	var lUpdateBankApiResp UpdateBankApiResp

	//checking for incoming request

	if r.Method==http.MethodPost{
		lbody,lErr01:=io.ReadAll(r.Body);
		if lErr01!=nil{
			log.Println("BAUBA:Error01",lErr01)
			lUpdateBankApiResp.Status="E"
			lUpdateBankApiResp.Msg=lErr01.Error()
	}else{
		//Unmarshall the request body to respective struct
		lErr02:=json.Unmarshal(lbody,&lUpdateBankApiReq)
			if lErr02!=nil{
				log.Println("BAUBA:Error02",lErr02)
				lUpdateBankApiResp.Status="E"
				lUpdateBankApiResp.Msg=lErr02.Error()

	}else{
		//change the data with respect to ifsc code
		for i:=range BankList{
			if BankList[i].IfscCode==lUpdateBankApiReq.IfscCode{
				BankList[i].BankName=lUpdateBankApiReq.BankName
				BankList[i].BranchName=lUpdateBankApiReq.BranchName
				BankList[i].Address=lUpdateBankApiReq.Address
				lUpdateBankApiResp.Status="S"
				lUpdateBankApiResp.Msg="Successfully updated"
				break
			}else{
				lUpdateBankApiResp.Status="E"
				lUpdateBankApiResp.Msg="Error:IFSC Code not Found"
			}
		}
		
	}
	}
}
//step5-marshall the structure

lResp,lErr02:=json.Marshal(lUpdateBankApiResp)
if lErr02!=nil{
	log.Println("BAUBA:Error2",lErr02)
}else{
	fmt.Fprintf(w,string(lResp))
}
log.Println("UpdateBankApi(-)")
}