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
	"google.golang.org/api/iterator"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := image.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleClient_UploadImage() {
	ctx := context.Background()
	c, err := image.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &imagepb.UploadImageRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/image/v1alpha1#UploadImageRequest.
	}
	resp, err := c.UploadImage(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ImportImage() {
	ctx := context.Background()
	c, err := image.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &imagepb.ImportImageRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/image/v1alpha1#ImportImageRequest.
	}
	resp, err := c.ImportImage(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetImage() {
	ctx := context.Background()
	c, err := image.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &imagepb.GetImageRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/image/v1alpha1#GetImageRequest.
	}
	resp, err := c.GetImage(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetAlbum() {
	ctx := context.Background()
	c, err := image.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &imagepb.GetAlbumRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/image/v1alpha1#GetAlbumRequest.
	}
	resp, err := c.GetAlbum(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListAlbums() {
	ctx := context.Background()
	c, err := image.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &imagepb.ListAlbumsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/image/v1alpha1#ListAlbumsRequest.
	}
	it := c.ListAlbums(ctx, req)
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

func ExampleClient_CreateAlbum() {
	ctx := context.Background()
	c, err := image.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &imagepb.CreateAlbumRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/image/v1alpha1#CreateAlbumRequest.
	}
	resp, err := c.CreateAlbum(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteAlbum() {
	ctx := context.Background()
	c, err := image.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &imagepb.DeleteAlbumRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/image/v1alpha1#DeleteAlbumRequest.
	}
	err = c.DeleteAlbum(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_GetAlbumSettings() {
	ctx := context.Background()
	c, err := image.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &imagepb.GetAlbumSettingsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/image/v1alpha1#GetAlbumSettingsRequest.
	}
	resp, err := c.GetAlbumSettings(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateAlbumSettings() {
	ctx := context.Background()
	c, err := image.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &imagepb.UpdateAlbumSettingsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/image/v1alpha1#UpdateAlbumSettingsRequest.
	}
	resp, err := c.UpdateAlbumSettings(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
