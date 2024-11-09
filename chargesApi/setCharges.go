package configchargesapis

import (
	"backendApis/common"
	dbconnection "backendApis/dbConnection"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type ChargesApiReq struct{
	ChargeType string `json:"chargeType"`
	ChargePercentage float64 `json:"chargePercentage"`
	EffectiveDate time.Time `json:"effectiveDate"`
	ChangingUser string `json:"changingUser"`
	ActiveStatus string `json:"activeStatus"`
}

func ChargesApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-control-Allow-origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	log.Println("ChargesApi(+)")

	var lresp common.RespStruct
	var lreq ChargesApiReq
	var lErr error

	if r.Method == http.MethodPost {
		//step1-->convert request body to byte
		lbody, lErr01 := io.ReadAll(r.Body)
		if lErr01 != nil {
			log.Println("CCACA01", lErr01)
			lresp.Status = common.Error
			lresp.Msg = lErr01.Error()
		} else {
			//step2-->unmarshall the request body(byte) to respective struct
			lErr02 := json.Unmarshal(lbody, &lreq)
			if lErr02 != nil {
				log.Println("CCACA02", lErr02)
				lresp.Status = common.Error
				lresp.Msg = lErr02.Error()
			} else {
				//step3-->call the method
				db, lErr03 := dbconnection.LocalDBConnect()
				if lErr03 != nil {
					log.Println("CCACA03", lErr03)
					lresp.Status = common.Error
					lresp.Msg = lErr03.Error()
				} else {
					defer db.Close()
					lErr = AddOrUpdate(db,lreq)
					if lErr != nil {
						log.Println("CCACA04", lErr)
						lresp.Status = common.Error
						lresp.Msg = lErr03.Error()
					} else {
						lresp.Status = common.Success
						lresp.Msg = " successfully done"
					}
				}

			}

		}
	} else {
		lresp.Status = common.Error
		lresp.Msg = "Method not allowed"
	}
	lResp, lFErr := json.Marshal(lresp)
	if lFErr != nil {
		log.Println("CCACA04", lFErr)
	} else {
		fmt.Fprint(w, string(lResp))
	}
	log.Println("ChargesApi(-)")
}

func InsertCharge(db *sql.DB,pChargesReq ChargesApiReq)error{
	log.Println("InsertCharge(+)")

	lInsertString:=`Insert into configChargesTable (chargeType,chargePercentage,effectiveDate,createdBy,createdAt,updatedBy,updatedAt) values(?,?,?,?,Now(),?,Now())`

	stmt,lerr01:=db.Prepare(lInsertString);

	if lerr01!=nil{
		log.Println("CCAIC01",lerr01)
		return lerr01
	}

	_,lerr02:=stmt.Exec(pChargesReq.ChargeType,pChargesReq.ChargePercentage,pChargesReq.EffectiveDate,pChargesReq.ChangingUser,pChargesReq.ChangingUser)

	if lerr02!=nil{
		log.Println("CCAIC02",lerr02)
		return lerr02
	}

	return nil
}

func UpdateCharge(db *sql.DB,pChargesReq ChargesApiReq)error{

	log.Println("UpdateCharge(+)")

	lUpdateString:=`update configChargesTable set chargePercentage=?,effectiveDate=?,activeStatus=? ,updatedBy=?,updatedAt=Now()where chargeType=? `

	stmt,lerr01:=db.Prepare(lUpdateString);

	if lerr01!=nil{
		log.Println("CCAUC01",lerr01)
		return lerr01
	}

	_,lErr02:=stmt.Exec(pChargesReq.ChargePercentage,pChargesReq.EffectiveDate,pChargesReq.ActiveStatus,pChargesReq.ChangingUser,pChargesReq.ChargeType)

	if lErr02!=nil{
		log.Println("CCAUC02",lErr02)
		return lErr02
	}

	return nil

}

func AddOrUpdate(db *sql.DB,preq ChargesApiReq)error{
	log.Println("AddOrUpdate(+)")

	var count int

	lSqlString:=`select count(*) from configChargesTable where chargeType=?`

	stmt,lErr01:=db.Prepare(lSqlString)

	if lErr01!=nil{
		log.Println("CCAAOU01",lErr01)
		return lErr01
	}

	lErr02:=stmt.QueryRow(preq.ChargeType).Scan(&count)

	if lErr02!=nil{
		log.Println("CCAAOU02",lErr02)
		return lErr02
	}

	if count==0{
		return InsertCharge(db,preq)
	}else{
		return UpdateCharge(db,preq)
	}


}