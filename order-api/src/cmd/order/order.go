package order

import (
	"OrderAPI/src/config"
	"OrderAPI/src/docs"
	"OrderAPI/src/internal/health_check"
	orderHandler "OrderAPI/src/internal/orders/handlers"
	orderRepository "OrderAPI/src/internal/orders/storages/mongo"
	fmMiddleware "OrderAPI/src/pkg/middlewares"
	"OrderAPI/src/pkg/utils"
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
	config.SetConfig(cfg)

	logrus.Info("Order API running on \"" + env + "\" environment.")
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: time.ANSIC})

	docs.SwaggerInfo.Host = utils.GetSwagHostEnv()
	e := echo.New()
	e.HideBanner = true
	baseGroup := e.Group("/order-api") // api routing

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
	orderRepository := orderRepository.NewRepository(mc.Database(cfg.OrderDBName).Collection(cfg.OrderCollectionName))
	orderHandler.NewHandler(baseGroup, orderRepository)

	baseGroup.GET("/swagger/*", echoSwagger.WrapHandler)
	baseGroup.GET("/health", health_check.HealthCheck)

	log.Fatal(e.Start(":4001"))
}
