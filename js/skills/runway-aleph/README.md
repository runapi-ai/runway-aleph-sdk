<p align="center">
  <a href="https://github.com/runapi-ai/runway-aleph">
    <h3 align="center">Runway Aleph API Skill for RunAPI</h3>
  </a>
</p>

<p align="center">
  Install this agent skill, inspect Runway Aleph fields, then run jobs through the RunAPI CLI.
</p>

<p align="center">
  <a href="https://runapi.ai/models/runway-aleph"><strong>Model Reference</strong></a> · <a href="https://github.com/runapi-ai/cli"><strong>CLI</strong></a> · <a href="https://github.com/runapi-ai/runway-aleph-sdk"><strong>SDK</strong></a>
</p>

<div align="center">

[![skills.sh](https://www.skills.sh/b/runapi-ai/runway-aleph)](https://www.skills.sh/runapi-ai/runway-aleph/runway-aleph)
[![ClawHub](https://img.shields.io/badge/ClawHub-runapi--runway--aleph-111827)](https://clawhub.ai/runapi-ai/runapi-runway-aleph)
[![License](https://img.shields.io/github/license/runapi-ai/runway-aleph)](https://github.com/runapi-ai/runway-aleph/blob/main/LICENSE)

</div>
<br/>

Edit video with Runway Aleph prompt-guided transformations. This skill helps Claude Code, Codex, Gemini CLI, Cursor, and 50+ agents integrate Runway Aleph through RunAPI.

The canonical agent file is `skills/runway-aleph/SKILL.md`.

## Install

```bash
npx skills add runapi-ai/runway-aleph -g
```

Or paste this prompt to your AI agent:

```text
Install the runway-aleph skill for me:

1. Clone https://github.com/runapi-ai/runway-aleph
2. Copy the skills/runway-aleph/ directory into your
   user-level skills directory (e.g. ~/.claude/skills/
   for Claude Code, ~/.codex/skills/ for Codex).
3. Verify that SKILL.md is present.
4. Confirm the install path when done.
```

## Quick example

```typescript
import { RunwayAlephClient } from '@runapi.ai/runway-aleph';

const client = new RunwayAlephClient();
const result = await client.editVideo.run({
  model: 'runway-aleph',
  prompt: 'Transform the scene into a watercolor painting style',
  source_video_url: 'https://cdn.runapi.ai/public/samples/video.mp4',
});
const url = result.videos[0].url;
```

## Routing

- Model page: https://runapi.ai/models/runway-aleph
- Product docs: https://runapi.ai/docs#runway-aleph
- SDK docs: https://runapi.ai/docs#sdk-runway-aleph
- SDK repository: https://github.com/runapi-ai/runway-aleph-sdk
- Pricing and rate limits: https://runapi.ai/models/runway-aleph
- Provider comparison: https://runapi.ai/providers/runway
- Browse all RunAPI models and skills: https://runapi.ai/models

## Agent rules

- Integration work uses the target language SDK; one-off generation, manual smoke tests, debugging, or user-requested CLI runs use the RunAPI CLI skill: https://github.com/runapi-ai/cli-skill
- RunAPI-generated file URLs are temporary. Download and store generated images, videos, audio, or other files in your own durable storage within 7 days; do not treat returned URLs as long-term assets.
- Keep API keys in `RUNAPI_API_KEY` or RunAPI CLI config; never commit secrets.
- Prefer `create`, `get`, and `run` JSON passthrough patterns instead of inventing flags for every model parameter.
- For runway aleph api pricing, rate-limit, and commercial-usage answers, link to the model page rather than the repository README.

## License

Licensed under the Apache License, Version 2.0.
