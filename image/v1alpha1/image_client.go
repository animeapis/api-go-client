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
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httpbodypb "google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	UploadImage []gax.CallOption
	ImportImage []gax.CallOption
	GetImage    []gax.CallOption
	GetCdnImage []gax.CallOption
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
		UploadImage: []gax.CallOption{},
		ImportImage: []gax.CallOption{},
		GetImage:    []gax.CallOption{},
		GetCdnImage: []gax.CallOption{},
	}
}

// internalClient is an interface that defines the methods availaible from Image API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	UploadImage(context.Context, *imagepb.UploadImageRequest, ...gax.CallOption) error
	ImportImage(context.Context, *imagepb.ImportImageRequest, ...gax.CallOption) error
	GetImage(context.Context, *imagepb.GetImageRequest, ...gax.CallOption) (*httpbodypb.HttpBody, error)
	GetCdnImage(context.Context, *imagepb.GetImageRequest, ...gax.CallOption) (*httpbodypb.HttpBody, error)
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

func (c *Client) UploadImage(ctx context.Context, req *imagepb.UploadImageRequest, opts ...gax.CallOption) error {
	return c.internalClient.UploadImage(ctx, req, opts...)
}

func (c *Client) ImportImage(ctx context.Context, req *imagepb.ImportImageRequest, opts ...gax.CallOption) error {
	return c.internalClient.ImportImage(ctx, req, opts...)
}

func (c *Client) GetImage(ctx context.Context, req *imagepb.GetImageRequest, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	return c.internalClient.GetImage(ctx, req, opts...)
}

func (c *Client) GetCdnImage(ctx context.Context, req *imagepb.GetImageRequest, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	return c.internalClient.GetCdnImage(ctx, req, opts...)
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

func (c *gRPCClient) UploadImage(ctx context.Context, req *imagepb.UploadImageRequest, opts ...gax.CallOption) error {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).UploadImage[0:len((*c.CallOptions).UploadImage):len((*c.CallOptions).UploadImage)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.UploadImage(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *gRPCClient) ImportImage(ctx context.Context, req *imagepb.ImportImageRequest, opts ...gax.CallOption) error {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ImportImage[0:len((*c.CallOptions).ImportImage):len((*c.CallOptions).ImportImage)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.ImportImage(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
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

func (c *gRPCClient) GetCdnImage(ctx context.Context, req *imagepb.GetImageRequest, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetCdnImage[0:len((*c.CallOptions).GetCdnImage):len((*c.CallOptions).GetCdnImage)], opts...)
	var resp *httpbodypb.HttpBody
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetCdnImage(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
