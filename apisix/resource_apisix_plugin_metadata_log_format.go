package apisix

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/model"
)

type ResourcePluginMetadataLogFormatType struct {
	p provider
}

func (r ResourcePluginMetadataLogFormatType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return ResourcePluginMetadataLogFormatType{
		p: *(p.(*provider)),
	}, nil
}

func (r ResourcePluginMetadataLogFormatType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return model.PluginMetadataHTTPLoggerSchema, nil
}

func (r ResourcePluginMetadataLogFormatType) Create(ctx context.Context, request tfsdk.CreateResourceRequest, response *tfsdk.CreateResourceResponse) {
	var plan model.PluginMetadataLogFormatType

	diags := request.Plan.Get(ctx, &plan)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	requestObjectJsonBytes, err := model.PluginMetadataHTTPLoggerTypeStateToMap(plan)
	if err != nil {
		response.Diagnostics.AddError(
			"Error in transformation from state to map",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	result, err := r.p.client.CreatePluginMetadataLogFormat(plan.Name.Value, requestObjectJsonBytes)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't create resource",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	newState, err := model.PluginMetadataHTTPLoggerTypeMapToState(result)
	newState.Name = types.String{Value: plan.Name.Value}

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

func (r ResourcePluginMetadataLogFormatType) Delete(ctx context.Context, request tfsdk.DeleteResourceRequest, response *tfsdk.DeleteResourceResponse) {
	var state model.PluginMetadataLogFormatType

	diags := request.State.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.p.client.DeletePluginMetadataLogFormat(state.Name.Value)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't delete resource",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	response.State.RemoveResource(ctx)
}

func (r ResourcePluginMetadataLogFormatType) Read(ctx context.Context, request tfsdk.ReadResourceRequest, response *tfsdk.ReadResourceResponse) {
	var state model.PluginMetadataLogFormatType

	diags := request.State.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	result, err := r.p.client.GetPluginMetadataLogFormat(state.Name.Value)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't read certificate resource",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	newState, err := model.PluginMetadataHTTPLoggerTypeMapToState(result)
	newState.Name = types.String{Value: state.Name.Value}

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

func (r ResourcePluginMetadataLogFormatType) Update(ctx context.Context, request tfsdk.UpdateResourceRequest, response *tfsdk.UpdateResourceResponse) {
	var state model.PluginMetadataLogFormatType

	diags := request.Plan.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	requestObjectJsonBytes, err := model.PluginMetadataHTTPLoggerTypeStateToMap(state)
	if err != nil {
		response.Diagnostics.AddError(
			"Error in transformation from state to map",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	result, err := r.p.client.UpdatePluginMetadataLogFormat(state.Name.Value, requestObjectJsonBytes)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't update resource",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	newState, err := model.PluginMetadataHTTPLoggerTypeMapToState(result)
	newState.Name = types.String{Value: state.Name.Value}
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

func (r ResourcePluginMetadataLogFormatType) ImportState(ctx context.Context, request tfsdk.ImportResourceStateRequest, response *tfsdk.ImportResourceStateResponse) {
	//TODO implement me
	panic("implement me")
}
