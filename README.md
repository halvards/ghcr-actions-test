# ghcr-actions-test

Experimenting with ghcr and actions.

## Run

Download and run the binary for your platform:

```sh
VERSION=v0.1.4

curl -Lo ghcr-actions-test "https://github.com/halvards/ghcr-actions-test/releases/download/$VERSION/ghcr-actions-test_$(uname -s)_$(uname -m)"

chmod +x ghcr-actions-test

./ghcr-actions-test
```

## Deploy using `kubectl`

Deploy the resources to your Kubernetes cluster using `kubectl`:

```sh
VERSION=v0.1.4

kubectl apply --kustomize https://github.com/halvards/ghcr-actions-test.git/manifests?ref=$VERSION
```

## Release

See the [release instructions](docs/release.md).
