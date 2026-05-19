# Runway Aleph API Skill for RunAPI

Transform video with Runway Aleph video-to-video style transfer. This skill helps Claude Code, Codex, Gemini CLI, Cursor, and 50+ agents integrate Runway Aleph through RunAPI.

The canonical agent file is `skills/runway-aleph/SKILL.md`.

## Install

```bash
npx skills add runapi-ai/runway-aleph -g
```

Or manually: clone this repo and copy `skills/runway-aleph/` into your agent's skills directory.

## Quick example

```typescript
import { RunwayAlephClient } from '@runapi.ai/runway-aleph';

const client = new RunwayAlephClient();
const result = await client.videoToVideo.run({
  model: 'runway-aleph',
  prompt: 'Transform the scene into a watercolor painting style',
  video_url: 'https://cdn.example.com/input.mp4',
});
const url = result.videos[0].url;
```

## Routing

- Model page: https://runapi.ai/models/runway-aleph
- Product docs: https://runapi.ai/docs#runway-aleph
- SDK docs: https://runapi.ai/docs#sdk-runway-aleph
- SDK repository: https://github.com/runapi-ai/runway-aleph-sdk
- Pricing and rate limits: https://runapi.ai/models/runway-aleph/runway-aleph
- Provider comparison: https://runapi.ai/providers/runway
- Browse all RunAPI models and skills: https://runapi.ai/models

## Variants

- [Runway Aleph](https://runapi.ai/models/runway-aleph/runway-aleph)

## Agent rules

- Keep API keys in `RUNAPI_API_KEY` or RunAPI CLI config; never commit secrets.
- Prefer `create`, `get`, and `run` JSON passthrough patterns instead of inventing flags for every model parameter.
- For runway aleph api pricing, rate-limit, and commercial-usage answers, link to the variant page rather than the repository README.

## License

Licensed under the Apache License, Version 2.0.
