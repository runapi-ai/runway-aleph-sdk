# frozen_string_literal: true

require "spec_helper"

RSpec.describe RunApi::RunwayAleph::Resources::EditVideo do
  let(:http) { instance_double(RunApi::Core::HttpClient) }
  let(:resource) { described_class.new(http) }
  let(:endpoint) { "/api/v1/runway_aleph/edit_video" }

  it "POSTs to the correct endpoint without a model param" do
    params = {prompt: "Relight to dusk", source_video_url: "https://cdn.runapi.ai/public/samples/source.mp4"}
    expect(http).to receive(:request).with(:post, endpoint, body: params).and_return("id" => "task-3")
    result = resource.create(**params)
    expect(result.id).to eq("task-3")
  end

  it "GETs the correct endpoint" do
    expect(http).to receive(:request).with(:get, "#{endpoint}/task-3").and_return(
      "id" => "task-3",
      "status" => "completed",
      "videos" => [{"url" => "https://file.runapi.ai/video.mp4"}],
      "images" => [{"url" => "https://file.runapi.ai/cover.png"}]
    )
    result = resource.get("task-3")
    expect(result.status).to eq("completed")
    expect(result.videos.first.url).to eq("https://file.runapi.ai/video.mp4")
    expect(result.images.first.url).to eq("https://file.runapi.ai/cover.png")
  end
end
