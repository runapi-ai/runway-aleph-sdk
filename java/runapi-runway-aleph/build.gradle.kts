plugins {
  `java-library`
  `maven-publish`
}

extra["runapiSlug"] = "runway-aleph"

description = "RunAPI Runway Aleph Java SDK for Runway Aleph workflows."

java {
  withSourcesJar()
  withJavadocJar()
}

dependencies {
  api("ai.runapi:runapi-core:0.2.6")

  testImplementation(platform("org.junit:junit-bom:5.10.3"))
  testImplementation("org.junit.jupiter:junit-jupiter")
}

publishing {
  publications {
    create<MavenPublication>("mavenJava") {
      from(components["java"])
      artifactId = "runapi-runway-aleph"
      pom {
        name = "RunAPI Runway Aleph Java SDK"
        description = "RunAPI Runway Aleph Java SDK for Runway Aleph workflows."
        url = "https://runapi.ai/models/runway-aleph"
        licenses {
          license {
            name = "Apache License, Version 2.0"
            url = "https://www.apache.org/licenses/LICENSE-2.0"
          }
        }
        developers {
          developer {
            id = "runapi"
            name = "RunAPI"
            email = "contact@runapi.ai"
          }
        }
        scm {
          url = "https://github.com/runapi-ai/runway-aleph-sdk"
          connection = "scm:git:https://github.com/runapi-ai/runway-aleph-sdk.git"
          developerConnection = "scm:git:ssh://git@github.com/runapi-ai/runway-aleph-sdk.git"
        }
      }
    }
  }
}
