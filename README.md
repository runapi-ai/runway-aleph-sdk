# Runway Aleph API SDK for RunAPI

The runway aleph api SDK packages JavaScript, Ruby, and Go clients for Runway Aleph on RunAPI. Use this runway aleph api SDK for video-to-video transformation workflows that need typed installs, JSON request bodies, task polling, and consistent RunAPI errors across services.

Runway Aleph belongs to the Runway catalog on RunAPI. The public model page is https://runapi.ai/models/runway-aleph; variant pages below carry pricing, rate-limit, and commercial-usage details. The public `runway-aleph-sdk` repository groups the JavaScript, Ruby, and Go packages for this model.

## Install

```bash
npm install @runapi.ai/runway-aleph
gem install runapi-runway_aleph
go get github.com/runapi-ai/runway-aleph-sdk/go@latest
```

## What you can build

- Build creative tools, agent pipelines, and production integrations with the runway aleph api SDK.
- Keep one model-specific repository while installing only the language package your app needs.
- Use `create` for submit-only jobs, `get` for status lookup, and `run` for submit-and-poll scripts.
- Handle authentication, validation, rate limits, insufficient credits, task failures, and polling timeouts through RunAPI SDK errors.

The JavaScript client exposes video to video resources, and the Ruby and Go packages mirror the same RunAPI task lifecycle.

## JavaScript quick start

```typescript
import { RunwayAlephClient } from '@runapi.ai/runway-aleph';

const client = new RunwayAlephClient();

const task = await client.videoToVideo.create({
  // Pass the Runway Aleph request body documented at https://runapi.ai/docs#runway-aleph.
});

const status = await client.videoToVideo.get(task.id);
```

For short scripts, use `run` with the same JSON body to create the task and wait for completion. For web request handlers, prefer `create` plus webhook or later `get` polling so the server does not hold a worker open.

## Repository layout

- `js/` publishes `@runapi.ai/runway-aleph`.
- `ruby/` publishes `runapi-runway_aleph` when RubyGems publishing resumes.
- `go/` publishes `github.com/runapi-ai/runway-aleph-sdk/go` and depends on `github.com/runapi-ai/core-sdk/go`.

## Public links

- Model page: https://runapi.ai/models/runway-aleph
- SDK docs: https://runapi.ai/docs#sdk-runway-aleph
- Product docs: https://runapi.ai/docs#runway-aleph
- SDK repository: https://github.com/runapi-ai/runway-aleph-sdk
- Skill repository: https://github.com/runapi-ai/runway-aleph
- Provider comparison: https://runapi.ai/providers/runway
- Full catalog: https://runapi.ai/models

## Pricing and variants

Use the most specific runway aleph api variant page for pricing, rate limits, and commercial usage:
- [Runway Aleph](https://runapi.ai/models/runway-aleph/runway-aleph)

Default pricing link for the runway aleph api SDK: https://runapi.ai/models/runway-aleph/runway-aleph

## FAQ

### Which package should I install for runway aleph api work?

Install the model package for your language: `@runapi.ai/runway-aleph`, `runapi-runway_aleph`, or `github.com/runapi-ai/runway-aleph-sdk/go`. Install core SDK packages only when you are building shared SDK infrastructure.

### Where should public links point?

Primary runway aleph api links point to https://runapi.ai/models/runway-aleph. Pricing and usage-policy links point to variant pages such as https://runapi.ai/models/runway-aleph/runway-aleph. Provider comparisons point to https://runapi.ai/providers/runway, and broad browsing points to https://runapi.ai/models.

## License

Licensed under the Apache License, Version 2.0.
