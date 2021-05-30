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

package grbac

import (
	"context"
	"fmt"
	"math"
	"net/url"

	grbacpb "github.com/animeapis/go-genproto/grbac/v1alpha1"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var newAuthorizerClientHook clientHook

// AuthorizerCallOptions contains the retry settings for each method of AuthorizerClient.
type AuthorizerCallOptions struct {
	Authorize      []gax.CallOption
	GetResource    []gax.CallOption
	CreateResource []gax.CallOption
	UpdateResource []gax.CallOption
	DeleteResource []gax.CallOption
	CreateSubject  []gax.CallOption
	GetGroup       []gax.CallOption
	CreateGroup    []gax.CallOption
	UpdateGroup    []gax.CallOption
	DeleteGroup    []gax.CallOption
	GetRole        []gax.CallOption
	CreateRole     []gax.CallOption
	UpdateRole     []gax.CallOption
	DeleteRole     []gax.CallOption
}

func defaultAuthorizerGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("127.0.0.1:9080"),
		internaloption.WithDefaultMTLSEndpoint("127.0.0.1:9080"),
		internaloption.WithDefaultAudience("https://127.0.0.1/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultAuthorizerCallOptions() *AuthorizerCallOptions {
	return &AuthorizerCallOptions{
		Authorize:      []gax.CallOption{},
		GetResource:    []gax.CallOption{},
		CreateResource: []gax.CallOption{},
		UpdateResource: []gax.CallOption{},
		DeleteResource: []gax.CallOption{},
		CreateSubject:  []gax.CallOption{},
		GetGroup:       []gax.CallOption{},
		CreateGroup:    []gax.CallOption{},
		UpdateGroup:    []gax.CallOption{},
		DeleteGroup:    []gax.CallOption{},
		GetRole:        []gax.CallOption{},
		CreateRole:     []gax.CallOption{},
		UpdateRole:     []gax.CallOption{},
		DeleteRole:     []gax.CallOption{},
	}
}

// internalAuthorizerClient is an interface that defines the methods availaible from gRBAC API.
type internalAuthorizerClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	Authorize(context.Context, *grbacpb.AuthorizeRequest, ...gax.CallOption) error
	GetResource(context.Context, *grbacpb.GetResourceRequest, ...gax.CallOption) (*grbacpb.Resource, error)
	CreateResource(context.Context, *grbacpb.CreateResourceRequest, ...gax.CallOption) (*grbacpb.Resource, error)
	UpdateResource(context.Context, *grbacpb.UpdateResourceRequest, ...gax.CallOption) (*grbacpb.Resource, error)
	DeleteResource(context.Context, *grbacpb.DeleteResourceRequest, ...gax.CallOption) error
	CreateSubject(context.Context, *grbacpb.CreateSubjectRequest, ...gax.CallOption) (*grbacpb.Subject, error)
	GetGroup(context.Context, *grbacpb.GetGroupRequest, ...gax.CallOption) (*grbacpb.Group, error)
	CreateGroup(context.Context, *grbacpb.CreateGroupRequest, ...gax.CallOption) (*grbacpb.Group, error)
	UpdateGroup(context.Context, *grbacpb.UpdateGroupRequest, ...gax.CallOption) (*grbacpb.Group, error)
	DeleteGroup(context.Context, *grbacpb.DeleteGroupRequest, ...gax.CallOption) error
	GetRole(context.Context, *grbacpb.GetRoleRequest, ...gax.CallOption) (*grbacpb.Role, error)
	CreateRole(context.Context, *grbacpb.CreateRoleRequest, ...gax.CallOption) (*grbacpb.Role, error)
	UpdateRole(context.Context, *grbacpb.UpdateRoleRequest, ...gax.CallOption) (*grbacpb.Role, error)
	DeleteRole(context.Context, *grbacpb.DeleteRoleRequest, ...gax.CallOption) error
}

// AuthorizerClient is a client for interacting with gRBAC API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Authorizer is the internal service used by Animeshon to enforce RBAC rules.
type AuthorizerClient struct {
	// The internal transport-dependent client.
	internalClient internalAuthorizerClient

	// The call options for this service.
	CallOptions *AuthorizerCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AuthorizerClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AuthorizerClient) setGoogleClientInfo(...string) {
	c.internalClient.setGoogleClientInfo()
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *AuthorizerClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// Authorize authorize returns whether a subject is allowed it perform an action on an
// object. If allowed the response will be OK (200), otherwise the response
// will be Unauthorized (403).
func (c *AuthorizerClient) Authorize(ctx context.Context, req *grbacpb.AuthorizeRequest, opts ...gax.CallOption) error {
	return c.internalClient.Authorize(ctx, req, opts...)
}

// GetResource getResource returns a resource.
func (c *AuthorizerClient) GetResource(ctx context.Context, req *grbacpb.GetResourceRequest, opts ...gax.CallOption) (*grbacpb.Resource, error) {
	return c.internalClient.GetResource(ctx, req, opts...)
}

// CreateResource createResource creates a new resource.
func (c *AuthorizerClient) CreateResource(ctx context.Context, req *grbacpb.CreateResourceRequest, opts ...gax.CallOption) (*grbacpb.Resource, error) {
	return c.internalClient.CreateResource(ctx, req, opts...)
}

// UpdateResource updateResource updates a resource with a field mask.
func (c *AuthorizerClient) UpdateResource(ctx context.Context, req *grbacpb.UpdateResourceRequest, opts ...gax.CallOption) (*grbacpb.Resource, error) {
	return c.internalClient.UpdateResource(ctx, req, opts...)
}

// DeleteResource deleteResource deletes a resource.
func (c *AuthorizerClient) DeleteResource(ctx context.Context, req *grbacpb.DeleteResourceRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteResource(ctx, req, opts...)
}

// CreateSubject createSubject creates a new subject.
func (c *AuthorizerClient) CreateSubject(ctx context.Context, req *grbacpb.CreateSubjectRequest, opts ...gax.CallOption) (*grbacpb.Subject, error) {
	return c.internalClient.CreateSubject(ctx, req, opts...)
}

// GetGroup getGroup returns a group.
func (c *AuthorizerClient) GetGroup(ctx context.Context, req *grbacpb.GetGroupRequest, opts ...gax.CallOption) (*grbacpb.Group, error) {
	return c.internalClient.GetGroup(ctx, req, opts...)
}

// CreateGroup createGroup creates a new group.
func (c *AuthorizerClient) CreateGroup(ctx context.Context, req *grbacpb.CreateGroupRequest, opts ...gax.CallOption) (*grbacpb.Group, error) {
	return c.internalClient.CreateGroup(ctx, req, opts...)
}

// UpdateGroup updateGroup updates a group with a field mask.
func (c *AuthorizerClient) UpdateGroup(ctx context.Context, req *grbacpb.UpdateGroupRequest, opts ...gax.CallOption) (*grbacpb.Group, error) {
	return c.internalClient.UpdateGroup(ctx, req, opts...)
}

// DeleteGroup deleteGroup deletes a group.
func (c *AuthorizerClient) DeleteGroup(ctx context.Context, req *grbacpb.DeleteGroupRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteGroup(ctx, req, opts...)
}

// GetRole getRole returns a role.
func (c *AuthorizerClient) GetRole(ctx context.Context, req *grbacpb.GetRoleRequest, opts ...gax.CallOption) (*grbacpb.Role, error) {
	return c.internalClient.GetRole(ctx, req, opts...)
}

// CreateRole createRole creates a new role.
func (c *AuthorizerClient) CreateRole(ctx context.Context, req *grbacpb.CreateRoleRequest, opts ...gax.CallOption) (*grbacpb.Role, error) {
	return c.internalClient.CreateRole(ctx, req, opts...)
}

// UpdateRole updateRole updates a role with a field mask.
func (c *AuthorizerClient) UpdateRole(ctx context.Context, req *grbacpb.UpdateRoleRequest, opts ...gax.CallOption) (*grbacpb.Role, error) {
	return c.internalClient.UpdateRole(ctx, req, opts...)
}

// DeleteRole deleteRole deletes a role.
func (c *AuthorizerClient) DeleteRole(ctx context.Context, req *grbacpb.DeleteRoleRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteRole(ctx, req, opts...)
}

// authorizerGRPCClient is a client for interacting with gRBAC API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type authorizerGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing AuthorizerClient
	CallOptions **AuthorizerCallOptions

	// The gRPC API client.
	authorizerClient grbacpb.AuthorizerClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewAuthorizerClient creates a new authorizer client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Authorizer is the internal service used by Animeshon to enforce RBAC rules.
func NewAuthorizerClient(ctx context.Context, opts ...option.ClientOption) (*AuthorizerClient, error) {
	clientOpts := defaultAuthorizerGRPCClientOptions()
	if newAuthorizerClientHook != nil {
		hookOpts, err := newAuthorizerClientHook(ctx, clientHookParams{})
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
	client := AuthorizerClient{CallOptions: defaultAuthorizerCallOptions()}

	c := &authorizerGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		authorizerClient: grbacpb.NewAuthorizerClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *authorizerGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *authorizerGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *authorizerGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *authorizerGRPCClient) Authorize(ctx context.Context, req *grbacpb.AuthorizeRequest, opts ...gax.CallOption) error {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).Authorize[0:len((*c.CallOptions).Authorize):len((*c.CallOptions).Authorize)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.authorizerClient.Authorize(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *authorizerGRPCClient) GetResource(ctx context.Context, req *grbacpb.GetResourceRequest, opts ...gax.CallOption) (*grbacpb.Resource, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetResource[0:len((*c.CallOptions).GetResource):len((*c.CallOptions).GetResource)], opts...)
	var resp *grbacpb.Resource
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.authorizerClient.GetResource(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *authorizerGRPCClient) CreateResource(ctx context.Context, req *grbacpb.CreateResourceRequest, opts ...gax.CallOption) (*grbacpb.Resource, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).CreateResource[0:len((*c.CallOptions).CreateResource):len((*c.CallOptions).CreateResource)], opts...)
	var resp *grbacpb.Resource
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.authorizerClient.CreateResource(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *authorizerGRPCClient) UpdateResource(ctx context.Context, req *grbacpb.UpdateResourceRequest, opts ...gax.CallOption) (*grbacpb.Resource, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource.name", url.QueryEscape(req.GetResource().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateResource[0:len((*c.CallOptions).UpdateResource):len((*c.CallOptions).UpdateResource)], opts...)
	var resp *grbacpb.Resource
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.authorizerClient.UpdateResource(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *authorizerGRPCClient) DeleteResource(ctx context.Context, req *grbacpb.DeleteResourceRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteResource[0:len((*c.CallOptions).DeleteResource):len((*c.CallOptions).DeleteResource)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.authorizerClient.DeleteResource(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *authorizerGRPCClient) CreateSubject(ctx context.Context, req *grbacpb.CreateSubjectRequest, opts ...gax.CallOption) (*grbacpb.Subject, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).CreateSubject[0:len((*c.CallOptions).CreateSubject):len((*c.CallOptions).CreateSubject)], opts...)
	var resp *grbacpb.Subject
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.authorizerClient.CreateSubject(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *authorizerGRPCClient) GetGroup(ctx context.Context, req *grbacpb.GetGroupRequest, opts ...gax.CallOption) (*grbacpb.Group, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetGroup[0:len((*c.CallOptions).GetGroup):len((*c.CallOptions).GetGroup)], opts...)
	var resp *grbacpb.Group
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.authorizerClient.GetGroup(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *authorizerGRPCClient) CreateGroup(ctx context.Context, req *grbacpb.CreateGroupRequest, opts ...gax.CallOption) (*grbacpb.Group, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).CreateGroup[0:len((*c.CallOptions).CreateGroup):len((*c.CallOptions).CreateGroup)], opts...)
	var resp *grbacpb.Group
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.authorizerClient.CreateGroup(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *authorizerGRPCClient) UpdateGroup(ctx context.Context, req *grbacpb.UpdateGroupRequest, opts ...gax.CallOption) (*grbacpb.Group, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "group.name", url.QueryEscape(req.GetGroup().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateGroup[0:len((*c.CallOptions).UpdateGroup):len((*c.CallOptions).UpdateGroup)], opts...)
	var resp *grbacpb.Group
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.authorizerClient.UpdateGroup(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *authorizerGRPCClient) DeleteGroup(ctx context.Context, req *grbacpb.DeleteGroupRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteGroup[0:len((*c.CallOptions).DeleteGroup):len((*c.CallOptions).DeleteGroup)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.authorizerClient.DeleteGroup(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *authorizerGRPCClient) GetRole(ctx context.Context, req *grbacpb.GetRoleRequest, opts ...gax.CallOption) (*grbacpb.Role, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetRole[0:len((*c.CallOptions).GetRole):len((*c.CallOptions).GetRole)], opts...)
	var resp *grbacpb.Role
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.authorizerClient.GetRole(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *authorizerGRPCClient) CreateRole(ctx context.Context, req *grbacpb.CreateRoleRequest, opts ...gax.CallOption) (*grbacpb.Role, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).CreateRole[0:len((*c.CallOptions).CreateRole):len((*c.CallOptions).CreateRole)], opts...)
	var resp *grbacpb.Role
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.authorizerClient.CreateRole(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *authorizerGRPCClient) UpdateRole(ctx context.Context, req *grbacpb.UpdateRoleRequest, opts ...gax.CallOption) (*grbacpb.Role, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "role.name", url.QueryEscape(req.GetRole().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateRole[0:len((*c.CallOptions).UpdateRole):len((*c.CallOptions).UpdateRole)], opts...)
	var resp *grbacpb.Role
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.authorizerClient.UpdateRole(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *authorizerGRPCClient) DeleteRole(ctx context.Context, req *grbacpb.DeleteRoleRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteRole[0:len((*c.CallOptions).DeleteRole):len((*c.CallOptions).DeleteRole)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.authorizerClient.DeleteRole(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}