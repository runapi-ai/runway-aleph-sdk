import type { HttpClient, PollingOptions, RequestOptions } from '@runapi.ai/core';
import { compactParams } from '@runapi.ai/core';
import { pollUntilComplete } from '@runapi.ai/core/internal';
import type { CompletedVideoToVideoResponse, TaskCreateResponse, VideoToVideoParams, VideoToVideoResponse } from '../types';

const ENDPOINT = '/api/v1/runway_aleph/video_to_video';

export class VideoToVideo {
  constructor(private readonly http: HttpClient) {}

  async run(params: VideoToVideoParams, options?: RequestOptions & PollingOptions): Promise<CompletedVideoToVideoResponse> {
    const { id } = await this.create(params, options);
    const response = await pollUntilComplete<VideoToVideoResponse>(() => this.get(id, options), {
      maxWaitMs: options?.maxWaitMs,
      pollIntervalMs: options?.pollIntervalMs,
    });
    return response as CompletedVideoToVideoResponse;
  }

  async create(params: VideoToVideoParams, options?: RequestOptions): Promise<TaskCreateResponse> {
    return this.http.request<TaskCreateResponse>('POST', ENDPOINT, { body: compactParams(params), ...options });
  }

  async get(id: string, options?: RequestOptions): Promise<VideoToVideoResponse> {
    return this.http.request<VideoToVideoResponse>('GET', `${ENDPOINT}/${id}`, options ?? {});
  }
}
