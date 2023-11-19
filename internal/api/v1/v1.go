package v1

import (
	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/harijar/geogame/internal/repo"
	"github.com/harijar/geogame/internal/repo/countries"
	"github.com/harijar/geogame/internal/service/prompts"
	"go.uber.org/zap"
	"time"
)

type ServerConfig struct {
	CookieDomain         string
	CookieSecure         bool
	CORSEnabled          bool
	CORSAllowAllOrigins  bool
	CORSOrigins          []string
	CORSAllowCredentials bool
	SameSite             int
}

type V1 struct {
	server       *gin.Engine
	countries    repo.Countries
	prompts      PromptsService
	triesLimit   int
	serverConfig *ServerConfig
	logger       *zap.Logger
}

type PromptsService interface {
	Gen(id int, c *countries.Country, prev []*prompts.Prompt) (*prompts.Prompt, error)
	GenRandom(c *countries.Country, prev []*prompts.Prompt) (*prompts.Prompt, error)
}

func New(countries repo.Countries, prompts PromptsService, triesLimit int, serverConfig *ServerConfig, logger *zap.Logger) *V1 {
	return &V1{
		server:       gin.New(),
		countries:    countries,
		prompts:      prompts,
		triesLimit:   triesLimit,
		serverConfig: serverConfig,
		logger:       logger,
	}
}

func (a *V1) Run(addr string) error {
	if a.serverConfig.CORSEnabled {
		config := cors.Config{
			AllowAllOrigins:  a.serverConfig.CORSAllowAllOrigins,
			AllowMethods:     []string{"GET", "POST"},
			AllowHeaders:     []string{"Origin", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length", "Content-Type"},
			AllowCredentials: a.serverConfig.CORSAllowCredentials,
			MaxAge:           12 * time.Hour,
		}
		if !a.serverConfig.CORSAllowAllOrigins {
			config.AllowOrigins = a.serverConfig.CORSOrigins
		}
		a.server.Use(cors.New(config))
	}
	a.server.Use(ginzap.Ginzap(a.logger, time.RFC3339, true))
	a.server.Use(ginzap.RecoveryWithZap(a.logger, true))
	a.registerRoutes()
	return a.server.Run(addr)
}
