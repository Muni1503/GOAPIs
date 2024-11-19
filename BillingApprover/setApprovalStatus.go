package billingapprover

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type SetApprovalReq struct {
	BillingId      int    `json:"billingId"`
	ApprovedBy     int    `json:"ApprovedBy"` //id of the person who approves bill by this we get name
	ApprovalStatus string `json:"approvalStatus"`
	ApprovalDate   string `json:"approvalDate"`
	Comments       string `json:"comments"`
}

func SetApprovalApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-control-Allow-origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	log.Println("SetApprovalApi(+)")

	var lSetApprovalReq SetApprovalReq
	var lSetApprovalResp  common.RespStruct

	if r.Method==http.MethodPost{
		lbody,lErr01:=io.ReadAll(r.Body);
		if lErr01!=nil{
			log.Println("BAUBA:Error01",lErr01)
			lSetApprovalResp.Status="E"
			lSetApprovalResp.Msg=lErr01.Error()
	}else{
		//Unmarshall the request body to respective struct
		lErr02:=json.Unmarshal(lbody,&lSetApprovalReq)
			if lErr02!=nil{
				log.Println("BAUBA:Error02",lErr02)
				lSetApprovalResp.Status="E"
				lSetApprovalResp.Msg=lErr02.Error()

	}else{
		//Establish the dbConnection
		db, lErr03 := dbconnection.LocalDBConnect()
				if lErr03 != nil {
					log.Println("CCACA03", lErr03)
					lSetApprovalResp.Status = common.Error
					lSetApprovalResp.Msg = lErr03.Error()
				} else {
					defer db.Close()
					lErr04:=SetApprovalStatus(db,lSetApprovalReq)

					if lErr04!=nil{
						log.Println("CCACA04", lErr04)
					lSetApprovalResp.Status = common.Error
					lSetApprovalResp.Msg = lErr04.Error()
					}else {
						lresp.Status = common.Success
						lresp.Msg = "Approval status Added successfully"
					}
				}
			}
		}
	}else{
		
		lSetApprovalResp.Status = common.Error
		lSetApprovalResp.Msg = "Method not allowed"
}
lResp, lFErr := json.Marshal(lSetApprovalResp)
if lFErr != nil {
	log.Println("UAUBA04", lFErr)
} else {
	fmt.Fprint(w, string(lResp))
}
log.Println("SetApprovalApi(-)")
}

func SetApprovalStatus (db *sql.DB,pSetApprovalReq SetApprovalReq)error{
	log.Println("SetApprovalStatus(+)")

	lInsertString:=`Insert into approvalstable (billingId,approvedBy,approvalStatus,approvalDate,createdBy,updatedBy,updatedAt,activeStatus) values(?,?,?,?,?,Now(),?,Now())`

	lStmt,lErr01:=db.Prepare(lInsertString)

	if lErr01!=nil{
		log.Println("BASAS01",lErr01)
		return lErr01
	}

	_,lErr02:=lStmt.Exec(pSetApprovalReq.BillingId,pSetApprovalReq.ApprovedBy,pSetApprovalReq.ApprovalStatus,pSetApprovalReq.ApprovalDate,//ned to write createAt,updatedAt
		)
	
	if lErr02!=nil{
		log.Println("BASAS01",lErr02)
		return lErr02
	}	

	log.Println("SetApprovalStatus(-)")
	return nil
}

