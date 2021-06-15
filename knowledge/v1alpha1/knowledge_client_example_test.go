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

package knowledge_test

import (
	"context"

	knowledge "github.com/animeapis/api-go-client/knowledge/v1alpha1"
	knowledgepb "github.com/animeapis/go-genproto/knowledge/v1alpha1"
	"google.golang.org/api/iterator"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := knowledge.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleClient_CreateContribution() {
	ctx := context.Background()
	c, err := knowledge.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &knowledgepb.CreateContributionRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateContribution(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListContributions() {
	ctx := context.Background()
	c, err := knowledge.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &knowledgepb.ListContributionsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListContributions(ctx, req)
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

func ExampleClient_GetContribution() {
	ctx := context.Background()
	c, err := knowledge.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &knowledgepb.GetContributionRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetContribution(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ApproveContribution() {
	ctx := context.Background()
	c, err := knowledge.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &knowledgepb.ApproveContributionRequest{
		// TODO: Fill request struct fields.
	}
	err = c.ApproveContribution(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_RejectContribution() {
	ctx := context.Background()
	c, err := knowledge.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &knowledgepb.RejectContributionRequest{
		// TODO: Fill request struct fields.
	}
	err = c.RejectContribution(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}