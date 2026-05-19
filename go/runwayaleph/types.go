package runwayaleph

type RunwayAlephAspectRatio string

type TaskStatus string

const (
	Aspect16x9 RunwayAlephAspectRatio = "16:9"
	Aspect9x16 RunwayAlephAspectRatio = "9:16"
	Aspect4x3  RunwayAlephAspectRatio = "4:3"
	Aspect3x4  RunwayAlephAspectRatio = "3:4"
	Aspect1x1  RunwayAlephAspectRatio = "1:1"
	Aspect21x9 RunwayAlephAspectRatio = "21:9"
)

type VideoToVideoParams struct {
	Prompt            string                 `json:"prompt" help:"required; transformation prompt"`
	VideoURL          string                 `json:"video_url" help:"required; source video URL"`
	CallbackURL       string                 `json:"callback_url,omitempty" help:"optional; webhook URL"`
	Watermark         string                 `json:"watermark,omitempty" help:"optional; watermark text"`
	UploadCN          *bool                  `json:"upload_cn,omitempty" help:"optional; use CN storage path"`
	AspectRatio       RunwayAlephAspectRatio `json:"aspect_ratio,omitempty" help:"optional; 16:9, 9:16, 4:3, 3:4, 1:1, 21:9"`
	Seed              *int                   `json:"seed,omitempty" help:"optional; random seed"`
	ReferenceImageURL string                 `json:"reference_image_url,omitempty" help:"optional; style reference image URL"`
}

type AsyncTaskResponse struct {
	ID     string     `json:"id"`
	Status TaskStatus `json:"status"`
	Error  string     `json:"error,omitempty"`
}

func (r AsyncTaskResponse) GetID() string     { return r.ID }
func (r AsyncTaskResponse) GetStatus() string { return string(r.Status) }
func (r AsyncTaskResponse) GetError() string  { return r.Error }

type Video struct {
	ID  string `json:"id,omitempty"`
	URL string `json:"url"`
}

type TaskResponse struct {
	AsyncTaskResponse
	Videos       []Video `json:"videos,omitempty"`
	ImageURL     string  `json:"image_url,omitempty"`
	ParentTaskID string  `json:"parent_task_id,omitempty"`
}
