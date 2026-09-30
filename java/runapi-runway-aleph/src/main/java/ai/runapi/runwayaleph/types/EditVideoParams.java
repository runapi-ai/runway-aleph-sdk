package ai.runapi.runwayaleph.types;

import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

/** Parameters for edit video operations. */
public final class EditVideoParams {
  private final String model;
  private final String prompt;
  private final String sourceVideoUrl;
  private final String callbackUrl;
  private final Boolean watermark;
  private final String aspectRatio;
  private final Integer seed;
  private final String referenceImageUrl;

  private EditVideoParams(Builder builder) {
    this.model = builder.model;
    this.prompt = builder.prompt;
    this.sourceVideoUrl = builder.sourceVideoUrl;
    this.callbackUrl = builder.callbackUrl;
    this.watermark = builder.watermark;
    this.aspectRatio = builder.aspectRatio;
    this.seed = builder.seed;
    this.referenceImageUrl = builder.referenceImageUrl;
  }

  /** Creates a new EditVideoParams builder. */
  public static Builder builder() {
    return new Builder();
  }

  /** Returns the RunAPI action key for this request. */
  public String action() {
    return "runway-aleph/edit-video";
  }

  /** Converts these parameters to the JSON request body shape. */
  public Map<String, Object> toMap() {
    Map<String, Object> raw = new LinkedHashMap<String, Object>();
    raw.put("model", RunwayalephParamUtils.wireValue(model));
    raw.put("prompt", RunwayalephParamUtils.wireValue(prompt));
    raw.put("source_video_url", RunwayalephParamUtils.wireValue(sourceVideoUrl));
    raw.put("callback_url", RunwayalephParamUtils.wireValue(callbackUrl));
    raw.put("watermark", RunwayalephParamUtils.wireValue(watermark));
    raw.put("aspect_ratio", RunwayalephParamUtils.wireValue(aspectRatio));
    raw.put("seed", RunwayalephParamUtils.wireValue(seed));
    raw.put("reference_image_url", RunwayalephParamUtils.wireValue(referenceImageUrl));
    return RunwayalephParamUtils.compact(raw);
  }



  /** Builder for {@link EditVideoParams}. */
  public static final class Builder {
    private String model;
    private String prompt;
    private String sourceVideoUrl;
    private String callbackUrl;
    private Boolean watermark;
    private String aspectRatio;
    private Integer seed;
    private String referenceImageUrl;

    private Builder() {}

    /** Sets the model slug using a typed model value. */
    public Builder model(EditVideoModel value) {
      this.model = java.util.Objects.requireNonNull(value, "model").value();
      return this;
    }

    /** Sets the model slug using a string value. */
    public Builder model(String value) {
      this.model = value;
      return this;
    }


    /** Sets the text prompt. */
    public Builder prompt(String value) {
      this.prompt = value;
      return this;
    }

    /** Sets the source video URL. */
    public Builder sourceVideoUrl(String value) {
      this.sourceVideoUrl = value;
      return this;
    }

    /** Sets the webhook URL for task completion notifications. */
    public Builder callbackUrl(String value) {
      this.callbackUrl = value;
      return this;
    }

    /** Sets the watermark toggle. */
    public Builder watermark(boolean value) {
      this.watermark = value;
      return this;
    }

    /** Sets the output aspect ratio. */
    public Builder aspectRatio(String value) {
      this.aspectRatio = value;
      return this;
    }

    /** Sets the random seed. */
    public Builder seed(int value) {
      this.seed = value;
      return this;
    }

    /** Sets the reference image URL. */
    public Builder referenceImageUrl(String value) {
      this.referenceImageUrl = value;
      return this;
    }

    /** Builds immutable edit video parameters. */
    public EditVideoParams build() {
      return new EditVideoParams(this);
    }
  }
}
