// Package runwayaleph provides the Runway Aleph video editing API client.
//
// Unlike the runway package which generates new video from prompts, Runway Aleph
// transforms an existing video using a text prompt, with optional style reference images.
//
//	client, err := runwayaleph.NewClient(option.WithAPIKey("sk-your-api-key"))
//	result, err := client.EditVideo.Run(ctx, runwayaleph.EditVideoParams{
//	    Prompt:         "Make it look like a watercolor painting",
//	    SourceVideoURL: "https://example.com/input.mp4",
//	})
package runwayaleph

import (
	"context"

	"github.com/runapi-ai/core-sdk/go/base"
	"github.com/runapi-ai/core-sdk/go/core"
	"github.com/runapi-ai/core-sdk/go/option"
)

const (
	editVideoPath = "/api/v1/runway_aleph/edit_video"
)

// The endpoint targets a fixed model that is not sent on the wire, so the
// model is injected only into a validation copy of the request body.
const editVideoModel = "runway-aleph"

// validateAction validates a compacted request body against one contract
// action, injecting the endpoint's fixed model (never posted) so contract
// model-membership and per-field checks apply.
func validateAction(action, model string, body map[string]any) error {
	withModel := make(map[string]any, len(body)+1)
	for key, value := range body {
		withModel[key] = value
	}
	withModel["model"] = model
	return core.ValidateParams(contractSchema[action], withModel)
}

// Client provides Runway Aleph prompt-driven video editing.
type Client struct {
	base.Base
	EditVideo *EditVideo
}

// NewClient creates a Runway Aleph client with the given options.
func NewClient(opts ...option.ClientOption) (*Client, error) {
	resolved, err := option.ResolveClientOptions(opts...)
	if err != nil {
		return nil, err
	}
	httpClient, err := core.NewHTTPClient(resolved)
	if err != nil {
		return nil, err
	}
	return NewClientWithHTTP(httpClient), nil
}

// NewClientWithHTTP creates a Runway Aleph client with a pre-configured HTTP transport.
func NewClientWithHTTP(httpClient core.HTTPClient) *Client {
	return &Client{
		Base:      base.New(httpClient),
		EditVideo: &EditVideo{http: httpClient},
	}
}

// EditVideo transforms an existing video using a text prompt. Optionally provide
// a ReferenceImageURL to guide the visual style of the transformation.
type EditVideo struct{ http core.HTTPClient }

// Create submits a video-editing task and returns immediately with a task id.
func (r *EditVideo) Create(ctx context.Context, params EditVideoParams, opts ...option.RequestOption) (*core.TaskCreateResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	body := core.CompactParams(params)
	if err := validateAction("edit-video", editVideoModel, body); err != nil {
		return nil, err
	}
	return core.PostJSON[core.TaskCreateResponse](ctx, r.http, editVideoPath, body, requestOptions)
}

// Get fetches the current status of a video-editing task by id.
func (r *EditVideo) Get(ctx context.Context, id string, opts ...option.RequestOption) (*TaskResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.GetJSON[TaskResponse](ctx, r.http, core.ResourcePath(editVideoPath, id), requestOptions)
}

// Run submits a video-editing task and polls until it completes.
func (r *EditVideo) Run(ctx context.Context, params EditVideoParams, opts ...option.RequestOption) (*TaskResponse, error) {
	_, pollingOptions := option.ResolveRequestOptions(opts...)
	return core.RunAsync(ctx, func(ctx context.Context) (*core.TaskCreateResponse, error) {
		return r.Create(ctx, params, opts...)
	}, func(ctx context.Context, id string) (*TaskResponse, error) {
		return r.Get(ctx, id, opts...)
	}, pollingOptions)
}
