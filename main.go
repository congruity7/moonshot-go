package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"github.com/congruity7/moonshot-go/pkg/api"
	"github.com/congruity7/moonshot-go/pkg/models"
	"github.com/congruity7/moonshot-go/pkg/service"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type options struct{}

// type Option func(*options)

// type authMiddleWare struct {
// 	opts *options
// }

// var _ negroni.Handler = (*authMiddleWare)(nil)

// func NewAuthMiddleware(opts ...Option) negroni.Handler {
// 	aOpts := &options{}
// 	for _, opt := range opts {
// 		opt(aOpts)
// 	}
// 	return &authMiddleWare{opts: aOpts}
// }

// func (m *authMiddleWare) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
// 	return
// }

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func Setup() {
	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s?parseTime=true",
		goDotEnvVariable("MOONSHOT_DB_USER"),
		goDotEnvVariable("MOONSHOT_DB_PASSWD"),
		goDotEnvVariable("MOONSHOT_DB_PROTOCOL"),
		goDotEnvVariable("MOONSHOT_DB_INSTANCE"),
		goDotEnvVariable("MOONSHOT_DB_NAME"))

	logrus.Info("DSN : ", dsn)

	dbInstance, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        dsn,
	}), &gorm.Config{})

	if err != nil {
		logrus.Fatal("connecting to database : ", err)
	}

	logrus.Info("success connecting to db")

	dbInstance.AutoMigrate(&models.User{}, &models.Wallet{}, &models.Bet{}, &models.PlacedBet{}, &models.Config{})

	var wg sync.WaitGroup
	var stopChan <-chan struct{}

	ds := service.NewDatabaseService(dbInstance)
	rs := service.NewRedisService(nil)
	logger := &logrus.Logger{}

	logger.SetOutput(os.Stdout)

	wg.Add(1)

	go StartAPI(&wg, ds, rs, logger, stopChan)

	wg.Wait()
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func StartAPI(wg *sync.WaitGroup, ds *service.DatabaseService, rs *service.RedisService, logger *logrus.Logger, stopChan <-chan struct{}) {
	defer wg.Done()

	router := httprouter.New()

	ac := api.NewAPIContext(ds, rs, logger)
	router.GET("/", Index)
	router.GET("/moonshot/v1/user/:user_id/", ac.GetUserByID)
	router.GET("/moonshot/v1/user", ac.GetUsers)
	router.POST("/moonshot/v1/user", ac.CreateUser)
	router.PUT("/moonshot/v1/user", ac.UpdateUser)
	router.DELETE("/moonshot/v1/user/:user_id/", ac.DeleteUserByID)

	router.GET("/moonshot/v1/wallet/:wallet_id/", ac.GetWalletByID)
	router.GET("/moonshot/v1/wallet", ac.GetWallets)
	router.POST("/moonshot/v1/wallet", ac.CreateWallet)
	router.PUT("/moonshot/v1/wallet", ac.UpdateWallet)
	router.DELETE("/moonshot/v1/wallet/:wallet/", ac.DeleteWalletByID)

	router.GET("/moonshot/v1/config", ac.GetConfig)
	router.POST("/moonshot/v1/config", ac.CreateConfig)
	router.PUT("/moonshot/v1/config", ac.UpdateConfig)

	// n := negroni.New(negroni.NewRecovery(),
	// 	NewAuthMiddleware())
	// n.UseHandler(router)

	//api := &http.Server{Addr: ":8000", Handler: n}

	go func() {
		if err := http.ListenAndServe(":8000", router); err != nil {
			logger.Error("starting api server", err)
		}
	}()

	logrus.Info("waiting for stop signal")

	<-stopChan

	logger.Info("shutting down api server")

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	// defer cancel()
	// api.Shutdown(ctx)

}

func main() {
	Setup()
}
