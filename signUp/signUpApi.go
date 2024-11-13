package signup

import (
	common "backendApis/common"
	dbConnection "backendApis/dbConnection"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type signupApiReq struct {
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phoneNumber"`
	PanCardNumber     string `json:"panCardNumber"`
	NomineeName       string `json:"nomineeName"`
	BankName          string `json:"bankName"`
	BranchName        string `json:"branchName"`
	IfscCode          string `json:"ifscCode"`
	BankAccountNumber string `json:"bankAccountNumber"`
	Address           string `json:"address"`
	KycStatus         string   `json:"kycStatus"`
}
type signupApiResp struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

func SignupApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-control-Allow-origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	log.Println("signUpApi(+)")

	var lSignupApiReq signupApiReq
	var lSignupApiResp signupApiResp

	//step1-check for incoming request
	if r.Method == http.MethodPost {
		//step-2 --convert the body into byte
		lbody, lErr01 := io.ReadAll(r.Body)
		if lErr01 != nil {
			log.Println("SUSUA01", lErr01)
			lSignupApiResp.Status = common.Error
			lSignupApiResp.Msg = "Error while converting to byte"
		} else {
			//step-3 unmarshall the request body(byte) to struct
			lErr02 := json.Unmarshal(lbody, &lSignupApiReq)
			if lErr02 != nil {
				log.Println("SUSUA:02", lErr02)
				lSignupApiResp.Status = common.Error
				lSignupApiResp.Msg = "Error while unmarshalling"
			} else {
				//step3-Establish the db connection
				db, lErr03 := dbConnection.LocalDBConnect()
				if lErr03 != nil {
					log.Println("SUSUA:03", lErr03)
					lSignupApiResp.Status = common.Error
					lSignupApiResp.Msg = "Error while connecting to db"
				} else {
					defer db.Close()
					//step4-Execute the function
					lErr04:=Signup(db, lSignupApiReq)
					if lErr04!=nil{
						log.Println("SUSUA04",lErr04)
						lSignupApiResp.Status = common.Error
						lSignupApiResp.Msg = lErr04.Error()
				}else{
					lSignupApiResp.Status = common.Success
					lSignupApiResp.Msg="Successfully inserted"
				}

			}
		}
	}
	}else{
		lSignupApiResp.Status = common.Error;
		lSignupApiResp.Msg="Method Not Allowed"
	}
	lResp, lErr05 := json.Marshal(lSignupApiResp)
	if lErr05 != nil {
		log.Println("SUSUA:05", lErr05)
	} else {
		fmt.Fprint(w, string(lResp))
	}
	log.Println("LoginApi(-)")

}

func Signup(db *sql.DB, pSignupApiReq signupApiReq) error {
	log.Println("SignUp(+)")

	duplicateEmail,lErr01:=CheckEmailDuplicate(db,pSignupApiReq.Email)

	if lErr01!=nil{
		log.Println("lErr01",lErr01)
		return lErr01
	}

	if duplicateEmail{
		return errors.New("Email Already Exists")
	}
	lInsertString := `INSERT INTO clienttable (firstName,lastName,email,phoneNumber,panCard,nomineeName,bankId,bankAccount,kycCompleted,createdBy,createdAt,updatedBy,updatedAt,activeStatus) values(?,?,?,?,?,?,?,?,?,?,Now(),?,Now(),'y')`

	_, lErr0301 := db.Exec(lInsertString, pSignupApiReq.FirstName, pSignupApiReq.LastName, pSignupApiReq.Email, pSignupApiReq.PhoneNumber, pSignupApiReq.PanCardNumber, pSignupApiReq.NomineeName, getBankId(db, pSignupApiReq.IfscCode), pSignupApiReq.BankAccountNumber, getKyc(pSignupApiReq.KycStatus), pSignupApiReq.Email, pSignupApiReq.Email)
	if lErr0301 != nil {
		log.Println("SUSUA:0301", lErr0301)
		return lErr0301
	}
	log.Println("SignUp(-)")
	return nil
}

func getBankId(db *sql.DB, pIfscCode string) int {
	var id int
	sqlString := `select bank_id from bank_master_table where ifsc_code=?`

	lrows, lerr := db.Query(sqlString, pIfscCode)

	if lerr != nil {
		log.Println("Error while executing the query", lerr)
	} else {
		for lrows.Next() {
			lerr = lrows.Scan(&id)
			if lerr != nil {
				log.Println("Error while scanning the result table", lerr)
			}
		}
	}
	return id
}

func getKyc(kyc string) int {
	if kyc=="Completed" {
		return 1
	}
	return 0
}

func CheckEmailDuplicate(db sql.DB,pEmail string)(bool,error){
	log.Println("CheckEmailDuplicate(+)")
	lSelString:=`select count(*) from clienttable where emailId=?`

	lStmt,lErr01:=db.Prepare(lSelString)

	if lErr01!=nil{
		log.Println("SUCED01",lErr01)
		return true,lErr01
	}
	var lduplicateCount int

	lErr02:=lStmt.QueryRow(pEmail).Scan(&lduplicateCount)

	if lErr02!=nil{
		log.Println("SUCED02",lErr02)
		return true,lErr02
	}

	if lduplicateCount!=0{
		return true,nil
	}
	log.Println("CheckEmailDuplicate(-)")
	return false,nil
}