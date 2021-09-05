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

package crossrefs_test

import (
	"context"

	crossrefs "github.com/animeapis/api-go-client/crossrefs/v1alpha1"
	crossrefspb "github.com/animeapis/go-genproto/crossrefs/v1alpha1"
	"google.golang.org/api/iterator"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func ExampleNewReferrerClient() {
	ctx := context.Background()
	c, err := crossrefs.NewReferrerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleReferrerClient_GetCrossRef() {
	ctx := context.Background()
	c, err := crossrefs.NewReferrerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &crossrefspb.GetCrossRefRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetCrossRef(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleReferrerClient_CreateCrossRef() {
	ctx := context.Background()
	c, err := crossrefs.NewReferrerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &crossrefspb.CreateCrossRefRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateCrossRef(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleReferrerClient_UpdateCrossRef() {
	ctx := context.Background()
	c, err := crossrefs.NewReferrerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &crossrefspb.UpdateCrossRefRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateCrossRef(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleReferrerClient_ListCrossRefs() {
	ctx := context.Background()
	c, err := crossrefs.NewReferrerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &crossrefspb.ListCrossRefsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListCrossRefs(ctx, req)
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

func ExampleReferrerClient_CountCrossRefs() {
	ctx := context.Background()
	c, err := crossrefs.NewReferrerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &crossrefspb.CountCrossRefsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CountCrossRefs(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleReferrerClient_AnalyzeCrossRefs() {
	ctx := context.Background()
	c, err := crossrefs.NewReferrerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &crossrefspb.AnalyzeCrossRefRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.AnalyzeCrossRefs(ctx, req)
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

func ExampleReferrerClient_ImportCrossRefs() {
	ctx := context.Background()
	c, err := crossrefs.NewReferrerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &crossrefspb.ImportCrossRefRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.ImportCrossRefs(ctx, req)
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

func ExampleReferrerClient_ExportCrossRefs() {
	ctx := context.Background()
	c, err := crossrefs.NewReferrerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &crossrefspb.ExportCrossRefRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.ExportCrossRefs(ctx, req)
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

func ExampleReferrerClient_InitializeCrossRefs() {
	ctx := context.Background()
	c, err := crossrefs.NewReferrerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &emptypb.Empty{
		// TODO: Fill request struct fields.
	}
	op, err := c.InitializeCrossRefs(ctx, req)
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

func ExampleReferrerClient_AnalyzeParodies() {
	ctx := context.Background()
	c, err := crossrefs.NewReferrerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &emptypb.Empty{
		// TODO: Fill request struct fields.
	}
	op, err := c.AnalyzeParodies(ctx, req)
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

func ExampleReferrerClient_ExportParodies() {
	ctx := context.Background()
	c, err := crossrefs.NewReferrerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &emptypb.Empty{
		// TODO: Fill request struct fields.
	}
	op, err := c.ExportParodies(ctx, req)
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

func ExampleReferrerClient_GetUniverse() {
	ctx := context.Background()
	c, err := crossrefs.NewReferrerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &crossrefspb.GetUniverseRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetUniverse(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleReferrerClient_UpdateUniverse() {
	ctx := context.Background()
	c, err := crossrefs.NewReferrerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &crossrefspb.UpdateUniverseRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateUniverse(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleReferrerClient_ExpandUniverse() {
	ctx := context.Background()
	c, err := crossrefs.NewReferrerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &crossrefspb.ExpandUniverseRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.ExpandUniverse(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleReferrerClient_GetWormhole() {
	ctx := context.Background()
	c, err := crossrefs.NewReferrerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &crossrefspb.GetWormholeRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetWormhole(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleReferrerClient_ListWormholeCrossRefs() {
	ctx := context.Background()
	c, err := crossrefs.NewReferrerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &crossrefspb.ListWormholeCrossRefsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.ListWormholeCrossRefs(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
