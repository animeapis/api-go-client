// Copyright 2023 Google LLC
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

package webpage_test

import (
	"context"
	"io"

	webpage "github.com/animeapis/api-go-client/webpage/v1alpha1"
	webpagepb "github.com/animeapis/go-genproto/webpage/v1alpha1"
	"google.golang.org/api/iterator"
)

func ExampleNewArchiveClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := webpage.NewArchiveClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleNewArchiveRESTClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := webpage.NewArchiveRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleArchiveClient_Query() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := webpage.NewArchiveClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()
	stream, err := c.Query(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	go func() {
		reqs := []*webpagepb.QueryRequest{
			// TODO: Create requests.
		}
		for _, req := range reqs {
			if err := stream.Send(req); err != nil {
				// TODO: Handle error.
			}
		}
		stream.CloseSend()
	}()
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			// TODO: handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleArchiveClient_GetPage() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := webpage.NewArchiveClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &webpagepb.GetPageRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/webpage/v1alpha1#GetPageRequest.
	}
	resp, err := c.GetPage(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleArchiveClient_ListPages() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := webpage.NewArchiveClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &webpagepb.ListPagesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/webpage/v1alpha1#ListPagesRequest.
	}
	it := c.ListPages(ctx, req)
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

func ExampleArchiveClient_QueryPage() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := webpage.NewArchiveClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &webpagepb.QueryPageRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/webpage/v1alpha1#QueryPageRequest.
	}
	resp, err := c.QueryPage(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleArchiveClient_CreatePage() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := webpage.NewArchiveClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &webpagepb.CreatePageRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/webpage/v1alpha1#CreatePageRequest.
	}
	resp, err := c.CreatePage(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleArchiveClient_ImportPage() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := webpage.NewArchiveClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &webpagepb.ImportPageRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/webpage/v1alpha1#ImportPageRequest.
	}
	resp, err := c.ImportPage(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleArchiveClient_DeletePage() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := webpage.NewArchiveClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &webpagepb.DeletePageRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/webpage/v1alpha1#DeletePageRequest.
	}
	err = c.DeletePage(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}
