# Release Process

This project ultimately would like to automate the generation and releasing of
the generated code it contains. Until the automation tooling is built to do
this, this document serves as an overview of how to make a new release. It
assumes the release is going to be made because upstream protobuf definitions
have been updated and the generated code needs to be updated.

## Create a Release Pull Request

1. Update the submodule and regenerate the code.

   ```sh
   make sync VERSION=<new-version>
   ```

2. Update [`versions.yaml`]. Ensure the correct modules versions are updated
   and the [`versions.yaml`] syntax is correct.

   ```sh
   make verify
   ```

3. Verify the changes.

   ```sh
   git diff main
   ```

4. If everything looks good, push the changes to GitHub and open a pull request.

   - Title: `Release {{VERSION}}`
   - Body:

      ```markdown
      Release of the [{{VERSION}}][otlp] version of the OTLP.

      [otlp]: https://github.com/open-telemetry/opentelemetry-proto/releases/tag/{{VERSION}}
      ```

## Tag Release

Once the pull request with all the generated code changes has been approved
and merged use the [`multimod`] utility to tag all modules according to
[`versions.yaml`].

1. For each module set that will be released, run the `add-tags` make target
   using the `<commit-hash>` of the commit on the main branch for the merged
   Pull Request.

   ```sh
   make add-tags MODSET=<module set> COMMIT=<commit hash>
   ```

   It should only be necessary to provide an explicit `COMMIT` value if the
   current `HEAD` of your working directory is not the correct commit.

2. Push tags to the upstream remote (not your fork:
   `github.com/open-telemetry/opentelemetry-go-contrib.git`). Make sure you
   push all sub-modules as well.

   ```sh
   export VERSION="<version>"
   for t in $( git tag -l | grep "$VERSION" ); do git push upstream "$t"; done
   ```

## Release

Create a GitHub release for the new `<new tag>` on GitHub.

- Title: `Release {{VERSION}}`
- Body:

   ```markdown
   Generated Go code for the [{{VERSION}}][otlp] version of the OTLP.

   [otlp]: https://github.com/open-telemetry/opentelemetry-proto/releases/tag/{{VERSION}}
   ```

[`versions.yaml`]: ./versions.yaml
[`multimod`]: https://pkg.go.dev/go.opentelemetry.io/build-tools/multimod
