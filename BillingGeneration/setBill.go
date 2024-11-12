package billinggeneration

import (
	"database/sql"
	"log"
)

type BillGenerationReq struct {
	ClientTransactionDetails ClientTransaction `json:"clientTransaction"`
	AdditionCharges          ExtraCharges      `json:"additionalCharges"`
	TransactionFees          float64           `json:"transactionFees"`
	TotalCharges             float64           `json:"totalCharges"`
}

func SetBill(db *sql.DB,pBillReq BillGenerationReq)error{
	log.Println("SetBill(+)")
	lInsertString:=`insert into billingtable (&backOfficeId,&brokerage,&stt,&dematCharges,&transactionFees,&totalAmount,&billingDate,&pdfGenerated) `

	lStmt,lErr01:=db.Prepare(lInsertString)

	if lErr01!=nil{
		log.Println("BGSB01",lErr01)
		log.Println("SetBill(-)")
		return lErr01
	}

	lStmt.Close()

	_,lErr02:=lStmt.Exec()

	if lErr02!=nil{
		log.Println("BGSB02",lErr02)
		log.Println("SetBill(-)")
		return lErr02
	}
	log.Println("SetBill(-)")
	return nil
}

func GetBackOfficeId(db *sql.DB,pTradeId int)(int,error){
	log.Println("GetBackOfficeId(+)")
	lSelstring:=`select backofficeId from backofficetable where tradeId=?`

	var lBackOfficeId int

	lStmt,lErr01:=db.Prepare(lSelstring)

	if lErr01!=nil{
		log.Println("BGGBOI01",lErr01)
		log.Println("GetBackOfficeId(-)")
		return lBackOfficeId,lErr01
	}
	defer lStmt.Close()
	lErr02:=lStmt.QueryRow(pTradeId).Scan(&lBackOfficeId)

	if lErr02!=nil{
		log.Println("BGGBOI02",lErr02)
		log.Println("GetBackOfficeId(-)")
		return lBackOfficeId,lErr02
	}
	log.Println("GetBackOfficeId(-)")
	return lBackOfficeId,nil
}