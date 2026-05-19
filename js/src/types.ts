import type { AsyncTaskStatus } from '@runapi.ai/core';

export type RunwayAlephAspectRatio = '16:9' | '9:16' | '4:3' | '3:4' | '1:1' | '21:9';

export interface Video {
  id?: string;
  url: string;
}

export interface TaskCreateResponse {
  id: string;
  status?: AsyncTaskStatus;
}

export interface VideoToVideoResponse {
  id: string;
  status: AsyncTaskStatus;
  videos?: Video[];
  image_url?: string;
  parent_task_id?: string;
  error?: string;
  [key: string]: unknown;
}

export type CompletedVideoToVideoResponse = VideoToVideoResponse & {
  status: 'completed';
  videos: Video[];
};

export interface VideoToVideoParams {
  prompt: string;
  video_url: string;
  callback_url?: string;
  watermark?: string;
  upload_cn?: boolean;
  aspect_ratio?: RunwayAlephAspectRatio;
  seed?: number;
  reference_image_url?: string;
}
