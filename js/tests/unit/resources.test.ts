import { beforeEach, describe, expect, it, vi } from 'vitest';
import type { HttpClient } from '@runapi.ai/core';
import { VideoToVideo } from '../../src/resources/video-to-video';

describe('Runway Aleph resources', () => {
  const mockHttp: HttpClient = { request: vi.fn() };

  beforeEach(() => {
    vi.clearAllMocks();
  });

  it('creates video-to-video task with video_url', async () => {
    vi.mocked(mockHttp.request).mockResolvedValueOnce({ id: 'task-3' });
    const videoToVideo = new VideoToVideo(mockHttp);

    await videoToVideo.create({ prompt: 'Regrade to dusk', video_url: 'https://example.com/source.mp4', upload_cn: false });

    expect(mockHttp.request).toHaveBeenCalledWith('POST', '/api/v1/runway_aleph/video_to_video', {
      body: { prompt: 'Regrade to dusk', video_url: 'https://example.com/source.mp4', upload_cn: false },
    });
  });

  it('gets video-to-video task by id', async () => {
    vi.mocked(mockHttp.request).mockResolvedValueOnce({ id: 'task-3', status: 'completed', videos: [{ url: 'https://file.runapi.ai/video.mp4' }] });
    const videoToVideo = new VideoToVideo(mockHttp);

    const result = await videoToVideo.get('task-3');

    expect(mockHttp.request).toHaveBeenCalledWith('GET', '/api/v1/runway_aleph/video_to_video/task-3', {});
    expect(result.videos?.[0]?.url).toBe('https://file.runapi.ai/video.mp4');
  });
});
