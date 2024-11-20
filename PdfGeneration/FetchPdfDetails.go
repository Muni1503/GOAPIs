package pdfGeneration

type FetchPdfDetails struct{
	BillingId int `json:"billingId"`
	ClientName string `json:"clientName"`
	TotalCharges float64 `json:"totalCharges"`
}

type FetchPdfResp struct{
	PdfDetails []FetchPdfDetails `json:"pdfDetails"`
	
}