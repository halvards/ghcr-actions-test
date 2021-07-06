# ghcr-actions-test

Experimenting with ghcr and actions.

## Run

Download and run the binary for your platform:

```sh
VERSION=v0.0.11

curl -Lo ghcr-actions-test "https://github.com/halvards/ghcr-actions-test/releases/download/$VERSION/ghcr-actions-test_$(uname -s)_$(uname -m)"

chmod +x ghcr-actions-test

./ghcr-actions-test
```

## Deploy

Deploy the resources to your Kubernetes cluster:

```sh
VERSION=v0.0.11

kpt pkg get https://github.com/halvards/ghcr-actions-test.git/manifests@$VERSION manifests

kpt live init manifests

kpt live apply manifests --reconcile-timeout=3m --output=table
```

## Release

See the [release instructions](docs/release.md).
