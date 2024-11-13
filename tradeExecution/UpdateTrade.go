package tradeexecution

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Trade struct {
	TradeId   int     `json:"tradeId"`
	StockName string  `json:"stockName"`
	Quantity  int     `json:"quantity"`
	TradeType string  `json:"tradeType"`
	Price     float64 `json:"price"`
	TradeDate string  `json:"tradeDate"`
	ActiveStatus string `json:"activeStatus"`
}

func UpdateTradeApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-control-Allow-origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("UpdateTradeApi(+)")

	var lTradeReq Trade
	var lUpdateTradeResp common.RespStruct


	if r.Method==http.MethodPost{
		lbody,lErr01:=io.ReadAll(r.Body);
		if lErr01!=nil{
			log.Println("BAUBA:Error01",lErr01)
			lUpdateTradeResp.Status="E"
			lUpdateTradeResp.Msg=lErr01.Error()
	}else{
		//Unmarshall the request body to respective struct
		lErr02:=json.Unmarshal(lbody,&lTradeReq)
			if lErr02!=nil{
				log.Println("BAUBA:Error02",lErr02)
				lUpdateTradeResp.Status="E"
				lUpdateTradeResp.Msg=lErr02.Error()

	}else{
		//Establish the dbConnection
		db, lErr03 := dbconnection.LocalDBConnect()
				if lErr03 != nil {
					log.Println("CCACA03", lErr03)
					lUpdateTradeResp.Status = common.Error
					lUpdateTradeResp.Msg = lErr03.Error()
				} else {
					defer db.Close()
					lErr04:=UpdateTrade(db,lTradeReq)

					if lErr04!=nil{
						log.Println("UTA04",lErr04)
						lUpdateTradeResp.Status = common.Error
						lUpdateTradeResp.Msg = lErr04.Error()
					}else{
						lUpdateTradeResp.Status = common.Success
						lUpdateTradeResp.Msg = "SucccessfullyUpdated"
					}
	}
}
	}
}else {
	lUpdateTradeResp.Status = common.Error
	lUpdateTradeResp.Msg = "Method not allowed"
}
lResp, lFErr := json.Marshal(lUpdateTradeResp)
if lFErr != nil {
	log.Println("CCACA04", lFErr)
} else {
	fmt.Fprint(w, string(lResp))
}
log.Println("UpdateTradeApi(-)")
}

func UpdateTrade(db *sql.DB,pTradeReq Trade)error{
	log.Println("UpdateTrade(+)")

	lUpdateString:=`update tradestable set stockName=?,quantity=?,tradeType=?,price=?,tradeDate=?,activeStatus=,updatedBy="Muni",updateAt=Now() where tradeId=?`

	lStmt,lErr01:=db.Prepare(lUpdateString)

	if lErr01!=nil{
		log.Println("DAUT01",lErr01)
		return lErr01
	}

	_,lErr02:=lStmt.Exec(pTradeReq.StockName,pTradeReq.Quantity,pTradeReq.TradeType,pTradeReq.Price,pTradeReq.TradeDate,pTradeReq.ActiveStatus,pTradeReq.TradeId)


	if lErr02!=nil{
		log.Println("DAUT02",lErr02)
		return lErr02
	}


	log.Println("UpdateTrade(-)")

	return nil
}