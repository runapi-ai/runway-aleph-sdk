# Runway Aleph API JavaScript SDK for RunAPI

The runway aleph api JavaScript SDK is the language-specific package for Runway Aleph on RunAPI. Use this runway aleph api package for prompt-guided video editing when your application needs JSON request bodies, task status lookup, and consistent RunAPI errors in JavaScript.

This runway aleph api README is the JavaScript package guide inside the public `runway-aleph-sdk` repository. For the repository overview, start at `../README.md`; for model details, use https://runapi.ai/models/runway-aleph; for API reference, use https://runapi.ai/docs#runway-aleph; for SDK docs, use https://runapi.ai/docs#sdk-runway-aleph.

## Install

```bash
npm install @runapi.ai/runway-aleph
```

## Quick start

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

Use `create` when you want to submit a task and return quickly, `get` when you need the latest task state, and `run` when a script should create and poll until completion. In web request handlers, prefer `create` plus webhook or later `get` polling so a worker is not held open.

## Language notes

Use the TypeScript types in `src/types.ts` and the resource classes under `src/resources` when building video applications. Keep `RUNAPI_API_KEY` in the environment or your secret manager; never commit API keys or callback secrets.

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
