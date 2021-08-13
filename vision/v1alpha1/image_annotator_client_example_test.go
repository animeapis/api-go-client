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

package vision_test

import (
	"context"

	vision "github.com/animeapis/api-go-client/vision/v1alpha1"
	visionpb "github.com/animeapis/go-genproto/vision/v1alpha1"
	"google.golang.org/api/iterator"
)

func ExampleNewImageAnnotatorClient() {
	ctx := context.Background()
	c, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleImageAnnotatorClient_AnalyzeImage() {
	ctx := context.Background()
	c, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &visionpb.AnalyzeImageRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.AnalyzeImage(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleImageAnnotatorClient_ListImageAnalyses() {
	ctx := context.Background()
	c, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &visionpb.ListImageAnalysesRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListImageAnalyses(ctx, req)
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

func ExampleImageAnnotatorClient_GetImageAnalysis() {
	ctx := context.Background()
	c, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &visionpb.GetImageAnalysisRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetImageAnalysis(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleImageAnnotatorClient_DeleteImageAnalysis() {
	ctx := context.Background()
	c, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &visionpb.DeleteImageAnalysisRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteImageAnalysis(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleImageAnnotatorClient_CreateImageAnnotation() {
	ctx := context.Background()
	c, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &visionpb.CreateImageAnnotationRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateImageAnnotation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleImageAnnotatorClient_ListImageAnnotations() {
	ctx := context.Background()
	c, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &visionpb.ListImageAnnotationsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListImageAnnotations(ctx, req)
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

func ExampleImageAnnotatorClient_GetImageAnnotation() {
	ctx := context.Background()
	c, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &visionpb.GetImageAnnotationRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetImageAnnotation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleImageAnnotatorClient_UpdateImageAnnotation() {
	ctx := context.Background()
	c, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &visionpb.UpdateImageAnnotationRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateImageAnnotation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleImageAnnotatorClient_DeleteImageAnnotation() {
	ctx := context.Background()
	c, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &visionpb.DeleteImageAnnotationRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteImageAnnotation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}
