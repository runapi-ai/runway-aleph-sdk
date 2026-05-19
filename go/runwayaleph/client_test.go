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

func TestVideoToVideoCreate(t *testing.T) {
	stub := &stubHTTPClient{response: json.RawMessage(`{"id":"task_aleph_123","status":"processing"}`)}
	client := NewClientWithHTTP(stub)
	uploadCN := false
	resp, err := client.VideoToVideo.Create(context.Background(), VideoToVideoParams{Prompt: "regrade to dusk", VideoURL: "https://example.com/source.mp4", UploadCN: &uploadCN})
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "POST" || stub.path != videoToVideoPath {
		t.Fatalf("unexpected request: %s %s", stub.method, stub.path)
	}
	body := stub.body.(map[string]any)
	if body["video_url"] != "https://example.com/source.mp4" || body["upload_cn"] != false {
		t.Fatalf("unexpected body: %v", body)
	}
	if resp.ID != "task_aleph_123" {
		t.Fatalf("unexpected id: %v", resp.ID)
	}
}

func TestVideoToVideoGet(t *testing.T) {
	stub := &stubHTTPClient{response: json.RawMessage(`{"id":"task_aleph_456","status":"completed","videos":[{"url":"https://file.runapi.ai/video.mp4"}]}`)}
	client := NewClientWithHTTP(stub)
	resp, err := client.VideoToVideo.Get(context.Background(), "task_aleph_456")
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "GET" || stub.path != videoToVideoPath+"/task_aleph_456" {
		t.Fatalf("unexpected request: %s %s", stub.method, stub.path)
	}
	if len(resp.Videos) != 1 || resp.Videos[0].URL != "https://file.runapi.ai/video.mp4" {
		t.Fatalf("unexpected response: %v", resp.Videos)
	}
}
