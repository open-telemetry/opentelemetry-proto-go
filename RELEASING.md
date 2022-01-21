# Release Process

This project ultimately would like to automate the generation and releasing of
the generated code it contains. Until the automation tooling is built to do
this, this document serves as an overview of how to make a new release. It
assumes the release is going to be made because upstream protobuf definitions
have been updated and the generated code needs to be updated.

## Create a Release Pull Request

Update the submodule and regenerate the code.

```sh
export VERSION=<new-version>
git config -f .gitmodules submodule.opentelemetry-proto.branch $VERSION
cd opentelemetry-proto
git checkout $VERSION
cd ..
make clean protobuf
```

Verify the changes.

```sh
git diff main
```

If everything looks good, push the changes to GitHub and open a pull request.

- Title: `Release {{VERSION}}`
- Body:

   ```markdown
   Release of the [{{VERSION}}][otlp] version of the OTLP.

   [otlp]: https://github.com/open-telemetry/opentelemetry-proto/releases/tag/{{VERSION}}
   ```

## Tag Release

Once the pull request with all the generated code changes has been approved
and merged tag the merged commit.

***IMPORTANT***: It is critical you use the same tag as the
opentelemetry-proto submodule. Failure to do so will leave things in a broken
state.

***IMPORTANT***: [There is currently no way to remove an incorrectly tagged
version of a Go module](https://github.com/golang/go/issues/34189). It is
critical you make sure the version you push upstream is correct. [Failure to
do so will lead to minor emergencies and tough to work
around](https://github.com/open-telemetry/opentelemetry-go/issues/331).

Run the tag.sh script using the `<commit-hash>` of the commit on the main
branch for the merged Pull Request.

```sh
./tag.sh $VERSION <commit-hash>
```

Push all generated tags to the upstream remote (not your fork!).

```sh
git push upstream $VERSION
git push upstream otlp/$VERSION
```

## Release

Create a GitHub release for the new `<new tag>` on GitHub.

- Title: `Release {{VERSION}}`
- Body:

   ```markdown
   Generated Go code for the [{{VERSION}}][otlp] version of the OTLP.

   [otlp]: https://github.com/open-telemetry/opentelemetry-proto/releases/tag/{{VERSION}}
   ```
