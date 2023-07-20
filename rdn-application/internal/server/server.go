package server

import (
	"RDN-application/internal/engine"
	"RDN-application/internal/ports"
	"RDN-application/internal/scheduler"
	"RDN-application/pkg/config"
	"RDN-application/pkg/logger"
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	scheduler        *scheduler.Scheduler
	dataService      ports.DataService
	notificationRepo ports.NotificationRepo
	collector        ports.DataCollector
	configuration    config.AppConfig
	logger           logger.Logger
}

func NewServer(cfg config.AppConfig, dataService ports.DataService, notificationRepo ports.NotificationRepo,
	collector ports.DataCollector, logger logger.Logger) *Server {
	return &Server{
		configuration:    cfg,
		scheduler:        scheduler.NewScheduler(cfg, logger),
		notificationRepo: notificationRepo,
		dataService:      dataService,
		collector:        collector,
		logger:           logger,
	}
}

func (s *Server) Run() error {
	s.logger.Info("Server started !")
	chanErr := make(chan error)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	ctx, shutdown := context.WithCancel(context.Background())

	go func() {
		s.logger.Infof("Server is listening on PORT: %s", s.configuration.GetServerConfig().Port)
		httpServer := s.StartHttpServer()
		s.MapHandlers(httpServer)
		s.logger.Infof("Http server initialized")
		err := httpServer.Run([]string{net.JoinHostPort("0.0.0.0", s.configuration.GetServerConfig().Port)}...)
		if err != nil {
			s.logger.Fatalf("Failed to start http server %v", err)
		}
	}()

	go func() {
		s.logger.Info("Scheduling task")
		task := engine.NewTaskProcessor(s.collector, s.dataService, s.notificationRepo, 3, s.configuration, s.logger)
		s.scheduler.AppendJob(task.ProcessCollectionTask(), ctx, chanErr)
		s.scheduler.ScheduleJob()
		task.WithFailureHandler(ctx, chanErr)
	}()

	sig := <-signals
	s.logger.Infof("Received signal %s", sig.String())
	s.logger.Infof("A graceful shutdown initiated")
	shutdown()

	return nil
}

func (s *Server) StartHttpServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(logger.GinLogger(s.logger), logger.GinRecovery(s.logger, true))

	if len(s.configuration.GetServerConfig().TrustedProxies) >= 1 {
		s.logger.Warnf("Trusted proxies set: %v", s.configuration.GetServerConfig().TrustedProxies)
		err := router.SetTrustedProxies(s.configuration.GetServerConfig().TrustedProxies)
		if err != nil {
			s.logger.Fatalf("Failed to set proxies: %v", s.configuration.GetServerConfig().TrustedProxies)
		}
	} else {
		err := router.SetTrustedProxies(nil)
		if err != nil {
			s.logger.Fatalf("Failed to initialize httpServer %v", err)
		}
	}

	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowAllOrigins = true

	router.Use(cors.New(config))

	return router
}
