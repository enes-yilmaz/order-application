package customer

import (
	"CustomerAPI/src/config"
	"CustomerAPI/src/docs"
	customerHandler "CustomerAPI/src/internal/handlers"
	customerRepo "CustomerAPI/src/internal/storages/mongo"
	fmMiddleware "CustomerAPI/src/pkg/middlewares"
	"CustomerAPI/src/pkg/utils"
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

func Execute(env string) {

	cfg := config.Config(env)

	logrus.Info("Customer API running on \"" + env + "\" environment.")
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: time.ANSIC})

	docs.SwaggerInfo.Host = utils.GetSwagHostEnv()
	e := echo.New()
	e.HideBanner = true
	baseGroup := e.Group("/customer-api") // api routing

	baseGroup.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))
	baseGroup.Use(fmMiddleware.PanicExceptionHandling())

	opts := options.Client().ApplyURI(cfg.MongoClientURI)
	ctx, cancel := context.WithTimeout(context.Background(), cfg.MongoClientDuration)
	defer cancel()
	mc, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = mc.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	customerRepository := customerRepo.NewRepository(mc.Database(cfg.CustomerDBName).Collection(cfg.CustomerCollectionName))
	customerHandler.NewHandler(baseGroup, customerRepository)

	baseGroup.GET("/swagger/*", echoSwagger.WrapHandler)
	baseGroup.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	log.Fatal(e.Start(":4000"))
}
