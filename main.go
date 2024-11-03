package main

import (
	"log"
	"net/http"
	bankApi "proj_api/BankApis"
)

func main() {
	log.Println("Server Started...")

	http.HandleFunc("/checkuser",bankApi.SimpleApi)
	http.HandleFunc("/addBank",bankApi.NewBankApi)
	http.HandleFunc("/getAllBanks",bankApi.GetAllBankApi)
	http.HandleFunc("/updateBank",bankApi.UpdateBankApi)


	http.ListenAndServe(":29053",nil)
	log.Println("Server Stopped")
}