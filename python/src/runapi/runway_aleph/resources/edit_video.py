"""Runway Aleph edit-video resource."""

from __future__ import annotations

from typing import Any, Optional

from runapi.core import Resource, RequestOptions

from ..contract_gen import CONTRACT
from ..types import (
    CompletedEditVideoResponse,
    TaskCreateResponse,
)


class EditVideo(Resource):
    """Prompt-guided video editing with Runway Aleph."""

    ENDPOINT = "/api/v1/runway_aleph/edit_video"

    RESPONSE_CLASS = TaskCreateResponse
    COMPLETED_RESPONSE_CLASS = CompletedEditVideoResponse

    MODEL = "runway-aleph"

    def run(self, options: Optional[RequestOptions] = None, **params: Any) -> Any:
        """Create a task and poll until it completes."""
        task = self.create(options=options, **params)
        return self._poll_until_complete(lambda: self.get(task.id, options=options))

    def create(self, options: Optional[RequestOptions] = None, **params: Any) -> Any:
        """Create an edit-video task and return immediately with an ``id``."""
        compacted = self._compact_params(params)
        self._validate_contract(CONTRACT["edit-video"], {**compacted, "model": self.MODEL})
        return self._request("post", self.ENDPOINT, body=compacted, options=options)

    def get(self, id: str, options: Optional[RequestOptions] = None) -> Any:
        """Fetch the current status of an edit-video task."""
        return self._request("get", f"{self.ENDPOINT}/{id}", options=options)
