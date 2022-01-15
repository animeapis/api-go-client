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

package hub

import (
	"context"
	"fmt"
	"math"
	"net/url"

	hubpb "github.com/animeapis/go-genproto/hub/v1alpha1"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httpbodypb "google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var newGitClientHook clientHook

// GitCallOptions contains the retry settings for each method of GitClient.
type GitCallOptions struct {
	AdvertiseReferences []gax.CallOption
	ReceivePack         []gax.CallOption
	UploadPack          []gax.CallOption
}

func defaultGitGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("hub.animeapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("hub.animeapis.com:443"),
		internaloption.WithDefaultAudience("https://hub.animeapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultGitCallOptions() *GitCallOptions {
	return &GitCallOptions{
		AdvertiseReferences: []gax.CallOption{},
		ReceivePack:         []gax.CallOption{},
		UploadPack:          []gax.CallOption{},
	}
}

// internalGitClient is an interface that defines the methods availaible from Hub API.
type internalGitClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AdvertiseReferences(context.Context, *hubpb.AdvertiseReferencesRequest, ...gax.CallOption) (*httpbodypb.HttpBody, error)
	ReceivePack(context.Context, *hubpb.ReceivePackRequest, ...gax.CallOption) (*httpbodypb.HttpBody, error)
	UploadPack(context.Context, *hubpb.UploadPackRequest, ...gax.CallOption) (*httpbodypb.HttpBody, error)
}

// GitClient is a client for interacting with Hub API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type GitClient struct {
	// The internal transport-dependent client.
	internalClient internalGitClient

	// The call options for this service.
	CallOptions *GitCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *GitClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *GitClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *GitClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

func (c *GitClient) AdvertiseReferences(ctx context.Context, req *hubpb.AdvertiseReferencesRequest, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	return c.internalClient.AdvertiseReferences(ctx, req, opts...)
}

func (c *GitClient) ReceivePack(ctx context.Context, req *hubpb.ReceivePackRequest, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	return c.internalClient.ReceivePack(ctx, req, opts...)
}

func (c *GitClient) UploadPack(ctx context.Context, req *hubpb.UploadPackRequest, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	return c.internalClient.UploadPack(ctx, req, opts...)
}

// gitGRPCClient is a client for interacting with Hub API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type gitGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing GitClient
	CallOptions **GitCallOptions

	// The gRPC API client.
	gitClient hubpb.GitClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewGitClient creates a new git client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
func NewGitClient(ctx context.Context, opts ...option.ClientOption) (*GitClient, error) {
	clientOpts := defaultGitGRPCClientOptions()
	if newGitClientHook != nil {
		hookOpts, err := newGitClientHook(ctx, clientHookParams{})
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
	client := GitClient{CallOptions: defaultGitCallOptions()}

	c := &gitGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		gitClient:        hubpb.NewGitClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *gitGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *gitGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *gitGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *gitGRPCClient) AdvertiseReferences(ctx context.Context, req *hubpb.AdvertiseReferencesRequest, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).AdvertiseReferences[0:len((*c.CallOptions).AdvertiseReferences):len((*c.CallOptions).AdvertiseReferences)], opts...)
	var resp *httpbodypb.HttpBody
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.gitClient.AdvertiseReferences(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gitGRPCClient) ReceivePack(ctx context.Context, req *hubpb.ReceivePackRequest, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ReceivePack[0:len((*c.CallOptions).ReceivePack):len((*c.CallOptions).ReceivePack)], opts...)
	var resp *httpbodypb.HttpBody
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.gitClient.ReceivePack(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gitGRPCClient) UploadPack(ctx context.Context, req *hubpb.UploadPackRequest, opts ...gax.CallOption) (*httpbodypb.HttpBody, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UploadPack[0:len((*c.CallOptions).UploadPack):len((*c.CallOptions).UploadPack)], opts...)
	var resp *httpbodypb.HttpBody
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.gitClient.UploadPack(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
