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

package release

import (
	"context"
	"fmt"
	"math"
	"net/url"

	releasepb "github.com/animeapis/go-genproto/release/v1alpha1"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newPublisherClientHook clientHook

// PublisherCallOptions contains the retry settings for each method of PublisherClient.
type PublisherCallOptions struct {
	GetRelease       []gax.CallOption
	ListReleases     []gax.CallOption
	CreateRelease    []gax.CallOption
	UpdateRelease    []gax.CallOption
	DeleteRelease    []gax.CallOption
	UndeleteRelease  []gax.CallOption
	PublishRelease   []gax.CallOption
	UnpublishRelease []gax.CallOption
	ScheduleRelease  []gax.CallOption
	CancelRelease    []gax.CallOption
	SuspendRelease   []gax.CallOption
}

func defaultPublisherGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("release.animeapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("release.animeapis.com:443"),
		internaloption.WithDefaultAudience("https://release.animeapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultPublisherCallOptions() *PublisherCallOptions {
	return &PublisherCallOptions{
		GetRelease:       []gax.CallOption{},
		ListReleases:     []gax.CallOption{},
		CreateRelease:    []gax.CallOption{},
		UpdateRelease:    []gax.CallOption{},
		DeleteRelease:    []gax.CallOption{},
		UndeleteRelease:  []gax.CallOption{},
		PublishRelease:   []gax.CallOption{},
		UnpublishRelease: []gax.CallOption{},
		ScheduleRelease:  []gax.CallOption{},
		CancelRelease:    []gax.CallOption{},
		SuspendRelease:   []gax.CallOption{},
	}
}

// internalPublisherClient is an interface that defines the methods availaible from Release API.
type internalPublisherClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetRelease(context.Context, *releasepb.GetReleaseRequest, ...gax.CallOption) (*releasepb.Release, error)
	ListReleases(context.Context, *releasepb.ListReleasesRequest, ...gax.CallOption) *ReleaseIterator
	CreateRelease(context.Context, *releasepb.CreateReleaseRequest, ...gax.CallOption) (*releasepb.Release, error)
	UpdateRelease(context.Context, *releasepb.UpdateReleaseRequest, ...gax.CallOption) (*releasepb.Release, error)
	DeleteRelease(context.Context, *releasepb.DeleteReleaseRequest, ...gax.CallOption) error
	UndeleteRelease(context.Context, *releasepb.UndeleteReleaseRequest, ...gax.CallOption) (*releasepb.Release, error)
	PublishRelease(context.Context, *releasepb.PublishReleaseRequest, ...gax.CallOption) (*releasepb.PublishReleaseResponse, error)
	UnpublishRelease(context.Context, *releasepb.UnpublishReleaseRequest, ...gax.CallOption) (*releasepb.UnpublishReleaseResponse, error)
	ScheduleRelease(context.Context, *releasepb.ScheduleReleaseRequest, ...gax.CallOption) (*releasepb.ScheduleReleaseResponse, error)
	CancelRelease(context.Context, *releasepb.CancelReleaseRequest, ...gax.CallOption) (*releasepb.CancelReleaseResponse, error)
	SuspendRelease(context.Context, *releasepb.SuspendReleaseRequest, ...gax.CallOption) (*releasepb.SuspendReleaseResponse, error)
}

// PublisherClient is a client for interacting with Release API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type PublisherClient struct {
	// The internal transport-dependent client.
	internalClient internalPublisherClient

	// The call options for this service.
	CallOptions *PublisherCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *PublisherClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *PublisherClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *PublisherClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

func (c *PublisherClient) GetRelease(ctx context.Context, req *releasepb.GetReleaseRequest, opts ...gax.CallOption) (*releasepb.Release, error) {
	return c.internalClient.GetRelease(ctx, req, opts...)
}

func (c *PublisherClient) ListReleases(ctx context.Context, req *releasepb.ListReleasesRequest, opts ...gax.CallOption) *ReleaseIterator {
	return c.internalClient.ListReleases(ctx, req, opts...)
}

func (c *PublisherClient) CreateRelease(ctx context.Context, req *releasepb.CreateReleaseRequest, opts ...gax.CallOption) (*releasepb.Release, error) {
	return c.internalClient.CreateRelease(ctx, req, opts...)
}

func (c *PublisherClient) UpdateRelease(ctx context.Context, req *releasepb.UpdateReleaseRequest, opts ...gax.CallOption) (*releasepb.Release, error) {
	return c.internalClient.UpdateRelease(ctx, req, opts...)
}

// DeleteRelease the release is soft-deleted and a grace period is granted before complete
// deletion. During this grace period the release can be recovered.
func (c *PublisherClient) DeleteRelease(ctx context.Context, req *releasepb.DeleteReleaseRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteRelease(ctx, req, opts...)
}

// UndeleteRelease this method allows to recover a release while still in the grace period.
func (c *PublisherClient) UndeleteRelease(ctx context.Context, req *releasepb.UndeleteReleaseRequest, opts ...gax.CallOption) (*releasepb.Release, error) {
	return c.internalClient.UndeleteRelease(ctx, req, opts...)
}

// PublishRelease the release is marked as immediately available to the public.
func (c *PublisherClient) PublishRelease(ctx context.Context, req *releasepb.PublishReleaseRequest, opts ...gax.CallOption) (*releasepb.PublishReleaseResponse, error) {
	return c.internalClient.PublishRelease(ctx, req, opts...)
}

// UnpublishRelease the release is unpublished and marked as a draft, associated
// non-authoritative will automatically be marked as suspended and hidden from
// the general public.
func (c *PublisherClient) UnpublishRelease(ctx context.Context, req *releasepb.UnpublishReleaseRequest, opts ...gax.CallOption) (*releasepb.UnpublishReleaseResponse, error) {
	return c.internalClient.UnpublishRelease(ctx, req, opts...)
}

// ScheduleRelease the release is scheduled to be released at a specific future date and time.
func (c *PublisherClient) ScheduleRelease(ctx context.Context, req *releasepb.ScheduleReleaseRequest, opts ...gax.CallOption) (*releasepb.ScheduleReleaseResponse, error) {
	return c.internalClient.ScheduleRelease(ctx, req, opts...)
}

// CancelRelease this method can only be called on scheduled releases. The scheduling is
// cancelled and the release is marked as a draft.
func (c *PublisherClient) CancelRelease(ctx context.Context, req *releasepb.CancelReleaseRequest, opts ...gax.CallOption) (*releasepb.CancelReleaseResponse, error) {
	return c.internalClient.CancelRelease(ctx, req, opts...)
}

// SuspendRelease this method can only be called on published releases marked as active. Any
// non-authoritative release associated to the specified release will also be
// automatically marked as suspended.
func (c *PublisherClient) SuspendRelease(ctx context.Context, req *releasepb.SuspendReleaseRequest, opts ...gax.CallOption) (*releasepb.SuspendReleaseResponse, error) {
	return c.internalClient.SuspendRelease(ctx, req, opts...)
}

// publisherGRPCClient is a client for interacting with Release API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type publisherGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing PublisherClient
	CallOptions **PublisherCallOptions

	// The gRPC API client.
	publisherClient releasepb.PublisherClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewPublisherClient creates a new publisher client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
func NewPublisherClient(ctx context.Context, opts ...option.ClientOption) (*PublisherClient, error) {
	clientOpts := defaultPublisherGRPCClientOptions()
	if newPublisherClientHook != nil {
		hookOpts, err := newPublisherClientHook(ctx, clientHookParams{})
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
	client := PublisherClient{CallOptions: defaultPublisherCallOptions()}

	c := &publisherGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		publisherClient:  releasepb.NewPublisherClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *publisherGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *publisherGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *publisherGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *publisherGRPCClient) GetRelease(ctx context.Context, req *releasepb.GetReleaseRequest, opts ...gax.CallOption) (*releasepb.Release, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetRelease[0:len((*c.CallOptions).GetRelease):len((*c.CallOptions).GetRelease)], opts...)
	var resp *releasepb.Release
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.publisherClient.GetRelease(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *publisherGRPCClient) ListReleases(ctx context.Context, req *releasepb.ListReleasesRequest, opts ...gax.CallOption) *ReleaseIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListReleases[0:len((*c.CallOptions).ListReleases):len((*c.CallOptions).ListReleases)], opts...)
	it := &ReleaseIterator{}
	req = proto.Clone(req).(*releasepb.ListReleasesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*releasepb.Release, string, error) {
		resp := &releasepb.ListReleasesResponse{}
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
			resp, err = c.publisherClient.ListReleases(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetReleases(), resp.GetNextPageToken(), nil
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

func (c *publisherGRPCClient) CreateRelease(ctx context.Context, req *releasepb.CreateReleaseRequest, opts ...gax.CallOption) (*releasepb.Release, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateRelease[0:len((*c.CallOptions).CreateRelease):len((*c.CallOptions).CreateRelease)], opts...)
	var resp *releasepb.Release
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.publisherClient.CreateRelease(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *publisherGRPCClient) UpdateRelease(ctx context.Context, req *releasepb.UpdateReleaseRequest, opts ...gax.CallOption) (*releasepb.Release, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "release.name", url.QueryEscape(req.GetRelease().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateRelease[0:len((*c.CallOptions).UpdateRelease):len((*c.CallOptions).UpdateRelease)], opts...)
	var resp *releasepb.Release
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.publisherClient.UpdateRelease(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *publisherGRPCClient) DeleteRelease(ctx context.Context, req *releasepb.DeleteReleaseRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteRelease[0:len((*c.CallOptions).DeleteRelease):len((*c.CallOptions).DeleteRelease)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.publisherClient.DeleteRelease(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *publisherGRPCClient) UndeleteRelease(ctx context.Context, req *releasepb.UndeleteReleaseRequest, opts ...gax.CallOption) (*releasepb.Release, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UndeleteRelease[0:len((*c.CallOptions).UndeleteRelease):len((*c.CallOptions).UndeleteRelease)], opts...)
	var resp *releasepb.Release
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.publisherClient.UndeleteRelease(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *publisherGRPCClient) PublishRelease(ctx context.Context, req *releasepb.PublishReleaseRequest, opts ...gax.CallOption) (*releasepb.PublishReleaseResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).PublishRelease[0:len((*c.CallOptions).PublishRelease):len((*c.CallOptions).PublishRelease)], opts...)
	var resp *releasepb.PublishReleaseResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.publisherClient.PublishRelease(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *publisherGRPCClient) UnpublishRelease(ctx context.Context, req *releasepb.UnpublishReleaseRequest, opts ...gax.CallOption) (*releasepb.UnpublishReleaseResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UnpublishRelease[0:len((*c.CallOptions).UnpublishRelease):len((*c.CallOptions).UnpublishRelease)], opts...)
	var resp *releasepb.UnpublishReleaseResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.publisherClient.UnpublishRelease(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *publisherGRPCClient) ScheduleRelease(ctx context.Context, req *releasepb.ScheduleReleaseRequest, opts ...gax.CallOption) (*releasepb.ScheduleReleaseResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ScheduleRelease[0:len((*c.CallOptions).ScheduleRelease):len((*c.CallOptions).ScheduleRelease)], opts...)
	var resp *releasepb.ScheduleReleaseResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.publisherClient.ScheduleRelease(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *publisherGRPCClient) CancelRelease(ctx context.Context, req *releasepb.CancelReleaseRequest, opts ...gax.CallOption) (*releasepb.CancelReleaseResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CancelRelease[0:len((*c.CallOptions).CancelRelease):len((*c.CallOptions).CancelRelease)], opts...)
	var resp *releasepb.CancelReleaseResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.publisherClient.CancelRelease(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *publisherGRPCClient) SuspendRelease(ctx context.Context, req *releasepb.SuspendReleaseRequest, opts ...gax.CallOption) (*releasepb.SuspendReleaseResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).SuspendRelease[0:len((*c.CallOptions).SuspendRelease):len((*c.CallOptions).SuspendRelease)], opts...)
	var resp *releasepb.SuspendReleaseResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.publisherClient.SuspendRelease(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ReleaseIterator manages a stream of *releasepb.Release.
type ReleaseIterator struct {
	items    []*releasepb.Release
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
	InternalFetch func(pageSize int, pageToken string) (results []*releasepb.Release, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ReleaseIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ReleaseIterator) Next() (*releasepb.Release, error) {
	var item *releasepb.Release
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ReleaseIterator) bufLen() int {
	return len(it.items)
}

func (it *ReleaseIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
