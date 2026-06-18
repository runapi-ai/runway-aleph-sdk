# Runway Aleph Python SDK for RunAPI

The Runway Aleph Python SDK is the language-specific package for Runway Aleph on RunAPI. Use this runway aleph package for prompt-guided video editing when your application needs JSON request bodies, task status lookup, and consistent RunAPI errors in Python.

This runway aleph README is the Python package guide inside the public `runway-aleph-sdk` repository. For the repository overview, start at `../README.md`; for model details, use https://runapi.ai/models/runway-aleph; for API reference, use https://runapi.ai/docs#runway-aleph; for SDK docs, use https://runapi.ai/docs#sdk-runway-aleph.

## Install

```bash
pip install runapi-runway-aleph
```

## Quick start

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

Use `create` when you want to submit a task and return quickly, `get` when you need the latest task state, and `run` when a script should create and poll until completion:

```python
result = client.edit_video.run(
    model="runway-aleph",
    prompt="Transform the scene into a watercolor painting style",
    source_video_url="https://cdn.runapi.ai/public/samples/video.mp4",
)
print(result.videos[0].url)
```

In web request handlers, prefer `create` plus webhook or later `get` polling so a worker is not held open.

RunAPI-generated file URLs are temporary. Download and store generated images, videos, audio, or other files in your own durable storage within 7 days; do not treat returned URLs as long-term assets.

## Language notes

Pass parameters as keyword arguments and catch the `runapi.runway_aleph` error classes when building video jobs, workers, or scripts. The available resource is `edit_video`. Keep `RUNAPI_API_KEY` in the environment or your secret manager; never commit API keys or callback secrets.

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
