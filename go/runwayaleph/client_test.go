package runwayaleph

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/runapi-ai/core-sdk/go/core"
)

type stubHTTPClient struct {
	method   string
	path     string
	body     any
	response json.RawMessage
}

func (s *stubHTTPClient) Request(_ context.Context, method, path string, opts *core.HTTPRequestOptions) (json.RawMessage, error) {
	s.method = method
	s.path = path
	if opts != nil {
		s.body = opts.Body
	}
	return s.response, nil
}

func TestEditVideoCreate(t *testing.T) {
	stub := &stubHTTPClient{response: json.RawMessage(`{"id":"task_aleph_123","status":"processing"}`)}
	client := NewClientWithHTTP(stub)
	resp, err := client.EditVideo.Create(context.Background(), EditVideoParams{Prompt: "regrade to dusk", SourceVideoURL: "https://cdn.runapi.ai/public/samples/source.mp4"})
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "POST" || stub.path != editVideoPath {
		t.Fatalf("unexpected request: %s %s", stub.method, stub.path)
	}
	body := stub.body.(map[string]any)
	if body["source_video_url"] != "https://cdn.runapi.ai/public/samples/source.mp4" {
		t.Fatalf("unexpected body: %v", body)
	}
	if _, ok := body["upload_cn"]; ok {
		t.Fatalf("expected request body to omit upload_cn key: %v", body)
	}
	if _, ok := body["video_url"]; ok {
		t.Fatalf("expected request body to omit provider video_url key: %v", body)
	}
	if resp.ID != "task_aleph_123" {
		t.Fatalf("unexpected id: %v", resp.ID)
	}
}

func TestEditVideoGet(t *testing.T) {
	stub := &stubHTTPClient{response: json.RawMessage(`{"id":"task_aleph_456","status":"completed","videos":[{"url":"https://file.runapi.ai/video.mp4"}],"images":[{"url":"https://file.runapi.ai/cover.png"}]}`)}
	client := NewClientWithHTTP(stub)
	resp, err := client.EditVideo.Get(context.Background(), "task_aleph_456")
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "GET" || stub.path != editVideoPath+"/task_aleph_456" {
		t.Fatalf("unexpected request: %s %s", stub.method, stub.path)
	}
	if len(resp.Videos) != 1 || resp.Videos[0].URL != "https://file.runapi.ai/video.mp4" {
		t.Fatalf("unexpected response: %v", resp.Videos)
	}
	if len(resp.Images) != 1 || resp.Images[0].URL != "https://file.runapi.ai/cover.png" {
		t.Fatalf("unexpected images: %v", resp.Images)
	}
}
