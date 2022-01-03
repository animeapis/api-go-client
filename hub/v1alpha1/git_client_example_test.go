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

package hub_test

import (
	"context"

	hub "github.com/animeapis/api-go-client/hub/v1alpha1"
	hubpb "github.com/animeapis/go-genproto/hub/v1alpha1"
)

func ExampleNewGitClient() {
	ctx := context.Background()
	c, err := hub.NewGitClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleGitClient_AdvertiseReferences() {
	ctx := context.Background()
	c, err := hub.NewGitClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &hubpb.AdvertiseReferencesRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.AdvertiseReferences(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleGitClient_ReceivePack() {
	ctx := context.Background()
	c, err := hub.NewGitClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &hubpb.ReceivePackRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.ReceivePack(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleGitClient_UploadPack() {
	ctx := context.Background()
	c, err := hub.NewGitClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &hubpb.UploadPackRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UploadPack(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
