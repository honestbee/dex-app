apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: {{ .CACert }}
    server: https://api.{{ .ClusterEndpoint }}
  name: {{ .ClusterEndpoint }}
contexts:
- context:
    cluster: {{ .ClusterEndpoint }}
    namespace: default
    user: k8s-user
  name: k8s-1.7-{{ .ClusterEndpoint }}
current-context: k8s-1.7-{{ .ClusterEndpoint }}
kind: Config
preferences: {}
users:
- name: k8s-user
  user:
    auth-provider:
      config:
        client-id: {{ .ClientID }}
        extra-scopes: groups
        id-token: {{ .IDToken }}
        idp-issuer-url: https://dex.honestbee.com
        refresh-token: {{ .RefreshToken }}
      name: oidc
