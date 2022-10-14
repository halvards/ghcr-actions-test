# ghcr-actions-test

Experimenting with ghcr and actions.

## Run

Download and run the binary for your platform:

```sh
VERSION=v0.1.9

curl -Lo ghcr-actions-test "https://github.com/halvards/ghcr-actions-test/releases/download/${VERSION}/ghcr-actions-test_$(uname -s)_$(uname -m)"

chmod +x ghcr-actions-test

./ghcr-actions-test
```

## Render manifests and apply

Render manifests from the [`manifests`](./manifests) directory, and apply the
resources to your Kubernetes cluster using `kubectl`:

```sh
VERSION=v0.1.9

kubectl apply --kustomize "https://github.com/halvards/ghcr-actions-test.git/manifests?ref=${VERSION}"
```

## Apply pre-rendered manifests

Apply a pre-rendered manifest from a GitHub release:

```sh
VERSION=v0.1.9

kubectl apply --filename "https://github.com/halvards/ghcr-actions-test/releases/download/${VERSION}/ghcr-actions-test_manifest.yaml"
```

## Fetch package

Fetch the manifest package using `kpt`:

```sh
VERSION=v0.1.9

kpt pkg get "https://github.com/halvards/ghcr-actions-test.git/manifests@${VERSION}" manifests
```

## Release

See the [release instructions](docs/release.md).
