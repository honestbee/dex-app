# This is the CoreOS Dex Example app

[![Build Status](https://drone.honestbee.com/api/badges/honestbee/dex-app/status.svg "Drone build status")](https://drone.honestbee.com/honestbee/dex-app)
[![Docker Repository on Quay](https://img.shields.io/badge/container-ready-brightgreen.svg "Docker Repository on Quay")](https://quay.io/repository/honestbee/dex-app?tab=tags)

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

|            ENV             |            Description             | Default |
| -------------------------- | ---------------------------------- | ------- |
| `STAGING_CA_CERT`          | CA certificate for staging cluster | ``      |
| `STAGING_CLUSTER_ENDPOINT` | staging cluster endpoint           | ``      |
| `SVC_CA_CERT`              | CA certificate for svc cluster     | ``      |
| `SVC_CLUSTER_ENDPOINT`     | svc cluster endpoint               | ``      |
| `SCOPES`                   | Scopes to fetch from Id provider   | ``      |
