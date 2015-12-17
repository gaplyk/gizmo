// Code generated by protoc-gen-go.
// source: semanticconcept.proto
// DO NOT EDIT!

package nyt

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SemanticConceptResponse struct {
	Status     string                   `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Copyright  string                   `protobuf:"bytes,2,opt,name=copyright" json:"copyright,omitempty"`
	NumResults uint32                   `protobuf:"varint,3,opt,name=num_results" json:"num_results,omitempty"`
	Results    []*SemanticConceptResult `protobuf:"bytes,4,rep,name=results" json:"results,omitempty"`
}

func (m *SemanticConceptResponse) Reset()                    { *m = SemanticConceptResponse{} }
func (m *SemanticConceptResponse) String() string            { return proto.CompactTextString(m) }
func (*SemanticConceptResponse) ProtoMessage()               {}
func (*SemanticConceptResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *SemanticConceptResponse) GetResults() []*SemanticConceptResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type SemanticConceptResult struct {
	ArticleList *SemanticConceptArticleList `protobuf:"bytes,1,opt,name=article_list" json:"article_list,omitempty"`
}

func (m *SemanticConceptResult) Reset()                    { *m = SemanticConceptResult{} }
func (m *SemanticConceptResult) String() string            { return proto.CompactTextString(m) }
func (*SemanticConceptResult) ProtoMessage()               {}
func (*SemanticConceptResult) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *SemanticConceptResult) GetArticleList() *SemanticConceptArticleList {
	if m != nil {
		return m.ArticleList
	}
	return nil
}

type SemanticConceptArticleList struct {
	Results []*SemanticConceptArticle `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
	Total   uint32                    `protobuf:"varint,2,opt,name=total" json:"total,omitempty"`
}

func (m *SemanticConceptArticleList) Reset()                    { *m = SemanticConceptArticleList{} }
func (m *SemanticConceptArticleList) String() string            { return proto.CompactTextString(m) }
func (*SemanticConceptArticleList) ProtoMessage()               {}
func (*SemanticConceptArticleList) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *SemanticConceptArticleList) GetResults() []*SemanticConceptArticle {
	if m != nil {
		return m.Results
	}
	return nil
}

type SemanticConceptArticle struct {
	Body   string `protobuf:"bytes,1,opt,name=body" json:"body,omitempty"`
	Byline string `protobuf:"bytes,2,opt,name=byline" json:"byline,omitempty"`
	Title  string `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Url    string `protobuf:"bytes,4,opt,name=url" json:"url,omitempty"`
}

func (m *SemanticConceptArticle) Reset()                    { *m = SemanticConceptArticle{} }
func (m *SemanticConceptArticle) String() string            { return proto.CompactTextString(m) }
func (*SemanticConceptArticle) ProtoMessage()               {}
func (*SemanticConceptArticle) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func init() {
	proto.RegisterType((*SemanticConceptResponse)(nil), "nyt.SemanticConceptResponse")
	proto.RegisterType((*SemanticConceptResult)(nil), "nyt.SemanticConceptResult")
	proto.RegisterType((*SemanticConceptArticleList)(nil), "nyt.SemanticConceptArticleList")
	proto.RegisterType((*SemanticConceptArticle)(nil), "nyt.SemanticConceptArticle")
}

var fileDescriptor1 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x89, 0xa9, 0x95, 0x4c, 0x1a, 0xc1, 0x95, 0x6a, 0xa8, 0x07, 0x4b, 0x4e, 0x05, 0x25,
	0x87, 0x8a, 0x3f, 0x40, 0xbc, 0x8a, 0x87, 0x7a, 0xf2, 0x54, 0x92, 0x38, 0xe8, 0xc2, 0x76, 0x37,
	0xec, 0x4c, 0x0e, 0xb9, 0xf8, 0xdb, 0x1d, 0xb7, 0x11, 0x0f, 0x36, 0x1e, 0xf3, 0xde, 0xe3, 0xcd,
	0xcb, 0xb7, 0x30, 0x27, 0xdc, 0x55, 0x96, 0x75, 0xd3, 0x38, 0xdb, 0x60, 0xcb, 0x65, 0xeb, 0x1d,
	0x3b, 0x15, 0xdb, 0x9e, 0x8b, 0x4f, 0xb8, 0x7c, 0x19, 0xdc, 0xc7, 0xbd, 0xbb, 0x41, 0x6a, 0x9d,
	0x25, 0x54, 0xa7, 0x30, 0x25, 0xae, 0xb8, 0xa3, 0x3c, 0x5a, 0x46, 0xab, 0x44, 0x9d, 0x41, 0xd2,
	0xb8, 0xb6, 0xf7, 0xfa, 0xfd, 0x83, 0xf3, 0xa3, 0x20, 0x9d, 0x43, 0x6a, 0xbb, 0xdd, 0xd6, 0x23,
	0x75, 0x86, 0x29, 0x8f, 0x45, 0xcc, 0xd4, 0x0d, 0x9c, 0xfc, 0x08, 0x93, 0x65, 0xbc, 0x4a, 0xd7,
	0x8b, 0x52, 0x2e, 0x95, 0x7f, 0xcf, 0x48, 0xa4, 0x78, 0x86, 0xf9, 0x41, 0x43, 0xdd, 0xc3, 0xac,
	0xf2, 0xa2, 0x1a, 0xdc, 0x1a, 0x4d, 0x1c, 0x36, 0xa4, 0xeb, 0xeb, 0x43, 0x55, 0x0f, 0xfb, 0xdc,
	0x93, 0xc4, 0x8a, 0x57, 0x58, 0x8c, 0xbb, 0xea, 0xf6, 0x77, 0x5a, 0x14, 0xa6, 0x5d, 0xfd, 0xd3,
	0xa7, 0x32, 0x38, 0x66, 0xc7, 0x95, 0x09, 0x3f, 0x9b, 0x15, 0x1b, 0xb8, 0x18, 0x09, 0xce, 0x60,
	0x52, 0xbb, 0xb7, 0x7e, 0xe0, 0x24, 0xdc, 0xea, 0xde, 0x68, 0x8b, 0x03, 0xa4, 0xef, 0x1a, 0xcd,
	0x06, 0x03, 0x9e, 0x44, 0xa5, 0x10, 0x77, 0xde, 0x08, 0x1a, 0xf9, 0xa8, 0xa7, 0xe1, 0x29, 0xee,
	0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x54, 0xe2, 0x0b, 0x29, 0xa3, 0x01, 0x00, 0x00,
}