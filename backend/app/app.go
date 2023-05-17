package app

import (
	"log"

	"Diploma/app/modules"
	"Diploma/pkg/config"
	"Diploma/pkg/database"
	"Diploma/pkg/errorPkg"
	"Diploma/pkg/storage"
	"Diploma/server"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jmoiron/sqlx"
)

func Run() {
	defer log.Print("Shutting down\n")

	//Logging initialisation
	Logger := log.Default()
	//if err := logging.LogInit(); err != nil {
	//	log.Fatal("Logging init error\n" + err.Error())
	//}

	//Config initialisation
	if err := config.Init(); err != nil {
		log.Fatalln("Config init error\n" + err.Error())
	}

	//Connecting to Postgres
	DatabaseClient, err := database.Setup()
	if err != nil {
		log.Fatalln("Db connect error\n", err.Error())
	}

	//Connecting to CloudStorage
	StorageClient, err := storage.Setup()
	if err != nil {
		log.Fatalln("Storage connect error", err.Error())
	}

	//SetupModules
	handlers := SetupModules(DatabaseClient, StorageClient, Logger)

	//Start handling httpService requests
	if err = server.Start(handlers); err != nil {
		log.Fatal("Server starting error\n" + err.Error())
	}
}

func SetupModules(DatabaseClient *sqlx.DB, StorageClient *s3.S3, logger *log.Logger) (HandlerModule modules.HandlerModule) {
	ErrorManager := errorPkg.InitErrorManager(logger)
	GatewayModule := modules.SetupGateway(DatabaseClient, StorageClient, ErrorManager.ErrorCreator)
	log.Println("Gateway module setup correctly")
	UseCaseModule := modules.SetupUseCase(GatewayModule, ErrorManager.ErrorCreator)
	log.Println("UseCase module setup correctly")
	DelegateModule := modules.SetupDelegate(UseCaseModule)
	log.Println("Delegate module setup correctly")
	HandlerModule = modules.SetupHandler(DelegateModule, ErrorManager)
	log.Println("Handler module setup correctly")
	return HandlerModule
}
