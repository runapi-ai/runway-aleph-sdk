import type { HttpClient, PollingOptions, RequestOptions, ActionSchema } from '@runapi.ai/core';
import { compactParams, validateParams } from '@runapi.ai/core';
import { pollUntilComplete } from '@runapi.ai/core/internal';
import { contract } from '../contract_gen';
import type { CompletedEditVideoResponse, EditVideoParams, EditVideoResponse, TaskCreateResponse } from '../types';

const ENDPOINT = '/api/v1/runway_aleph/edit_video';

// Fixed endpoint model, injected only for contract validation (never sent on the wire).
const MODEL = 'runway-aleph';

/** Transform an existing video using a text prompt. Optionally provide a reference_image_url to guide the visual style of the transformation. */
export class EditVideo {
  constructor(private readonly http: HttpClient) {}

  /**
   * Create an edit video task and wait until complete.
   * @param params Edit video parameters.
   * @param options Per-request and polling overrides.
   * @returns The completed edit video response.
   */
  async run(params: EditVideoParams, options?: RequestOptions & PollingOptions): Promise<CompletedEditVideoResponse> {
    const { id } = await this.create(params, options);
    const response = await pollUntilComplete<EditVideoResponse>(() => this.get(id, options), {
      maxWaitMs: options?.maxWaitMs,
      pollIntervalMs: options?.pollIntervalMs,
    });
    return response as CompletedEditVideoResponse;
  }

  /**
   * Create an edit video task; returns immediately with a task id.
   * @param params Edit video parameters.
   * @param options Per-request overrides.
   * @returns The task creation result.
   */
  async create(params: EditVideoParams, options?: RequestOptions): Promise<TaskCreateResponse> {
    const body = compactParams(params);
    validateParams(contract['edit-video'] as ActionSchema, { ...body, model: MODEL } as Record<string, unknown>);
    return this.http.request<TaskCreateResponse>('POST', ENDPOINT, { body, ...options });
  }

  /**
   * Fetch the current status of an edit video task.
   * @param id The task id.
   * @param options Per-request overrides.
   * @returns The current edit video task status.
   */
  async get(id: string, options?: RequestOptions): Promise<EditVideoResponse> {
    return this.http.request<EditVideoResponse>('GET', `${ENDPOINT}/${id}`, options ?? {});
  }
}
