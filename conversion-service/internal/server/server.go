package server

import (
	"ConversionService/config"
	"ConversionService/internal/controller"
	"ConversionService/internal/converter"
	"ConversionService/internal/converter/custom_logger"
	"ConversionService/internal/converter/html_filler"
	"ConversionService/internal/converter/html_to_pdf_converter"
	"ConversionService/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Server struct {
	config      *config.Config
	logger      logger.Logger
	ginLogger   gin.HandlerFunc
	ginRecovery gin.HandlerFunc
}

func NewServer(config *config.Config, logger logger.Logger) *Server {
	return &Server{
		config:      config,
		logger:      logger,
		ginLogger:   custom_logger.GinLogger(logger),
		ginRecovery: custom_logger.GinRecovery(logger, true),
	}
}

func (receiver *Server) newEngine(httpClient html_to_pdf_converter.HttpClient) *gin.Engine {
	engine := receiver.initializeServer()
	htmlToPdfConverter := html_to_pdf_converter.NewHtmlToPdfConverter(receiver.config.HtmlToPdfConverter, receiver.logger, httpClient)
	htmlFiller := html_filler.NewHtmlFiller(receiver.config.Service.DebugMode, receiver.logger)
	converter := converter.NewConverter(htmlToPdfConverter, htmlFiller, *receiver.config, receiver.logger)
	controller := controller.NewController(converter, receiver.logger)

	receiver.logger.Infof("Conversion Service components initialized")

	engine.POST("/api/convert-pdf", controller.ConvertPdf)
	return engine
}

func (receiver *Server) Run() error {
	engine := receiver.newEngine(&http.Client{Timeout: receiver.config.HtmlToPdfConverter.Timeout * time.Second})

	err := engine.Run()
	if err != nil {
		receiver.logger.Fatalf("Fatal error ! SHUTTING DOWN !!", err)
		return err
	}

	return nil
}

func (receiver *Server) initializeServer() *gin.Engine {
	if receiver.config.Service.ReleaseMode == true {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(receiver.ginLogger, receiver.ginRecovery)
	serviceConfig := receiver.config.Service

	if len(serviceConfig.TrustedProxies) >= 1 {
		err := router.SetTrustedProxies(serviceConfig.TrustedProxies)
		if err != nil {
			receiver.logger.Errorf("Invalid proxies config !! Current config: %v", serviceConfig.TrustedProxies)
		}
	}
	receiver.logger.Infof("HTTP Server initialized")
	return router
}
