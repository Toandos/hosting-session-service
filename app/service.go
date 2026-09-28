package main

import (
	"context"

	"connectrpc.com/connect"
	hostingv1alpha1 "github.com/Toandos/apis-go/generated/toando/hosting/v1alpha1"
	api "github.com/Toandos/apis-go/generated/toando/hosting/v1alpha1/hostingv1alpha1connect"
)

type ServiceHandler struct{}

var _ api.AppServiceHandler = (*ServiceHandler)(nil)

// CreateApp implements [hostingv1alpha1connect.AppServiceHandler].
func (a *ServiceHandler) CreateApp(context.Context, *connect.Request[hostingv1alpha1.CreateAppRequest]) (*connect.Response[hostingv1alpha1.App], error) {
	panic("unimplemented")
}

// DeleteApp implements [hostingv1alpha1connect.AppServiceHandler].
func (a *ServiceHandler) DeleteApp(context.Context, *connect.Request[hostingv1alpha1.DeleteAppRequest]) (*connect.Response[hostingv1alpha1.DeleteAppResponse], error) {
	panic("unimplemented")
}

// GetApp implements [hostingv1alpha1connect.AppServiceHandler].
func (a *ServiceHandler) GetApp(context.Context, *connect.Request[hostingv1alpha1.GetAppRequest]) (*connect.Response[hostingv1alpha1.App], error) {
	panic("unimplemented")
}

// ListApps implements [hostingv1alpha1connect.AppServiceHandler].
func (a *ServiceHandler) ListApps(context.Context, *connect.Request[hostingv1alpha1.ListAppsRequest]) (*connect.Response[hostingv1alpha1.ListAppsResponse], error) {
	panic("unimplemented")
}
