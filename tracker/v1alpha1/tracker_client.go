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

package tracker

import (
	"context"
	"fmt"
	"math"
	"net/url"

	trackerpb "github.com/animeapis/go-genproto/tracker/v1alpha1"
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
	GetTracker    []gax.CallOption
	ListTrackers  []gax.CallOption
	CreateTracker []gax.CallOption
	UpdateTracker []gax.CallOption
	DeleteTracker []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("tracker.animeapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("tracker.animeapis.com:443"),
		internaloption.WithDefaultAudience("https://tracker.animeapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		GetTracker:    []gax.CallOption{},
		ListTrackers:  []gax.CallOption{},
		CreateTracker: []gax.CallOption{},
		UpdateTracker: []gax.CallOption{},
		DeleteTracker: []gax.CallOption{},
	}
}

// internalClient is an interface that defines the methods availaible from Tracker API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetTracker(context.Context, *trackerpb.GetTrackerRequest, ...gax.CallOption) (*trackerpb.Tracker, error)
	ListTrackers(context.Context, *trackerpb.ListTrackersRequest, ...gax.CallOption) *TrackerIterator
	CreateTracker(context.Context, *trackerpb.CreateTrackerRequest, ...gax.CallOption) (*trackerpb.Tracker, error)
	UpdateTracker(context.Context, *trackerpb.UpdateTrackerRequest, ...gax.CallOption) (*trackerpb.Tracker, error)
	DeleteTracker(context.Context, *trackerpb.DeleteTrackerRequest, ...gax.CallOption) error
}

// Client is a client for interacting with Tracker API.
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

func (c *Client) GetTracker(ctx context.Context, req *trackerpb.GetTrackerRequest, opts ...gax.CallOption) (*trackerpb.Tracker, error) {
	return c.internalClient.GetTracker(ctx, req, opts...)
}

func (c *Client) ListTrackers(ctx context.Context, req *trackerpb.ListTrackersRequest, opts ...gax.CallOption) *TrackerIterator {
	return c.internalClient.ListTrackers(ctx, req, opts...)
}

func (c *Client) CreateTracker(ctx context.Context, req *trackerpb.CreateTrackerRequest, opts ...gax.CallOption) (*trackerpb.Tracker, error) {
	return c.internalClient.CreateTracker(ctx, req, opts...)
}

func (c *Client) UpdateTracker(ctx context.Context, req *trackerpb.UpdateTrackerRequest, opts ...gax.CallOption) (*trackerpb.Tracker, error) {
	return c.internalClient.UpdateTracker(ctx, req, opts...)
}

func (c *Client) DeleteTracker(ctx context.Context, req *trackerpb.DeleteTrackerRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteTracker(ctx, req, opts...)
}

// gRPCClient is a client for interacting with Tracker API over gRPC transport.
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
	client trackerpb.TrackerServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new tracker service client based on gRPC.
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
		client:           trackerpb.NewTrackerServiceClient(connPool),
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

func (c *gRPCClient) GetTracker(ctx context.Context, req *trackerpb.GetTrackerRequest, opts ...gax.CallOption) (*trackerpb.Tracker, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetTracker[0:len((*c.CallOptions).GetTracker):len((*c.CallOptions).GetTracker)], opts...)
	var resp *trackerpb.Tracker
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetTracker(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListTrackers(ctx context.Context, req *trackerpb.ListTrackersRequest, opts ...gax.CallOption) *TrackerIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListTrackers[0:len((*c.CallOptions).ListTrackers):len((*c.CallOptions).ListTrackers)], opts...)
	it := &TrackerIterator{}
	req = proto.Clone(req).(*trackerpb.ListTrackersRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*trackerpb.Tracker, string, error) {
		var resp *trackerpb.ListTrackersResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListTrackers(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetTrackers(), resp.GetNextPageToken(), nil
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

func (c *gRPCClient) CreateTracker(ctx context.Context, req *trackerpb.CreateTrackerRequest, opts ...gax.CallOption) (*trackerpb.Tracker, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateTracker[0:len((*c.CallOptions).CreateTracker):len((*c.CallOptions).CreateTracker)], opts...)
	var resp *trackerpb.Tracker
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateTracker(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) UpdateTracker(ctx context.Context, req *trackerpb.UpdateTrackerRequest, opts ...gax.CallOption) (*trackerpb.Tracker, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "tracker.name", url.QueryEscape(req.GetTracker().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateTracker[0:len((*c.CallOptions).UpdateTracker):len((*c.CallOptions).UpdateTracker)], opts...)
	var resp *trackerpb.Tracker
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UpdateTracker(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) DeleteTracker(ctx context.Context, req *trackerpb.DeleteTrackerRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteTracker[0:len((*c.CallOptions).DeleteTracker):len((*c.CallOptions).DeleteTracker)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.DeleteTracker(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// TrackerIterator manages a stream of *trackerpb.Tracker.
type TrackerIterator struct {
	items    []*trackerpb.Tracker
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
	InternalFetch func(pageSize int, pageToken string) (results []*trackerpb.Tracker, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *TrackerIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *TrackerIterator) Next() (*trackerpb.Tracker, error) {
	var item *trackerpb.Tracker
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *TrackerIterator) bufLen() int {
	return len(it.items)
}

func (it *TrackerIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
