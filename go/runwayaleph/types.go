// Package runwayaleph provides the Runway Aleph video editing API client.
package runwayaleph

// RunwayAlephAspectRatio controls the output video aspect ratio.
// Runway Aleph supports an additional ultra-wide 21:9 ratio beyond the standard set.
type RunwayAlephAspectRatio string

// TaskStatus is the async task lifecycle state (e.g. "processing", "completed", "failed").
type TaskStatus string

const (
	Aspect16x9 RunwayAlephAspectRatio = "16:9"
	Aspect9x16 RunwayAlephAspectRatio = "9:16"
	Aspect4x3  RunwayAlephAspectRatio = "4:3"
	Aspect3x4  RunwayAlephAspectRatio = "3:4"
	Aspect1x1  RunwayAlephAspectRatio = "1:1"
	// Aspect21x9 is the ultra-wide cinematic ratio.
	Aspect21x9 RunwayAlephAspectRatio = "21:9"
)

// EditVideoParams configures prompt-driven video editing.
// Set ReferenceImageURL to guide the visual style with a reference image (style transfer).
// Set Seed for reproducible output across identical inputs.
type EditVideoParams struct {
	Prompt            string                 `json:"prompt" help:"required; transformation prompt"`
	SourceVideoURL    string                 `json:"source_video_url" help:"required; source video URL"`
	CallbackURL       string                 `json:"callback_url,omitempty" help:"optional; webhook URL"`
	Watermark         string                 `json:"watermark,omitempty" help:"optional; watermark text"`
	AspectRatio       RunwayAlephAspectRatio `json:"aspect_ratio,omitempty" help:"optional; output aspect ratio"`
	Seed              *int                   `json:"seed,omitempty" help:"optional; random seed"`
	ReferenceImageURL string                 `json:"reference_image_url,omitempty" help:"optional; style reference image URL"`
}

// AsyncTaskResponse carries the task ID, lifecycle status, and error for Runway Aleph async operations.
type AsyncTaskResponse struct {
	ID     string     `json:"id"`
	Status TaskStatus `json:"status"`
	Error  string     `json:"error,omitempty"`
}

func (r AsyncTaskResponse) GetID() string     { return r.ID }
func (r AsyncTaskResponse) GetStatus() string { return string(r.Status) }
func (r AsyncTaskResponse) GetError() string  { return r.Error }

// Video holds a URL to a generated or edited video file.
type Video struct {
	ID  string `json:"id,omitempty"`
	URL string `json:"url"`
}

// Image holds a URL to an image.
type Image struct {
	URL string `json:"url"`
}

// TaskResponse is the result of a completed EditVideo task.
type TaskResponse struct {
	AsyncTaskResponse
	Videos []Video `json:"videos,omitempty"`
	Images []Image `json:"images,omitempty"`
}
