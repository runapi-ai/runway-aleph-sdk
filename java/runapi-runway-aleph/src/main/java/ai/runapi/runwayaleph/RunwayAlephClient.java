package ai.runapi.runwayaleph;

import ai.runapi.core.BaseClient;
import ai.runapi.core.ClientOptions;
import ai.runapi.core.http.HttpTransport;
import java.net.URI;
import ai.runapi.runwayaleph.resources.EditVideoResource;

/** RunwayAleph model-family Java SDK client. */
public final class RunwayAlephClient extends BaseClient {
  private final EditVideoResource editVideo;

  private RunwayAlephClient(ClientOptions options) {
    super(options);
    this.editVideo = new EditVideoResource(transport(), options());
  }

  /** Creates a new RunwayAlephClient builder. */
  public static Builder builder() {
    return new Builder();
  }

  /** Edit Video operations. */
  public EditVideoResource editVideo() {
    return editVideo;
  }

  /** Builder for {@link RunwayAlephClient}. */
  public static final class Builder extends BaseClient.Builder<Builder> {
    private Builder() {}

    /** Sets the API key. If omitted, the SDK reads {@code RUNAPI_API_KEY}. */
    @Override
    public Builder apiKey(String value) {
      return super.apiKey(value);
    }

    /** Sets the RunAPI base URL. If omitted, the SDK reads {@code RUNAPI_BASE_URL}. */
    @Override
    public Builder baseUrl(String value) {
      return super.baseUrl(value);
    }

    /** Sets the RunAPI base URL from a URI. */
    @Override
    public Builder baseUrl(URI value) {
      return super.baseUrl(value);
    }

    /** Sets a custom HTTP transport. User-provided transports are not closed by SDK clients. */
    @Override
    public Builder transport(HttpTransport value) {
      return super.transport(value);
    }

    /** Builds an immutable RunwayAlephClient. */
    @Override
    public RunwayAlephClient build() {
      return new RunwayAlephClient(options.build());
    }
  }
}
