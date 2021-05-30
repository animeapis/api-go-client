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

package grbac_test

import (
	"context"

	grbac "github.com/animeapis/api-go-client/grbac/v1alpha1"
	grbacpb "github.com/animeapis/go-genproto/grbac/v1alpha1"
)

func ExampleNewAuthorizerClient() {
	ctx := context.Background()
	c, err := grbac.NewAuthorizerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleAuthorizerClient_Authorize() {
	ctx := context.Background()
	c, err := grbac.NewAuthorizerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &grbacpb.AuthorizeRequest{
		// TODO: Fill request struct fields.
	}
	err = c.Authorize(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleAuthorizerClient_GetResource() {
	ctx := context.Background()
	c, err := grbac.NewAuthorizerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &grbacpb.GetResourceRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetResource(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAuthorizerClient_CreateResource() {
	ctx := context.Background()
	c, err := grbac.NewAuthorizerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &grbacpb.CreateResourceRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateResource(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAuthorizerClient_UpdateResource() {
	ctx := context.Background()
	c, err := grbac.NewAuthorizerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &grbacpb.UpdateResourceRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateResource(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAuthorizerClient_DeleteResource() {
	ctx := context.Background()
	c, err := grbac.NewAuthorizerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &grbacpb.DeleteResourceRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteResource(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleAuthorizerClient_CreateSubject() {
	ctx := context.Background()
	c, err := grbac.NewAuthorizerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &grbacpb.CreateSubjectRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateSubject(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAuthorizerClient_GetGroup() {
	ctx := context.Background()
	c, err := grbac.NewAuthorizerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &grbacpb.GetGroupRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetGroup(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAuthorizerClient_CreateGroup() {
	ctx := context.Background()
	c, err := grbac.NewAuthorizerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &grbacpb.CreateGroupRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateGroup(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAuthorizerClient_UpdateGroup() {
	ctx := context.Background()
	c, err := grbac.NewAuthorizerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &grbacpb.UpdateGroupRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateGroup(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAuthorizerClient_DeleteGroup() {
	ctx := context.Background()
	c, err := grbac.NewAuthorizerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &grbacpb.DeleteGroupRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteGroup(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleAuthorizerClient_GetRole() {
	ctx := context.Background()
	c, err := grbac.NewAuthorizerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &grbacpb.GetRoleRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetRole(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAuthorizerClient_CreateRole() {
	ctx := context.Background()
	c, err := grbac.NewAuthorizerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &grbacpb.CreateRoleRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateRole(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAuthorizerClient_UpdateRole() {
	ctx := context.Background()
	c, err := grbac.NewAuthorizerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &grbacpb.UpdateRoleRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateRole(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAuthorizerClient_DeleteRole() {
	ctx := context.Background()
	c, err := grbac.NewAuthorizerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &grbacpb.DeleteRoleRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteRole(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}