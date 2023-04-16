package main

import (
	_ "github.com/go-sql-driver/mysql"
	"test-sql/pkg/web"
)

func main() {
	//var logger *zap.Logger
	//var loggerErr error
	//
	//if logger, loggerErr = zap.NewDevelopment(); loggerErr != nil {
	//	log.Fatalln(loggerErr)
	//}
	//
	//defer logger.Sync()
	//
	//zap.ReplaceGlobals(logger)

	//
	srv := web.CreateWebServer()
	if err := srv.StartListeningForRequests(":3000"); err != nil {
		//zap.S().Errorf("Error while connecting to server : %v", err)
		panic(err)
	}
}
