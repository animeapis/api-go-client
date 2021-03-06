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

package webpage

import (
	"context"
	"fmt"
	"math"
	"net/url"

	webpagepb "github.com/animeapis/go-genproto/webpage/v1alpha1"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newArchiveClientHook clientHook

// ArchiveCallOptions contains the retry settings for each method of ArchiveClient.
type ArchiveCallOptions struct {
	Query      []gax.CallOption
	GetPage    []gax.CallOption
	ListPages  []gax.CallOption
	QueryPage  []gax.CallOption
	CreatePage []gax.CallOption
	ImportPage []gax.CallOption
	DeletePage []gax.CallOption
}

func defaultArchiveGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("webpage.animeapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("webpage.animeapis.com:443"),
		internaloption.WithDefaultAudience("https://webpage.animeapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultArchiveCallOptions() *ArchiveCallOptions {
	return &ArchiveCallOptions{
		Query:      []gax.CallOption{},
		GetPage:    []gax.CallOption{},
		ListPages:  []gax.CallOption{},
		QueryPage:  []gax.CallOption{},
		CreatePage: []gax.CallOption{},
		ImportPage: []gax.CallOption{},
		DeletePage: []gax.CallOption{},
	}
}

// internalArchiveClient is an interface that defines the methods availaible from WebPage API.
type internalArchiveClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	Query(context.Context, ...gax.CallOption) (webpagepb.Archive_QueryClient, error)
	GetPage(context.Context, *webpagepb.GetPageRequest, ...gax.CallOption) (*webpagepb.Page, error)
	ListPages(context.Context, *webpagepb.ListPagesRequest, ...gax.CallOption) *PageIterator
	QueryPage(context.Context, *webpagepb.QueryPageRequest, ...gax.CallOption) (*webpagepb.QueryPageResponse, error)
	CreatePage(context.Context, *webpagepb.CreatePageRequest, ...gax.CallOption) (*webpagepb.Page, error)
	ImportPage(context.Context, *webpagepb.ImportPageRequest, ...gax.CallOption) (*webpagepb.ImportPageResponse, error)
	DeletePage(context.Context, *webpagepb.DeletePageRequest, ...gax.CallOption) error
}

// ArchiveClient is a client for interacting with WebPage API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type ArchiveClient struct {
	// The internal transport-dependent client.
	internalClient internalArchiveClient

	// The call options for this service.
	CallOptions *ArchiveCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ArchiveClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ArchiveClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ArchiveClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

func (c *ArchiveClient) Query(ctx context.Context, opts ...gax.CallOption) (webpagepb.Archive_QueryClient, error) {
	return c.internalClient.Query(ctx, opts...)
}

func (c *ArchiveClient) GetPage(ctx context.Context, req *webpagepb.GetPageRequest, opts ...gax.CallOption) (*webpagepb.Page, error) {
	return c.internalClient.GetPage(ctx, req, opts...)
}

func (c *ArchiveClient) ListPages(ctx context.Context, req *webpagepb.ListPagesRequest, opts ...gax.CallOption) *PageIterator {
	return c.internalClient.ListPages(ctx, req, opts...)
}

func (c *ArchiveClient) QueryPage(ctx context.Context, req *webpagepb.QueryPageRequest, opts ...gax.CallOption) (*webpagepb.QueryPageResponse, error) {
	return c.internalClient.QueryPage(ctx, req, opts...)
}

func (c *ArchiveClient) CreatePage(ctx context.Context, req *webpagepb.CreatePageRequest, opts ...gax.CallOption) (*webpagepb.Page, error) {
	return c.internalClient.CreatePage(ctx, req, opts...)
}

func (c *ArchiveClient) ImportPage(ctx context.Context, req *webpagepb.ImportPageRequest, opts ...gax.CallOption) (*webpagepb.ImportPageResponse, error) {
	return c.internalClient.ImportPage(ctx, req, opts...)
}

func (c *ArchiveClient) DeletePage(ctx context.Context, req *webpagepb.DeletePageRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeletePage(ctx, req, opts...)
}

// archiveGRPCClient is a client for interacting with WebPage API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type archiveGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ArchiveClient
	CallOptions **ArchiveCallOptions

	// The gRPC API client.
	archiveClient webpagepb.ArchiveClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewArchiveClient creates a new archive client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
func NewArchiveClient(ctx context.Context, opts ...option.ClientOption) (*ArchiveClient, error) {
	clientOpts := defaultArchiveGRPCClientOptions()
	if newArchiveClientHook != nil {
		hookOpts, err := newArchiveClientHook(ctx, clientHookParams{})
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
	client := ArchiveClient{CallOptions: defaultArchiveCallOptions()}

	c := &archiveGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		archiveClient:    webpagepb.NewArchiveClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *archiveGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *archiveGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *archiveGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *archiveGRPCClient) Query(ctx context.Context, opts ...gax.CallOption) (webpagepb.Archive_QueryClient, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	var resp webpagepb.Archive_QueryClient
	opts = append((*c.CallOptions).Query[0:len((*c.CallOptions).Query):len((*c.CallOptions).Query)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.archiveClient.Query(ctx, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *archiveGRPCClient) GetPage(ctx context.Context, req *webpagepb.GetPageRequest, opts ...gax.CallOption) (*webpagepb.Page, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetPage[0:len((*c.CallOptions).GetPage):len((*c.CallOptions).GetPage)], opts...)
	var resp *webpagepb.Page
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.archiveClient.GetPage(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *archiveGRPCClient) ListPages(ctx context.Context, req *webpagepb.ListPagesRequest, opts ...gax.CallOption) *PageIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListPages[0:len((*c.CallOptions).ListPages):len((*c.CallOptions).ListPages)], opts...)
	it := &PageIterator{}
	req = proto.Clone(req).(*webpagepb.ListPagesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*webpagepb.Page, string, error) {
		resp := &webpagepb.ListPagesResponse{}
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
			resp, err = c.archiveClient.ListPages(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetPages(), resp.GetNextPageToken(), nil
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

func (c *archiveGRPCClient) QueryPage(ctx context.Context, req *webpagepb.QueryPageRequest, opts ...gax.CallOption) (*webpagepb.QueryPageResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).QueryPage[0:len((*c.CallOptions).QueryPage):len((*c.CallOptions).QueryPage)], opts...)
	var resp *webpagepb.QueryPageResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.archiveClient.QueryPage(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *archiveGRPCClient) CreatePage(ctx context.Context, req *webpagepb.CreatePageRequest, opts ...gax.CallOption) (*webpagepb.Page, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreatePage[0:len((*c.CallOptions).CreatePage):len((*c.CallOptions).CreatePage)], opts...)
	var resp *webpagepb.Page
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.archiveClient.CreatePage(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *archiveGRPCClient) ImportPage(ctx context.Context, req *webpagepb.ImportPageRequest, opts ...gax.CallOption) (*webpagepb.ImportPageResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ImportPage[0:len((*c.CallOptions).ImportPage):len((*c.CallOptions).ImportPage)], opts...)
	var resp *webpagepb.ImportPageResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.archiveClient.ImportPage(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *archiveGRPCClient) DeletePage(ctx context.Context, req *webpagepb.DeletePageRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeletePage[0:len((*c.CallOptions).DeletePage):len((*c.CallOptions).DeletePage)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.archiveClient.DeletePage(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// PageIterator manages a stream of *webpagepb.Page.
type PageIterator struct {
	items    []*webpagepb.Page
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
	InternalFetch func(pageSize int, pageToken string) (results []*webpagepb.Page, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *PageIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *PageIterator) Next() (*webpagepb.Page, error) {
	var item *webpagepb.Page
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *PageIterator) bufLen() int {
	return len(it.items)
}

func (it *PageIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
