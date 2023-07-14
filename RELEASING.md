# Release Process

This project ultimately would like to automate the generation and releasing of
the generated code it contains. Until the automation tooling is built to do
this, this document serves as an overview of how to make a new release. It
assumes the release is going to be made because upstream protobuf definitions
have been updated and the generated code needs to be updated.

## Create a Release Pull Request

1. Update the [`opentelemetry-proto`] submodule and regenerate the code.

   ```sh
   make sync VERSION=<new-version>
   ```

2. Edit [`versions.yaml`] with the new version number.  Then, ensure the correct modules versions
   are updated and the [`versions.yaml`] syntax is correct.

   ```sh
   make verify-versions
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

```sh
make push-tags REMOTE=<upstream> COMMIT=<hash>
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
[`opentelemetry-proto`]: https://github.com/open-telemetry/opentelemetry-proto
