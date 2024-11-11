package billinggeneration

import (
	"database/sql"
	"log"
	"net/http"
	"time"
)

type ClientTransaction struct  {
	TransactionDate time.Time `json:"transactionDate"`
	ClientId int `json:"clientId"`
	TradeId int `json:"tradeId"`
	StockName string `json:"stockName"`
	Quantity int `json:"quantity"`
	TradeType string `json:"tradeType"`
	
}

type GetClientTransactionsResp struct{
	ClientTransactions []ClientTransaction `json:"clientTransactions"`
	Status string `json:"status"`
	Msg string `json:"msg"`
}

func GetClientTransactionsApi(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-control-Allow-origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	log.Println("GetClientTransactions(+)")

	var lGetClientTransactionsResp GetClientTransactionsResp
	var lErr error
	
	if r.Method==http.MethodGet{
		db, lErr := dbconnection.LocalDBConnect()
				if lErr != nil {
					log.Println("BGGCTA01", lErr)
					lGetClientTransactionsResp.Status = common.Error
					lGetClientTransactionsResp.Msg = lErr03.Error()
				}else{
					
					lClientTransactions,lErr:=GetClientTransactions(db)
					if lErr!=nil{
						log.Println("BGGCTA01", lErr)
						lGetClientTransactionsResp.Status = common.Error
						lGetClientTransactionsResp.Msg = lErr03.Error()
					}else{
						lGetClientTransactionsResp.ClientTransactions=lClientTransactions
						lGetClientTransactionsResp.Status = common.Success
						lGetClientTransactionsResp.Msg = "ClientTransactions fetched successfully"

					}
				}
		

	}else{
		lGetClientTransactionsResp.Status=common.ErrStatus
		lGetClientTransactionsResp.Msg="Method Not Allowed"
	}
}

func GetClientTransactions(db *sql.DB)([]ClientTransaction,error){
	log.Println("GetClientTransactions(+)")
	var lClientTransactions []ClientTransaction
	
	lSelectString:=`select s.stockName,t.quantity,t.tradeType,t.clientId,t.tradeId from tradestable t join scriptMasterTable s on t.scriptId=s.scriptId`

	lStmt,lErr01:=db.Prepare(lSelectString)

	if lErr01!=nil{
		log.Println("BGGCT01",lErr01)
		return lClientTransactions,lErr01
	}

	lRows,lErr02:=lStmt.Query()

	if lErr02!=nil{
		log.Println("BGGCT02",lErr02)
		return lClientTransactions,lErr02
	}

	for lRows.Next(){
		var lClientTransaction ClientTransaction
		lErr03:=lRows.Scan(&lClientTransaction.StockName,&lClientTransaction.Quantity,&lClientTransaction.TradeType,&lClientTransaction.ClientId,&lClientTransaction.TradeId,&lClientTransaction.TradeType)

		if lErr03!=nil{
			log.Println("BGGCT03",lErr03)
			return lClientTransactions,lErr03
		}

		lClientTransactions=append(lClientTransactions, lClientTransaction)
	}
	log.Println("GetClientTransactions(-)")

	return lClientTransactions,nil


}