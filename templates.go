package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type tokenTmplData struct {
	CACert          string
	ClientID        string
	ClusterEndpoint string
	IDToken         string
	RefreshToken    string
	RedirectURL     string
	Claims          string
}

var indexTmpl = template.Must(template.New("index.html").ParseFiles("templates/index.html"))

var tokenTmpl = template.Must(template.New("token.html").ParseFiles("templates/token.html"))

var kubeConfigTmpl = template.Must(template.New("kubeconfig.tpl").ParseFiles("templates/kubeconfig.tpl"))

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

func renderKubeConfig(w http.ResponseWriter, ClientID, caCert, clusterEndpoint, idToken, refreshToken string) {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", "kubeconfig"))
	renderTemplate(w, kubeConfigTmpl, tokenTmplData{
		CACert:          caCert,
		ClientID:        ClientID,
		ClusterEndpoint: clusterEndpoint,
		IDToken:         idToken,
		RefreshToken:    refreshToken,
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
