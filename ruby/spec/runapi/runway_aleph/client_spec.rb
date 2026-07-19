# frozen_string_literal: true

require "spec_helper"

RSpec.describe RunApi::RunwayAleph::Client do
  after { RunApi.api_key = nil }

  it "accepts api_key as parameter" do
    expect(described_class.new(api_key: "param-key")).to be_a(described_class)
  end

  it "falls back to global RunApi.api_key" do
    RunApi.api_key = "global-key"
    expect(described_class.new).to be_a(described_class)
  end

  it "raises AuthenticationError without api_key" do
    expect { described_class.new }.to raise_error(RunApi::Core::AuthenticationError, /API key is required/)
  end

  it "exposes edit_video accessor" do
    client = described_class.new(api_key: "test-key")
    expect(client.edit_video).to be_a(RunApi::RunwayAleph::Resources::EditVideo)
  end
end
