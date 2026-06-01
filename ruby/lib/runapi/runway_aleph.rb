# frozen_string_literal: true

require "runapi/core"
require_relative "runway_aleph/types"
require_relative "runway_aleph/resources/edit_video"
require_relative "runway_aleph/client"

module RunApi
  module RunwayAleph
    AuthenticationError = RunApi::Core::AuthenticationError
    RateLimitError = RunApi::Core::RateLimitError
    InsufficientCreditsError = RunApi::Core::InsufficientCreditsError
    NotFoundError = RunApi::Core::NotFoundError
    ValidationError = RunApi::Core::ValidationError
    TaskFailedError = RunApi::Core::TaskFailedError
    TaskTimeoutError = RunApi::Core::TaskTimeoutError
  end
end
