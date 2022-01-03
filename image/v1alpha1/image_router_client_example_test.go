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

package image_test

import (
	"context"

	image "github.com/animeapis/api-go-client/image/v1alpha1"
	imagepb "github.com/animeapis/go-genproto/image/v1alpha1"
)

func ExampleNewImageRouterClient() {
	ctx := context.Background()
	c, err := image.NewImageRouterClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleImageRouterClient_GetImageRoute() {
	ctx := context.Background()
	c, err := image.NewImageRouterClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &imagepb.GetImageRouteRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetImageRoute(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleImageRouterClient_RouteImage() {
	ctx := context.Background()
	c, err := image.NewImageRouterClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &imagepb.RouteImageRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.RouteImage(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
