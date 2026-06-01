# frozen_string_literal: true

module RunApi
  module RunwayAleph
    module Resources
      class EditVideo
        include RunApi::Core::ResourceHelpers

        ENDPOINT = "/api/v1/runway_aleph/edit_video"
        RESPONSE_CLASS = Types::TaskCreateResponse
        COMPLETED_RESPONSE_CLASS = Types::CompletedEditVideoResponse

        def initialize(http)
          @http = http
        end

        def run(**params)
          task = create(**params)
          poll_until_complete { get(task.id) }
        end

        def create(**params)
          params = compact_params(params)
          validate_params!(params)
          request(:post, ENDPOINT, body: params)
        end

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
