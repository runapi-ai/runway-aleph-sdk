import { createHttpClient, type ClientOptions } from '@runapi.ai/core';
import { VideoToVideo } from './resources/video-to-video';

export class RunwayAlephClient {
  public readonly videoToVideo: VideoToVideo;

  constructor(options: ClientOptions = {}) {
    const http = createHttpClient(options);
    this.videoToVideo = new VideoToVideo(http);
  }
}
