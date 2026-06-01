import { beforeEach, describe, expect, it, vi } from 'vitest';
import type { HttpClient } from '@runapi.ai/core';
import { EditVideo } from '../../src/resources/edit-video';

describe('Runway Aleph resources', () => {
  const mockHttp: HttpClient = { request: vi.fn() };

  beforeEach(() => {
    vi.clearAllMocks();
  });

  it('creates edit-video task with source_video_url', async () => {
    vi.mocked(mockHttp.request).mockResolvedValueOnce({ id: 'task-3' });
    const editVideo = new EditVideo(mockHttp);

    await editVideo.create({ prompt: 'Regrade to dusk', source_video_url: 'https://cdn.runapi.ai/public/samples/source.mp4' });

    expect(mockHttp.request).toHaveBeenCalledWith('POST', '/api/v1/runway_aleph/edit_video', {
      body: { prompt: 'Regrade to dusk', source_video_url: 'https://cdn.runapi.ai/public/samples/source.mp4' },
    });
  });

  it('gets edit-video task by id', async () => {
    vi.mocked(mockHttp.request).mockResolvedValueOnce({
      id: 'task-3',
      status: 'completed',
      videos: [{ url: 'https://file.runapi.ai/video.mp4' }],
      images: [{ url: 'https://file.runapi.ai/cover.png' }],
    });
    const editVideo = new EditVideo(mockHttp);

    const result = await editVideo.get('task-3');

    expect(mockHttp.request).toHaveBeenCalledWith('GET', '/api/v1/runway_aleph/edit_video/task-3', {});
    expect(result.videos?.[0]?.url).toBe('https://file.runapi.ai/video.mp4');
    expect(result.images?.[0]?.url).toBe('https://file.runapi.ai/cover.png');
  });
});
