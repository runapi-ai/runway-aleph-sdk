import pytest

from runapi.core import config
from runapi.core.errors import AuthenticationError, ValidationError
from runapi.runway_aleph import RunwayAlephClient
from runapi.runway_aleph.resources.edit_video import EditVideo
from runapi.runway_aleph.types import CompletedEditVideoResponse, EditVideoResponse


class FakeHttp:
    """Records (method, path, body) and replays preset responses by call order."""

    def __init__(self, *responses):
        self._responses = list(responses)
        self.calls = []

    def request(self, method, path, body=None, options=None):
        self.calls.append((method, path, body))
        if self._responses:
            return self._responses.pop(0)
        return {"id": "task_1", "status": "pending"}


@pytest.fixture(autouse=True)
def reset_config(monkeypatch):
    monkeypatch.delenv("RUNAPI_API_KEY", raising=False)
    monkeypatch.setattr(config, "api_key", None)
    yield


# --- authentication -------------------------------------------------------


def test_accepts_api_key_parameter():
    assert isinstance(
        RunwayAlephClient(api_key="param-key", http_client=FakeHttp()), RunwayAlephClient
    )


def test_falls_back_to_global(monkeypatch):
    monkeypatch.setattr(config, "api_key", "global-key")
    assert isinstance(RunwayAlephClient(http_client=FakeHttp()), RunwayAlephClient)


def test_falls_back_to_env(monkeypatch):
    monkeypatch.setenv("RUNAPI_API_KEY", "env-key")
    assert isinstance(RunwayAlephClient(http_client=FakeHttp()), RunwayAlephClient)


def test_raises_without_api_key():
    with pytest.raises(AuthenticationError, match="API key is required"):
        RunwayAlephClient()


# --- transport injection / accessors --------------------------------------


def test_uses_injected_http_client():
    fake = FakeHttp()
    client = RunwayAlephClient(api_key="k", http_client=fake)
    assert client.edit_video._http is fake


def test_exposes_resource_accessors():
    client = RunwayAlephClient(api_key="k", http_client=FakeHttp())
    assert isinstance(client.edit_video, EditVideo)


# --- request shapes -------------------------------------------------------


def test_create_posts_compacted_body():
    fake = FakeHttp({"id": "t1", "status": "pending"})
    client = RunwayAlephClient(api_key="k", http_client=fake)
    result = client.edit_video.create(
        prompt="hello",
        source_video_url="https://example.com/v.mp4",
        aspect_ratio="16:9",
        seed=None,
    )
    assert fake.calls == [
        (
            "post",
            "/api/v1/runway_aleph/edit_video",
            {
                "prompt": "hello",
                "source_video_url": "https://example.com/v.mp4",
                "aspect_ratio": "16:9",
            },
        ),
    ]
    _, _, body = fake.calls[0]
    assert "model" not in body
    assert isinstance(result, EditVideoResponse)
    assert result.id == "t1"


def test_get_fetches_by_id():
    fake = FakeHttp({"id": "t1", "status": "processing"})
    client = RunwayAlephClient(api_key="k", http_client=fake)
    client.edit_video.get("t1")
    assert fake.calls == [("get", "/api/v1/runway_aleph/edit_video/t1", None)]


def test_run_polls_and_narrows_completed_type():
    fake = FakeHttp(
        {"id": "t1", "status": "pending"},
        {"id": "t1", "status": "completed", "videos": [{"url": "https://x/y.mp4"}]},
    )
    client = RunwayAlephClient(api_key="k", http_client=fake)
    result = client.edit_video.run(
        model="runway-aleph",
        prompt="hi",
        source_video_url="https://example.com/v.mp4",
    )

    assert isinstance(result, CompletedEditVideoResponse)
    assert result.videos[0].url == "https://x/y.mp4"
    assert [call[0] for call in fake.calls] == ["post", "get"]


# --- validation -----------------------------------------------------------


def test_create_requires_prompt():
    client = RunwayAlephClient(api_key="k", http_client=FakeHttp())
    with pytest.raises(ValidationError, match="prompt is required"):
        client.edit_video.create(
            model="runway-aleph", source_video_url="https://example.com/v.mp4"
        )


def test_create_requires_source_video_url():
    client = RunwayAlephClient(api_key="k", http_client=FakeHttp())
    with pytest.raises(ValidationError, match="source_video_url is required"):
        client.edit_video.create(model="runway-aleph", prompt="hi")


def test_create_rejects_invalid_aspect_ratio():
    client = RunwayAlephClient(api_key="k", http_client=FakeHttp())
    with pytest.raises(ValidationError, match="aspect_ratio must be one of"):
        client.edit_video.create(
            model="runway-aleph",
            prompt="hi",
            source_video_url="https://example.com/v.mp4",
            aspect_ratio="99:1",
        )


def test_create_accepts_valid_aspect_ratio():
    fake = FakeHttp({"id": "t1", "status": "pending"})
    client = RunwayAlephClient(api_key="k", http_client=fake)
    client.edit_video.create(
        model="runway-aleph",
        prompt="hi",
        source_video_url="https://example.com/v.mp4",
        aspect_ratio="21:9",
    )
    _, path, body = fake.calls[0]
    assert path == "/api/v1/runway_aleph/edit_video"
    assert body["aspect_ratio"] == "21:9"
