import { createHttpClient, type ClientOptions } from '@runapi.ai/core';
import { EditVideo } from './resources/edit-video';

export class RunwayAlephClient {
  public readonly editVideo: EditVideo;

  constructor(options: ClientOptions = {}) {
    const http = createHttpClient(options);
    this.editVideo = new EditVideo(http);
  }
}
