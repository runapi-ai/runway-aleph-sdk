# Runway Aleph Python SDK for RunAPI

The Runway Aleph Python SDK is the language-specific package for Runway Aleph on RunAPI. Use this package to query historical video editing tasks and handle RunAPI errors in Python.

This README is the Python package guide inside the public `runway-aleph-sdk` repository. For the repository overview, start at `../README.md`; for model details, use https://runapi.ai/models/runway-aleph; for API reference, use https://runapi.ai/docs/api/runway-aleph/edit-video; for SDK docs, use https://runapi.ai/docs/resources/sdks.

## Retirement

`runway-aleph` was retired on 2026-10-11. New `edit_video` requests return HTTP 410 before task creation or balance deduction. Existing task IDs remain queryable. This model ID is not mapped to another model.

## Install

```bash
pip install runapi-runway-aleph
```

## Historical request reference

The request examples below describe the retired API. Do not submit new tasks. Use `get` with an existing task ID to retrieve its status.

```python
from runapi.runway_aleph import RunwayAlephClient

client = RunwayAlephClient()  # reads RUNAPI_API_KEY, or pass api_key="sk-..."

task = client.edit_video.create(
    model="runway-aleph",
    prompt="Transform the scene into a watercolor painting style",
    source_video_url="https://cdn.runapi.ai/public/samples/video.mp4",
)
status = client.edit_video.get(task.id)
```

Use the task-status lookup with an existing task ID. Task creation and submit-and-poll helpers return HTTP 410 for this retired model.

```python
result = client.edit_video.run(
    model="runway-aleph",
    prompt="Transform the scene into a watercolor painting style",
    source_video_url="https://cdn.runapi.ai/public/samples/video.mp4",
)
print(result.videos[0].url)
```

Do not call `create` or `run` for new tasks.

RunAPI-generated file URLs are temporary. Download and store generated images, videos, audio, or other files in your own durable storage within 7 days; do not treat returned URLs as long-term assets.

## Language notes

Pass parameters as keyword arguments and catch the `runapi.runway_aleph` error classes when building video jobs, workers, or scripts. The available resource is `edit_video`. Keep `RUNAPI_API_KEY` in the environment or your secret manager; never commit API keys or callback secrets.

## Links

- Model page: https://runapi.ai/models/runway-aleph
- SDK docs: https://runapi.ai/docs/resources/sdks
- Product docs: https://runapi.ai/docs/api/runway-aleph/edit-video
- Pricing and rate limits: https://runapi.ai/models/runway-aleph
- Provider comparison: https://runapi.ai/providers/runway
- Full catalog: https://runapi.ai/models
- Repository: https://github.com/runapi-ai/runway-aleph-sdk

## License

Licensed under the Apache License, Version 2.0.
