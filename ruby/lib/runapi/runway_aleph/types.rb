# frozen_string_literal: true

module RunApi
  module RunwayAleph
    # Type definitions and constants for Runway Aleph video editing.
    module Types
      # A generated output video.
      class Video < RunApi::Core::BaseModel
        optional :id, String
        required :url, String
      end

      # A reference or extracted image.
      class Image < RunApi::Core::BaseModel
        required :url, String
      end

      # Task status response for a video editing operation.
      # Includes output videos and images when the task completes.
      class EditVideoResponse < RunApi::Core::TaskResponse
        required :id, String
        optional :status, String, enum: -> { RunApi::Core::TaskResponse::Status::ALL }
        optional :videos, [-> { Video }]
        optional :images, [-> { Image }]
        optional :error, String
      end

      # Initial response when a video editing task is created.
      class TaskCreateResponse < EditVideoResponse; end

      # Completed video editing response with guaranteed output videos.
      class CompletedEditVideoResponse < EditVideoResponse
        required :videos, [-> { Video }]
      end
    end
  end
end
