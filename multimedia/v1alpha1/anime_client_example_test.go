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

func ExampleNewAnimeClient() {
	ctx := context.Background()
	c, err := multimedia.NewAnimeClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleAnimeClient_GetAnime() {
	ctx := context.Background()
	c, err := multimedia.NewAnimeClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &multimediapb.GetAnimeRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/multimedia/v1alpha1#GetAnimeRequest.
	}
	resp, err := c.GetAnime(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAnimeClient_ListAnimes() {
	ctx := context.Background()
	c, err := multimedia.NewAnimeClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &multimediapb.ListAnimesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/multimedia/v1alpha1#ListAnimesRequest.
	}
	it := c.ListAnimes(ctx, req)
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

func ExampleAnimeClient_CreateAnime() {
	ctx := context.Background()
	c, err := multimedia.NewAnimeClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &multimediapb.CreateAnimeRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/multimedia/v1alpha1#CreateAnimeRequest.
	}
	resp, err := c.CreateAnime(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAnimeClient_UpdateAnime() {
	ctx := context.Background()
	c, err := multimedia.NewAnimeClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &multimediapb.UpdateAnimeRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/multimedia/v1alpha1#UpdateAnimeRequest.
	}
	resp, err := c.UpdateAnime(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAnimeClient_DeleteAnime() {
	ctx := context.Background()
	c, err := multimedia.NewAnimeClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &multimediapb.DeleteAnimeRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/multimedia/v1alpha1#DeleteAnimeRequest.
	}
	err = c.DeleteAnime(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleAnimeClient_ReconcileAnimes() {
	ctx := context.Background()
	c, err := multimedia.NewAnimeClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &multimediapb.ReconcileAnimesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/multimedia/v1alpha1#ReconcileAnimesRequest.
	}
	op, err := c.ReconcileAnimes(ctx, req)
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
