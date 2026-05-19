---
name: runway-aleph
description: Transform video with Runway Aleph video-to-video through RunAPI.ai using the @runapi.ai/runway-aleph Node/TypeScript SDK. Use when the user asks for video-to-video transformation, Runway Aleph, or writes against @runapi.ai/runway-aleph. Triggers on "runway aleph", "video-to-video", "video transform", "@runapi.ai/runway-aleph".
documentation: https://runapi.ai/models/runway-aleph
provider_page: https://runapi.ai/providers/runway
catalog: https://runapi.ai/models
---
# @runapi.ai/runway-aleph -- RunAPI.ai Runway Aleph video-to-video

Build Node / TypeScript integrations that transform video with Runway Aleph through RunAPI.ai.

## Setup

Requires **Node 18+** (global `fetch`).

```bash
npm install @runapi.ai/runway-aleph
```

```dotenv
# .env
RUNAPI_API_KEY=runapi_xxx   # get one at https://runapi.ai/settings/api_keys
```

```ts
import { RunwayAlephClient } from '@runapi.ai/runway-aleph';

const client = new RunwayAlephClient();
```

Pass `{ apiKey }` explicitly if you manage secrets differently. `baseUrl` defaults to `https://runapi.ai`; override only for local development.

## Resource

`client.videoToVideo` uses the async task contract:

```ts
const { id } = await client.videoToVideo.create({ ... });
const status = await client.videoToVideo.get(id);
const result = await client.videoToVideo.run({ ... });
```

## Video to video

```ts
const result = await client.videoToVideo.run({
  model: 'runway-aleph',
  prompt: 'Transform the scene into a watercolor painting style',
  video_url: 'https://cdn.example.com/input.mp4',
});

const url = result.videos[0].url;
```

## Errors

All errors are re-exported from `@runapi.ai/core`. Use `instanceof` checks instead of string-matching messages. For long-running tasks, prefer `create()` plus webhook or `get(id)` in request handlers, and reserve `run()` for jobs / CLI.

## RunAPI public routing

runway aleph api public links use the API-379 catalog route map. The main runway aleph api page is https://runapi.ai/models/runway-aleph. SDK docs live at https://runapi.ai/docs#sdk-runway-aleph and product docs live at https://runapi.ai/docs#runway-aleph.

Pricing, rate limits, and commercial usage for runway aleph api should point to the most specific variant page:
- [Runway Aleph](https://runapi.ai/models/runway-aleph/runway-aleph)

Compare Runway Aleph with other Runway models at https://runapi.ai/providers/runway. Browse every RunAPI model and skill at https://runapi.ai/models. SDK repository: https://github.com/runapi-ai/runway-aleph-sdk. Skill repository: https://github.com/runapi-ai/runway-aleph.
