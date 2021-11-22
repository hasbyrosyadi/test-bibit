package usecase

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	pb "stockbit/gen/proto"
)
type OmdpServiceServer struct {
	//query *gorm.DB
	pb.UnimplementedOmdpServiceServer
}

//func NewServiceServer(db *gorm.DB) *OmdpServiceServer {
//	return &OmdpServiceServer{query: db, service: pb.UnimplementedOmdpServiceServer{}}
//}

func (s *OmdpServiceServer) ListMovie(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	var urlRequest = "http://www.omdbapi.com/?apikey=faf7e5bb&s=" + req.Searchword + "&page=" + req.Pagination
	resp, err := http.Get(urlRequest)
	if err != nil {
		log.Fatalln(err)
	}
	data, _ := ioutil.ReadAll(resp.Body)

	var dataResp *pb.Response
	if err := json.Unmarshal(data, &dataResp); err != nil {
		return nil, err
	}
	return dataResp, nil
}

func (s *OmdpServiceServer) DetailMovie(ctx context.Context, req *pb.RequestDetail) (*pb.ResponseDetail, error) {
	var urlRequest = "http://www.omdbapi.com/?apikey=faf7e5bb&i=" + req.ImdbID
	resp, err := http.Get(urlRequest)
	if err != nil {
		log.Fatalln(err)
	}
	data, _ := ioutil.ReadAll(resp.Body)

	var dataResp *pb.ResponseDetail
	if err := json.Unmarshal(data, &dataResp); err != nil {
		return nil, err
	}
	return dataResp, nil
}

