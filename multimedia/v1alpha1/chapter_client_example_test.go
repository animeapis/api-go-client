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

func ExampleNewChapterClient() {
	ctx := context.Background()
	c, err := multimedia.NewChapterClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleChapterClient_GetChapter() {
	ctx := context.Background()
	c, err := multimedia.NewChapterClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &multimediapb.GetChapterRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/multimedia/v1alpha1#GetChapterRequest.
	}
	resp, err := c.GetChapter(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleChapterClient_ListChapters() {
	ctx := context.Background()
	c, err := multimedia.NewChapterClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &multimediapb.ListChaptersRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/multimedia/v1alpha1#ListChaptersRequest.
	}
	it := c.ListChapters(ctx, req)
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

func ExampleChapterClient_CreateChapter() {
	ctx := context.Background()
	c, err := multimedia.NewChapterClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &multimediapb.CreateChapterRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/multimedia/v1alpha1#CreateChapterRequest.
	}
	resp, err := c.CreateChapter(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleChapterClient_UpdateChapter() {
	ctx := context.Background()
	c, err := multimedia.NewChapterClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &multimediapb.UpdateChapterRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/multimedia/v1alpha1#UpdateChapterRequest.
	}
	resp, err := c.UpdateChapter(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleChapterClient_DeleteChapter() {
	ctx := context.Background()
	c, err := multimedia.NewChapterClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &multimediapb.DeleteChapterRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/multimedia/v1alpha1#DeleteChapterRequest.
	}
	err = c.DeleteChapter(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}