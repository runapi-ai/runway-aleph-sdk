"""Runway Aleph client."""

from __future__ import annotations

from typing import Any, Optional

from runapi.core import ClientOptions, HttpClient, resolve_api_key

from .resources.edit_video import EditVideo


class RunwayAlephClient:
    """Runway Aleph prompt-guided video editing client.

    Example::

        client = RunwayAlephClient(api_key="sk-...")
        result = client.edit_video.run(
            model="runway-aleph",
            prompt="Transform the scene into a watercolor painting style",
            source_video_url="https://cdn.runapi.ai/public/samples/video.mp4",
        )
    """

    def __init__(self, api_key: Optional[str] = None, **options: Any) -> None:
        resolved_api_key = resolve_api_key(api_key)
        client_options = ClientOptions(api_key=resolved_api_key, **options)
        http = client_options.http_client or HttpClient(client_options)
        self.edit_video = EditVideo(http)
