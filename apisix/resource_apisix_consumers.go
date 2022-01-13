package apisix

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/webbankir/terraform-provider-apisix/apisix/model"
)

type ResourceConsumerType struct {
	p provider
}

func (r ResourceConsumerType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return ResourceConsumerType{
		p: *(p.(*provider)),
	}, nil
}

func (r ResourceConsumerType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return model.ConsumerSchema, nil
}

func (r ResourceConsumerType) Create(ctx context.Context, request tfsdk.CreateResourceRequest, response *tfsdk.CreateResourceResponse) {
	var plan model.ConsumerType
	diags := request.Plan.Get(ctx, &plan)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	requestObjectJsonBytes, err := model.ConsumerTypeStateToMap(plan)
	if err != nil {
		response.Diagnostics.AddError(
			"Error in transformation from state to map",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	result, err := r.p.client.CreateConsumer(requestObjectJsonBytes)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't create new consumer resource",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	newState, err := model.ConsumerTypeMapToState(result)

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

func (r ResourceConsumerType) Read(ctx context.Context, request tfsdk.ReadResourceRequest, response *tfsdk.ReadResourceResponse) {

	var state model.ConsumerType

	diags := request.State.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	result, err := r.p.client.GetConsumer(state.Username.Value)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't read consumer resource",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	newState, err := model.ConsumerTypeMapToState(result)

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

func (r ResourceConsumerType) Update(ctx context.Context, request tfsdk.UpdateResourceRequest, response *tfsdk.UpdateResourceResponse) {
	var plan model.ConsumerType

	diags := request.Plan.Get(ctx, &plan)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	requestObjectJsonBytes, err := model.ConsumerTypeStateToMap(plan)
	if err != nil {
		response.Diagnostics.AddError(
			"Error in transformation from state to map",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	result, err := r.p.client.UpdateConsumer(requestObjectJsonBytes)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't update consumer resource",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	newState, err := model.ConsumerTypeMapToState(result)

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

func (r ResourceConsumerType) Delete(ctx context.Context, request tfsdk.DeleteResourceRequest, response *tfsdk.DeleteResourceResponse) {
	var state model.ConsumerType

	diags := request.State.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.p.client.DeleteConsumer(state.Username.Value)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't delete consumer resource",
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

func (r ResourceConsumerType) ImportState(ctx context.Context, request tfsdk.ImportResourceStateRequest, response *tfsdk.ImportResourceStateResponse) {
	result, err := r.p.client.GetConsumer(request.ID)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't read consumer resource",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	newState, err := model.ConsumerTypeMapToState(result)

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
