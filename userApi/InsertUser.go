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
	"time"
)

type UserApireq struct{
	UserName string `json:"userId"` 
	Role string `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	Status string `json:"status"`
	ActiveStatus string `json:"activeStatus"`
}

func InsertUserApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-control-Allow-origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	log.Println("InsertUserApi(+)")

	var lresp common.RespStruct
	var lreq userApireq
	var lErr error

	if r.Method == http.MethodPost {
		//step1-->convert request body to byte
		lbody, lErr01 := io.ReadAll(r.Body)
		if lErr01 != nil {
			log.Println("UAIUA01", lErr01)
			lresp.Status = common.Error
			lresp.Msg = lErr01.Error()
		} else {
			//step2-->unmarshall the request body(byte) to respective struct
			lErr02 := json.Unmarshal(lbody, &lreq)
			if lErr02 != nil {
				log.Println("UAIUA02", lErr02)
				lresp.Status = common.Error
				lresp.Msg = lErr02.Error()
			} else {
				//step3-->call the method
				db, lErr03 := dbconnection.LocalDBConnect()
				if lErr03 != nil {
					log.Println("UAIUA03", lErr03)
					lresp.Status = common.Error
					lresp.Msg = lErr03.Error()
				} else {
					defer db.Close()
					lErr = AddUser(db, lreq)
					if lErr != nil {
						log.Println("UAIUA04", lErr)
						lresp.Status = common.Error
						lresp.Msg = lErr03.Error()
					} else {
						lresp.Status = common.Success
						lresp.Msg = "User Added successfully"
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
		log.Println("UAIUA04", lFErr)
	} else {
		fmt.Fprint(w, string(lResp))
	}
	log.Println("InsertUserApi(-)")
}

func AddUser(db *sql.DB, pUserDetails UserApireq) error {
	log.Println("AddUser(+)")


	lInsertString := `insert into config_users_table (userName,role,createdAt,status,createdBy,updatedBy,updatedAt,activeStatus) values(?,?,Now(),?,?,?,Now(),?)`

	_, lErr01 := db.Exec(lInsertString, pUserDetails.UserName, pUserDetails.Role, pUserDetails.Status,  pUserDetails.UserName,pUserDetails.UserName,pUserDetails.ActiveStatus)
	if lErr01 != nil {
		log.Println("BFUB02", lErr01)
		return lErr01
	}


	return nil
}