import type { AsyncTaskStatus } from '@runapi.ai/core';

export type RunwayAlephAspectRatio = '16:9' | '9:16' | '4:3' | '3:4' | '1:1' | '21:9';

export interface Video {
  id?: string;
  url: string;
}

export interface Image {
  url: string;
}

export interface TaskCreateResponse {
  id: string;
  status?: AsyncTaskStatus;
}

export interface EditVideoResponse {
  id: string;
  status: AsyncTaskStatus;
  videos?: Video[];
  images?: Image[];
  error?: string;
  [key: string]: unknown;
}

export type CompletedEditVideoResponse = EditVideoResponse & {
  status: 'completed';
  videos: Video[];
};

export interface EditVideoParams {
  prompt: string;
  source_video_url: string;
  callback_url?: string;
  watermark?: string;
  aspect_ratio?: RunwayAlephAspectRatio;
  seed?: number;
  reference_image_url?: string;
}
