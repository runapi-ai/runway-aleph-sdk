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
        MODEL = "runway-aleph"

        def initialize(http)
          @http = http
        end

        # Transform a video and wait until complete.
        #
        # @param params [Hash] edit parameters
        # @return [RunApi::RunwayAleph::Types::CompletedEditVideoResponse] completed task with videos
        def run(options: nil, **params)
          task = create(options: options, **params)
          poll_until_complete { get(task.id, options: options) }
        end

        # Start a video editing task.
        #
        # @param params [Hash] edit parameters
        # @return [RunApi::RunwayAleph::Types::TaskCreateResponse] task creation result with id
        def create(options: nil, **params)
          params = compact_params(params)
          validate_contract!(CONTRACT["edit-video"], params.merge(model: MODEL))
          request(:post, ENDPOINT, body: params, options: options)
        end

        # Get video editing task status by task ID.
        #
        # @param id [String] task ID
        # @return [RunApi::RunwayAleph::Types::EditVideoResponse] current task status
        def get(id, options: nil)
          request(:get, "#{ENDPOINT}/#{id}", options: options)
        end
      end
    end
  end
end
