package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mrdhira/210622-kuncie-tchnl-test/internals/domains/repository"
	"github.com/mrdhira/210622-kuncie-tchnl-test/internals/domains/services"
	Router "github.com/mrdhira/210622-kuncie-tchnl-test/internals/ports/http"
	"github.com/mrdhira/210622-kuncie-tchnl-test/internals/ports/http/controllers"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// serveHttp add command
var serveHttp = &cobra.Command{
	Use:   "serveHttp",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
	examples and usage of using your application. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Init DB Connection
		postgresClient, err := gorm.Open(postgres.Open(config.PostgresConfigs.DSN), &gorm.Config{})
		if err != nil {
			log.Fatalf("error when try to connect to postgres: %v", err)
			panic(err)
		}

		// Init Repository
		repository := repository.InitRepository(postgresClient)

		// Init Services
		cartServices := services.InitCartServices(repository)
		checkoutServices := services.InitCheckoutServices(repository)

		// Init Controllers
		cartController := controllers.InitCartControllers(cartServices)
		checkoutController := controllers.InitCheckoutControllers(checkoutServices)

		// Init Router
		router := Router.InitHttp(&config.HttpConfigs, cartController, checkoutController)

		var gracefulStop = make(chan os.Signal)
		signal.Notify(gracefulStop, syscall.SIGTERM)
		signal.Notify(gracefulStop, syscall.SIGINT)

		HTTPServer := &http.Server{
			Handler:      router,
			Addr:         "0.0.0.0:8000",
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
		}

		go func() {
			if err := HTTPServer.ListenAndServe(); err != http.ErrServerClosed {
				log.Fatalf("error on listen and serve: %v", err)
				panic(err)
			}
		}()

		<-gracefulStop
		if err := HTTPServer.Shutdown(context.TODO()); err != nil {
			panic(err)
		}
		fmt.Println("http closed")
	},
}

func init() {
	rootCmd.AddCommand(serveHttp)
}
