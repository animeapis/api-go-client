// Copyright 2021 Google LLC
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

package webcache

import (
	"context"
	"fmt"
	"math"
	"net/url"

	webcachepb "github.com/animeapis/go-genproto/webcache/v1alpha1"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	CreateCache []gax.CallOption
	ListCaches  []gax.CallOption
	GetCache    []gax.CallOption
	DeleteCache []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("webcache.animeapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("webcache.animeapis.com:443"),
		internaloption.WithDefaultAudience("https://webcache.animeapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		CreateCache: []gax.CallOption{},
		ListCaches:  []gax.CallOption{},
		GetCache:    []gax.CallOption{},
		DeleteCache: []gax.CallOption{},
	}
}

// internalClient is an interface that defines the methods availaible from WebCache API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateCache(context.Context, *webcachepb.CreateCacheRequest, ...gax.CallOption) (*webcachepb.Cache, error)
	ListCaches(context.Context, *webcachepb.ListCachesRequest, ...gax.CallOption) *CacheIterator
	GetCache(context.Context, *webcachepb.GetCacheRequest, ...gax.CallOption) (*webcachepb.Cache, error)
	DeleteCache(context.Context, *webcachepb.DeleteCacheRequest, ...gax.CallOption) error
}

// Client is a client for interacting with WebCache API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type Client struct {
	// The internal transport-dependent client.
	internalClient internalClient

	// The call options for this service.
	CallOptions *CallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

func (c *Client) CreateCache(ctx context.Context, req *webcachepb.CreateCacheRequest, opts ...gax.CallOption) (*webcachepb.Cache, error) {
	return c.internalClient.CreateCache(ctx, req, opts...)
}

func (c *Client) ListCaches(ctx context.Context, req *webcachepb.ListCachesRequest, opts ...gax.CallOption) *CacheIterator {
	return c.internalClient.ListCaches(ctx, req, opts...)
}

func (c *Client) GetCache(ctx context.Context, req *webcachepb.GetCacheRequest, opts ...gax.CallOption) (*webcachepb.Cache, error) {
	return c.internalClient.GetCache(ctx, req, opts...)
}

func (c *Client) DeleteCache(ctx context.Context, req *webcachepb.DeleteCacheRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteCache(ctx, req, opts...)
}

// gRPCClient is a client for interacting with WebCache API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type gRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	// The gRPC API client.
	client webcachepb.WebCacheClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new web cache client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := defaultGRPCClientOptions()
	if newClientHook != nil {
		hookOpts, err := newClientHook(ctx, clientHookParams{})
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
	client := Client{CallOptions: defaultCallOptions()}

	c := &gRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		client:           webcachepb.NewWebCacheClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *gRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *gRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *gRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *gRPCClient) CreateCache(ctx context.Context, req *webcachepb.CreateCacheRequest, opts ...gax.CallOption) (*webcachepb.Cache, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).CreateCache[0:len((*c.CallOptions).CreateCache):len((*c.CallOptions).CreateCache)], opts...)
	var resp *webcachepb.Cache
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateCache(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListCaches(ctx context.Context, req *webcachepb.ListCachesRequest, opts ...gax.CallOption) *CacheIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ListCaches[0:len((*c.CallOptions).ListCaches):len((*c.CallOptions).ListCaches)], opts...)
	it := &CacheIterator{}
	req = proto.Clone(req).(*webcachepb.ListCachesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*webcachepb.Cache, string, error) {
		var resp *webcachepb.ListCachesResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListCaches(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetCaches(), resp.GetNextPageToken(), nil
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

func (c *gRPCClient) GetCache(ctx context.Context, req *webcachepb.GetCacheRequest, opts ...gax.CallOption) (*webcachepb.Cache, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetCache[0:len((*c.CallOptions).GetCache):len((*c.CallOptions).GetCache)], opts...)
	var resp *webcachepb.Cache
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetCache(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) DeleteCache(ctx context.Context, req *webcachepb.DeleteCacheRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteCache[0:len((*c.CallOptions).DeleteCache):len((*c.CallOptions).DeleteCache)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.DeleteCache(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// CacheIterator manages a stream of *webcachepb.Cache.
type CacheIterator struct {
	items    []*webcachepb.Cache
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
	InternalFetch func(pageSize int, pageToken string) (results []*webcachepb.Cache, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *CacheIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *CacheIterator) Next() (*webcachepb.Cache, error) {
	var item *webcachepb.Cache
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *CacheIterator) bufLen() int {
	return len(it.items)
}

func (it *CacheIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
