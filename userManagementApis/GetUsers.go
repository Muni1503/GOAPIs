package userApi

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type users struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
	Role     string `json:"rule"`
	Status   string `json:"status"`
}
type GetUserResp struct{
	UserDetails []users `json:"userDetails"`
	Status   string  `json:"status"`
	Msg      string  `json:"msg"`
}

func GetUserDetailsApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-control-Allow-origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("GetUserDetailsApi(+)")

	
	var lGetUserResp GetUserResp

	if r.Method==http.MethodGet{
		db, lErr01 := dbconnection.LocalDBConnect()
		if lErr01!=nil{
			log.Println("UAGUDA01",lErr01)
			lGetUserResp.Status=common.Error
			lGetUserResp.Msg=lErr01.error()
		}else{
			defer db.Close()

			lUsers,lErr02:=GetUsers(db)

			if lErr02!=nil{
				log.Println("UAGUDA02",lErr02)
				lGetUserResp.Status=common.Error
			lGetUserResp.Msg=lErr02.error()
			}else{
				lGetUserResp.UserDetails=lUsers
				lGetUserResp.Status=common.Success
				lGetUserResp.Msg="Fetched Successfully"
			}
			
		}

	}else{
		
			lGetUserResp.Status = common.Error
			lGetUserResp.Msg = "Method not allowed"
	}
	lResp, lFErr := json.Marshal(lGetUserResp)
	if lFErr != nil {
		log.Println("UAUBA04", lFErr)
	} else {
		fmt.Fprint(w, string(lResp))
	}
	log.Println("GetUserDetailsApi(-)")
}

func GetUsers(db *sql.DB)([]users,error){
	log.Println("GetUsers(+)")

	var lUsers []users

	lSelString:=`select userId,userName,role,status from configuserstable`

	lStmt,lErr01:=db.Prepare(lSelString)

	if lErr01!=nil{
		log.Println("UAGU01",lErr01)
		return lUsers,lErr01
	}

	lRows,lErr02:=lStmt.Query()

	if lErr02!=nil{
		log.Println("UAGU02",lErr02)
		return lUsers,lErr02
	}

	for lRows.Next(){
		var lUser users

		lErr03:=lRows.Scan(&lUser.UserId,&lUser.UserName,&lUser.Role,&lUser.Status)

		if lErr03!=nil{
			log.Println("UAGU03",lErr03)
		return lUsers,lErr03
		}
		lUsers = append(lUsers, lUser)
	}
	return lUsers,nil
}
