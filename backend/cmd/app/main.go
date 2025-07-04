package main

import (
	"backend/internal/delivery"
	"backend/internal/infra/postgres"
	"backend/internal/usecases"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	//repo := mdb.NewIssueRepository()
	repo := postgres.NewIssueRepository()
	usecase := usecases.NewIssueService(repo)
	handler := delivery.NewIssueHandler(usecase)

	r := gin.Default()
	r.Use(cors.Default())
	delivery.RegisterRoutes(r, handler)

	r.Run()
}
