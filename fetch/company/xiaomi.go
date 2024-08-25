package company

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aeilang/fetch/types"
)

type Host struct {
	Company   string
	Href      string
	URL       string
	Body      string
	Limit     int
	Offset    int
	Count     int
	Result    []types.AddPostPayload
	ReteLimit time.Duration
}

// url := fmt.Sprintf("https://xiaomi.jobs.f.mioffice.cn/api/v1/search/job/posts?keyword=&limit=%d&offset=%d&job_category_id_list=&tag_id_list=&location_code_list=&subject_id_list=&recruitment_id_list=&portal_type=6&job_function_id_list=&portal_entrance=1&_signature=O3PUmgAAAABMbbN38dBVcjtz1IAAF3L", limit, offset)

// bodyStr := fmt.Sprintf(`{"keyword":"","limit":%d,"offset":%d,"job_category_id_list":[],"tag_id_list":[],"location_code_list":[],"subject_id_list":[],"recruitment_id_list":[],"portal_type":6,"job_function_id_list":[],"portal_entrance":1}`, limit, offset)

func (h *Host) getReqeust() *http.Request {
	url := fmt.Sprintf(h.URL, h.Limit, h.Offset)

	bodyStr := fmt.Sprintf(h.Body, h.Limit, h.Offset)

	// 创建一个新的请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(bodyStr)))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// 设置请求头
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "device-id=; locale=zh-CN; channel=saas-career; platform=pc; s_v_web_id=verify_lzohnh56_FotoSIbG_JcBw_4mDw_8twU_EH4qR9GQ5Vsl; atsx-csrf-token=p8kzZL_clf7mZTHSemoQ0OX568cwDxu9j4z5sBXrvBY%3D")
	req.Header.Set("Env", "undefined")
	req.Header.Set("Origin", "https://xiaomi.jobs.f.mioffice.cn")
	req.Header.Set("Portal-Channel", "saas-career")
	req.Header.Set("Portal-Platform", "pc")
	req.Header.Set("Priority", "u=1, i")
	req.Header.Set("Referer", "https://xiaomi.jobs.f.mioffice.cn/index/?keywords=&category=&location=&project=&type=&job_hot_flag=&current=1&limit=10&functionCategory=&tag=")
	req.Header.Set("Sec-CH-UA", `"Not)A;Brand";v="99", "Google Chrome";v="127", "Chromium";v="127"`)
	req.Header.Set("Sec-CH-UA-Mobile", "?0")
	req.Header.Set("Sec-CH-UA-Platform", `"Linux"`)
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36")
	req.Header.Set("Website-Path", "index")
	req.Header.Set("X-Csrf-Token", "p8kzZL_clf7mZTHSemoQ0OX568cwDxu9j4z5sBXrvBY=")

	return req
}

// Href:        fmt.Sprintf(`https://xiaomi.jobs.f.mioffice.cn/index/position/%s/detail`, job.Id),

func (h *Host) transToAddPostPayload(job JobPost) types.AddPostPayload {
	pl := types.AddPostPayload{
		Href:        fmt.Sprintf(h.Href, job.Id),
		CompanyName: h.Company,
		Title:       job.Title,
		City:        job.City.Name,
		Fulltime:    job.Recruit.Fulltime == "全职",
		JobType:     job.Recruit.JobType.Name,
		Description: job.Description,
		Requirement: job.Requirement,
	}

	return pl
}

func (h *Host) Pull() {
	client := http.Client{}
	for i := 1; h.Offset < h.Count; i++ {
		req := h.getReqeust()
		resp, err := client.Do(req)
		if err != nil {
			log.Println(h.Company, i, "faild", err)
			h.Offset = i * h.Limit
			time.Sleep(h.ReteLimit)
			continue
		}
		defer resp.Body.Close()
		var p Payload
		if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
			log.Println("failed to Decoder json:", i, err)
			h.Offset = i * h.Limit
			time.Sleep(h.ReteLimit)
			continue
		}

		for _, job := range p.Data.JobPostList {
			h.Result = append(h.Result, h.transToAddPostPayload(job))
		}

		h.Offset = i * h.Limit
		h.Count = p.Data.Count
		log.Println("finish ", h.Company, i)
		time.Sleep(h.ReteLimit)
	}
}

func (h *Host) Push() {
	if len(h.Result) == 0 {
		return
	}

	client := http.Client{}
	for i, job := range h.Result {
		body, err := json.Marshal(job)
		if err != nil {
			log.Printf("pushing %d failed", i)
			continue
		}
		req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/post", bytes.NewBuffer(body))
		if err != nil {
			log.Printf("faield %d", i)
			continue
		}
		req.Header.Set("Authorization", "ILoveFSY")
		resp, err := client.Do(req)
		if err != nil || resp.StatusCode != http.StatusOK {
			log.Printf("failed to push %d", i)
			continue
		}

		resp.Body.Close()
		fmt.Printf("push %d sucessful", i)
	}
}
func (h *Host) ToFile() {
	data, err := json.Marshal(h.Result)
	if err != nil {
		return
	}

	os.WriteFile("./data.json", data, 0644)
}

type Payload struct {
	Data Data `json:"data"`
}

type Data struct {
	Count       int       `json:"count"`
	JobPostList []JobPost `json:"job_post_list"`
}

type JobPost struct {
	Id          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Requirement string  `json:"requirement"`
	City        City    `json:"city_info"`
	Recruit     Recruit `json:"recruit_type"`
}

type City struct {
	Name string `json:"name"`
}

type Recruit struct {
	Fulltime string `json:"name"`
	JobType  Parent `json:"parent"`
}

type Parent struct {
	Name string `json:"name"`
}
