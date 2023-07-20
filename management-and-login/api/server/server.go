package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/handler"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/handler/impl"
	conf "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/configuration"
	"net"
)

var (
	logger = logging.MustGetLogger("http_server")
)

var (
	authHandler                      impl.AuthHandler
	providerHandler                  handler.Handler
	workerHandler                    handler.Handler
	workerHandlerImpl                impl.WorkerHandler
	workerUserHandler                handler.Handler
	userHandler                      handler.Handler
	userHandlerImpl                  impl.UserHandler
	customerAccountHandler           handler.Handler
	workerCustomerAccountHandlerImpl impl.WorkerCustomerAccountHandler
	customerUserHandler              handler.Handler
	serviceAccessPointHandler        handler.Handler
	serviceAccessPointHandlerImpl    impl.ServiceAccessPointHandler
)

type HTTPServer interface {
	Run()
}

type httpServer struct {
	hf     handler.HandlerFactory
	engine *gin.Engine
}

func NewHttpServer(hf handler.HandlerFactory) *httpServer {
	return &httpServer{hf: hf}
}

func (h *httpServer) Run() {
	httpConf := conf.GetHttpConfig()
	h.engine = gin.New()
	h.engine.RouterGroup.Group(httpConf.ApiPrefix)

	h.engine.Use(logRequest)
	h.engine.Use(attachHeaders)
	h.engine.Use(handleOptions)
	h.engine.Use(verifyRBAC)

	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowAllOrigins = true
	h.engine.Use(cors.New(config))
	h.initializeHandlers()
	h.setUpRoutes(httpConf.ApiPrefix)

	h.engine.Run(net.JoinHostPort("0.0.0.0", httpConf.Port))
}

func (h *httpServer) setUpRoutes(basePath string) {
	g := h.engine.Group(basePath)

	// AUTH
	// /authenticate
	g.POST("/authenticate", authHandler.Authenticate)
	g.POST("/refresh", authHandler.RefreshToken)
	g.POST("/logout", authHandler.Logout)

	// ACTIVATE
	// /activate/:activationCode
	g.PUT("/activate/:activationCode", userHandlerImpl.Activate)
	g.PUT("/resetPassword/:resetCode", userHandlerImpl.ResetPassword)
	g.GET("/forgotPassword/:email", userHandlerImpl.ForgotPassword)
	g.GET("/activationLink/:customerAccountId", userHandlerImpl.SendActivationLink)

	// PROVIDERS
	// /providers
	g.POST("/providers", providerHandler.Create)
	g.GET("/providers", providerHandler.List)
	// /providers/check
	g.GET("/providers/check", providerHandler.CheckIfExistWithFilter)
	// /providers/:providerId
	g.GET("/providers/:providerId", providerHandler.GetByID)
	g.PUT("/providers/:providerId", providerHandler.UpdateByID)
	g.DELETE("/providers/:providerId", providerHandler.DeleteByID)
	// /providers/details
	g.GET("/providers/details", providerHandler.GetDetails)
	// /providers/:providerId/users/administrators
	g.GET("/providers/:providerId/administrators", userHandlerImpl.ListAdministrators)
	g.GET("/providers/query", providerHandler.Query)

	// USERS
	// /users
	g.GET("/users", userHandler.List)
	g.POST("/users", userHandler.Create)
	// /users/check
	g.GET("/users/check", userHandler.CheckIfExistWithFilter)
	// /users/:userId
	g.GET("/users/:userId", userHandler.GetByID)
	g.PUT("/users/:userId", userHandler.UpdateByID)
	g.DELETE("/users/:userId", userHandler.DeleteByID)
	// /users/superAdmins
	g.GET("/users/superAdmins", userHandlerImpl.ListSuperAdmins)
	// /users/details
	g.GET("/users/details", userHandler.GetDetails)
	g.GET("/users/query", userHandler.Query)

	// CUSTOMER_ACCOUNTS
	// /customerAccounts
	g.GET("/customerAccounts", customerAccountHandler.List)
	g.POST("/customerAccounts", customerAccountHandler.Create)
	// /customerAccounts/check
	g.GET("/customerAccounts/check", customerUserHandler.CheckIfExistWithFilter)
	// /customerAccounts/:customerAccountId
	g.GET("/customerAccounts/:customerAccountId", customerAccountHandler.GetByID)
	g.PUT("/customerAccounts/:customerAccountId", customerAccountHandler.UpdateByID)
	g.DELETE("/customerAccounts/:customerAccountId", customerAccountHandler.DeleteByID)
	// /customerAccounts/details
	g.GET("/customerAccounts/details", customerAccountHandler.GetDetails)
	g.GET("/customerAccounts/query", workerCustomerAccountHandlerImpl.ListWorkerCustomerAccounts)
	// /customerAccounts/serviceAccessPoints
	g.GET("/customerAccounts/:customerAccountId/serviceAccessPoints", serviceAccessPointHandlerImpl.ListForCustomer)

	// SERVICE_ACCESS_POINTS
	// /serviceAccessPoints
	g.GET("/serviceAccessPoints", serviceAccessPointHandler.List)
	g.POST("/serviceAccessPoints", serviceAccessPointHandler.Create)
	// /serviceAccessPoints/check
	g.GET("/serviceAccessPoints/check", serviceAccessPointHandler.CheckIfExistWithFilter)
	// /serviceAccessPoints/:serviceAccessPointId
	g.GET("/serviceAccessPoints/:serviceAccessPointId", serviceAccessPointHandler.GetByID)
	g.PUT("/serviceAccessPoints/:serviceAccessPointId", serviceAccessPointHandler.UpdateByID)
	g.DELETE("/serviceAccessPoints/:serviceAccessPointId", serviceAccessPointHandler.DeleteByID)
	g.GET("/serviceAccessPoints/query", workerCustomerAccountHandlerImpl.ListWorkerServiceAccessPoints)

	// WORKERS
	// /workers
	g.GET("/workers", workerHandler.List)
	g.POST("/workers", workerHandler.Create)
	// /workers/check
	g.GET("/workers/check", workerUserHandler.CheckIfExistWithFilter)
	// /workers/:workerId
	g.GET("/workers/:workerId", workerHandler.GetByID)
	g.PUT("/workers/:workerId", workerHandler.UpdateByID)
	g.DELETE("/workers/:workerId", workerHandler.DeleteByID)
	// /workers/details
	g.GET("/workers/details", workerHandler.GetDetails)
	g.GET("/workers/query", workerHandlerImpl.ListWorkers)
	// /workers/customerAccounts
	g.GET("/workers/customerAccounts", workerCustomerAccountHandlerImpl.ListWorkerCustomerAccounts)

	// CUSTOMER_USERS
	// /customerUsers
	g.POST("/customerUsers", customerUserHandler.Create)

	// WORKER_USERS
	// /workerUsers
	g.POST("/workerUsers", workerUserHandler.Create)
}

func (h *httpServer) initializeHandlers() {
	// login
	authHandler = (h.hf.New(handler.AUTH)).(impl.AuthHandler)

	// provider
	providerHandler = h.hf.New(handler.PROVIDER)

	// worker
	workerHandler = h.hf.New(handler.WORKER)
	// impl
	workerHandlerImpl = workerHandler.(impl.WorkerHandler)

	// workerUser
	workerUserHandler = h.hf.New(handler.WORKER_USER)

	// user
	userHandler = h.hf.New(handler.USER)
	// impl
	userHandlerImpl = userHandler.(impl.UserHandler)

	// customer_account
	customerAccountHandler = h.hf.New(handler.CUSTOMER_ACCOUNT)

	// impl
	workerCustomerAccountHandlerImpl = h.hf.New(handler.WORKER_CUSTOMER_ACCOUNT).(impl.WorkerCustomerAccountHandler)

	// customerUser
	customerUserHandler = h.hf.New(handler.CUSTOMER_USER)

	// serviceAccessPoint
	serviceAccessPointHandler = h.hf.New(handler.SERVICE_ACCESS_POINT)
	// impl
	serviceAccessPointHandlerImpl = serviceAccessPointHandler.(impl.ServiceAccessPointHandler)
}
