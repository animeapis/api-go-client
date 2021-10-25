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

package graph_test

import (
	"context"

	graph "github.com/animeapis/api-go-client/graph/v1alpha1"
	graphpb "github.com/animeapis/go-genproto/graph/v1alpha1"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := graph.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleClient_MigrateGraph() {
	ctx := context.Background()
	c, err := graph.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &graphpb.MigrateGraphRequest{
		// TODO: Fill request struct fields.
	}
	err = c.MigrateGraph(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_DeleteGraph() {
	ctx := context.Background()
	c, err := graph.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &graphpb.DeleteGraphRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteGraph(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}
