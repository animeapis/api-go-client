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

package admin_test

import (
	"context"

	admin "github.com/animeapis/api-go-client/iam/admin/v1alpha1"
	adminpb "github.com/animeapis/go-genproto/iam/admin/v1alpha1"
	"google.golang.org/api/iterator"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
)

func ExampleNewIamClient() {
	ctx := context.Background()
	c, err := admin.NewIamClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleIamClient_GetServiceAccount() {
	ctx := context.Background()
	c, err := admin.NewIamClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &adminpb.GetServiceAccountRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetServiceAccount(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleIamClient_ListServiceAccounts() {
	ctx := context.Background()
	c, err := admin.NewIamClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &adminpb.ListServiceAccountsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListServiceAccounts(ctx, req)
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

func ExampleIamClient_CreateServiceAccount() {
	ctx := context.Background()
	c, err := admin.NewIamClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &adminpb.CreateServiceAccountRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateServiceAccount(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleIamClient_UpdateServiceAccount() {
	ctx := context.Background()
	c, err := admin.NewIamClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &adminpb.UpdateServiceAccountRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateServiceAccount(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleIamClient_DeleteServiceAccount() {
	ctx := context.Background()
	c, err := admin.NewIamClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &adminpb.DeleteServiceAccountRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteServiceAccount(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleIamClient_GetIamPolicy() {
	ctx := context.Background()
	c, err := admin.NewIamClient(ctx)
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

func ExampleIamClient_SetIamPolicy() {
	ctx := context.Background()
	c, err := admin.NewIamClient(ctx)
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

func ExampleIamClient_TestIamPermissions() {
	ctx := context.Background()
	c, err := admin.NewIamClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.TestIamPermissionsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.TestIamPermissions(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleIamClient_GetRole() {
	ctx := context.Background()
	c, err := admin.NewIamClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &adminpb.GetRoleRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetRole(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleIamClient_ListRoles() {
	ctx := context.Background()
	c, err := admin.NewIamClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &adminpb.ListRolesRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListRoles(ctx, req)
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

func ExampleIamClient_CreateRole() {
	ctx := context.Background()
	c, err := admin.NewIamClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &adminpb.CreateRoleRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateRole(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleIamClient_UpdateRole() {
	ctx := context.Background()
	c, err := admin.NewIamClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &adminpb.UpdateRoleRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateRole(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleIamClient_DeleteRole() {
	ctx := context.Background()
	c, err := admin.NewIamClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &adminpb.DeleteRoleRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.DeleteRole(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleIamClient_GetPermission() {
	ctx := context.Background()
	c, err := admin.NewIamClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &adminpb.GetPermissionRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetPermission(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleIamClient_ListPermissions() {
	ctx := context.Background()
	c, err := admin.NewIamClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &adminpb.ListPermissionsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListPermissions(ctx, req)
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

func ExampleIamClient_CreatePermission() {
	ctx := context.Background()
	c, err := admin.NewIamClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &adminpb.CreatePermissionRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreatePermission(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleIamClient_UpdatePermission() {
	ctx := context.Background()
	c, err := admin.NewIamClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &adminpb.UpdatePermissionRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdatePermission(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleIamClient_DeletePermission() {
	ctx := context.Background()
	c, err := admin.NewIamClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &adminpb.DeletePermissionRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.DeletePermission(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}