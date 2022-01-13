package apisix

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/webbankir/terraform-provider-apisix/apisix/model"
)

type ResourceStreamRouteType struct {
	p provider
}

func (r ResourceStreamRouteType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return ResourceStreamRouteType{
		p: *(p.(*provider)),
	}, nil
}

func (r ResourceStreamRouteType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return model.StreamRouteTypeSchema, nil
}

func (r ResourceStreamRouteType) Create(ctx context.Context, request tfsdk.CreateResourceRequest, response *tfsdk.CreateResourceResponse) {
	var plan model.StreamRouteType

	diags := request.Plan.Get(ctx, &plan)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	requestObjectJsonBytes, err := model.StreamRouteTypeStateToMap(plan)
	if err != nil {
		response.Diagnostics.AddError(
			"Error in transformation from state to map",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	result, err := r.p.client.CreateStreamRoute(requestObjectJsonBytes)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't create new stream route resource",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	newState, err := model.StreamRouteTypeMapToState(result)

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

func (r ResourceStreamRouteType) Delete(ctx context.Context, request tfsdk.DeleteResourceRequest, response *tfsdk.DeleteResourceResponse) {
	var state model.StreamRouteType

	diags := request.State.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.p.client.DeleteStreamRoute(state.ID.Value)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't delete stream route resource",
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

func (r ResourceStreamRouteType) Read(ctx context.Context, request tfsdk.ReadResourceRequest, response *tfsdk.ReadResourceResponse) {
	var state model.StreamRouteType

	diags := request.State.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	if state.ID.Null {
		response.Diagnostics.AddError(
			"Can't read stream route resource, ID is null",
			"Unexpected error",
		)
		return
	}

	result, err := r.p.client.GetStreamRoute(state.ID.Value)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't read ssl certificate resource",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	newState, err := model.StreamRouteTypeMapToState(result)

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

func (r ResourceStreamRouteType) Update(ctx context.Context, request tfsdk.UpdateResourceRequest, response *tfsdk.UpdateResourceResponse) {
	var state model.StreamRouteType

	diags := request.Plan.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	requestObjectJsonBytes, err := model.StreamRouteTypeStateToMap(state)
	if err != nil {
		response.Diagnostics.AddError(
			"Error in transformation from state to map",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	result, err := r.p.client.UpdateStreamRoute(state.ID.Value, requestObjectJsonBytes)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't update ssl certificate resource",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	newState, err := model.StreamRouteTypeMapToState(result)

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

func (r ResourceStreamRouteType) ImportState(ctx context.Context, request tfsdk.ImportResourceStateRequest, response *tfsdk.ImportResourceStateResponse) {
	result, err := r.p.client.GetStreamRoute(request.ID)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't read ssl certificate resource",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	newState, err := model.StreamRouteTypeMapToState(result)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't transform json to state",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	diags := response.State.Set(ctx, &newState)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}
}
