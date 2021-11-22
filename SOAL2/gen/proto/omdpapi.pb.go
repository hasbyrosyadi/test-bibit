// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: omdpapi.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Searchword string `protobuf:"bytes,1,opt,name=searchword,proto3" json:"searchword,omitempty"`
	Pagination string `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdpapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_omdpapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_omdpapi_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetSearchword() string {
	if x != nil {
		return x.Searchword
	}
	return ""
}

func (x *Request) GetPagination() string {
	if x != nil {
		return x.Pagination
	}
	return ""
}

type DataOmdp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Year   string `protobuf:"bytes,2,opt,name=Year,proto3" json:"Year,omitempty"`
	ImdbID string `protobuf:"bytes,3,opt,name=imdbID,proto3" json:"imdbID,omitempty"`
	Type   string `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	Poster string `protobuf:"bytes,5,opt,name=Poster,proto3" json:"Poster,omitempty"`
}

func (x *DataOmdp) Reset() {
	*x = DataOmdp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdpapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataOmdp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataOmdp) ProtoMessage() {}

func (x *DataOmdp) ProtoReflect() protoreflect.Message {
	mi := &file_omdpapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataOmdp.ProtoReflect.Descriptor instead.
func (*DataOmdp) Descriptor() ([]byte, []int) {
	return file_omdpapi_proto_rawDescGZIP(), []int{1}
}

func (x *DataOmdp) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *DataOmdp) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *DataOmdp) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *DataOmdp) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DataOmdp) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search       []*DataOmdp `protobuf:"bytes,1,rep,name=Search,proto3" json:"Search,omitempty"`
	TotalResults string      `protobuf:"bytes,2,opt,name=totalResults,proto3" json:"totalResults,omitempty"`
	Response     string      `protobuf:"bytes,3,opt,name=Response,proto3" json:"Response,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdpapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_omdpapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_omdpapi_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetSearch() []*DataOmdp {
	if x != nil {
		return x.Search
	}
	return nil
}

func (x *Response) GetTotalResults() string {
	if x != nil {
		return x.TotalResults
	}
	return ""
}

func (x *Response) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type RequestDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImdbID string `protobuf:"bytes,1,opt,name=imdbID,proto3" json:"imdbID,omitempty"`
}

func (x *RequestDetail) Reset() {
	*x = RequestDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdpapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestDetail) ProtoMessage() {}

func (x *RequestDetail) ProtoReflect() protoreflect.Message {
	mi := &file_omdpapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestDetail.ProtoReflect.Descriptor instead.
func (*RequestDetail) Descriptor() ([]byte, []int) {
	return file_omdpapi_proto_rawDescGZIP(), []int{3}
}

func (x *RequestDetail) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

type Rating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string `protobuf:"bytes,1,opt,name=Source,proto3" json:"Source,omitempty"`
	Value  string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *Rating) Reset() {
	*x = Rating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdpapi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rating) ProtoMessage() {}

func (x *Rating) ProtoReflect() protoreflect.Message {
	mi := &file_omdpapi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rating.ProtoReflect.Descriptor instead.
func (*Rating) Descriptor() ([]byte, []int) {
	return file_omdpapi_proto_rawDescGZIP(), []int{4}
}

func (x *Rating) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Rating) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ResponseDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string    `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Year       string    `protobuf:"bytes,2,opt,name=Year,proto3" json:"Year,omitempty"`
	Rated      string    `protobuf:"bytes,3,opt,name=Rated,proto3" json:"Rated,omitempty"`
	Released   string    `protobuf:"bytes,4,opt,name=Released,proto3" json:"Released,omitempty"`
	Runtime    string    `protobuf:"bytes,5,opt,name=Runtime,proto3" json:"Runtime,omitempty"`
	Genre      string    `protobuf:"bytes,6,opt,name=Genre,proto3" json:"Genre,omitempty"`
	Director   string    `protobuf:"bytes,7,opt,name=Director,proto3" json:"Director,omitempty"`
	Writer     string    `protobuf:"bytes,8,opt,name=Writer,proto3" json:"Writer,omitempty"`
	Actors     string    `protobuf:"bytes,9,opt,name=Actors,proto3" json:"Actors,omitempty"`
	Plot       string    `protobuf:"bytes,10,opt,name=Plot,proto3" json:"Plot,omitempty"`
	Language   string    `protobuf:"bytes,11,opt,name=Language,proto3" json:"Language,omitempty"`
	Country    string    `protobuf:"bytes,12,opt,name=Country,proto3" json:"Country,omitempty"`
	Awards     string    `protobuf:"bytes,13,opt,name=Awards,proto3" json:"Awards,omitempty"`
	Poster     string    `protobuf:"bytes,14,opt,name=Poster,proto3" json:"Poster,omitempty"`
	Ratings    []*Rating `protobuf:"bytes,15,rep,name=Ratings,proto3" json:"Ratings,omitempty"`
	Metascore  string    `protobuf:"bytes,16,opt,name=Metascore,proto3" json:"Metascore,omitempty"`
	ImdbRating string    `protobuf:"bytes,17,opt,name=imdbRating,proto3" json:"imdbRating,omitempty"`
	ImdbVotes  string    `protobuf:"bytes,18,opt,name=imdbVotes,proto3" json:"imdbVotes,omitempty"`
	ImdbID     string    `protobuf:"bytes,19,opt,name=imdbID,proto3" json:"imdbID,omitempty"`
	Type       string    `protobuf:"bytes,20,opt,name=Type,proto3" json:"Type,omitempty"`
	DVD        string    `protobuf:"bytes,21,opt,name=DVD,proto3" json:"DVD,omitempty"`
	BoxOffice  string    `protobuf:"bytes,22,opt,name=BoxOffice,proto3" json:"BoxOffice,omitempty"`
	Production string    `protobuf:"bytes,23,opt,name=Production,proto3" json:"Production,omitempty"`
	Website    string    `protobuf:"bytes,24,opt,name=Website,proto3" json:"Website,omitempty"`
	Response   string    `protobuf:"bytes,25,opt,name=Response,proto3" json:"Response,omitempty"`
}

func (x *ResponseDetail) Reset() {
	*x = ResponseDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdpapi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseDetail) ProtoMessage() {}

func (x *ResponseDetail) ProtoReflect() protoreflect.Message {
	mi := &file_omdpapi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseDetail.ProtoReflect.Descriptor instead.
func (*ResponseDetail) Descriptor() ([]byte, []int) {
	return file_omdpapi_proto_rawDescGZIP(), []int{5}
}

func (x *ResponseDetail) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ResponseDetail) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *ResponseDetail) GetRated() string {
	if x != nil {
		return x.Rated
	}
	return ""
}

func (x *ResponseDetail) GetReleased() string {
	if x != nil {
		return x.Released
	}
	return ""
}

func (x *ResponseDetail) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *ResponseDetail) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *ResponseDetail) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

func (x *ResponseDetail) GetWriter() string {
	if x != nil {
		return x.Writer
	}
	return ""
}

func (x *ResponseDetail) GetActors() string {
	if x != nil {
		return x.Actors
	}
	return ""
}

func (x *ResponseDetail) GetPlot() string {
	if x != nil {
		return x.Plot
	}
	return ""
}

func (x *ResponseDetail) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *ResponseDetail) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *ResponseDetail) GetAwards() string {
	if x != nil {
		return x.Awards
	}
	return ""
}

func (x *ResponseDetail) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

func (x *ResponseDetail) GetRatings() []*Rating {
	if x != nil {
		return x.Ratings
	}
	return nil
}

func (x *ResponseDetail) GetMetascore() string {
	if x != nil {
		return x.Metascore
	}
	return ""
}

func (x *ResponseDetail) GetImdbRating() string {
	if x != nil {
		return x.ImdbRating
	}
	return ""
}

func (x *ResponseDetail) GetImdbVotes() string {
	if x != nil {
		return x.ImdbVotes
	}
	return ""
}

func (x *ResponseDetail) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *ResponseDetail) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ResponseDetail) GetDVD() string {
	if x != nil {
		return x.DVD
	}
	return ""
}

func (x *ResponseDetail) GetBoxOffice() string {
	if x != nil {
		return x.BoxOffice
	}
	return ""
}

func (x *ResponseDetail) GetProduction() string {
	if x != nil {
		return x.Production
	}
	return ""
}

func (x *ResponseDetail) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *ResponseDetail) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_omdpapi_proto protoreflect.FileDescriptor

var file_omdpapi_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6f, 0x6d, 0x64, 0x70, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x78, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x4f, 0x6d, 0x64, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x59, 0x65, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x22, 0x73, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x4f, 0x6d, 0x64, 0x70, 0x52, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x22,
	0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27,
	0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x69, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x22, 0x36, 0x0a, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x99, 0x05, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x59, 0x65, 0x61, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x59, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x52, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x52, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x47, 0x65, 0x6e, 0x72,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6c,
	0x6f, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x6c, 0x6f, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6f,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x07, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69,
	0x6d, 0x64, 0x62, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x69, 0x6d, 0x64, 0x62, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x69,
	0x6d, 0x64, 0x62, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x69, 0x6d, 0x64, 0x62, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x64,
	0x62, 0x49, 0x44, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49,
	0x44, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x44, 0x56, 0x44, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x44, 0x56, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x6f, 0x78, 0x4f, 0x66,
	0x66, 0x69, 0x63, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x42, 0x6f, 0x78, 0x4f,
	0x66, 0x66, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65,
	0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xc4, 0x01, 0x0a, 0x0b,
	0x4f, 0x6d, 0x64, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x09, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x26, 0x12, 0x24, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x2f, 0x7b, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x77, 0x6f, 0x72, 0x64, 0x7d, 0x2f, 0x7b, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x12, 0x59, 0x0a, 0x0b, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x2f, 0x7b, 0x69, 0x6d, 0x64, 0x62, 0x49,
	0x44, 0x7d, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_omdpapi_proto_rawDescOnce sync.Once
	file_omdpapi_proto_rawDescData = file_omdpapi_proto_rawDesc
)

func file_omdpapi_proto_rawDescGZIP() []byte {
	file_omdpapi_proto_rawDescOnce.Do(func() {
		file_omdpapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_omdpapi_proto_rawDescData)
	})
	return file_omdpapi_proto_rawDescData
}

var file_omdpapi_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_omdpapi_proto_goTypes = []interface{}{
	(*Request)(nil),        // 0: proto.Request
	(*DataOmdp)(nil),       // 1: proto.DataOmdp
	(*Response)(nil),       // 2: proto.Response
	(*RequestDetail)(nil),  // 3: proto.RequestDetail
	(*Rating)(nil),         // 4: proto.Rating
	(*ResponseDetail)(nil), // 5: proto.ResponseDetail
}
var file_omdpapi_proto_depIdxs = []int32{
	1, // 0: proto.Response.Search:type_name -> proto.DataOmdp
	4, // 1: proto.ResponseDetail.Ratings:type_name -> proto.Rating
	0, // 2: proto.OmdpService.ListMovie:input_type -> proto.Request
	3, // 3: proto.OmdpService.DetailMovie:input_type -> proto.RequestDetail
	2, // 4: proto.OmdpService.ListMovie:output_type -> proto.Response
	5, // 5: proto.OmdpService.DetailMovie:output_type -> proto.ResponseDetail
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_omdpapi_proto_init() }
func file_omdpapi_proto_init() {
	if File_omdpapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_omdpapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_omdpapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataOmdp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_omdpapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_omdpapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestDetail); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_omdpapi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rating); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_omdpapi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseDetail); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_omdpapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_omdpapi_proto_goTypes,
		DependencyIndexes: file_omdpapi_proto_depIdxs,
		MessageInfos:      file_omdpapi_proto_msgTypes,
	}.Build()
	File_omdpapi_proto = out.File
	file_omdpapi_proto_rawDesc = nil
	file_omdpapi_proto_goTypes = nil
	file_omdpapi_proto_depIdxs = nil
}