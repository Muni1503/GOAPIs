package bankApi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type BankApiReq struct {
	BankName   string `json:"bankName"`
	BranchName string `json:"branchName"`
	IfscCode   string `json:"ifscCode"`
	Address    string `json:"address"`
}

type BankApiResp struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

type Banks struct {
	BankName   string `json:"bankName"`
	BranchName string `json:"branchName"`
	IfscCode   string `json:"ifscCode"`
	Address    string `json:"address"`
}

var BankList = []Banks{
	{BankName: "Karnataka", BranchName: "Thiruvallur", IfscCode: "KARB000321", Address: "kakallur Byepass Thiruvallur"},
	{BankName: "Karnataka", BranchName: "Chennai", IfscCode: "KARB000432", Address: "No. 23, Anna Salai, Chennai"},
	{BankName: "Karnataka", BranchName: "Bangalore", IfscCode: "KARB000874", Address: "No. 78, MG Road, Bangalore"},
	{BankName: "Karnataka", BranchName: "Coimbatore", IfscCode: "KARB000541", Address: "Avinashi Road, Coimbatore"},
	{BankName: "SBI", BranchName: "Thiruvallur", IfscCode: "SBIN001234", Address: "Main Road, Thiruvallur"},
	{BankName: "SBI", BranchName: "Hyderabad", IfscCode: "SBIN000789", Address: "Charminar Road, Hyderabad"},
	{BankName: "HDFC", BranchName: "Pune", IfscCode: "HDFC000345", Address: "MG Road, Pune"},
	{BankName: "HDFC", BranchName: "Kolkata", IfscCode: "HDFC000678", Address: "Park Street, Kolkata"},
	{BankName: "ICICI", BranchName: "Mumbai", IfscCode: "ICIC000123", Address: "Nariman Point, Mumbai"},
	{BankName: "ICICI", BranchName: "Jaipur", IfscCode: "ICIC000456", Address: "JLN Marg, Jaipur"},
	{BankName: "ICICI", BranchName: "Ahmedabad", IfscCode: "ICIC000789", Address: "CG Road, Ahmedabad"},
}

func NewBankApi(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Access-control-Allow-origin","*")
	w.Header().Set("Access-Control-Allow-Methods","POST")
	w.Header().Set("Access-Control-Allow-Headers","Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	log.Println("NewBankApi(+)")

	var lBankApiReq BankApiReq
	var lBankApiResp BankApiResp

	//checking for incooming request is post or other
	if r.Method==http.MethodPost{
		//step1-->convert request body to byte
		lbody,lErr01:=io.ReadAll(r.Body);
		if lErr01!=nil{
			log.Println("BANBA:Error01",lErr01)
			lBankApiResp.Status="E"
			lBankApiResp.Msg=lErr01.Error()
		}else{
			//step2-->unmarshall the request body(byte ) to respective struct
			lErr02:=json.Unmarshal(lbody,&lBankApiReq)
			if lErr02!=nil{
				log.Println("BANBA:Error02",lErr02)
				lBankApiResp.Status="E"
				lBankApiResp.Msg=lErr02.Error()
			}else{
				//step3-->push the bank data into BankList
				log.Println(len(BankList))
				newBank :=Banks{
					BankName:   lBankApiReq.BankName,
					BranchName: lBankApiReq.BranchName,
					IfscCode:   lBankApiReq.IfscCode,
					Address:    lBankApiReq.Address,
				}
				BankList=append(BankList,newBank)
				lBankApiResp.Status="S"
				lBankApiResp.Msg="Bank inserted successfully"
				log.Println(len(BankList))
			}
		}
	}
	//step4-Marshall the Structure Which is return as an response
	lResp,lErr04:=json.Marshal(lBankApiResp)
	if lErr04!=nil{
		log.Println("BANBA:Error4",lErr04)
	}else{
		fmt.Fprintf(w,string(lResp))
	}
	log.Println("BankApi(-)")
}