package apisix

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/webbankir/terraform-provider-apisix/apisix/model"
)

type ResourceUpstreamType struct {
	p provider
}

func (r ResourceUpstreamType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {

	return ResourceUpstreamType{
		p: *(p.(*provider)),
	}, nil
}

func (r ResourceUpstreamType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: model.UpstreamSchemaAttribute.Attributes.GetAttributes(),
	}, nil
}

func (r ResourceUpstreamType) Create(ctx context.Context, request tfsdk.CreateResourceRequest, response *tfsdk.CreateResourceResponse) {
	var plan model.UpstreamType
	diags := request.Plan.Get(ctx, &plan)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	requestObjectJsonBytes, err := model.UpstreamTypeStateToMap(&plan, false)
	if err != nil {
		response.Diagnostics.AddError(
			"Error in transformation from state to map",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	result, err := r.p.client.CreateUpstream(requestObjectJsonBytes)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't create new upstream resource",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	newState, err := model.UpstreamTypeMapToState(map[string]interface{}{"upstream": result})

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

func (r ResourceUpstreamType) Read(ctx context.Context, request tfsdk.ReadResourceRequest, response *tfsdk.ReadResourceResponse) {

	var state model.UpstreamType

	diags := request.State.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	if state.ID.Null {
		response.Diagnostics.AddError(
			"Can't read upstream resource, ID is null",
			"Unexpected error",
		)
		return
	}
	result, err := r.p.client.GetUpstream(state.ID.Value)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't read upstream resource",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	newState, err := model.UpstreamTypeMapToState(map[string]interface{}{"upstream": result})

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

func (r ResourceUpstreamType) Update(ctx context.Context, request tfsdk.UpdateResourceRequest, response *tfsdk.UpdateResourceResponse) {
	var state model.UpstreamType

	diags := request.Plan.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	requestObjectJsonBytes, err := model.UpstreamTypeStateToMap(&state, true)
	if err != nil {
		response.Diagnostics.AddError(
			"Error in transformation from state to map",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	result, err := r.p.client.UpdateUpstream(state.ID.Value, requestObjectJsonBytes)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't update upstream resource",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	newState, err := model.UpstreamTypeMapToState(map[string]interface{}{"upstream": result})

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

func (r ResourceUpstreamType) Delete(ctx context.Context, request tfsdk.DeleteResourceRequest, response *tfsdk.DeleteResourceResponse) {
	var state model.UpstreamType

	diags := request.State.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.p.client.DeleteUpstream(state.ID.Value)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't delete upstream resource",
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

func (r ResourceUpstreamType) ImportState(ctx context.Context, request tfsdk.ImportResourceStateRequest, response *tfsdk.ImportResourceStateResponse) {
	//TODO implement me
	panic("implement me")
}
