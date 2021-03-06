// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package product

import (
	"context"
	"fmt"
	"math"
	"net/url"

	productpb "github.com/animeapis/go-genproto/product/v1alpha1"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newChapterClientHook clientHook

// ChapterCallOptions contains the retry settings for each method of ChapterClient.
type ChapterCallOptions struct {
	GetChapter    []gax.CallOption
	ListChapters  []gax.CallOption
	CreateChapter []gax.CallOption
	UpdateChapter []gax.CallOption
	DeleteChapter []gax.CallOption
}

func defaultChapterGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("product.animeapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("product.animeapis.com:443"),
		internaloption.WithDefaultAudience("https://product.animeapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultChapterCallOptions() *ChapterCallOptions {
	return &ChapterCallOptions{
		GetChapter:    []gax.CallOption{},
		ListChapters:  []gax.CallOption{},
		CreateChapter: []gax.CallOption{},
		UpdateChapter: []gax.CallOption{},
		DeleteChapter: []gax.CallOption{},
	}
}

// internalChapterClient is an interface that defines the methods availaible from Product API.
type internalChapterClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetChapter(context.Context, *productpb.GetChapterRequest, ...gax.CallOption) (*productpb.Chapter, error)
	ListChapters(context.Context, *productpb.ListChaptersRequest, ...gax.CallOption) *ChapterIterator
	CreateChapter(context.Context, *productpb.CreateChapterRequest, ...gax.CallOption) (*productpb.Chapter, error)
	UpdateChapter(context.Context, *productpb.UpdateChapterRequest, ...gax.CallOption) (*productpb.Chapter, error)
	DeleteChapter(context.Context, *productpb.DeleteChapterRequest, ...gax.CallOption) error
}

// ChapterClient is a client for interacting with Product API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type ChapterClient struct {
	// The internal transport-dependent client.
	internalClient internalChapterClient

	// The call options for this service.
	CallOptions *ChapterCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ChapterClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ChapterClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ChapterClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

func (c *ChapterClient) GetChapter(ctx context.Context, req *productpb.GetChapterRequest, opts ...gax.CallOption) (*productpb.Chapter, error) {
	return c.internalClient.GetChapter(ctx, req, opts...)
}

func (c *ChapterClient) ListChapters(ctx context.Context, req *productpb.ListChaptersRequest, opts ...gax.CallOption) *ChapterIterator {
	return c.internalClient.ListChapters(ctx, req, opts...)
}

func (c *ChapterClient) CreateChapter(ctx context.Context, req *productpb.CreateChapterRequest, opts ...gax.CallOption) (*productpb.Chapter, error) {
	return c.internalClient.CreateChapter(ctx, req, opts...)
}

func (c *ChapterClient) UpdateChapter(ctx context.Context, req *productpb.UpdateChapterRequest, opts ...gax.CallOption) (*productpb.Chapter, error) {
	return c.internalClient.UpdateChapter(ctx, req, opts...)
}

func (c *ChapterClient) DeleteChapter(ctx context.Context, req *productpb.DeleteChapterRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteChapter(ctx, req, opts...)
}

// chapterGRPCClient is a client for interacting with Product API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type chapterGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ChapterClient
	CallOptions **ChapterCallOptions

	// The gRPC API client.
	chapterClient productpb.ChapterServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewChapterClient creates a new chapter service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
func NewChapterClient(ctx context.Context, opts ...option.ClientOption) (*ChapterClient, error) {
	clientOpts := defaultChapterGRPCClientOptions()
	if newChapterClientHook != nil {
		hookOpts, err := newChapterClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ChapterClient{CallOptions: defaultChapterCallOptions()}

	c := &chapterGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		chapterClient:    productpb.NewChapterServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *chapterGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *chapterGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *chapterGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *chapterGRPCClient) GetChapter(ctx context.Context, req *productpb.GetChapterRequest, opts ...gax.CallOption) (*productpb.Chapter, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetChapter[0:len((*c.CallOptions).GetChapter):len((*c.CallOptions).GetChapter)], opts...)
	var resp *productpb.Chapter
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.chapterClient.GetChapter(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *chapterGRPCClient) ListChapters(ctx context.Context, req *productpb.ListChaptersRequest, opts ...gax.CallOption) *ChapterIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListChapters[0:len((*c.CallOptions).ListChapters):len((*c.CallOptions).ListChapters)], opts...)
	it := &ChapterIterator{}
	req = proto.Clone(req).(*productpb.ListChaptersRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*productpb.Chapter, string, error) {
		resp := &productpb.ListChaptersResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.chapterClient.ListChapters(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetChapters(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *chapterGRPCClient) CreateChapter(ctx context.Context, req *productpb.CreateChapterRequest, opts ...gax.CallOption) (*productpb.Chapter, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateChapter[0:len((*c.CallOptions).CreateChapter):len((*c.CallOptions).CreateChapter)], opts...)
	var resp *productpb.Chapter
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.chapterClient.CreateChapter(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *chapterGRPCClient) UpdateChapter(ctx context.Context, req *productpb.UpdateChapterRequest, opts ...gax.CallOption) (*productpb.Chapter, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "chapter.name", url.QueryEscape(req.GetChapter().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateChapter[0:len((*c.CallOptions).UpdateChapter):len((*c.CallOptions).UpdateChapter)], opts...)
	var resp *productpb.Chapter
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.chapterClient.UpdateChapter(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *chapterGRPCClient) DeleteChapter(ctx context.Context, req *productpb.DeleteChapterRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteChapter[0:len((*c.CallOptions).DeleteChapter):len((*c.CallOptions).DeleteChapter)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.chapterClient.DeleteChapter(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// ChapterIterator manages a stream of *productpb.Chapter.
type ChapterIterator struct {
	items    []*productpb.Chapter
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*productpb.Chapter, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ChapterIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ChapterIterator) Next() (*productpb.Chapter, error) {
	var item *productpb.Chapter
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ChapterIterator) bufLen() int {
	return len(it.items)
}

func (it *ChapterIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
