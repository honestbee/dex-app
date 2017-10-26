# This is the CoreOS Dex Example app

[![Docker Repository on Quay](https://quay.io/repository/vincentdesmet/dex-app/status "Docker Repository on Quay")](https://quay.io/repository/vincentdesmet/dex-app)

[Ref](https://github.com/coreos/dex/blob/master/Documentation/using-dex.md#writing-apps-that-use-dex)

Ensure to add the following to Dex config before deploying Example app

```
staticClients:
- id: example-app
  secret: ZXhhbXBsZS1hcHAtc2VjcmV0
  name: 'Example App'
  # Where the app will be running.
  redirectURIs:
  - 'http://127.0.0.1:5555/callback'
```

## Usage:
```
./dex-app --issuer https://dex.honestbee.com --client-id example-app --client-secret ZXhhbXBsZS1hcHAtc2VjcmV0
```
