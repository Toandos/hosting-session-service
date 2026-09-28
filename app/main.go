package main

import (
	"net/http"

	api "github.com/Toandos/apis-go/generated/toando/hosting/v1alpha1/hostingv1alpha1connect"
)

func main() {
	http.Handle(api.NewAppServiceHandler(&ServiceHandler{}))
	http.ListenAndServe(":8080", nil)
}
