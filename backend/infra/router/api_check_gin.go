// internal/infrastructure/router/gin.go
package router

import (
	"context"
	emailadapter "corporation-site/adapter/email"
	forkadapter "corporation-site/adapter/forklift"
	"corporation-site/domain"
	"corporation-site/infra/db/nosql"
	forkrepo "corporation-site/repository/nosql/forklift"
	"corporation-site/service"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type CheckApiGinAdapter interface {
	RegisterRoutes(router interface{})
}

type CheckApiGinServer struct {
	port        Port
	dbNoSQL     nosql.NoSQLClient
	emailClient domain.EmailClient
	logger      domain.Logger
	validator   domain.Validator
	router      *gin.Engine
	adapters    []CheckApiGinAdapter
}

func NewCheckApiGinServer(
	port Port,
	dbNoSQL nosql.NoSQLClient,
	emailClint domain.EmailClient,
	logger domain.Logger,
	validator domain.Validator,
) *CheckApiGinServer {
	return &CheckApiGinServer{
		port:        port,
		dbNoSQL:     dbNoSQL,
		emailClient: emailClint,
		logger:      logger,
		validator:   validator,
		router:      gin.New(),
	}
}

func (g *CheckApiGinServer) AddAdapter(adapter CheckApiGinAdapter) {
	g.adapters = append(g.adapters, adapter)
}

func (g *CheckApiGinServer) Listen() error {
	if err := g.SetupRouter(); err != nil {
		return fmt.Errorf("failed to setup router: %w", err)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", g.port),
		Handler: g.router,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		g.logger.Infof("Starting HTTP Server on port %d", g.port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			g.logger.Fatalln("Error starting HTTP server:", err)
		}
	}()

	<-stop

	g.logger.Infof("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server forced to shutdown: %w", err)
	}

	g.logger.Infof("Server exiting")
	return nil
}

func (g *CheckApiGinServer) GetHandler() http.Handler {
	return g.router
}

func (g *CheckApiGinServer) SetupRouter() error {
	g.router.Use(gin.Recovery())
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	g.router.Use(cors.New(config))

	// forklift関連
	forkliftRepo, err := forkrepo.NewForkliftRepository(g.dbNoSQL)
	if err != nil {
		return fmt.Errorf("failed to create forklift repository: %w", err)
	}
	forkliftService := service.NewForkliftService(forkliftRepo)
	forkliftAdapter := forkadapter.NewForkliftAdapter(forkliftService)
	g.AddAdapter(forkliftAdapter)

	// email関連
	emailService := service.NewEmailService(g.emailClient)
	emailAdapter := emailadapter.NewEmailAdapter(emailService, g.validator)
	g.AddAdapter(emailAdapter)

	// adapterの登録
	for _, adapter := range g.adapters {
		adapter.RegisterRoutes(g.router)
	}
	// Gin特有のカスタムルート登録
	g.router.GET("/health", g.healthCheck)
	return nil
}

func (g *CheckApiGinServer) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
