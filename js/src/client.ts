import { BaseClient, type ClientOptions } from '@runapi.ai/core';
import { EditVideo } from './resources/edit-video';

/**
 * Runway Aleph prompt-driven video editing API client.
 * Unlike generation from scratch, Runway Aleph transforms an existing video
 * using a text prompt, with optional style reference images.
 *
 * @example
 * ```typescript
 * const client = new RunwayAlephClient({ apiKey: 'your-api-key' });
 *
 * const result = await client.editVideo.run({
 *   prompt: 'Make it look like a watercolor painting',
 *   source_video_url: 'https://example.com/input.mp4',
 * });
 * ```
 */
export class RunwayAlephClient extends BaseClient {
  /** Transform an existing video using a text prompt, with optional style reference image. */
  public readonly editVideo: EditVideo;

  constructor(options: ClientOptions = {}) {
    super(options);
    this.editVideo = new EditVideo(this.http);
  }
}
