package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type JobPayload struct {
	Href         string `json:"href" validate:"required"`
	CompanyName  string `json:"company_name" validate:"required"`
	CompanyImage string `json:"company_image" validate:"required"`
	Title        string `json:"title" validate:"required"`
	Keyword      string `json:"keyword" validate:"required"`
	City         string `json:"city" validate:"required"`
	Fulltime     bool   `json:"fulltime" validate:"required"`
	JobType      string `json:"job_type" validate:"required"`
	Description  string `json:"description" validate:"required"`
	Requirement  string `json:"requirement" validate:"required"`
	PriceDown    int32  `json:"price_down" validate:"required"`
	PriceUp      int32  `json:"price_up" validate:"required"`
}

type Job struct {
	Href        string `json:"href"`
	CompanyName string `json:"company_name"`
	Title       string `json:"title"`
	City        string `json:"city"`
	Fulltime    bool   `json:"fulltime"`
	JobType     string `json:"job_type"`
	Description string `json:"description"`
	Requirement string `json:"requirement"`
	PriceDown   int    `json:"price_down"`
	PriceUp     int    `json:"price_up"`
}

func main() {
	data, err := os.ReadFile("./data.json")
	if err != nil {
		panic(err)
	}

	var jobs []Job
	if err := json.Unmarshal(data, &jobs); err != nil {
		panic(err)
	}

	client := http.Client{}

	for i, j := range jobs {
		fmt.Println("begin", i)
		pyload := JobPayload{
			Href:         j.Href,
			CompanyName:  j.CompanyName,
			CompanyImage: "https://sf1-lark-tos.f.mioffice.cn/obj/static-atsx-online-ee-tob/3c46b0f71765aa018256901bcff58378/e7714c27714b83e4c36ef45c69ee49dc6e44d32540880490f6af9a34f47d52f0.png",
			Title:        j.Title,
			Keyword:      "关键词1，关键词2，关键词3",
			City:         j.City,
			Fulltime:     j.Fulltime,
			JobType:      j.JobType,
			Description:  j.Description,
			Requirement:  j.Requirement,
			PriceDown:    int32(j.PriceDown),
			PriceUp:      int32(j.PriceUp),
		}

		body, err := json.Marshal(pyload)
		if err != nil {
			continue
		}

		req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/job", bytes.NewReader(body))
		if err != nil {
			continue
		}

		resp, err := client.Do(req)
		if err != nil || resp.StatusCode != http.StatusOK {
			continue
		}

		defer resp.Body.Close()

		fmt.Println("finished ", i)
	}
}
