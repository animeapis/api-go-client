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

package identity_test

import (
	"context"

	identity "github.com/animeapis/api-go-client/identity/v1alpha1"
	identitypb "github.com/animeapis/go-genproto/identity/v1alpha1"
	"google.golang.org/api/iterator"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := identity.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleClient_GetUserProfile() {
	ctx := context.Background()
	c, err := identity.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &identitypb.GetUserProfileRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetUserProfile(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetUser() {
	ctx := context.Background()
	c, err := identity.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &identitypb.GetUserRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetUser(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListUsers() {
	ctx := context.Background()
	c, err := identity.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &identitypb.ListUsersRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListUsers(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_CreateUser() {
	ctx := context.Background()
	c, err := identity.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &identitypb.CreateUserRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateUser(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateUser() {
	ctx := context.Background()
	c, err := identity.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &identitypb.UpdateUserRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateUser(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteUser() {
	ctx := context.Background()
	c, err := identity.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &identitypb.DeleteUserRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteUser(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_GetUserSettings() {
	ctx := context.Background()
	c, err := identity.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &identitypb.GetUserSettingsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetUserSettings(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateUserSettings() {
	ctx := context.Background()
	c, err := identity.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &identitypb.UpdateUserSettingsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateUserSettings(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetGroup() {
	ctx := context.Background()
	c, err := identity.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &identitypb.GetGroupRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetGroup(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListGroups() {
	ctx := context.Background()
	c, err := identity.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &identitypb.ListGroupsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListGroups(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_CreateGroup() {
	ctx := context.Background()
	c, err := identity.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &identitypb.CreateGroupRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateGroup(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateGroup() {
	ctx := context.Background()
	c, err := identity.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &identitypb.UpdateGroupRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateGroup(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteGroup() {
	ctx := context.Background()
	c, err := identity.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &identitypb.DeleteGroupRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteGroup(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}
