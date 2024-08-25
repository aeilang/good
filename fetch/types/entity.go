package types

type AddPostPayload struct {
	Href        string `json:"href"`
	CompanyName string `json:"company_name"`
	Title       string `json:"title"`
	City        string `json:"city"`
	Fulltime    bool   `json:"fulltime"`
	JobType     string `json:"job_type"`
	Description string `json:"description"`
	Requirement string `json:"requirement"`
	PriceDown   int32  `json:"price_down"`
	PriceUp     int32  `json:"price_up"`
}
