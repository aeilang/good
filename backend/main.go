package main

import (
	"context"
	"database/sql"

	"net/http"
	"time"

	"github.com/aeilang/backend/config"
	"github.com/aeilang/backend/internal/handler/v1"
	"github.com/aeilang/backend/internal/service"
	"github.com/aeilang/backend/internal/store/jobStore"
	"github.com/aeilang/backend/logger"
	"github.com/aeilang/backend/middleware"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func main() {
	logger.StartLogger()
	cfg := config.GetConfig()

	// 连接postgres数据库
	db, err := sql.Open("postgres", cfg.DBSource)
	if err != nil || db.Ping() != nil {
		log.Fatal().Caller().Err(err).Msg("db connecting failed!")
	}

	// 初始化Prepare
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	jobQuery, err := jobStore.Prepare(ctx, db)

	if err != nil {
		log.Fatal().Caller().Err(err).Msg("create queries failed")
	}

	server := service.NewJobService(jobQuery)

	jobHandler := handler.NewJobHandler(server)

	mux := http.NewServeMux()
	jobHandler.Register(mux)

	srv := http.Server{
		Addr:    cfg.HTTPServerAddress,
		Handler: chain(mux, middleware.TimeOut(cfg.HTTPTimeOut), middleware.Logger, middleware.CORS),
	}

	log.Info().Msgf("listen to port %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal().Caller().Err(err).Msg("listen procese failed")
	}
}

func chain(mux http.Handler, fns ...func(http.Handler) http.Handler) http.Handler {
	for _, fn := range fns {
		mux = fn(mux)
	}
	return mux
}
