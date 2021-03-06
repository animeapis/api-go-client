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

package library

import (
	"context"
	"fmt"
	"math"
	"net/url"

	librarypb "github.com/animeapis/go-genproto/library/v1alpha1"
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
	GetPlaylist              []gax.CallOption
	ListPlaylists            []gax.CallOption
	CreatePlaylist           []gax.CallOption
	UpdatePlaylist           []gax.CallOption
	DeletePlaylist           []gax.CallOption
	ListPlaylistItems        []gax.CallOption
	CreatePlaylistItem       []gax.CallOption
	BatchCreatePlaylistItems []gax.CallOption
	DeletePlaylistItem       []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("library.animeapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("library.animeapis.com:443"),
		internaloption.WithDefaultAudience("https://library.animeapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		GetPlaylist:              []gax.CallOption{},
		ListPlaylists:            []gax.CallOption{},
		CreatePlaylist:           []gax.CallOption{},
		UpdatePlaylist:           []gax.CallOption{},
		DeletePlaylist:           []gax.CallOption{},
		ListPlaylistItems:        []gax.CallOption{},
		CreatePlaylistItem:       []gax.CallOption{},
		BatchCreatePlaylistItems: []gax.CallOption{},
		DeletePlaylistItem:       []gax.CallOption{},
	}
}

// internalClient is an interface that defines the methods availaible from Library API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetPlaylist(context.Context, *librarypb.GetPlaylistRequest, ...gax.CallOption) (*librarypb.Playlist, error)
	ListPlaylists(context.Context, *librarypb.ListPlaylistsRequest, ...gax.CallOption) *PlaylistIterator
	CreatePlaylist(context.Context, *librarypb.CreatePlaylistRequest, ...gax.CallOption) (*librarypb.Playlist, error)
	UpdatePlaylist(context.Context, *librarypb.UpdatePlaylistRequest, ...gax.CallOption) (*librarypb.Playlist, error)
	DeletePlaylist(context.Context, *librarypb.DeletePlaylistRequest, ...gax.CallOption) error
	ListPlaylistItems(context.Context, *librarypb.ListPlaylistItemsRequest, ...gax.CallOption) *PlaylistItemIterator
	CreatePlaylistItem(context.Context, *librarypb.CreatePlaylistItemRequest, ...gax.CallOption) (*librarypb.PlaylistItem, error)
	BatchCreatePlaylistItems(context.Context, *librarypb.BatchCreatePlaylistItemsRequest, ...gax.CallOption) (*librarypb.BatchCreatePlaylistItemsResponse, error)
	DeletePlaylistItem(context.Context, *librarypb.DeletePlaylistItemRequest, ...gax.CallOption) error
}

// Client is a client for interacting with Library API.
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

func (c *Client) GetPlaylist(ctx context.Context, req *librarypb.GetPlaylistRequest, opts ...gax.CallOption) (*librarypb.Playlist, error) {
	return c.internalClient.GetPlaylist(ctx, req, opts...)
}

func (c *Client) ListPlaylists(ctx context.Context, req *librarypb.ListPlaylistsRequest, opts ...gax.CallOption) *PlaylistIterator {
	return c.internalClient.ListPlaylists(ctx, req, opts...)
}

func (c *Client) CreatePlaylist(ctx context.Context, req *librarypb.CreatePlaylistRequest, opts ...gax.CallOption) (*librarypb.Playlist, error) {
	return c.internalClient.CreatePlaylist(ctx, req, opts...)
}

func (c *Client) UpdatePlaylist(ctx context.Context, req *librarypb.UpdatePlaylistRequest, opts ...gax.CallOption) (*librarypb.Playlist, error) {
	return c.internalClient.UpdatePlaylist(ctx, req, opts...)
}

func (c *Client) DeletePlaylist(ctx context.Context, req *librarypb.DeletePlaylistRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeletePlaylist(ctx, req, opts...)
}

func (c *Client) ListPlaylistItems(ctx context.Context, req *librarypb.ListPlaylistItemsRequest, opts ...gax.CallOption) *PlaylistItemIterator {
	return c.internalClient.ListPlaylistItems(ctx, req, opts...)
}

func (c *Client) CreatePlaylistItem(ctx context.Context, req *librarypb.CreatePlaylistItemRequest, opts ...gax.CallOption) (*librarypb.PlaylistItem, error) {
	return c.internalClient.CreatePlaylistItem(ctx, req, opts...)
}

func (c *Client) BatchCreatePlaylistItems(ctx context.Context, req *librarypb.BatchCreatePlaylistItemsRequest, opts ...gax.CallOption) (*librarypb.BatchCreatePlaylistItemsResponse, error) {
	return c.internalClient.BatchCreatePlaylistItems(ctx, req, opts...)
}

func (c *Client) DeletePlaylistItem(ctx context.Context, req *librarypb.DeletePlaylistItemRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeletePlaylistItem(ctx, req, opts...)
}

// gRPCClient is a client for interacting with Library API over gRPC transport.
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
	client librarypb.LibraryClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new library client based on gRPC.
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
		client:           librarypb.NewLibraryClient(connPool),
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

func (c *gRPCClient) GetPlaylist(ctx context.Context, req *librarypb.GetPlaylistRequest, opts ...gax.CallOption) (*librarypb.Playlist, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetPlaylist[0:len((*c.CallOptions).GetPlaylist):len((*c.CallOptions).GetPlaylist)], opts...)
	var resp *librarypb.Playlist
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetPlaylist(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListPlaylists(ctx context.Context, req *librarypb.ListPlaylistsRequest, opts ...gax.CallOption) *PlaylistIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListPlaylists[0:len((*c.CallOptions).ListPlaylists):len((*c.CallOptions).ListPlaylists)], opts...)
	it := &PlaylistIterator{}
	req = proto.Clone(req).(*librarypb.ListPlaylistsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*librarypb.Playlist, string, error) {
		resp := &librarypb.ListPlaylistsResponse{}
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
			resp, err = c.client.ListPlaylists(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetPlaylists(), resp.GetNextPageToken(), nil
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

func (c *gRPCClient) CreatePlaylist(ctx context.Context, req *librarypb.CreatePlaylistRequest, opts ...gax.CallOption) (*librarypb.Playlist, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreatePlaylist[0:len((*c.CallOptions).CreatePlaylist):len((*c.CallOptions).CreatePlaylist)], opts...)
	var resp *librarypb.Playlist
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreatePlaylist(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) UpdatePlaylist(ctx context.Context, req *librarypb.UpdatePlaylistRequest, opts ...gax.CallOption) (*librarypb.Playlist, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "playlist.name", url.QueryEscape(req.GetPlaylist().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdatePlaylist[0:len((*c.CallOptions).UpdatePlaylist):len((*c.CallOptions).UpdatePlaylist)], opts...)
	var resp *librarypb.Playlist
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UpdatePlaylist(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) DeletePlaylist(ctx context.Context, req *librarypb.DeletePlaylistRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeletePlaylist[0:len((*c.CallOptions).DeletePlaylist):len((*c.CallOptions).DeletePlaylist)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.DeletePlaylist(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *gRPCClient) ListPlaylistItems(ctx context.Context, req *librarypb.ListPlaylistItemsRequest, opts ...gax.CallOption) *PlaylistItemIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListPlaylistItems[0:len((*c.CallOptions).ListPlaylistItems):len((*c.CallOptions).ListPlaylistItems)], opts...)
	it := &PlaylistItemIterator{}
	req = proto.Clone(req).(*librarypb.ListPlaylistItemsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*librarypb.PlaylistItem, string, error) {
		resp := &librarypb.ListPlaylistItemsResponse{}
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
			resp, err = c.client.ListPlaylistItems(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetItems(), resp.GetNextPageToken(), nil
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

func (c *gRPCClient) CreatePlaylistItem(ctx context.Context, req *librarypb.CreatePlaylistItemRequest, opts ...gax.CallOption) (*librarypb.PlaylistItem, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreatePlaylistItem[0:len((*c.CallOptions).CreatePlaylistItem):len((*c.CallOptions).CreatePlaylistItem)], opts...)
	var resp *librarypb.PlaylistItem
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreatePlaylistItem(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) BatchCreatePlaylistItems(ctx context.Context, req *librarypb.BatchCreatePlaylistItemsRequest, opts ...gax.CallOption) (*librarypb.BatchCreatePlaylistItemsResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).BatchCreatePlaylistItems[0:len((*c.CallOptions).BatchCreatePlaylistItems):len((*c.CallOptions).BatchCreatePlaylistItems)], opts...)
	var resp *librarypb.BatchCreatePlaylistItemsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.BatchCreatePlaylistItems(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) DeletePlaylistItem(ctx context.Context, req *librarypb.DeletePlaylistItemRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeletePlaylistItem[0:len((*c.CallOptions).DeletePlaylistItem):len((*c.CallOptions).DeletePlaylistItem)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.DeletePlaylistItem(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// PlaylistItemIterator manages a stream of *librarypb.PlaylistItem.
type PlaylistItemIterator struct {
	items    []*librarypb.PlaylistItem
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
	InternalFetch func(pageSize int, pageToken string) (results []*librarypb.PlaylistItem, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *PlaylistItemIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *PlaylistItemIterator) Next() (*librarypb.PlaylistItem, error) {
	var item *librarypb.PlaylistItem
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *PlaylistItemIterator) bufLen() int {
	return len(it.items)
}

func (it *PlaylistItemIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// PlaylistIterator manages a stream of *librarypb.Playlist.
type PlaylistIterator struct {
	items    []*librarypb.Playlist
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
	InternalFetch func(pageSize int, pageToken string) (results []*librarypb.Playlist, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *PlaylistIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *PlaylistIterator) Next() (*librarypb.Playlist, error) {
	var item *librarypb.Playlist
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *PlaylistIterator) bufLen() int {
	return len(it.items)
}

func (it *PlaylistIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
