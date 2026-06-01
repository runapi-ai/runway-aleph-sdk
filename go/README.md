# Runway Aleph API Go SDK for RunAPI

The runway aleph api Go SDK is the language-specific package for Runway Aleph on RunAPI. Use this runway aleph api package for prompt-guided video editing when your application needs JSON request bodies, task status lookup, and consistent RunAPI errors in Go.

This runway aleph api README is the Go package guide inside the public `runway-aleph-sdk` repository. For the repository overview, start at `../README.md`; for model details, use https://runapi.ai/models/runway-aleph; for API reference, use https://runapi.ai/docs#runway-aleph; for SDK docs, use https://runapi.ai/docs#sdk-runway-aleph.

## Install

```bash
go get github.com/runapi-ai/runway-aleph-sdk/go@latest
```

## Quick start

```go
import (
  "context"

  "github.com/runapi-ai/runway-aleph-sdk/go/runwayaleph"
)

client, err := runwayaleph.NewClient()
task, err := client.EditVideo.Create(context.Background(), runwayaleph.EditVideoParams{
  Model:    "runway-aleph",
  Prompt:   "Transform the scene into a watercolor painting style",
  VideoURL: "https://cdn.runapi.ai/public/samples/video.mp4",
})
status, err := client.EditVideo.Get(context.Background(), task.ID)
```

Use `Create` when you want to submit a task and return quickly, `Get` when you need the latest task state, and `Run` when a script should create and poll until completion. In web request handlers, prefer `Create` plus webhook or later `Get` polling so a worker is not held open.

## Language notes

Use the public Go module with `github.com/runapi-ai/core-sdk/go` options when building video services, CLIs, or workers. Keep `RUNAPI_API_KEY` in the environment or your secret manager; never commit API keys or callback secrets.

## Links

- Model page: https://runapi.ai/models/runway-aleph
- SDK docs: https://runapi.ai/docs#sdk-runway-aleph
- Product docs: https://runapi.ai/docs#runway-aleph
- Pricing and rate limits: https://runapi.ai/models/runway-aleph
- Provider comparison: https://runapi.ai/providers/runway
- Full catalog: https://runapi.ai/models
- Repository: https://github.com/runapi-ai/runway-aleph-sdk

## License

Licensed under the Apache License, Version 2.0.
