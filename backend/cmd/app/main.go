package main

import (
	"backend/internal/config"
	"backend/internal/delivery"
	mdb "backend/internal/infra/MDB"
	"backend/internal/usecases"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewConfig()

	repo := mdb.NewIssueRepository()
	//repo := postgres.NewIssueRepository()
	usecase := usecases.NewIssueService(repo)
	handler := delivery.NewIssueHandler(usecase)

	r := gin.Default()
	r.Use(cors.Default())
	delivery.RegisterRoutes(r, handler)
	r.Run(cfg.BackendAPI)
}
