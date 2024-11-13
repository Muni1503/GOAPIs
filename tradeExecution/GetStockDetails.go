package tradeexecution

import (
	"database/sql"
	"log"
)

type StockDetail struct{
	ScriptId int `json:"scriptId"`
	StockName string `json:"stockName"`
	Isin string `json:"isin"`
	SegmentCode string `json:"segmentCode"`
	Price float64 `json:"price"`
}

func GetStockDetails(db *sql.DB)([]StockDetail,error){
	log.Println("GetStockDetails(+)")

	var lStockDetails []StockDetail

	lSelString:=`select scriptId,stockName,isin,segmentCode,price from scriptMasterTable`

	lStmt,lErr01:=db.Prepare(lSelString)

	if lErr01!=nil{
		log.Println("TEGSD01",lErr01)
		return lStockDetails,lErr01
	}

	lRows,lErr02:=lStmt.Query()

	if lErr02!=nil{
		log.Println("TEGSD02",lErr02)
		return lStockDetails,lErr02
	}

	for lRows.Next(){
		var lStockDetail StockDetail
		lErr03:=lRows.Scan(&lStockDetail.ScriptId,&lStockDetail.StockName,&lStockDetail.Isin,&lStockDetail.SegmentCode,lStockDetail.Price)

		if lErr03!=nil{
			log.Println("TEGSD03",lErr03)
		return lStockDetails,lErr03
		}

		lStockDetails=append(lStockDetails, lStockDetail)
	}
	log.Println("GetStockDetails(-)")
	return lStockDetails,nil
}