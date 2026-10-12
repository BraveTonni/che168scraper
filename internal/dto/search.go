package dto

type SearchCar struct {
	BrandName   string `json:"brandname"`
	CarName     string `json:"carname"`
	CertStatus  int    `json:"cert_status"`
	ClaimStatus int    `json:"claim_status"`
	CName       string `json:"c_name"`
	DealerId    int    `json:"dealer_id"`
	FuelName    string `json:"fuel_name"`
	ImageUrl    string `json:"image_url"`
	InfoId      int    `json:"info_id"`
	Mileage     string `json:"mileage"`
	MrStatus    int    `json:"mr_status"`
	Price       string `json:"price"`
	RegDate     string `json:"regdate"`
	SeriesName  string `json:"seriesname"`
	SpecName    string `json:"specname"`
	WeChatURL   string `json:"wechat_url"`
	WhatsAppURL string `json:"whatsapp_url"`
}

type SearchResult struct {
	CarList []SearchCar `json:"carlist"`

	CarsCount  int    `json:"pagecount"`
	PageIndex  int    `json:"pageindex"`
	PageSize   int    `json:"pagesize"`
	QueryId    string `json:"queryid"`
	TotalCount int    `json:"totalcount"`
}

type SearchResponse struct {
	Message    string        `json:"message"`
	Result     *SearchResult `json:"result"`
	ReturnCode int           `json:"returncode"`
}
