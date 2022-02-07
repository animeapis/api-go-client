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

package multimedia

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	multimediapb "github.com/animeapis/go-genproto/multimedia/v1alpha1"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newAnimeClientHook clientHook

// AnimeCallOptions contains the retry settings for each method of AnimeClient.
type AnimeCallOptions struct {
	GetAnime        []gax.CallOption
	ListAnimes      []gax.CallOption
	CreateAnime     []gax.CallOption
	UpdateAnime     []gax.CallOption
	DeleteAnime     []gax.CallOption
	ReconcileAnimes []gax.CallOption
}

func defaultAnimeGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("multimedia.animeapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("multimedia.animeapis.com:443"),
		internaloption.WithDefaultAudience("https://multimedia.animeapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultAnimeCallOptions() *AnimeCallOptions {
	return &AnimeCallOptions{
		GetAnime:        []gax.CallOption{},
		ListAnimes:      []gax.CallOption{},
		CreateAnime:     []gax.CallOption{},
		UpdateAnime:     []gax.CallOption{},
		DeleteAnime:     []gax.CallOption{},
		ReconcileAnimes: []gax.CallOption{},
	}
}

// internalAnimeClient is an interface that defines the methods availaible from Multimedia API.
type internalAnimeClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetAnime(context.Context, *multimediapb.GetAnimeRequest, ...gax.CallOption) (*multimediapb.Anime, error)
	ListAnimes(context.Context, *multimediapb.ListAnimesRequest, ...gax.CallOption) *AnimeIterator
	CreateAnime(context.Context, *multimediapb.CreateAnimeRequest, ...gax.CallOption) (*multimediapb.Anime, error)
	UpdateAnime(context.Context, *multimediapb.UpdateAnimeRequest, ...gax.CallOption) (*multimediapb.Anime, error)
	DeleteAnime(context.Context, *multimediapb.DeleteAnimeRequest, ...gax.CallOption) error
	ReconcileAnimes(context.Context, *multimediapb.ReconcileAnimesRequest, ...gax.CallOption) (*ReconcileAnimesOperation, error)
	ReconcileAnimesOperation(name string) *ReconcileAnimesOperation
}

// AnimeClient is a client for interacting with Multimedia API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type AnimeClient struct {
	// The internal transport-dependent client.
	internalClient internalAnimeClient

	// The call options for this service.
	CallOptions *AnimeCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AnimeClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AnimeClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *AnimeClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

func (c *AnimeClient) GetAnime(ctx context.Context, req *multimediapb.GetAnimeRequest, opts ...gax.CallOption) (*multimediapb.Anime, error) {
	return c.internalClient.GetAnime(ctx, req, opts...)
}

func (c *AnimeClient) ListAnimes(ctx context.Context, req *multimediapb.ListAnimesRequest, opts ...gax.CallOption) *AnimeIterator {
	return c.internalClient.ListAnimes(ctx, req, opts...)
}

func (c *AnimeClient) CreateAnime(ctx context.Context, req *multimediapb.CreateAnimeRequest, opts ...gax.CallOption) (*multimediapb.Anime, error) {
	return c.internalClient.CreateAnime(ctx, req, opts...)
}

func (c *AnimeClient) UpdateAnime(ctx context.Context, req *multimediapb.UpdateAnimeRequest, opts ...gax.CallOption) (*multimediapb.Anime, error) {
	return c.internalClient.UpdateAnime(ctx, req, opts...)
}

func (c *AnimeClient) DeleteAnime(ctx context.Context, req *multimediapb.DeleteAnimeRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteAnime(ctx, req, opts...)
}

// ReconcileAnimes reconcile animes with the search and knowledge base.
func (c *AnimeClient) ReconcileAnimes(ctx context.Context, req *multimediapb.ReconcileAnimesRequest, opts ...gax.CallOption) (*ReconcileAnimesOperation, error) {
	return c.internalClient.ReconcileAnimes(ctx, req, opts...)
}

// ReconcileAnimesOperation returns a new ReconcileAnimesOperation from a given name.
// The name must be that of a previously created ReconcileAnimesOperation, possibly from a different process.
func (c *AnimeClient) ReconcileAnimesOperation(name string) *ReconcileAnimesOperation {
	return c.internalClient.ReconcileAnimesOperation(name)
}

// animeGRPCClient is a client for interacting with Multimedia API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type animeGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing AnimeClient
	CallOptions **AnimeCallOptions

	// The gRPC API client.
	animeClient multimediapb.AnimeServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewAnimeClient creates a new anime service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
func NewAnimeClient(ctx context.Context, opts ...option.ClientOption) (*AnimeClient, error) {
	clientOpts := defaultAnimeGRPCClientOptions()
	if newAnimeClientHook != nil {
		hookOpts, err := newAnimeClientHook(ctx, clientHookParams{})
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
	client := AnimeClient{CallOptions: defaultAnimeCallOptions()}

	c := &animeGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		animeClient:      multimediapb.NewAnimeServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *animeGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *animeGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *animeGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *animeGRPCClient) GetAnime(ctx context.Context, req *multimediapb.GetAnimeRequest, opts ...gax.CallOption) (*multimediapb.Anime, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetAnime[0:len((*c.CallOptions).GetAnime):len((*c.CallOptions).GetAnime)], opts...)
	var resp *multimediapb.Anime
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.animeClient.GetAnime(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *animeGRPCClient) ListAnimes(ctx context.Context, req *multimediapb.ListAnimesRequest, opts ...gax.CallOption) *AnimeIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ListAnimes[0:len((*c.CallOptions).ListAnimes):len((*c.CallOptions).ListAnimes)], opts...)
	it := &AnimeIterator{}
	req = proto.Clone(req).(*multimediapb.ListAnimesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*multimediapb.Anime, string, error) {
		resp := &multimediapb.ListAnimesResponse{}
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
			resp, err = c.animeClient.ListAnimes(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetAnimes(), resp.GetNextPageToken(), nil
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

func (c *animeGRPCClient) CreateAnime(ctx context.Context, req *multimediapb.CreateAnimeRequest, opts ...gax.CallOption) (*multimediapb.Anime, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).CreateAnime[0:len((*c.CallOptions).CreateAnime):len((*c.CallOptions).CreateAnime)], opts...)
	var resp *multimediapb.Anime
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.animeClient.CreateAnime(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *animeGRPCClient) UpdateAnime(ctx context.Context, req *multimediapb.UpdateAnimeRequest, opts ...gax.CallOption) (*multimediapb.Anime, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "anime.name", url.QueryEscape(req.GetAnime().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateAnime[0:len((*c.CallOptions).UpdateAnime):len((*c.CallOptions).UpdateAnime)], opts...)
	var resp *multimediapb.Anime
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.animeClient.UpdateAnime(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *animeGRPCClient) DeleteAnime(ctx context.Context, req *multimediapb.DeleteAnimeRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteAnime[0:len((*c.CallOptions).DeleteAnime):len((*c.CallOptions).DeleteAnime)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.animeClient.DeleteAnime(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *animeGRPCClient) ReconcileAnimes(ctx context.Context, req *multimediapb.ReconcileAnimesRequest, opts ...gax.CallOption) (*ReconcileAnimesOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ReconcileAnimes[0:len((*c.CallOptions).ReconcileAnimes):len((*c.CallOptions).ReconcileAnimes)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.animeClient.ReconcileAnimes(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &ReconcileAnimesOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// ReconcileAnimesOperation manages a long-running operation from ReconcileAnimes.
type ReconcileAnimesOperation struct {
	lro *longrunning.Operation
}

// ReconcileAnimesOperation returns a new ReconcileAnimesOperation from a given name.
// The name must be that of a previously created ReconcileAnimesOperation, possibly from a different process.
func (c *animeGRPCClient) ReconcileAnimesOperation(name string) *ReconcileAnimesOperation {
	return &ReconcileAnimesOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *ReconcileAnimesOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*multimediapb.ReconcileAnimesResponse, error) {
	var resp multimediapb.ReconcileAnimesResponse
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *ReconcileAnimesOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*multimediapb.ReconcileAnimesResponse, error) {
	var resp multimediapb.ReconcileAnimesResponse
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *ReconcileAnimesOperation) Metadata() (*multimediapb.OperationMetadata, error) {
	var meta multimediapb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *ReconcileAnimesOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *ReconcileAnimesOperation) Name() string {
	return op.lro.Name()
}

// AnimeIterator manages a stream of *multimediapb.Anime.
type AnimeIterator struct {
	items    []*multimediapb.Anime
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
	InternalFetch func(pageSize int, pageToken string) (results []*multimediapb.Anime, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *AnimeIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *AnimeIterator) Next() (*multimediapb.Anime, error) {
	var item *multimediapb.Anime
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *AnimeIterator) bufLen() int {
	return len(it.items)
}

func (it *AnimeIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
