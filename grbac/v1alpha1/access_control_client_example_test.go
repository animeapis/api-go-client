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

package grbac_test

import (
	"context"

	grbac "github.com/animeapis/api-go-client/grbac/v1alpha1"
	grbacpb "github.com/animeapis/go-genproto/grbac/v1alpha1"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
)

func ExampleNewAccessControlClient() {
	ctx := context.Background()
	c, err := grbac.NewAccessControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleAccessControlClient_TestIamPolicy() {
	ctx := context.Background()
	c, err := grbac.NewAccessControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &grbacpb.TestIamPolicyRequest{
		// TODO: Fill request struct fields.
	}
	err = c.TestIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleAccessControlClient_GetIamPolicy() {
	ctx := context.Background()
	c, err := grbac.NewAccessControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.GetIamPolicyRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAccessControlClient_SetIamPolicy() {
	ctx := context.Background()
	c, err := grbac.NewAccessControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.SetIamPolicyRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.SetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAccessControlClient_GetResource() {
	ctx := context.Background()
	c, err := grbac.NewAccessControlClient(ctx)
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

func ExampleAccessControlClient_CreateResource() {
	ctx := context.Background()
	c, err := grbac.NewAccessControlClient(ctx)
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

func ExampleAccessControlClient_TransferResource() {
	ctx := context.Background()
	c, err := grbac.NewAccessControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &grbacpb.TransferResourceRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.TransferResource(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAccessControlClient_DeleteResource() {
	ctx := context.Background()
	c, err := grbac.NewAccessControlClient(ctx)
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

func ExampleAccessControlClient_CreateSubject() {
	ctx := context.Background()
	c, err := grbac.NewAccessControlClient(ctx)
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

func ExampleAccessControlClient_DeleteSubject() {
	ctx := context.Background()
	c, err := grbac.NewAccessControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &grbacpb.DeleteSubjectRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteSubject(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleAccessControlClient_GetGroup() {
	ctx := context.Background()
	c, err := grbac.NewAccessControlClient(ctx)
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

func ExampleAccessControlClient_CreateGroup() {
	ctx := context.Background()
	c, err := grbac.NewAccessControlClient(ctx)
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

func ExampleAccessControlClient_UpdateGroup() {
	ctx := context.Background()
	c, err := grbac.NewAccessControlClient(ctx)
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

func ExampleAccessControlClient_AddGroupMember() {
	ctx := context.Background()
	c, err := grbac.NewAccessControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &grbacpb.AddGroupMemberRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.AddGroupMember(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAccessControlClient_RemoveGroupMember() {
	ctx := context.Background()
	c, err := grbac.NewAccessControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &grbacpb.RemoveGroupMemberRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.RemoveGroupMember(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAccessControlClient_DeleteGroup() {
	ctx := context.Background()
	c, err := grbac.NewAccessControlClient(ctx)
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

func ExampleAccessControlClient_CreatePermission() {
	ctx := context.Background()
	c, err := grbac.NewAccessControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &grbacpb.CreatePermissionRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreatePermission(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAccessControlClient_DeletePermission() {
	ctx := context.Background()
	c, err := grbac.NewAccessControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &grbacpb.DeletePermissionRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeletePermission(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleAccessControlClient_GetRole() {
	ctx := context.Background()
	c, err := grbac.NewAccessControlClient(ctx)
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

func ExampleAccessControlClient_CreateRole() {
	ctx := context.Background()
	c, err := grbac.NewAccessControlClient(ctx)
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

func ExampleAccessControlClient_UpdateRole() {
	ctx := context.Background()
	c, err := grbac.NewAccessControlClient(ctx)
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

func ExampleAccessControlClient_DeleteRole() {
	ctx := context.Background()
	c, err := grbac.NewAccessControlClient(ctx)
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
