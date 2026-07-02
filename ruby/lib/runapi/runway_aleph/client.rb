# frozen_string_literal: true

module RunApi
  module RunwayAleph
    # Runway Aleph prompt-driven video editing API client.
    # Unlike generation from scratch, Runway Aleph transforms an existing video
    # using a text prompt, with optional style reference images.
    #
    # @example
    #   client = RunApi::RunwayAleph::Client.new(api_key: "your-api-key")
    #   result = client.edit_video.run(
    #     prompt: "Make it look like a watercolor painting",
    #     source_video_url: "https://cdn.runapi.ai/public/samples/video.mp4"
    #   )
    class Client < RunApi::Core::Client
      # @return [Resources::EditVideo] Prompt-driven video editing with optional style reference.
      attr_reader :edit_video

      def initialize(api_key: nil, **options)
        super
        @edit_video = Resources::EditVideo.new(http)
      end
    end
  end
end
