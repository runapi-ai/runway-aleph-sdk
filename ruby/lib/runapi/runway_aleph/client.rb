# frozen_string_literal: true

module RunApi
  module RunwayAleph
    class Client
      attr_reader :edit_video

      def initialize(api_key: nil, **options)
        @api_key = Core::Auth.resolve_api_key(api_key)

        client_options = Core::ClientOptions.new(api_key: @api_key, **options)
        http = client_options.http_client || Core::HttpClient.new(client_options)
        @edit_video = Resources::EditVideo.new(http)
      end
    end
  end
end
