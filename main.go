package main

import (
	"spos/auth/controllers"
	"spos/auth/repository"
	"spos/auth/routes"
	"spos/auth/usecase"

	"github.com/s-pos/go-utils/adapter"
	"github.com/s-pos/go-utils/config"
	"github.com/s-pos/go-utils/middleware"
	"github.com/s-pos/go-utils/utils/server"
)

func init() {
	serviceName := "authentication"

	config.Load(serviceName)
}

func main() {
	log := config.Logrus()
	timezone := config.Timezone()

	db := adapter.DBConnection()
	redis := adapter.GetClientRedis()

	mdl := middleware.NewMiddleware(redis, log, timezone)

	// repository will be here
	baseRepo := repository.New(db, redis, timezone)

	// all usecase will be here
	baseUsecase := usecase.New(baseRepo, timezone)

	// all controller will be here
	baseController := controllers.New(baseUsecase)

	// init router
	router := routes.NewRouter(mdl, baseController)

	// run server
	log.Fatalln(server.Wrapper(router.Router()))
}
