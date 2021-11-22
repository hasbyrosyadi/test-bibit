package main

import (
	"context"
	"log"
	"net/http"
	pb "stockbit/gen/proto"
	"stockbit/server/usecase"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func main() {
	//dbHost := "localhost"
	//dbPort := "3306"
	//dbUser := "root"
	//dbPass := ""
	//dbName := "stockbit"
	//
	//var dsn = dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=true&loc=Local"
	//
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	fmt.Println(err)
	//}

	mux := runtime.NewServeMux()

	go func() {
		var server = usecase.OmdpServiceServer{}
		pb.RegisterOmdpServiceHandlerServer(context.Background(), mux, &server)
	}()

	log.Fatal(http.ListenAndServe("localhost:8081", mux))
}
