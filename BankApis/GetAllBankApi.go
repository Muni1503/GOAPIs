package bankApi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type GetAllBankApiResp struct {
	BanksArr []Banks `json:"bankArr"`
	Status   string  `json:"status"`
	Msg      string  `json:"msg"`
}

func GetAllBankApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-control-Allow-origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	log.Println("GetAllBankApi(+)")

	var lGetAllBankApiResp GetAllBankApiResp

	//check for the incoming Request

	if r.Method==http.MethodGet{
		for _,lBankDetails:=range BankList{
			lGetAllBankApiResp.BanksArr = append(lGetAllBankApiResp.BanksArr, lBankDetails)
		}
		lGetAllBankApiResp.Status="S"
		lGetAllBankApiResp.Msg="Successfully fetched"
	}

	lResp,lErr01:=json.Marshal(lGetAllBankApiResp)
	if lErr01!=nil{
		log.Println("ASA:Error4",lErr01)
	}else{
		fmt.Fprintf(w,string(lResp))
	}
	log.Println("GetAllBankApi(-)")

}