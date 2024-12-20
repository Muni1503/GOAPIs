package billingapprover

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type BillApproverTradeDetails struct{
	BillingId int `json:"billingId"`
	ClientId int `json:"clientId"`
	TradeId int `json:"tradeId"`
	TotalAmount float64 `json:"total"`
}

type BillApprovalResp struct{
	BillApproverTradeDetailsArray []BillApproverTradeDetails `json:"billApproverTradeDetailsArray"`
	Status string `json:"status"`
	Msg string `json:"msg"`
}

func GetApprovaDetailsApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-control-Allow-origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("GetApprovaDetailsApi(+)")

	var lGetApprovalResp BillApprovalResp

	if r.Method==http.MethodGet{
		db, lErr01 := dbconnection.LocalDBConnect()
		if lErr01!=nil{
			log.Println("BAGADA01",lErr01)
			lGetApprovalResp.Status=common.Error
			lGetApprovalResp.Msg=lErr01.Error()
		}else{
			defer db.Close()

			lGetApproverDetailSArray,lErr02:=getAllBillDetails(db)

			if lErr02!=nil{
				log.Println("BAGADA02",lErr02)
			lGetApprovalResp.Status=common.Error
			lGetApprovalResp.Msg=lErr02.Error()
			}else{
				lGetApprovalResp.BillApproverTradeDetailsArray=lGetApproverDetailSArray
				lGetApprovalResp.Status=common.Success
				lGetApprovalResp.Msg="Fetched Successfully"
			}
		}

	}else{
		
			lGetApprovalResp.Status = common.Error
			lGetApprovalResp.Msg = "Method not allowed"
	}
	lResp, lFErr := json.Marshal(lGetApprovalResp)
	if lFErr != nil {
		log.Println("UAUBA04", lFErr)
	} else {
		fmt.Fprint(w, string(lResp))
	}
	log.Println("GetUserDetailsApi(-)")
}

func getAllBillDetails(db *sql.DB)([]BillApproverTradeDetails,error){
	log.Println("GetAllBillDetails(+)")

	var lBillApproverTradeDetailsArray []BillApproverTradeDetails

	

	lSelString:=`select b.billingId,bo.clientId,bo.tradeId,b.totalAmount from billingtable b join backofficetable bo on b.backOfficeId=bo.backOfficeId`

	lStmt,lErr01:=db.Prepare(lSelString)

	if lErr01!=nil{
		log.Println("BAGABD01",lErr01)
		return lBillApproverTradeDetailsArray,lErr01
	}

	lRows,lErr02:=lStmt.Query()

	if lErr02!=nil{
		log.Println("BAGABD02",lErr02)
		return lBillApproverTradeDetailsArray,lErr02
	}

	for lRows.Next(){
		var lBillApproverTradeDetails BillApproverTradeDetails
		lErr03:=lRows.Scan(&lBillApproverTradeDetails.BillingId,&lBillApproverTradeDetails.ClientId,lBillApproverTradeDetails.TradeId,lBillApproverTradeDetails.TotalAmount)

		if lErr03!=nil{
			log.Println("BAGABD03",lErr03)
		return lBillApproverTradeDetailsArray,lErr03
		}
		lBillApproverTradeDetailsArray=append(lBillApproverTradeDetailsArray, lBillApproverTradeDetails)

	}

	return lBillApproverTradeDetailsArray,nil

}