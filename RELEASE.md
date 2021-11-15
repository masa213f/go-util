# Release procedure

## Versioning

Follow [semantic versioning 2.0.0][semver] to choose the new version number.

## Bump version

1. Determine a new version number.
    ```bash
    VERSION=x.y.z
    ```

2. Edit `CHANGELOG.md` for the new version.

3. Commit the change and push it.
    ```bash
    git checkout main
    git pull
    git commit -a -m "Bump version to $VERSION"
    git push origin
    ```

4. Add a git tag and push it.
    ```bash
    git checkout main
    git pull
    git tag -a -m "Release v$VERSION" "v$VERSION"
    git push origin "v$VERSION"
    ```

[semver]: https://semver.org/spec/v2.0.0.html