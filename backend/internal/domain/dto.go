package domain

type AddJobPayload struct {
	Href         string `json:"href" validate:"required"`
	CompanyName  string `json:"company_name" validate:"required"`
	CompanyImage string `json:"company_image" validate:"required"`
	Title        string `json:"title" validate:"required"`
	Keyword      string `json:"keyword" validate:"required"`
	City         string `json:"city"`
	Fulltime     bool   `json:"fulltime" validate:"required"`
	JobType      string `json:"job_type" validate:"required"`
	Description  string `json:"description" validate:"required"`
	Requirement  string `json:"requirement" validate:"required"`
	PriceDown    int32  `json:"price_down"`
	PriceUp      int32  `json:"price_up"`
}

type GetJobsPayload struct {
	Search      string
	PriceDown   int
	PriceUp     int
	CompanyName string
	City        string
	JobType     string
	Page        int
	PageSize    int
	OrderBy     string
}
