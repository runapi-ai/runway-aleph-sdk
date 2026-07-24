"""Runway Aleph client."""

from __future__ import annotations

from typing import Any, Optional

from runapi.core import ProviderClient

from .resources.edit_video import EditVideo


class RunwayAlephClient(ProviderClient):
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
        super().__init__(api_key, **options)
        http = self._http
        self.edit_video = EditVideo(http)
