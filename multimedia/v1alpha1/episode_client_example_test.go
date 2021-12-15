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

package multimedia_test

import (
	"context"

	multimedia "github.com/animeapis/api-go-client/multimedia/v1alpha1"
	multimediapb "github.com/animeapis/go-genproto/multimedia/v1alpha1"
	"google.golang.org/api/iterator"
)

func ExampleNewEpisodeClient() {
	ctx := context.Background()
	c, err := multimedia.NewEpisodeClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleEpisodeClient_GetEpisode() {
	ctx := context.Background()
	c, err := multimedia.NewEpisodeClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &multimediapb.GetEpisodeRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetEpisode(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleEpisodeClient_ListEpisodes() {
	ctx := context.Background()
	c, err := multimedia.NewEpisodeClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &multimediapb.ListEpisodesRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListEpisodes(ctx, req)
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

func ExampleEpisodeClient_CreateEpisode() {
	ctx := context.Background()
	c, err := multimedia.NewEpisodeClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &multimediapb.CreateEpisodeRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateEpisode(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleEpisodeClient_UpdateEpisode() {
	ctx := context.Background()
	c, err := multimedia.NewEpisodeClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &multimediapb.UpdateEpisodeRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateEpisode(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleEpisodeClient_DeleteEpisode() {
	ctx := context.Background()
	c, err := multimedia.NewEpisodeClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &multimediapb.DeleteEpisodeRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteEpisode(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}
