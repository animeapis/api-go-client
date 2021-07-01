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

package vision

import (
	"context"
	"fmt"
	"math"
	"net/url"

	visionpb "github.com/animeapis/go-genproto/vision/v1alpha1"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var newImageAnnotatorClientHook clientHook

// ImageAnnotatorCallOptions contains the retry settings for each method of ImageAnnotatorClient.
type ImageAnnotatorCallOptions struct {
	CreateImageReport         []gax.CallOption
	ListImageReports          []gax.CallOption
	GetImageReport            []gax.CallOption
	DeleteImageReport         []gax.CallOption
	CreateImageAnnotationHint []gax.CallOption
	ListImageAnnotationHints  []gax.CallOption
	GetImageAnnotationHint    []gax.CallOption
	UpdateImageAnnotationHint []gax.CallOption
	DeleteImageAnnotationHint []gax.CallOption
}

func defaultImageAnnotatorGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("vision.animeapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("vision.animeapis.com:443"),
		internaloption.WithDefaultAudience("https://vision.animeapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultImageAnnotatorCallOptions() *ImageAnnotatorCallOptions {
	return &ImageAnnotatorCallOptions{
		CreateImageReport:         []gax.CallOption{},
		ListImageReports:          []gax.CallOption{},
		GetImageReport:            []gax.CallOption{},
		DeleteImageReport:         []gax.CallOption{},
		CreateImageAnnotationHint: []gax.CallOption{},
		ListImageAnnotationHints:  []gax.CallOption{},
		GetImageAnnotationHint:    []gax.CallOption{},
		UpdateImageAnnotationHint: []gax.CallOption{},
		DeleteImageAnnotationHint: []gax.CallOption{},
	}
}

// internalImageAnnotatorClient is an interface that defines the methods availaible from Vision API.
type internalImageAnnotatorClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateImageReport(context.Context, *visionpb.CreateImageReportRequest, ...gax.CallOption) (*visionpb.ImageReport, error)
	ListImageReports(context.Context, *visionpb.ListImageReportsRequest, ...gax.CallOption) (*visionpb.ListImageReportsResponse, error)
	GetImageReport(context.Context, *visionpb.GetImageReportRequest, ...gax.CallOption) (*visionpb.ImageReport, error)
	DeleteImageReport(context.Context, *visionpb.DeleteImageReportRequest, ...gax.CallOption) error
	CreateImageAnnotationHint(context.Context, *visionpb.CreateImageAnnotationHintRequest, ...gax.CallOption) (*visionpb.ImageAnnotationHint, error)
	ListImageAnnotationHints(context.Context, *visionpb.ListImageAnnotationHintsRequest, ...gax.CallOption) (*visionpb.ListImageAnnotationHintsResponse, error)
	GetImageAnnotationHint(context.Context, *visionpb.GetImageAnnotationHintRequest, ...gax.CallOption) (*visionpb.ImageAnnotationHint, error)
	UpdateImageAnnotationHint(context.Context, *visionpb.UpdateImageAnnotationHintRequest, ...gax.CallOption) (*visionpb.ImageAnnotationHint, error)
	DeleteImageAnnotationHint(context.Context, *visionpb.DeleteImageAnnotationHintRequest, ...gax.CallOption) error
}

// ImageAnnotatorClient is a client for interacting with Vision API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type ImageAnnotatorClient struct {
	// The internal transport-dependent client.
	internalClient internalImageAnnotatorClient

	// The call options for this service.
	CallOptions *ImageAnnotatorCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ImageAnnotatorClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ImageAnnotatorClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ImageAnnotatorClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

func (c *ImageAnnotatorClient) CreateImageReport(ctx context.Context, req *visionpb.CreateImageReportRequest, opts ...gax.CallOption) (*visionpb.ImageReport, error) {
	return c.internalClient.CreateImageReport(ctx, req, opts...)
}

func (c *ImageAnnotatorClient) ListImageReports(ctx context.Context, req *visionpb.ListImageReportsRequest, opts ...gax.CallOption) (*visionpb.ListImageReportsResponse, error) {
	return c.internalClient.ListImageReports(ctx, req, opts...)
}

func (c *ImageAnnotatorClient) GetImageReport(ctx context.Context, req *visionpb.GetImageReportRequest, opts ...gax.CallOption) (*visionpb.ImageReport, error) {
	return c.internalClient.GetImageReport(ctx, req, opts...)
}

func (c *ImageAnnotatorClient) DeleteImageReport(ctx context.Context, req *visionpb.DeleteImageReportRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteImageReport(ctx, req, opts...)
}

func (c *ImageAnnotatorClient) CreateImageAnnotationHint(ctx context.Context, req *visionpb.CreateImageAnnotationHintRequest, opts ...gax.CallOption) (*visionpb.ImageAnnotationHint, error) {
	return c.internalClient.CreateImageAnnotationHint(ctx, req, opts...)
}

func (c *ImageAnnotatorClient) ListImageAnnotationHints(ctx context.Context, req *visionpb.ListImageAnnotationHintsRequest, opts ...gax.CallOption) (*visionpb.ListImageAnnotationHintsResponse, error) {
	return c.internalClient.ListImageAnnotationHints(ctx, req, opts...)
}

func (c *ImageAnnotatorClient) GetImageAnnotationHint(ctx context.Context, req *visionpb.GetImageAnnotationHintRequest, opts ...gax.CallOption) (*visionpb.ImageAnnotationHint, error) {
	return c.internalClient.GetImageAnnotationHint(ctx, req, opts...)
}

func (c *ImageAnnotatorClient) UpdateImageAnnotationHint(ctx context.Context, req *visionpb.UpdateImageAnnotationHintRequest, opts ...gax.CallOption) (*visionpb.ImageAnnotationHint, error) {
	return c.internalClient.UpdateImageAnnotationHint(ctx, req, opts...)
}

func (c *ImageAnnotatorClient) DeleteImageAnnotationHint(ctx context.Context, req *visionpb.DeleteImageAnnotationHintRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteImageAnnotationHint(ctx, req, opts...)
}

// imageAnnotatorGRPCClient is a client for interacting with Vision API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type imageAnnotatorGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ImageAnnotatorClient
	CallOptions **ImageAnnotatorCallOptions

	// The gRPC API client.
	imageAnnotatorClient visionpb.ImageAnnotatorClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewImageAnnotatorClient creates a new image annotator client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
func NewImageAnnotatorClient(ctx context.Context, opts ...option.ClientOption) (*ImageAnnotatorClient, error) {
	clientOpts := defaultImageAnnotatorGRPCClientOptions()
	if newImageAnnotatorClientHook != nil {
		hookOpts, err := newImageAnnotatorClientHook(ctx, clientHookParams{})
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
	client := ImageAnnotatorClient{CallOptions: defaultImageAnnotatorCallOptions()}

	c := &imageAnnotatorGRPCClient{
		connPool:             connPool,
		disableDeadlines:     disableDeadlines,
		imageAnnotatorClient: visionpb.NewImageAnnotatorClient(connPool),
		CallOptions:          &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *imageAnnotatorGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *imageAnnotatorGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *imageAnnotatorGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *imageAnnotatorGRPCClient) CreateImageReport(ctx context.Context, req *visionpb.CreateImageReportRequest, opts ...gax.CallOption) (*visionpb.ImageReport, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateImageReport[0:len((*c.CallOptions).CreateImageReport):len((*c.CallOptions).CreateImageReport)], opts...)
	var resp *visionpb.ImageReport
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.imageAnnotatorClient.CreateImageReport(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *imageAnnotatorGRPCClient) ListImageReports(ctx context.Context, req *visionpb.ListImageReportsRequest, opts ...gax.CallOption) (*visionpb.ListImageReportsResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListImageReports[0:len((*c.CallOptions).ListImageReports):len((*c.CallOptions).ListImageReports)], opts...)
	var resp *visionpb.ListImageReportsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.imageAnnotatorClient.ListImageReports(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *imageAnnotatorGRPCClient) GetImageReport(ctx context.Context, req *visionpb.GetImageReportRequest, opts ...gax.CallOption) (*visionpb.ImageReport, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetImageReport[0:len((*c.CallOptions).GetImageReport):len((*c.CallOptions).GetImageReport)], opts...)
	var resp *visionpb.ImageReport
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.imageAnnotatorClient.GetImageReport(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *imageAnnotatorGRPCClient) DeleteImageReport(ctx context.Context, req *visionpb.DeleteImageReportRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteImageReport[0:len((*c.CallOptions).DeleteImageReport):len((*c.CallOptions).DeleteImageReport)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.imageAnnotatorClient.DeleteImageReport(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *imageAnnotatorGRPCClient) CreateImageAnnotationHint(ctx context.Context, req *visionpb.CreateImageAnnotationHintRequest, opts ...gax.CallOption) (*visionpb.ImageAnnotationHint, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateImageAnnotationHint[0:len((*c.CallOptions).CreateImageAnnotationHint):len((*c.CallOptions).CreateImageAnnotationHint)], opts...)
	var resp *visionpb.ImageAnnotationHint
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.imageAnnotatorClient.CreateImageAnnotationHint(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *imageAnnotatorGRPCClient) ListImageAnnotationHints(ctx context.Context, req *visionpb.ListImageAnnotationHintsRequest, opts ...gax.CallOption) (*visionpb.ListImageAnnotationHintsResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListImageAnnotationHints[0:len((*c.CallOptions).ListImageAnnotationHints):len((*c.CallOptions).ListImageAnnotationHints)], opts...)
	var resp *visionpb.ListImageAnnotationHintsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.imageAnnotatorClient.ListImageAnnotationHints(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *imageAnnotatorGRPCClient) GetImageAnnotationHint(ctx context.Context, req *visionpb.GetImageAnnotationHintRequest, opts ...gax.CallOption) (*visionpb.ImageAnnotationHint, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetImageAnnotationHint[0:len((*c.CallOptions).GetImageAnnotationHint):len((*c.CallOptions).GetImageAnnotationHint)], opts...)
	var resp *visionpb.ImageAnnotationHint
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.imageAnnotatorClient.GetImageAnnotationHint(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *imageAnnotatorGRPCClient) UpdateImageAnnotationHint(ctx context.Context, req *visionpb.UpdateImageAnnotationHintRequest, opts ...gax.CallOption) (*visionpb.ImageAnnotationHint, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "hint.name", url.QueryEscape(req.GetHint().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateImageAnnotationHint[0:len((*c.CallOptions).UpdateImageAnnotationHint):len((*c.CallOptions).UpdateImageAnnotationHint)], opts...)
	var resp *visionpb.ImageAnnotationHint
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.imageAnnotatorClient.UpdateImageAnnotationHint(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *imageAnnotatorGRPCClient) DeleteImageAnnotationHint(ctx context.Context, req *visionpb.DeleteImageAnnotationHintRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteImageAnnotationHint[0:len((*c.CallOptions).DeleteImageAnnotationHint):len((*c.CallOptions).DeleteImageAnnotationHint)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.imageAnnotatorClient.DeleteImageAnnotationHint(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}