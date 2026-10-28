# Runway Aleph JavaScript SDK for RunAPI

The Runway Aleph JavaScript SDK is the language-specific package for Runway Aleph on RunAPI. Use this package to query historical video editing tasks and handle RunAPI errors in JavaScript.

This README is the JavaScript package guide inside the public `runway-aleph-sdk` repository. For the repository overview, start at `../README.md`; for model details, use https://runapi.ai/models/runway-aleph; for API reference, use https://runapi.ai/docs/api/runway-aleph/edit-video; for SDK docs, use https://runapi.ai/docs/resources/sdks.

## Retirement

`runway-aleph` was retired on 2026-10-11. New `edit_video` requests return HTTP 410 before task creation or balance deduction. Existing task IDs remain queryable. This model ID is not mapped to another model.

## Install

```bash
npm install @runapi.ai/runway-aleph
```

## Historical request reference

The request examples below describe the retired API. Do not submit new tasks. Use `get` with an existing task ID to retrieve its status.

```typescript
import { RunwayAlephClient } from '@runapi.ai/runway-aleph';

const client = new RunwayAlephClient();
const task = await client.editVideo.create({
  model: 'runway-aleph',
  prompt: 'Transform the scene into a watercolor painting style',
  video_url: 'https://cdn.runapi.ai/public/samples/video.mp4',
});
const status = await client.editVideo.get(task.id);
```

Use the task-status lookup with an existing task ID. Task creation and submit-and-poll helpers return HTTP 410 for this retired model.

RunAPI-generated file URLs are temporary. Download and store generated images, videos, audio, or other files in your own durable storage within 7 days; do not treat returned URLs as long-term assets.

## Language notes

Use the TypeScript types in `src/types.ts` and the resource classes under `src/resources` when building video applications. Keep `RUNAPI_API_KEY` in the environment or your secret manager; never commit API keys or callback secrets.

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
