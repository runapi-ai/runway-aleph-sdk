package runwayaleph

import (
	"context"

	"github.com/runapi-ai/core-sdk/go/core"
	"github.com/runapi-ai/core-sdk/go/option"
)

const (
	videoToVideoPath = "/api/v1/runway_aleph/video_to_video"
)

type Client struct {
	VideoToVideo *VideoToVideo
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
		VideoToVideo: &VideoToVideo{http: httpClient},
	}
}

type VideoToVideo struct{ http core.HTTPClient }

func (r *VideoToVideo) Create(ctx context.Context, params VideoToVideoParams, opts ...option.RequestOption) (*core.TaskCreateResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.PostJSON[core.TaskCreateResponse](ctx, r.http, videoToVideoPath, core.CompactParams(params), requestOptions)
}

func (r *VideoToVideo) Get(ctx context.Context, id string, opts ...option.RequestOption) (*TaskResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.GetJSON[TaskResponse](ctx, r.http, core.ResourcePath(videoToVideoPath, id), requestOptions)
}

func (r *VideoToVideo) Run(ctx context.Context, params VideoToVideoParams, opts ...option.RequestOption) (*TaskResponse, error) {
	_, pollingOptions := option.ResolveRequestOptions(opts...)
	return core.RunAsync(ctx, func(ctx context.Context) (*core.TaskCreateResponse, error) {
		return r.Create(ctx, params, opts...)
	}, func(ctx context.Context, id string) (*TaskResponse, error) {
		return r.Get(ctx, id, opts...)
	}, pollingOptions)
}
