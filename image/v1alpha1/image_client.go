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

package image

import (
	"context"
	"fmt"
	"math"
	"net/url"

	imagepb "github.com/animeapis/go-genproto/image/v1alpha1"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httpbodypb "google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	UploadImage         []gax.CallOption
	ImportImage         []gax.CallOption
	GetImage            []gax.CallOption
	GetAlbum            []gax.CallOption
	ListAlbums          []gax.CallOption
	CreateAlbum         []gax.CallOption
	DeleteAlbum         []gax.CallOption
	GetAlbumSettings    []gax.CallOption
	UpdateAlbumSettings []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("image.animeapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("image.animeapis.com:443"),
		internaloption.WithDefaultAudience("https://image.animeapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		UploadImage:         []gax.CallOption{},
		ImportImage:         []gax.CallOption{},
		GetImage:            []gax.CallOption{},
		GetAlbum:            []gax.CallOption{},
		ListAlbums:          []gax.CallOption{},
		CreateAlbum:         []gax.CallOption{},
		DeleteAlbum:         []gax.CallOption{},
		GetAlbumSettings:    []gax.CallOption{},
		UpdateAlbumSettings: []gax.CallOption{},
	}
}

// internalClient is an interface that defines the methods availaible from Image API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	UploadImage(context.Context, *imagepb.UploadImageRequest, ...gax.CallOption) (*imagepb.UploadImageResponse, error)
	ImportImage(context.Context, *imagepb.ImportImageRequest, ...gax.CallOption) (*imagepb.ImportImageResponse, error)
	GetImage(context.Context, *imagepb.GetImageRequest, ...gax.CallOption) (*httpbodypb.HttpBody, error)
	GetAlbum(context.Context, *imagepb.GetAlbumRequest, ...gax.CallOption) (*imagepb.Album, error)
	ListAlbums(context.Context, *imagepb.ListAlbumsRequest, ...gax.CallOption) *AlbumIterator
	CreateAlbum(context.Context, *imagepb.CreateAlbumRequest, ...gax.CallOption) (*imagepb.Album, error)
	DeleteAlbum(context.Context, *imagepb.DeleteAlbumRequest, ...gax.CallOption) error
	GetAlbumSettings(context.Context, *imagepb.GetAlbumSettingsRequest, ...gax.CallOption) (*imagepb.AlbumSettings, error)
	UpdateAlbumSettings(context.Context, *imagepb.UpdateAlbumSettingsRequest, ...gax.CallOption) (*imagepb.AlbumSettings, error)
}

// Client is a client for interacting with Image API.
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

// UploadImage uploads an image through the request HttpBody.
func (c *Client) UploadImage(ctx context.Context, req *imagepb.UploadImageRequest, opts ...gax.CallOption) (*imagepb.UploadImageResponse, error) {
	return c.internalClient.UploadImage(ctx, req, opts...)
}

// ImportImage imports an image from a remote web address.
func (c *Client) ImportImage(ctx context.Context, req *imagepb.ImportImageRequest, opts ...gax.CallOption) (*imagepb.ImportImageResponse, error) {
	return c.internalClient.ImportImage(ctx, req, opts...)
}

// GetImage gets an image in binary representation with the format and size requested.
func (c *Client) GetImage(ctx context.Context, req *imagepb.GetImageRequest, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	return c.internalClient.GetImage(ctx, req, opts...)
}

// GetAlbum gets an image album.
func (c *Client) GetAlbum(ctx context.Context, req *imagepb.GetAlbumRequest, opts ...gax.CallOption) (*imagepb.Album, error) {
	return c.internalClient.GetAlbum(ctx, req, opts...)
}

// ListAlbums lists image albums with pagination.
func (c *Client) ListAlbums(ctx context.Context, req *imagepb.ListAlbumsRequest, opts ...gax.CallOption) *AlbumIterator {
	return c.internalClient.ListAlbums(ctx, req, opts...)
}

// CreateAlbum creates a new image album.
func (c *Client) CreateAlbum(ctx context.Context, req *imagepb.CreateAlbumRequest, opts ...gax.CallOption) (*imagepb.Album, error) {
	return c.internalClient.CreateAlbum(ctx, req, opts...)
}

// DeleteAlbum deletes an existing image album.
func (c *Client) DeleteAlbum(ctx context.Context, req *imagepb.DeleteAlbumRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteAlbum(ctx, req, opts...)
}

// GetAlbumSettings gets the settings of an image album.
func (c *Client) GetAlbumSettings(ctx context.Context, req *imagepb.GetAlbumSettingsRequest, opts ...gax.CallOption) (*imagepb.AlbumSettings, error) {
	return c.internalClient.GetAlbumSettings(ctx, req, opts...)
}

// UpdateAlbumSettings updates the settings of an image album.
func (c *Client) UpdateAlbumSettings(ctx context.Context, req *imagepb.UpdateAlbumSettingsRequest, opts ...gax.CallOption) (*imagepb.AlbumSettings, error) {
	return c.internalClient.UpdateAlbumSettings(ctx, req, opts...)
}

// gRPCClient is a client for interacting with Image API over gRPC transport.
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
	client imagepb.ImageClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new image client based on gRPC.
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
		client:           imagepb.NewImageClient(connPool),
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

func (c *gRPCClient) UploadImage(ctx context.Context, req *imagepb.UploadImageRequest, opts ...gax.CallOption) (*imagepb.UploadImageResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UploadImage[0:len((*c.CallOptions).UploadImage):len((*c.CallOptions).UploadImage)], opts...)
	var resp *imagepb.UploadImageResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UploadImage(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ImportImage(ctx context.Context, req *imagepb.ImportImageRequest, opts ...gax.CallOption) (*imagepb.ImportImageResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ImportImage[0:len((*c.CallOptions).ImportImage):len((*c.CallOptions).ImportImage)], opts...)
	var resp *imagepb.ImportImageResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ImportImage(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) GetImage(ctx context.Context, req *imagepb.GetImageRequest, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetImage[0:len((*c.CallOptions).GetImage):len((*c.CallOptions).GetImage)], opts...)
	var resp *httpbodypb.HttpBody
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetImage(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) GetAlbum(ctx context.Context, req *imagepb.GetAlbumRequest, opts ...gax.CallOption) (*imagepb.Album, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetAlbum[0:len((*c.CallOptions).GetAlbum):len((*c.CallOptions).GetAlbum)], opts...)
	var resp *imagepb.Album
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetAlbum(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListAlbums(ctx context.Context, req *imagepb.ListAlbumsRequest, opts ...gax.CallOption) *AlbumIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListAlbums[0:len((*c.CallOptions).ListAlbums):len((*c.CallOptions).ListAlbums)], opts...)
	it := &AlbumIterator{}
	req = proto.Clone(req).(*imagepb.ListAlbumsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*imagepb.Album, string, error) {
		var resp *imagepb.ListAlbumsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListAlbums(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetAlbums(), resp.GetNextPageToken(), nil
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

func (c *gRPCClient) CreateAlbum(ctx context.Context, req *imagepb.CreateAlbumRequest, opts ...gax.CallOption) (*imagepb.Album, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateAlbum[0:len((*c.CallOptions).CreateAlbum):len((*c.CallOptions).CreateAlbum)], opts...)
	var resp *imagepb.Album
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateAlbum(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) DeleteAlbum(ctx context.Context, req *imagepb.DeleteAlbumRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteAlbum[0:len((*c.CallOptions).DeleteAlbum):len((*c.CallOptions).DeleteAlbum)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.DeleteAlbum(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *gRPCClient) GetAlbumSettings(ctx context.Context, req *imagepb.GetAlbumSettingsRequest, opts ...gax.CallOption) (*imagepb.AlbumSettings, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetAlbumSettings[0:len((*c.CallOptions).GetAlbumSettings):len((*c.CallOptions).GetAlbumSettings)], opts...)
	var resp *imagepb.AlbumSettings
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetAlbumSettings(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) UpdateAlbumSettings(ctx context.Context, req *imagepb.UpdateAlbumSettingsRequest, opts ...gax.CallOption) (*imagepb.AlbumSettings, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "settings.name", url.QueryEscape(req.GetSettings().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateAlbumSettings[0:len((*c.CallOptions).UpdateAlbumSettings):len((*c.CallOptions).UpdateAlbumSettings)], opts...)
	var resp *imagepb.AlbumSettings
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UpdateAlbumSettings(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// AlbumIterator manages a stream of *imagepb.Album.
type AlbumIterator struct {
	items    []*imagepb.Album
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
	InternalFetch func(pageSize int, pageToken string) (results []*imagepb.Album, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *AlbumIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *AlbumIterator) Next() (*imagepb.Album, error) {
	var item *imagepb.Album
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *AlbumIterator) bufLen() int {
	return len(it.items)
}

func (it *AlbumIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
