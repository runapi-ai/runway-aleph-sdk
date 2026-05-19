# frozen_string_literal: true

module RunApi
  module RunwayAleph
    module Types
      ASPECT_RATIOS = %w[16:9 9:16 4:3 3:4 1:1 21:9].freeze

      class Video < RunApi::Core::BaseModel
        optional :id, String
        required :url, String
      end

      class VideoToVideoResponse < RunApi::Core::TaskResponse
        required :id, String
        optional :status, String, enum: -> { RunApi::Core::TaskResponse::Status::ALL }
        optional :videos, [ -> { Video } ]
        optional :image_url, String
        optional :parent_task_id, String
        optional :error, String
      end

      class TaskCreateResponse < VideoToVideoResponse; end

      class CompletedVideoToVideoResponse < VideoToVideoResponse
        required :videos, [ -> { Video } ]
      end
    end
  end
end
