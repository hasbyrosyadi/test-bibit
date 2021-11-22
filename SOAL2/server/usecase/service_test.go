package usecase

import (
	"context"
	pb "stockbit/gen/proto"
	"testing"
)

type listMovieTest struct {
	ctx context.Context
	searchword string
	pagination string
	expected string
}

type detailMovieTest struct {
	ctx context.Context
	imdbID string
	expected string
}

var listMovieTests = []listMovieTest{
	listMovieTest{context.Background(),"Batman", "2", "True"},
}

var detailMovieTests = []detailMovieTest{
	detailMovieTest{context.Background(),"tt0106364", "True"},
}
func TestListMovie(t *testing.T) {

	for _, test := range listMovieTests{
		var req = &pb.Request{
			Searchword: test.searchword,
			Pagination: test.pagination,
		}
		var server = OmdpServiceServer{}
		if output, _ := pb.OmdpServiceServer.ListMovie(&server, test.ctx, req); output.Response != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func TestDetailMovie(t *testing.T) {

	for _, test := range detailMovieTests{
		var req = &pb.RequestDetail{
			ImdbID: test.imdbID,
		}
		var server = OmdpServiceServer{}
		if output, _ := pb.OmdpServiceServer.DetailMovie(&server, test.ctx, req); output.Response != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}