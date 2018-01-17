apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: {{ .CACert }}
    server: https://api.{{ .ClusterEndpoint }}
  name: {{ .ClusterEndpoint }}
contexts:
- context:
    cluster: {{ .ClusterEndpoint }}
    namespace: {{ .Namespace }}
    user: {{ .ClientID }}-user
  name: {{ .ClientID }}-dex
current-context: {{ .ClientID }}-dex
kind: Config
preferences: {}
users:
- name: {{ .ClientID }}-user
  user:
    auth-provider:
      config:
        client-id: {{ .ClientID }}
        extra-scopes: groups
        id-token: {{ .IDToken }}
        idp-issuer-url: https://dex.honestbee.com
        refresh-token: {{ .RefreshToken }}
      name: oidc
