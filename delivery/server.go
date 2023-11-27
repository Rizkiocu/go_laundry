package delivery

import (
	"fmt"
	"go_laundry/config"
	"go_laundry/delivery/controller/api"
	"go_laundry/delivery/middleware"
	"go_laundry/manager"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	ucManager manager.UsecaseManager
	engine    *gin.Engine
	host      string
	log       *logrus.Logger
}

func (s *Server) Run() {
	s.initControllers()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}
func (s *Server) initControllers() {
	s.engine.Use(middleware.LogRequestMiddleware(s.log))
	rg := s.engine.Group("/api/v1")
	// Inisialisasi Controller
	api.NewUomController(s.ucManager.UomUC(), rg).Route()
	api.NewEmployeeController(s.ucManager.EmployeeUC(), rg).Route()
	api.NewAuthController(s.ucManager.UserUC(), s.ucManager.AuthUC(), rg).Route()
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}
	infraManager, err := manager.NewInfraManager(cfg)
	if err != nil {
		fmt.Println(err)
	}

	rm := manager.NewRepoManager(infraManager)
	uom := manager.NewUseCaseManager(rm)

	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	log := logrus.New()

	engine := gin.Default()
	server := Server{
		ucManager: uom,
		engine:    engine,
		host:      host,
		log:       log,
	}

	return &server
}
