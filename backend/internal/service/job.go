package service

import (
	"context"

	"github.com/aeilang/backend/internal/store/jobStore"
	"github.com/google/uuid"
)

type JobService struct {
	store jobStore.Querier
}

func NewJobService(store jobStore.Querier) *JobService {
	return &JobService{
		store: store,
	}
}

func (j *JobService) GetJobById(ctx context.Context, id uuid.UUID) (jobStore.GetJobByIdRow, error) {
	return j.store.GetJobById(ctx, id)
}

func (j *JobService) GetJobs(ctx context.Context, filter jobStore.GetAllJobsParams) ([]jobStore.GetAllJobsRow, error) {
	return j.store.GetAllJobs(ctx, filter)
}

func (j *JobService) AddJob(ctx context.Context, job jobStore.CreateJobParams) error {
	return j.store.CreateJob(ctx, job)
}
