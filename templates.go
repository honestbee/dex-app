package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type indexTmplData struct {
	ClusterList []string
}

type tokenTmplData struct {
	CACert          string
	ClientID        string
	ClusterEndpoint string
	IDToken         string
	RefreshToken    string
	RedirectURL     string
	Claims          string
	NamespaceList   []string
	Namespace       string
}

var indexTmpl = template.Must(template.New("index.html").ParseFiles("templates/index.html"))

var tokenTmpl = template.Must(template.New("token.html").ParseFiles("templates/token.html"))

var kubeConfigTmpl = template.Must(template.New("kubeconfig.tpl").ParseFiles("templates/kubeconfig.tpl"))

func renderIndex(w http.ResponseWriter, ClientClusters map[string]map[string]string) {
	clusterList := []string{}
	for k := range ClientClusters {
		clusterList = append(clusterList, k)
	}
	renderTemplate(w, indexTmpl, indexTmplData{
		ClusterList: clusterList,
	})
}

func renderToken(w http.ResponseWriter, caCert, clientID, clusterEndpoint, redirectURL, idToken, refreshToken string, claims []byte, namespaceList []string) {
	renderTemplate(w, tokenTmpl, tokenTmplData{
		CACert:          caCert,
		ClientID:        clientID,
		ClusterEndpoint: clusterEndpoint,
		IDToken:         idToken,
		RefreshToken:    refreshToken,
		RedirectURL:     redirectURL,
		Claims:          string(claims),
		NamespaceList:   namespaceList,
	})
}

func renderKubeConfig(w http.ResponseWriter, ClientID, caCert, clusterEndpoint, idToken, refreshToken string, namespace string) {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s_%s_%s\"", "kubeconfig", ClientID, namespace))
	renderTemplate(w, kubeConfigTmpl, tokenTmplData{
		CACert:          caCert,
		ClientID:        ClientID,
		ClusterEndpoint: clusterEndpoint,
		IDToken:         idToken,
		RefreshToken:    refreshToken,
		Namespace:       namespace,
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
