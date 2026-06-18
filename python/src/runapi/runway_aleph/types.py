"""Runway Aleph enums and response models."""

from __future__ import annotations

from runapi.core import BaseModel, TaskResponse, optional, required

ASPECT_RATIOS = ["16:9", "9:16", "4:3", "3:4", "1:1", "21:9"]


class Video(BaseModel):
    id = optional(str)
    url = required(str)


class Image(BaseModel):
    url = required(str)


class EditVideoResponse(TaskResponse):
    """Runway Aleph video editing task status response."""

    id = required(str)
    status = optional(str, enum=lambda: TaskResponse.Status.ALL)
    videos = optional([lambda: Video])
    images = optional([lambda: Image])
    error = optional(str)


class TaskCreateResponse(EditVideoResponse):
    """Runway Aleph task creation response with an id."""


class CompletedEditVideoResponse(EditVideoResponse):
    """Returned by ``edit_video.run()`` once polling observes completion.

    ``videos`` is required so callers never have to null-check it on success.
    """

    videos = required([lambda: Video])
