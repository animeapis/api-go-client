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

package crossrefs

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	crossrefspb "github.com/animeapis/go-genproto/crossrefs/v1alpha1"
	"github.com/golang/protobuf/proto"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var newReferrerClientHook clientHook

// ReferrerCallOptions contains the retry settings for each method of ReferrerClient.
type ReferrerCallOptions struct {
	GetCrossRef      []gax.CallOption
	CreateCrossRef   []gax.CallOption
	UpdateCrossRef   []gax.CallOption
	ListCrossRefs    []gax.CallOption
	AnalyzeCrossRefs []gax.CallOption
	ImportCrossRefs  []gax.CallOption
	ExportCrossRefs  []gax.CallOption
	AnalyzeParodies  []gax.CallOption
	ExportParodies   []gax.CallOption
	GetUniverse      []gax.CallOption
	UpdateUniverse   []gax.CallOption
	ExpandUniverse   []gax.CallOption
}

func defaultReferrerGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("crossrefs.animeapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("crossrefs.animeapis.com:443"),
		internaloption.WithDefaultAudience("https://crossrefs.animeapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultReferrerCallOptions() *ReferrerCallOptions {
	return &ReferrerCallOptions{
		GetCrossRef:      []gax.CallOption{},
		CreateCrossRef:   []gax.CallOption{},
		UpdateCrossRef:   []gax.CallOption{},
		ListCrossRefs:    []gax.CallOption{},
		AnalyzeCrossRefs: []gax.CallOption{},
		ImportCrossRefs:  []gax.CallOption{},
		ExportCrossRefs:  []gax.CallOption{},
		AnalyzeParodies:  []gax.CallOption{},
		ExportParodies:   []gax.CallOption{},
		GetUniverse:      []gax.CallOption{},
		UpdateUniverse:   []gax.CallOption{},
		ExpandUniverse:   []gax.CallOption{},
	}
}

// internalReferrerClient is an interface that defines the methods availaible from CrossRefs API.
type internalReferrerClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetCrossRef(context.Context, *crossrefspb.GetCrossRefRequest, ...gax.CallOption) (*crossrefspb.CrossRef, error)
	CreateCrossRef(context.Context, *crossrefspb.CreateCrossRefRequest, ...gax.CallOption) (*crossrefspb.CrossRef, error)
	UpdateCrossRef(context.Context, *crossrefspb.UpdateCrossRefRequest, ...gax.CallOption) (*crossrefspb.CrossRef, error)
	ListCrossRefs(context.Context, *crossrefspb.ListCrossRefsRequest, ...gax.CallOption) *CrossRefIterator
	AnalyzeCrossRefs(context.Context, *crossrefspb.AnalyzeCrossRefRequest, ...gax.CallOption) (*AnalyzeCrossRefsOperation, error)
	AnalyzeCrossRefsOperation(name string) *AnalyzeCrossRefsOperation
	ImportCrossRefs(context.Context, *crossrefspb.ImportCrossRefRequest, ...gax.CallOption) (*ImportCrossRefsOperation, error)
	ImportCrossRefsOperation(name string) *ImportCrossRefsOperation
	ExportCrossRefs(context.Context, *emptypb.Empty, ...gax.CallOption) (*ExportCrossRefsOperation, error)
	ExportCrossRefsOperation(name string) *ExportCrossRefsOperation
	AnalyzeParodies(context.Context, *emptypb.Empty, ...gax.CallOption) (*AnalyzeParodiesOperation, error)
	AnalyzeParodiesOperation(name string) *AnalyzeParodiesOperation
	ExportParodies(context.Context, *emptypb.Empty, ...gax.CallOption) (*ExportParodiesOperation, error)
	ExportParodiesOperation(name string) *ExportParodiesOperation
	GetUniverse(context.Context, *crossrefspb.GetUniverseRequest, ...gax.CallOption) (*crossrefspb.Universe, error)
	UpdateUniverse(context.Context, *crossrefspb.UpdateUniverseRequest, ...gax.CallOption) (*crossrefspb.Universe, error)
	ExpandUniverse(context.Context, *crossrefspb.ExpandUniverseRequest, ...gax.CallOption) (*crossrefspb.ExpandUniverseResponse, error)
}

// ReferrerClient is a client for interacting with CrossRefs API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type ReferrerClient struct {
	// The internal transport-dependent client.
	internalClient internalReferrerClient

	// The call options for this service.
	CallOptions *ReferrerCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ReferrerClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ReferrerClient) setGoogleClientInfo(...string) {
	c.internalClient.setGoogleClientInfo()
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ReferrerClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetCrossRef getCrossRef returns a crossref.
func (c *ReferrerClient) GetCrossRef(ctx context.Context, req *crossrefspb.GetCrossRefRequest, opts ...gax.CallOption) (*crossrefspb.CrossRef, error) {
	return c.internalClient.GetCrossRef(ctx, req, opts...)
}

// CreateCrossRef createCrossRef creates a new crossref.
func (c *ReferrerClient) CreateCrossRef(ctx context.Context, req *crossrefspb.CreateCrossRefRequest, opts ...gax.CallOption) (*crossrefspb.CrossRef, error) {
	return c.internalClient.CreateCrossRef(ctx, req, opts...)
}

func (c *ReferrerClient) UpdateCrossRef(ctx context.Context, req *crossrefspb.UpdateCrossRefRequest, opts ...gax.CallOption) (*crossrefspb.CrossRef, error) {
	return c.internalClient.UpdateCrossRef(ctx, req, opts...)
}

func (c *ReferrerClient) ListCrossRefs(ctx context.Context, req *crossrefspb.ListCrossRefsRequest, opts ...gax.CallOption) *CrossRefIterator {
	return c.internalClient.ListCrossRefs(ctx, req, opts...)
}

// AnalyzeCrossRefs analyzes and proposes new cross-references according to their similarity.
func (c *ReferrerClient) AnalyzeCrossRefs(ctx context.Context, req *crossrefspb.AnalyzeCrossRefRequest, opts ...gax.CallOption) (*AnalyzeCrossRefsOperation, error) {
	return c.internalClient.AnalyzeCrossRefs(ctx, req, opts...)
}

// AnalyzeCrossRefsOperation returns a new AnalyzeCrossRefsOperation from a given name.
// The name must be that of a previously created AnalyzeCrossRefsOperation, possibly from a different process.
func (c *ReferrerClient) AnalyzeCrossRefsOperation(name string) *AnalyzeCrossRefsOperation {
	return c.internalClient.AnalyzeCrossRefsOperation(name)
}

// ImportCrossRefs imports already existing cross-references from third-parties.
func (c *ReferrerClient) ImportCrossRefs(ctx context.Context, req *crossrefspb.ImportCrossRefRequest, opts ...gax.CallOption) (*ImportCrossRefsOperation, error) {
	return c.internalClient.ImportCrossRefs(ctx, req, opts...)
}

// ImportCrossRefsOperation returns a new ImportCrossRefsOperation from a given name.
// The name must be that of a previously created ImportCrossRefsOperation, possibly from a different process.
func (c *ReferrerClient) ImportCrossRefsOperation(name string) *ImportCrossRefsOperation {
	return c.internalClient.ImportCrossRefsOperation(name)
}

// ExportCrossRefs exports the cross-references to Cloud Pub/Sub for a full synchronization.
// This operation is usually called after a new import with a clean database.
func (c *ReferrerClient) ExportCrossRefs(ctx context.Context, req *emptypb.Empty, opts ...gax.CallOption) (*ExportCrossRefsOperation, error) {
	return c.internalClient.ExportCrossRefs(ctx, req, opts...)
}

// ExportCrossRefsOperation returns a new ExportCrossRefsOperation from a given name.
// The name must be that of a previously created ExportCrossRefsOperation, possibly from a different process.
func (c *ReferrerClient) ExportCrossRefsOperation(name string) *ExportCrossRefsOperation {
	return c.internalClient.ExportCrossRefsOperation(name)
}

func (c *ReferrerClient) AnalyzeParodies(ctx context.Context, req *emptypb.Empty, opts ...gax.CallOption) (*AnalyzeParodiesOperation, error) {
	return c.internalClient.AnalyzeParodies(ctx, req, opts...)
}

// AnalyzeParodiesOperation returns a new AnalyzeParodiesOperation from a given name.
// The name must be that of a previously created AnalyzeParodiesOperation, possibly from a different process.
func (c *ReferrerClient) AnalyzeParodiesOperation(name string) *AnalyzeParodiesOperation {
	return c.internalClient.AnalyzeParodiesOperation(name)
}

func (c *ReferrerClient) ExportParodies(ctx context.Context, req *emptypb.Empty, opts ...gax.CallOption) (*ExportParodiesOperation, error) {
	return c.internalClient.ExportParodies(ctx, req, opts...)
}

// ExportParodiesOperation returns a new ExportParodiesOperation from a given name.
// The name must be that of a previously created ExportParodiesOperation, possibly from a different process.
func (c *ReferrerClient) ExportParodiesOperation(name string) *ExportParodiesOperation {
	return c.internalClient.ExportParodiesOperation(name)
}

func (c *ReferrerClient) GetUniverse(ctx context.Context, req *crossrefspb.GetUniverseRequest, opts ...gax.CallOption) (*crossrefspb.Universe, error) {
	return c.internalClient.GetUniverse(ctx, req, opts...)
}

func (c *ReferrerClient) UpdateUniverse(ctx context.Context, req *crossrefspb.UpdateUniverseRequest, opts ...gax.CallOption) (*crossrefspb.Universe, error) {
	return c.internalClient.UpdateUniverse(ctx, req, opts...)
}

func (c *ReferrerClient) ExpandUniverse(ctx context.Context, req *crossrefspb.ExpandUniverseRequest, opts ...gax.CallOption) (*crossrefspb.ExpandUniverseResponse, error) {
	return c.internalClient.ExpandUniverse(ctx, req, opts...)
}

// referrerGRPCClient is a client for interacting with CrossRefs API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type referrerGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ReferrerClient
	CallOptions **ReferrerCallOptions

	// The gRPC API client.
	referrerClient crossrefspb.ReferrerClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewReferrerClient creates a new referrer client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
func NewReferrerClient(ctx context.Context, opts ...option.ClientOption) (*ReferrerClient, error) {
	clientOpts := defaultReferrerGRPCClientOptions()
	if newReferrerClientHook != nil {
		hookOpts, err := newReferrerClientHook(ctx, clientHookParams{})
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
	client := ReferrerClient{CallOptions: defaultReferrerCallOptions()}

	c := &referrerGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		referrerClient:   crossrefspb.NewReferrerClient(connPool),
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
func (c *referrerGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *referrerGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *referrerGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *referrerGRPCClient) GetCrossRef(ctx context.Context, req *crossrefspb.GetCrossRefRequest, opts ...gax.CallOption) (*crossrefspb.CrossRef, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetCrossRef[0:len((*c.CallOptions).GetCrossRef):len((*c.CallOptions).GetCrossRef)], opts...)
	var resp *crossrefspb.CrossRef
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.referrerClient.GetCrossRef(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *referrerGRPCClient) CreateCrossRef(ctx context.Context, req *crossrefspb.CreateCrossRefRequest, opts ...gax.CallOption) (*crossrefspb.CrossRef, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).CreateCrossRef[0:len((*c.CallOptions).CreateCrossRef):len((*c.CallOptions).CreateCrossRef)], opts...)
	var resp *crossrefspb.CrossRef
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.referrerClient.CreateCrossRef(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *referrerGRPCClient) UpdateCrossRef(ctx context.Context, req *crossrefspb.UpdateCrossRefRequest, opts ...gax.CallOption) (*crossrefspb.CrossRef, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "crossref.name", url.QueryEscape(req.GetCrossref().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateCrossRef[0:len((*c.CallOptions).UpdateCrossRef):len((*c.CallOptions).UpdateCrossRef)], opts...)
	var resp *crossrefspb.CrossRef
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.referrerClient.UpdateCrossRef(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *referrerGRPCClient) ListCrossRefs(ctx context.Context, req *crossrefspb.ListCrossRefsRequest, opts ...gax.CallOption) *CrossRefIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ListCrossRefs[0:len((*c.CallOptions).ListCrossRefs):len((*c.CallOptions).ListCrossRefs)], opts...)
	it := &CrossRefIterator{}
	req = proto.Clone(req).(*crossrefspb.ListCrossRefsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*crossrefspb.CrossRef, string, error) {
		var resp *crossrefspb.ListCrossRefsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.referrerClient.ListCrossRefs(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetCrossrefs(), resp.GetNextPageToken(), nil
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

func (c *referrerGRPCClient) AnalyzeCrossRefs(ctx context.Context, req *crossrefspb.AnalyzeCrossRefRequest, opts ...gax.CallOption) (*AnalyzeCrossRefsOperation, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).AnalyzeCrossRefs[0:len((*c.CallOptions).AnalyzeCrossRefs):len((*c.CallOptions).AnalyzeCrossRefs)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.referrerClient.AnalyzeCrossRefs(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &AnalyzeCrossRefsOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *referrerGRPCClient) ImportCrossRefs(ctx context.Context, req *crossrefspb.ImportCrossRefRequest, opts ...gax.CallOption) (*ImportCrossRefsOperation, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ImportCrossRefs[0:len((*c.CallOptions).ImportCrossRefs):len((*c.CallOptions).ImportCrossRefs)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.referrerClient.ImportCrossRefs(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &ImportCrossRefsOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *referrerGRPCClient) ExportCrossRefs(ctx context.Context, req *emptypb.Empty, opts ...gax.CallOption) (*ExportCrossRefsOperation, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ExportCrossRefs[0:len((*c.CallOptions).ExportCrossRefs):len((*c.CallOptions).ExportCrossRefs)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.referrerClient.ExportCrossRefs(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &ExportCrossRefsOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *referrerGRPCClient) AnalyzeParodies(ctx context.Context, req *emptypb.Empty, opts ...gax.CallOption) (*AnalyzeParodiesOperation, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).AnalyzeParodies[0:len((*c.CallOptions).AnalyzeParodies):len((*c.CallOptions).AnalyzeParodies)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.referrerClient.AnalyzeParodies(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &AnalyzeParodiesOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *referrerGRPCClient) ExportParodies(ctx context.Context, req *emptypb.Empty, opts ...gax.CallOption) (*ExportParodiesOperation, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ExportParodies[0:len((*c.CallOptions).ExportParodies):len((*c.CallOptions).ExportParodies)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.referrerClient.ExportParodies(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &ExportParodiesOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *referrerGRPCClient) GetUniverse(ctx context.Context, req *crossrefspb.GetUniverseRequest, opts ...gax.CallOption) (*crossrefspb.Universe, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetUniverse[0:len((*c.CallOptions).GetUniverse):len((*c.CallOptions).GetUniverse)], opts...)
	var resp *crossrefspb.Universe
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.referrerClient.GetUniverse(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *referrerGRPCClient) UpdateUniverse(ctx context.Context, req *crossrefspb.UpdateUniverseRequest, opts ...gax.CallOption) (*crossrefspb.Universe, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "universe.name", url.QueryEscape(req.GetUniverse().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateUniverse[0:len((*c.CallOptions).UpdateUniverse):len((*c.CallOptions).UpdateUniverse)], opts...)
	var resp *crossrefspb.Universe
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.referrerClient.UpdateUniverse(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *referrerGRPCClient) ExpandUniverse(ctx context.Context, req *crossrefspb.ExpandUniverseRequest, opts ...gax.CallOption) (*crossrefspb.ExpandUniverseResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ExpandUniverse[0:len((*c.CallOptions).ExpandUniverse):len((*c.CallOptions).ExpandUniverse)], opts...)
	var resp *crossrefspb.ExpandUniverseResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.referrerClient.ExpandUniverse(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// AnalyzeCrossRefsOperation manages a long-running operation from AnalyzeCrossRefs.
type AnalyzeCrossRefsOperation struct {
	lro *longrunning.Operation
}

// AnalyzeCrossRefsOperation returns a new AnalyzeCrossRefsOperation from a given name.
// The name must be that of a previously created AnalyzeCrossRefsOperation, possibly from a different process.
func (c *referrerGRPCClient) AnalyzeCrossRefsOperation(name string) *AnalyzeCrossRefsOperation {
	return &AnalyzeCrossRefsOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *AnalyzeCrossRefsOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*crossrefspb.AnalyzeCrossRefsResponse, error) {
	var resp crossrefspb.AnalyzeCrossRefsResponse
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
func (op *AnalyzeCrossRefsOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*crossrefspb.AnalyzeCrossRefsResponse, error) {
	var resp crossrefspb.AnalyzeCrossRefsResponse
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
func (op *AnalyzeCrossRefsOperation) Metadata() (*crossrefspb.OperationMetadata, error) {
	var meta crossrefspb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *AnalyzeCrossRefsOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *AnalyzeCrossRefsOperation) Name() string {
	return op.lro.Name()
}

// AnalyzeParodiesOperation manages a long-running operation from AnalyzeParodies.
type AnalyzeParodiesOperation struct {
	lro *longrunning.Operation
}

// AnalyzeParodiesOperation returns a new AnalyzeParodiesOperation from a given name.
// The name must be that of a previously created AnalyzeParodiesOperation, possibly from a different process.
func (c *referrerGRPCClient) AnalyzeParodiesOperation(name string) *AnalyzeParodiesOperation {
	return &AnalyzeParodiesOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *AnalyzeParodiesOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*crossrefspb.AnalyzeParodiesResponse, error) {
	var resp crossrefspb.AnalyzeParodiesResponse
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
func (op *AnalyzeParodiesOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*crossrefspb.AnalyzeParodiesResponse, error) {
	var resp crossrefspb.AnalyzeParodiesResponse
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
func (op *AnalyzeParodiesOperation) Metadata() (*crossrefspb.OperationMetadata, error) {
	var meta crossrefspb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *AnalyzeParodiesOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *AnalyzeParodiesOperation) Name() string {
	return op.lro.Name()
}

// ExportCrossRefsOperation manages a long-running operation from ExportCrossRefs.
type ExportCrossRefsOperation struct {
	lro *longrunning.Operation
}

// ExportCrossRefsOperation returns a new ExportCrossRefsOperation from a given name.
// The name must be that of a previously created ExportCrossRefsOperation, possibly from a different process.
func (c *referrerGRPCClient) ExportCrossRefsOperation(name string) *ExportCrossRefsOperation {
	return &ExportCrossRefsOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *ExportCrossRefsOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*crossrefspb.ExportCrossRefsResponse, error) {
	var resp crossrefspb.ExportCrossRefsResponse
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
func (op *ExportCrossRefsOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*crossrefspb.ExportCrossRefsResponse, error) {
	var resp crossrefspb.ExportCrossRefsResponse
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
func (op *ExportCrossRefsOperation) Metadata() (*crossrefspb.OperationMetadata, error) {
	var meta crossrefspb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *ExportCrossRefsOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *ExportCrossRefsOperation) Name() string {
	return op.lro.Name()
}

// ExportParodiesOperation manages a long-running operation from ExportParodies.
type ExportParodiesOperation struct {
	lro *longrunning.Operation
}

// ExportParodiesOperation returns a new ExportParodiesOperation from a given name.
// The name must be that of a previously created ExportParodiesOperation, possibly from a different process.
func (c *referrerGRPCClient) ExportParodiesOperation(name string) *ExportParodiesOperation {
	return &ExportParodiesOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *ExportParodiesOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*crossrefspb.ExportParodiesResponse, error) {
	var resp crossrefspb.ExportParodiesResponse
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
func (op *ExportParodiesOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*crossrefspb.ExportParodiesResponse, error) {
	var resp crossrefspb.ExportParodiesResponse
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
func (op *ExportParodiesOperation) Metadata() (*crossrefspb.OperationMetadata, error) {
	var meta crossrefspb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *ExportParodiesOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *ExportParodiesOperation) Name() string {
	return op.lro.Name()
}

// ImportCrossRefsOperation manages a long-running operation from ImportCrossRefs.
type ImportCrossRefsOperation struct {
	lro *longrunning.Operation
}

// ImportCrossRefsOperation returns a new ImportCrossRefsOperation from a given name.
// The name must be that of a previously created ImportCrossRefsOperation, possibly from a different process.
func (c *referrerGRPCClient) ImportCrossRefsOperation(name string) *ImportCrossRefsOperation {
	return &ImportCrossRefsOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *ImportCrossRefsOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*crossrefspb.ImportCrossRefsResponse, error) {
	var resp crossrefspb.ImportCrossRefsResponse
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
func (op *ImportCrossRefsOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*crossrefspb.ImportCrossRefsResponse, error) {
	var resp crossrefspb.ImportCrossRefsResponse
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
func (op *ImportCrossRefsOperation) Metadata() (*crossrefspb.OperationMetadata, error) {
	var meta crossrefspb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *ImportCrossRefsOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *ImportCrossRefsOperation) Name() string {
	return op.lro.Name()
}

// CrossRefIterator manages a stream of *crossrefspb.CrossRef.
type CrossRefIterator struct {
	items    []*crossrefspb.CrossRef
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
	InternalFetch func(pageSize int, pageToken string) (results []*crossrefspb.CrossRef, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *CrossRefIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *CrossRefIterator) Next() (*crossrefspb.CrossRef, error) {
	var item *crossrefspb.CrossRef
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *CrossRefIterator) bufLen() int {
	return len(it.items)
}

func (it *CrossRefIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
