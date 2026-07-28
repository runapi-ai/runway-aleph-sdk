<p align="center">
  <a href="https://runapi.ai"><img src="https://runapi.ai/icon.svg" height="56" alt="RunAPI"></a>
</p>

<h3 align="center">
  <a href="https://github.com/runapi-ai/runway-aleph-sdk">Runway Aleph API SDK for RunAPI</a>
</h3>

<p align="center">
  Runway Aleph API SDKs for JavaScript, Python, Ruby, Go, Java, and PHP on RunAPI.
</p>

<div align="center">

[![npm](https://img.shields.io/npm/v/@runapi.ai/runway-aleph)](https://www.npmjs.com/package/@runapi.ai/runway-aleph)
[![PyPI](https://img.shields.io/pypi/v/runapi-runway-aleph)](https://pypi.org/project/runapi-runway-aleph/)
[![RubyGems](https://img.shields.io/gem/v/runapi-runway-aleph)](https://rubygems.org/gems/runapi-runway-aleph)
[![Go Reference](https://pkg.go.dev/badge/github.com/runapi-ai/runway-aleph-sdk/go.svg)](https://pkg.go.dev/github.com/runapi-ai/runway-aleph-sdk/go)
[![Maven Central](https://img.shields.io/maven-central/v/ai.runapi/runapi-runway-aleph)](https://central.sonatype.com/artifact/ai.runapi/runapi-runway-aleph)
[![License](https://img.shields.io/github/license/runapi-ai/runway-aleph-sdk)](https://github.com/runapi-ai/runway-aleph-sdk/blob/main/LICENSE)

</div>
<br/>

The Runway Aleph API SDK packages JavaScript, Python, Ruby, Go, Java, and PHP clients for Runway Aleph on RunAPI. Use it for prompt-guided video editing workflows when your app needs typed request builders, predictable task polling, file upload helpers, account helpers, and consistent RunAPI errors.

Runway Aleph is listed in the RunAPI model catalog at https://runapi.ai/models/runway-aleph. Variant pages below carry pricing, rate-limit, and commercial-usage details. The public `runway-aleph-sdk` repository groups the non-PHP language packages, examples, CI, and release tags for this model. The PHP package is released from a split Composer repository.

## Install

```bash
npm install @runapi.ai/runway-aleph
pip install runapi-runway-aleph
gem install runapi-runway-aleph
go get github.com/runapi-ai/runway-aleph-sdk/go@latest
```

Gradle:

```kotlin
dependencies {
  implementation("ai.runapi:runapi-runway-aleph:0.1.1")
}
```

Maven:

```xml
<dependency>
  <groupId>ai.runapi</groupId>
  <artifactId>runapi-runway-aleph</artifactId>
  <version>0.1.1</version>
</dependency>
```

Use the Java BOM when installing multiple RunAPI Java modules:

```kotlin
dependencies {
  implementation(platform("ai.runapi:runapi-bom:0.2.7"))
  implementation("ai.runapi:runapi-runway-aleph")
}
```

The PHP package is published from the split Composer repository as `runapi-ai/runway-aleph`; see https://github.com/runapi-ai/runway-aleph-php for PHP install and examples.

## What you can build

- Build apps, agent workflows, batch jobs, and production services around Runway Aleph requests.
- Install only the language package your app needs while keeping one model-specific repository for docs and releases.
- Use `create` for submit-only jobs, `get` for status lookup, and `run` for submit-and-poll scripts.
- Upload local files, URL files, or base64 files through shared RunAPI file helpers.
- Handle validation, authentication, rate limits, insufficient credits, task failures, and polling timeouts through RunAPI SDK errors.

## Java quick start

```java
import ai.runapi.runwayaleph.RunwayAlephClient;
import ai.runapi.runwayaleph.types.EditVideoParams;
import ai.runapi.runwayaleph.types.CompletedEditVideoResponse;
import ai.runapi.runwayaleph.types.EditVideoModel;

RunwayAlephClient client = RunwayAlephClient.builder()
    .apiKey(System.getenv("RUNAPI_API_KEY"))
    .build();

CompletedEditVideoResponse result = client.editVideo().run(
    EditVideoParams.builder()
        .model(EditVideoModel.RUNWAY_ALEPH)
        .prompt("Change the scene to golden hour")
        .sourceVideoUrl("https://cdn.runapi.ai/public/samples/video.mp4")
        .aspectRatio("16:9")
        .build()
);
```

Java packages target Java 8 bytecode and are tested on Java 8, 11, 17, and 21. Each model artifact depends on `ai.runapi:runapi-core`, so application code normally installs only `ai.runapi:runapi-runway-aleph`.

## Task lifecycle

Most media endpoints are asynchronous. `create()` submits a task and returns its id, `get(id)` fetches the latest task state, and `run(params)` creates the task and polls until it reaches a terminal state. In web request handlers, prefer `create()` plus webhook or later `get()` polling so the server does not hold a worker open.

## Repository layout

- `js/` publishes `@runapi.ai/runway-aleph`.
- `python/` publishes `runapi-runway-aleph`.
- `ruby/` publishes `runapi-runway-aleph`.
- `go/` publishes `github.com/runapi-ai/runway-aleph-sdk/go` and depends on `github.com/runapi-ai/core-sdk/go`.
- `java/` publishes `ai.runapi:runapi-runway-aleph` and depends on `ai.runapi:runapi-core`.

## Public links

- Model page: https://runapi.ai/models/runway-aleph
- SDK docs: https://runapi.ai/docs#sdk-runway-aleph
- Product docs: https://runapi.ai/docs#runway-aleph
- SDK repository: https://github.com/runapi-ai/runway-aleph-sdk
- PHP package repository: https://github.com/runapi-ai/runway-aleph-php
- Skill repository: https://github.com/runapi-ai/runway-aleph
- Provider comparison: https://runapi.ai/providers/runway
- Full catalog: https://runapi.ai/models

## Pricing and variants

Use the most specific Runway Aleph variant page for pricing, rate limits, and commercial usage:
- [Runway Aleph](https://runapi.ai/models/runway-aleph)

Default pricing link for the Runway Aleph SDK: https://runapi.ai/models/runway-aleph

## File storage

RunAPI-generated file URLs are temporary. Download and store generated images, videos, audio, or other files in your own durable storage within 7 days; do not treat returned URLs as long-term assets.

## FAQ

### Which package should I install for Runway Aleph work?

Install the model package for your language: `@runapi.ai/runway-aleph` on npm, `runapi-runway-aleph` on PyPI, `runapi-runway-aleph` on RubyGems, `github.com/runapi-ai/runway-aleph-sdk/go`, `ai.runapi:runapi-runway-aleph` on Maven Central, or `runapi-ai/runway-aleph` on Packagist. Install core SDK packages only when you are building shared SDK infrastructure.

### Where should public links point?

Primary Runway Aleph links point to https://runapi.ai/models/runway-aleph. Pricing and usage-policy links point to variant pages such as https://runapi.ai/models/runway-aleph. Provider comparisons point to https://runapi.ai/providers/runway, and broad browsing points to https://runapi.ai/models.

## License

Licensed under the Apache License, Version 2.0.
