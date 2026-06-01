package runwayaleph

import (
	"context"

	"github.com/runapi-ai/core-sdk/go/core"
	"github.com/runapi-ai/core-sdk/go/option"
)

const (
	editVideoPath = "/api/v1/runway_aleph/edit_video"
)

type Client struct {
	EditVideo *EditVideo
}

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

func NewClientWithHTTP(httpClient core.HTTPClient) *Client {
	return &Client{
		EditVideo: &EditVideo{http: httpClient},
	}
}

type EditVideo struct{ http core.HTTPClient }

func (r *EditVideo) Create(ctx context.Context, params EditVideoParams, opts ...option.RequestOption) (*core.TaskCreateResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.PostJSON[core.TaskCreateResponse](ctx, r.http, editVideoPath, core.CompactParams(params), requestOptions)
}

func (r *EditVideo) Get(ctx context.Context, id string, opts ...option.RequestOption) (*TaskResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.GetJSON[TaskResponse](ctx, r.http, core.ResourcePath(editVideoPath, id), requestOptions)
}

func (r *EditVideo) Run(ctx context.Context, params EditVideoParams, opts ...option.RequestOption) (*TaskResponse, error) {
	_, pollingOptions := option.ResolveRequestOptions(opts...)
	return core.RunAsync(ctx, func(ctx context.Context) (*core.TaskCreateResponse, error) {
		return r.Create(ctx, params, opts...)
	}, func(ctx context.Context, id string) (*TaskResponse, error) {
		return r.Get(ctx, id, opts...)
	}, pollingOptions)
}
