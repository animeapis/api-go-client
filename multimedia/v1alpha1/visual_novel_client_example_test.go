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

package multimedia_test

import (
	"context"

	multimedia "github.com/animeapis/api-go-client/multimedia/v1alpha1"
	multimediapb "github.com/animeapis/go-genproto/multimedia/v1alpha1"
	"google.golang.org/api/iterator"
)

func ExampleNewVisualNovelClient() {
	ctx := context.Background()
	c, err := multimedia.NewVisualNovelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleNewVisualNovelRESTClient() {
	ctx := context.Background()
	c, err := multimedia.NewVisualNovelRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleVisualNovelClient_GetVisualNovel() {
	ctx := context.Background()
	c, err := multimedia.NewVisualNovelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &multimediapb.GetVisualNovelRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/multimedia/v1alpha1#GetVisualNovelRequest.
	}
	resp, err := c.GetVisualNovel(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleVisualNovelClient_ListVisualNovels() {
	ctx := context.Background()
	c, err := multimedia.NewVisualNovelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &multimediapb.ListVisualNovelsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/multimedia/v1alpha1#ListVisualNovelsRequest.
	}
	it := c.ListVisualNovels(ctx, req)
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

func ExampleVisualNovelClient_CreateVisualNovel() {
	ctx := context.Background()
	c, err := multimedia.NewVisualNovelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &multimediapb.CreateVisualNovelRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/multimedia/v1alpha1#CreateVisualNovelRequest.
	}
	resp, err := c.CreateVisualNovel(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleVisualNovelClient_UpdateVisualNovel() {
	ctx := context.Background()
	c, err := multimedia.NewVisualNovelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &multimediapb.UpdateVisualNovelRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/multimedia/v1alpha1#UpdateVisualNovelRequest.
	}
	resp, err := c.UpdateVisualNovel(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleVisualNovelClient_DeleteVisualNovel() {
	ctx := context.Background()
	c, err := multimedia.NewVisualNovelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &multimediapb.DeleteVisualNovelRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/multimedia/v1alpha1#DeleteVisualNovelRequest.
	}
	err = c.DeleteVisualNovel(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleVisualNovelClient_ReconcileVisualNovels() {
	ctx := context.Background()
	c, err := multimedia.NewVisualNovelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &multimediapb.ReconcileVisualNovelsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/multimedia/v1alpha1#ReconcileVisualNovelsRequest.
	}
	op, err := c.ReconcileVisualNovels(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
