package billingapprover

import (
	"database/sql"
	"log"
)

type BillApproverTradeDetails struct{
	BillingId int `json:"billingId"`
	ClientId int `json:"clientId"`
	TradeId int `json:"tradeId"`
	TotalAmount float64 `json:"total"`
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