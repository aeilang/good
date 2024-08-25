package handler

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/aeilang/backend/internal/domain"
	"github.com/aeilang/backend/internal/store/jobStore"
	"github.com/aeilang/backend/utils"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type JobHandler struct {
	srv domain.JobService
}

func NewJobHandler(srv domain.JobService) *JobHandler {
	return &JobHandler{
		srv: srv,
	}
}

func (j *JobHandler) Register(mux *http.ServeMux) {
	r := http.NewServeMux()
	r.HandleFunc("GET /job/{id}", j.GetJob)
	r.HandleFunc("GET /jobs", j.GetJobs)
	r.HandleFunc("POST /job", j.AddJob)

	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", r))
}

func (j *JobHandler) GetJob(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("uuid 解析错误")
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row, err := j.srv.GetJobById(ctx, uid)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("无法从数据库中取出id对应的post")
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, row)
}

func (j *JobHandler) GetJobs(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	pl := domain.GetJobsPayload{
		Search:      ParseString("search", q, ""),
		PriceDown:   ParseInt("price_down", q, 0),
		PriceUp:     ParseInt("price_up", q, 0),
		CompanyName: ParseString("company_name", q, ""),
		City:        ParseString("city", q, ""),
		JobType:     ParseString("job_type", q, ""),
		Page:        ParseInt("page", q, 1),
		PageSize:    ParseInt("page_size", q, 20),
		OrderBy:     ParseString("order_by", q, "created_at"),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	jobs, err := j.srv.GetJobs(ctx, jobStore.GetAllJobsParams{
		Limit:       int32(pl.PageSize),
		Offset:      int32((pl.Page - 1) * pl.PageSize),
		PriceDown:   int32(pl.PriceDown),
		PriceUp:     int32(pl.PriceUp),
		Companyname: pl.CompanyName,
		City:        pl.City,
		Jobtype:     pl.JobType,
		Title:       pl.Search,
		Oderby:      pl.OrderBy,
	})

	if err != nil {
		log.Warn().Err(err).Msg("数据库请求getJobs失败")
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, jobs)
}

func (j *JobHandler) AddJob(w http.ResponseWriter, r *http.Request) {
	job, err := utils.DecodeFromRequst[domain.AddJobPayload](r)
	if err != nil {
		log.Warn().Err(err).Msg("无法反序列化AddPostPayLoad")
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validator().Struct(job); err != nil {
		log.Info().Err(err).Msg("数据不匹配")
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	uid, err := uuid.NewRandom()
	if err != nil {
		log.Warn().Caller().Err(err).Msg("uuid无法生成")
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	params := jobStore.CreateJobParams{
		ID:           uid,
		Href:         job.Href,
		CompanyName:  job.CompanyName,
		CompanyImage: job.CompanyImage,
		Title:        job.Title,
		Keyword:      strings.Split(job.Keyword, "，"),
		City:         job.City,
		Fulltime:     job.Fulltime,
		JobType:      job.JobType,
		Description:  job.Description,
		Requirement:  job.Requirement,
		PriceDown:    job.PriceDown,
		PriceUp:      job.PriceUp,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = j.srv.AddJob(ctx, params)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("addJob无法加入数据库")
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"msg": "success",
	})
}
