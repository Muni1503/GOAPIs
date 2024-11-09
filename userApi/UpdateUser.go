package userApi

import (
	"backendApis/common"
	dbconnection "backendApis/dbConnection"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type UpdateUSerApiReq struct{
	UserId int `json:"userId"`
	UserApiReqInstance UserApireq `json:"userApiReqInstance"`
}
func UpdateBankApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-control-Allow-origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	log.Println("InsertBankApi(+)")

	var lresp common.RespStruct
	var lreq UpdateUSerApiReq
	var lErr error

	if r.Method == http.MethodPost {
		//step1-->convert request body to byte
		lbody, lErr01 := io.ReadAll(r.Body)
		if lErr01 != nil {
			log.Println("UAUUA01", lErr01)
			lresp.Status = common.Error
			lresp.Msg = lErr01.Error()
		} else {
			//step2-->unmarshall the request body(byte) to respective struct
			lErr02 := json.Unmarshal(lbody, &lreq)
			if lErr02 != nil {
				log.Println("UAUUA02", lErr02)
				lresp.Status = common.Error
				lresp.Msg = lErr02.Error()
			} else {
				//step3-->call the method
				db, lErr03 := dbconnection.LocalDBConnect()
				if lErr03 != nil {
					log.Println("UAUUA03", lErr03)
					lresp.Status = common.Error
					lresp.Msg = lErr03.Error()
				} else {
					defer db.Close()
					lErr = UpdateUser(db, lreq)
					if lErr != nil {
						log.Println("UAUUA04", lErr)
						lresp.Status = common.Error
						lresp.Msg = lErr03.Error()
					} else {
						lresp.Status = common.Success
						lresp.Msg = "Bank updated successfully"
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
		log.Println("UAUBA04", lFErr)
	} else {
		fmt.Fprint(w, string(lResp))
	}
	log.Println("UpdateBankApi(-)")
}

func UpdateUser(db *sql.DB, pUpdateUserDetails UpdateUSerApiReq) error {
	log.Println("UpdateBank(+)")


	lUpdateString := `Update config_users_table set userName=?,role=?,status=?, updatedBy='admin',updatedAt=Now() where userId=?`

	_, lErr01 := db.Exec(lUpdateString, pUpdateUserDetails.UserApiReqInstance.UserName,pUpdateUserDetails.UserApiReqInstance.Role ,pUpdateUserDetails.UserApiReqInstance.Status,  pUpdateUserDetails.UserId)
	if lErr01 != nil {
		log.Println("UAUA02", lErr01)
		return lErr01
	}


	return nil
}