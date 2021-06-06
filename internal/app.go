package internal

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	chi_middleware "github.com/go-chi/chi/middleware"

	"github.com/gopher-translator-service/internal/cache"
	"github.com/gopher-translator-service/internal/handlers"
	"github.com/gopher-translator-service/pkg/logger"
	"github.com/gopher-translator-service/pkg/middleware"
	"github.com/gopher-translator-service/pkg/util"
)

var (
	TryingToRunUninitializedAppErr = errors.New("trying to run uninitialized app")
	AppInitCfgFailedValidationErr  = errors.New("init configurations failed validation")
)

const defaultRequestTimeout = 60 * time.Second

type App struct {
	port             int
	router           chi.Router
	log              logger.Logger
	translationCache cache.KeyValueCache // translation cache

	initialized bool
	reqTimeout  time.Duration
}

type AppInitCfg struct {
	Log  logger.Logger
	Port int
}

func (aif *AppInitCfg) Validate() error {
	if aif == nil {
		return AppInitCfgFailedValidationErr
	} else if aif.Port <= 0 {
		return AppInitCfgFailedValidationErr
	} else if util.IsInterfaceNil(aif.Log) {
		return AppInitCfgFailedValidationErr
	}

	return nil
}

func (a *App) Init(cfg *AppInitCfg) error {
	if err := cfg.Validate(); err != nil {
		return err
	}

	translationCache := cache.NewMapCache()

	// TODO: separate the router logic:
	router := chi.NewRouter()
	router.Use(chi_middleware.RequestID)
	router.Use(chi_middleware.CleanPath)
	router.Use(chi_middleware.Recoverer)
	router.Use(chi_middleware.StripSlashes)
	router.Use(chi_middleware.AllowContentType("application/json"))
	router.Use(chi_middleware.Timeout(defaultRequestTimeout))
	router.Use(middleware.LogIncommingReq(cfg.Log))

	router.Post("/word", handlers.WordHandler(cfg.Log, translationCache))
	router.Post("/sentence", handlers.SentenceHandler(cfg.Log, translationCache))
	router.Get("/history", handlers.HistoryHandler(cfg.Log, translationCache))

	a.port = cfg.Port
	a.router = router
	a.initialized = true
	a.log = cfg.Log
	a.reqTimeout = defaultRequestTimeout
	a.translationCache = translationCache
	return nil
}

func (a *App) Run() error {
	if a.initialized == false {
		return TryingToRunUninitializedAppErr
	}

	address := fmt.Sprintf(":%d", a.port)
	http.ListenAndServe(address, a.router)

	return nil
}

func (a *App) Shutdown() error {
	a.log.Flush()
	return nil
}
