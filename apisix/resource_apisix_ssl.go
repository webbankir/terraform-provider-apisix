package apisix

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/model"
)

type ResourceSslCertificateType struct {
	p provider
}

func (r ResourceSslCertificateType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return ResourceSslCertificateType{
		p: *(p.(*provider)),
	}, nil
}

func (r ResourceSslCertificateType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return model.SslCertificateSchema, nil
}

func (r ResourceSslCertificateType) Create(ctx context.Context, request tfsdk.CreateResourceRequest, response *tfsdk.CreateResourceResponse) {
	var plan model.SslCertificateType

	diags := request.Plan.Get(ctx, &plan)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	requestObjectJsonBytes, err := model.SslCertificateTypeStateToMap(plan)
	if err != nil {
		response.Diagnostics.AddError(
			"Error in transformation from state to map",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	result, err := r.p.client.CreateSslCertificate(requestObjectJsonBytes)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't create new ssl resource",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	newState, err := model.SslCertificateTypeMapToState(result)
	newState.PrivateKey = types.String{Value: plan.PrivateKey.Value}

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

func (r ResourceSslCertificateType) Delete(ctx context.Context, request tfsdk.DeleteResourceRequest, response *tfsdk.DeleteResourceResponse) {
	var state model.SslCertificateType

	diags := request.State.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.p.client.DeleteSslCertificate(state.ID.Value)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't delete ssl resource",
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

func (r ResourceSslCertificateType) Read(ctx context.Context, request tfsdk.ReadResourceRequest, response *tfsdk.ReadResourceResponse) {
	var state model.SslCertificateType

	diags := request.State.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	if state.ID.Null {
		response.Diagnostics.AddError(
			"Can't read ssl certificate resource, ID is null",
			"Unexpected error",
		)
		return
	}

	result, err := r.p.client.GetSslCertificate(state.ID.Value)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't read ssl certificate resource",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	newState, err := model.SslCertificateTypeMapToState(result)
	newState.PrivateKey = types.String{Value: state.PrivateKey.Value}

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

func (r ResourceSslCertificateType) Update(ctx context.Context, request tfsdk.UpdateResourceRequest, response *tfsdk.UpdateResourceResponse) {
	var state model.SslCertificateType

	diags := request.Plan.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	requestObjectJsonBytes, err := model.SslCertificateTypeStateToMap(state)
	if err != nil {
		response.Diagnostics.AddError(
			"Error in transformation from state to map",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	result, err := r.p.client.UpdateSslCertificate(state.ID.Value, requestObjectJsonBytes)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't update ssl certificate resource",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	newState, err := model.SslCertificateTypeMapToState(result)
	newState.PrivateKey = types.String{Value: state.PrivateKey.Value}

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

func (r ResourceSslCertificateType) ImportState(ctx context.Context, request tfsdk.ImportResourceStateRequest, response *tfsdk.ImportResourceStateResponse) {

	result, err := r.p.client.GetSslCertificate(request.ID)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't read ssl certificate resource",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	newState, err := model.SslCertificateTypeMapToState(result)

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
