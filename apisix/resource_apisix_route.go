package apisix

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/webbankir/terraform-provider-apisix/apisix/model"
)

type ResourceRouteType struct {
	p      provider
	client ApiClient
}

func (r ResourceRouteType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return ResourceRouteType{
		p:      *(p.(*provider)),
		client: getCl(),
	}, nil
}

func (r ResourceRouteType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return model.RouteSchema, nil
}

func (r ResourceRouteType) Create(ctx context.Context, request tfsdk.CreateResourceRequest, response *tfsdk.CreateResourceResponse) {
	var plan model.RouteType
	diags := request.Plan.Get(ctx, &plan)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	requestObjectJsonBytes, err := model.RouteTypeStateToMap(plan)
	if err != nil {
		response.Diagnostics.AddError(
			"Error in transformation from state to map",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	result, err := r.client.CreateRoute(requestObjectJsonBytes)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't create new route resource",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	newState, err := model.RouteTypeMapToState(result)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't transform json to state",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	diags = response.State.Set(ctx, &newState)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (r ResourceRouteType) Read(ctx context.Context, request tfsdk.ReadResourceRequest, response *tfsdk.ReadResourceResponse) {

	var state model.RouteType

	diags := request.State.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	if state.ID.Null {
		response.Diagnostics.AddError(
			"Can't read route resource, ID is null",
			"Unexpected error",
		)
		return
	}
	result, err := r.client.GetRoute(state.ID.Value)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't read route resource",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	newState, err := model.RouteTypeMapToState(result)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't transform json to state",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	diags = response.State.Set(ctx, &newState)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (r ResourceRouteType) Update(ctx context.Context, request tfsdk.UpdateResourceRequest, response *tfsdk.UpdateResourceResponse) {
	var state model.RouteType

	diags := request.Plan.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	requestObjectJsonBytes, err := model.RouteTypeStateToMap(state)
	if err != nil {
		response.Diagnostics.AddError(
			"Error in transformation from state to map",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	result, err := r.client.UpdateRoute(state.ID.Value, requestObjectJsonBytes)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't update route resource",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	newState, err := model.RouteTypeMapToState(result)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't convert json to state",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	diags = response.State.Set(ctx, &newState)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (r ResourceRouteType) Delete(ctx context.Context, request tfsdk.DeleteResourceRequest, response *tfsdk.DeleteResourceResponse) {
	var state model.RouteType

	diags := request.State.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteRoute(state.ID.Value)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't delete route resource",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	response.State.RemoveResource(ctx)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (r ResourceRouteType) ImportState(ctx context.Context, request tfsdk.ImportResourceStateRequest, response *tfsdk.ImportResourceStateResponse) {
	//TODO implement me
	panic("implement me")
}
