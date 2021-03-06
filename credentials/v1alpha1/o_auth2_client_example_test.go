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

package credentials_test

import (
	"context"

	credentials "github.com/animeapis/api-go-client/credentials/v1alpha1"
	credentialspb "github.com/animeapis/go-genproto/credentials/v1alpha1"
)

func ExampleNewOAuth2Client() {
	ctx := context.Background()
	c, err := credentials.NewOAuth2Client(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleOAuth2Client_SignIn() {
	ctx := context.Background()
	c, err := credentials.NewOAuth2Client(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &credentialspb.SignInRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/credentials/v1alpha1#SignInRequest.
	}
	resp, err := c.SignIn(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleOAuth2Client_Exchange() {
	ctx := context.Background()
	c, err := credentials.NewOAuth2Client(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &credentialspb.ExchangeRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/animeapis/go-genproto/credentials/v1alpha1#ExchangeRequest.
	}
	resp, err := c.Exchange(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
