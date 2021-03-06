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

package credentials_test

import (
	"context"

	credentials "github.com/animeapis/api-go-client/credentials/v1alpha1"
	credentialspb "github.com/animeapis/go-genproto/credentials/v1alpha1"
	"google.golang.org/api/iterator"
)

func ExampleNewKeeperClient() {
	ctx := context.Background()
	c, err := credentials.NewKeeperClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleKeeperClient_GetCredentials() {
	ctx := context.Background()
	c, err := credentials.NewKeeperClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &credentialspb.GetCredentialsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/credentials/v1alpha1#GetCredentialsRequest.
	}
	resp, err := c.GetCredentials(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleKeeperClient_ListCredentials() {
	ctx := context.Background()
	c, err := credentials.NewKeeperClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &credentialspb.ListCredentialsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/credentials/v1alpha1#ListCredentialsRequest.
	}
	it := c.ListCredentials(ctx, req)
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

func ExampleKeeperClient_CreateCredentials() {
	ctx := context.Background()
	c, err := credentials.NewKeeperClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &credentialspb.CreateCredentialsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/credentials/v1alpha1#CreateCredentialsRequest.
	}
	resp, err := c.CreateCredentials(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleKeeperClient_DeleteCredentials() {
	ctx := context.Background()
	c, err := credentials.NewKeeperClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &credentialspb.DeleteCredentialsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/credentials/v1alpha1#DeleteCredentialsRequest.
	}
	err = c.DeleteCredentials(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleKeeperClient_ActAsCredentials() {
	ctx := context.Background()
	c, err := credentials.NewKeeperClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &credentialspb.ActAsCredentialsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/credentials/v1alpha1#ActAsCredentialsRequest.
	}
	resp, err := c.ActAsCredentials(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
