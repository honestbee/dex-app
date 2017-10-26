apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: REDACTED
    server: https://api.ap-southeast-1a.staging.k8s.honestbee.com
  name: ap-southeast-1a.staging.k8s.honestbee.com
contexts:
- context:
    cluster: ap-southeast-1a.staging.k8s.honestbee.com
    namespace: default
    user: k8s-user
  name: k8s-1.7-staging
current-context: k8s-1.7-staging
kind: Config
preferences: {}
users:
- name: k8s-user
  user:
    auth-provider:
      config:
        client-id: kubernetes
        client-secret: AtMUIzMy00ODg0LTkwMDQtME
        extra-scopes: groups
        id-token: {{ .IDToken }}
        idp-issuer-url: https://dex.honestbee.com
        refresh-token: {{ .RefreshToken }}
      name: oidc
