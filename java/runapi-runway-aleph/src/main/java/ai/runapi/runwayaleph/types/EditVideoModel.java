package ai.runapi.runwayaleph.types;

import com.fasterxml.jackson.annotation.JsonCreator;

/** Model slug for edit video operations. */
public final class EditVideoModel extends RunwayalephValue {
  /** runway-aleph model slug. */
  public static final EditVideoModel RUNWAY_ALEPH = new EditVideoModel("runway-aleph");

  /** Creates a model value from a literal model slug. */
  @JsonCreator
  public EditVideoModel(String value) {
    super(value);
  }
}
