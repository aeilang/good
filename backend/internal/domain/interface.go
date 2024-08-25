package domain

import (
	"context"

	"github.com/aeilang/backend/internal/store/jobStore"
	"github.com/google/uuid"
)

type JobService interface {
	GetJobById(ctx context.Context, id uuid.UUID) (jobStore.GetJobByIdRow, error)
	GetJobs(ctx context.Context, filter jobStore.GetAllJobsParams) ([]jobStore.GetAllJobsRow, error)
	AddJob(ctx context.Context, job jobStore.CreateJobParams) error
}
