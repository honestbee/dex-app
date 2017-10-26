package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type tokenTmplData struct {
	IDToken      string
	RefreshToken string
	RedirectURL  string
	Claims       string
}

var indexTmpl = template.Must(template.New("index.html").Parse(`<html>
  <body>
    <form action="/login" method="post">
       <p>
         Authenticate for:<input type="text" name="cross_client" placeholder="list of client-ids">
       </p>
       <p>
         Extra scopes:<input type="text" name="extra_scopes" placeholder="list of scopes" value="groups">
       </p>
	   <p>
	     Request offline access:<input type="checkbox" name="offline_access" value="yes" checked>
       </p>
       <input type="submit" value="Login">
    </form>
  </body>
</html>`))

var tokenTmpl = template.Must(template.New("token.html").Parse(`<html>
  <head>
    <style>
/* make pre wrap */
pre {
 white-space: pre-wrap;       /* css-3 */
 white-space: -moz-pre-wrap;  /* Mozilla, since 1999 */
 white-space: -pre-wrap;      /* Opera 4-6 */
 white-space: -o-pre-wrap;    /* Opera 7 */
 word-wrap: break-word;       /* Internet Explorer 5.5+ */
}
    </style>
  </head>
	<body>
	<h5> Copy & paste the following commands: </h5>
	<p> kubectl config set-credentials k8s-user \<br/>
	--auth-provider=oidc \<br/>
	--auth-provider-arg=idp-issuer-url=https://dex.honestbee.com  \<br/>
	--auth-provider-arg=client-id=kubernetes \<br/>
	--auth-provider-arg=client-secret=AtMUIzMy00ODg0LTkwMDQtME \<br/>
	{{ if .RefreshToken }}
	--auth-provider-arg=refresh-token=<code>{{ .RefreshToken }}</code> \<br/>
	{{ end }}
	--auth-provider-arg=extra-scopes=groups \<br/>
	--auth-provider-arg=id-token=<code>{{ .IDToken }}</code></p>

	<p>kubectl config set-context k8s-1.7-staging --namespace=default --user=k8s-user --cluster=ap-southeast-1a.staging.k8s.honestbee.com</p>
	<p>kubectl config use-context k8s-1.7-staging</p>

	{{ if .RefreshToken }}
	<form action="{{ .RedirectURL }}" method="post">
	  <input type="hidden" name="refresh_token" value="{{ .RefreshToken }}">
	  <input type="submit" value="Redeem refresh token">
    </form>
	{{ end }}

	<form action="/download" method="post">
		<input type="hidden" name="refresh_token" value="{{ .RefreshToken }}">
		<input type="hidden" name="id_token" value="{{ .IDToken }}">
		<input type="submit" value="Download Kubeconfig">
	</form>
  </body>
</html>
`))

var kubeConfigTmpl = template.Must(template.New("kubeconfig.tpl").ParseFiles("kubeconfig.tpl"))

func renderIndex(w http.ResponseWriter) {
	renderTemplate(w, indexTmpl, nil)
}

func renderToken(w http.ResponseWriter, redirectURL, idToken, refreshToken string, claims []byte) {
	renderTemplate(w, tokenTmpl, tokenTmplData{
		IDToken:      idToken,
		RefreshToken: refreshToken,
		RedirectURL:  redirectURL,
		Claims:       string(claims),
	})
}

func renderKubeConfig(w http.ResponseWriter, idToken, refreshToken string) {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", "kubeconfig"))
	renderTemplate(w, kubeConfigTmpl, tokenTmplData{
		IDToken:      idToken,
		RefreshToken: refreshToken,
	})
}

func renderTemplate(w http.ResponseWriter, tmpl *template.Template, data interface{}) {
	err := tmpl.Execute(w, data)
	if err == nil {
		return
	}

	switch err := err.(type) {
	case *template.Error:
		// An ExecError guarantees that Execute has not written to the underlying reader.
		log.Printf("Error rendering template %s: %s", tmpl.Name(), err)

		// TODO(ericchiang): replace with better internal server error.
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	default:
		// An error with the underlying write, such as the connection being
		// dropped. Ignore for now.
	}
}
