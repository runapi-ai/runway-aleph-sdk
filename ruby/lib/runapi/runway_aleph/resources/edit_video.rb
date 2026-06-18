# frozen_string_literal: true

module RunApi
  module RunwayAleph
    module Resources
      # Runway Aleph video editing resource.
      # Transform an existing video using a text prompt and optional style reference image.
      class EditVideo
        include RunApi::Core::ResourceHelpers

        ENDPOINT = "/api/v1/runway_aleph/edit_video"
        RESPONSE_CLASS = Types::TaskCreateResponse
        COMPLETED_RESPONSE_CLASS = Types::CompletedEditVideoResponse

        def initialize(http)
          @http = http
        end

        # Transform a video and wait until complete.
        #
        # @param params [Hash] edit parameters
        # @return [RunApi::RunwayAleph::Types::CompletedEditVideoResponse] completed task with videos
        def run(**params)
          task = create(**params)
          poll_until_complete { get(task.id) }
        end

        # Start a video editing task.
        #
        # @param params [Hash] edit parameters
        # @return [RunApi::RunwayAleph::Types::TaskCreateResponse] task creation result with id
        def create(**params)
          params = compact_params(params)
          validate_params!(params)
          request(:post, ENDPOINT, body: params)
        end

        # Get video editing task status by task ID.
        #
        # @param id [String] task ID
        # @return [RunApi::RunwayAleph::Types::EditVideoResponse] current task status
        def get(id)
          request(:get, "#{ENDPOINT}/#{id}")
        end

        private

        def validate_params!(params)
          raise Core::ValidationError, "prompt is required" unless param(params, :prompt)
          raise Core::ValidationError, "source_video_url is required" unless param(params, :source_video_url)
          validate_optional!(params, :aspect_ratio, Types::ASPECT_RATIOS)
        end
      end
    end
  end
end
