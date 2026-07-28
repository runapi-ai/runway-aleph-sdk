import type { AsyncTaskStatus, TaskBillingResponse, TaskResponse } from '@runapi.ai/core';

/** Runway Aleph model slug. */
export type RunwayAlephModel = 'runway-aleph';
/** Output aspect ratio. Includes 21:9 ultra-wide for cinematic letterbox output. */
export type RunwayAlephAspectRatio = '16:9' | '9:16' | '4:3' | '3:4' | '1:1' | '21:9';

/** A generated output video. */
export interface Video {
  id?: string;
  url: string;
}

/** A reference or extracted image. */
export interface Image {
  url: string;
}

/** Initial response when a video editing task is created. */
export interface TaskCreateResponse extends TaskBillingResponse {
  id: string;
  status?: AsyncTaskStatus;
}

/** Task status response for a video editing operation. Includes output videos and images when complete. */
export interface EditVideoResponse extends TaskResponse {
  id: string;
  status: AsyncTaskStatus;
  /** Edited output video(s), populated when the task completes. */
  videos?: Video[];
  /** Reference or extracted images, if applicable. */
  images?: Image[];
  /** Human-readable error description when the task fails. */
  error?: string;
  [key: string]: unknown;
}

/** Completed video editing response with guaranteed output videos. */
export type CompletedEditVideoResponse = EditVideoResponse & {
  status: 'completed';
  videos: Video[];
};

/**
 * Parameters for prompt-driven video editing. Provide a text prompt describing the
 * desired transformation and a source video URL. Optionally supply a reference image
 * to guide the visual style of the output.
 */
export interface EditVideoParams {
  /** Model slug. */
  model: RunwayAlephModel;
  /** Text description of the desired transformation (e.g. "Make it look like a watercolor painting"). */
  prompt: string;
  /** Publicly accessible URL of the source video to transform. */
  source_video_url: string;
  /** URL to receive a webhook notification when the task completes. */
  callback_url?: string;
  /** Watermark mode for the output video. */
  watermark?: string;
  /** Output aspect ratio. Includes 21:9 ultra-wide cinematic option. */
  aspect_ratio?: RunwayAlephAspectRatio;
  /** Fixed random seed for reproducible output across identical inputs. */
  seed?: number;
  /** Image URL used as a visual style reference to guide the transformation. */
  reference_image_url?: string;
}
