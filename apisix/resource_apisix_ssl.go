package apisix

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ResourceSslType struct {
	p      provider
	client ApiClient
}

type SslCertificate struct {
	ID          types.String `tfsdk:"id"`
	Certificate types.String `tfsdk:"certificate"`
	PrivateKey  types.String `tfsdk:"private_key"`
}

func (r ResourceSslType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return ResourceSslType{
		p:      *(p.(*provider)),
		client: getCl(),
	}, nil
}

func (r ResourceSslType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.StringType,
				Computed: true,
			},
			"certificate": {
				Type:     types.StringType,
				Required: true,
			},
			"private_key": {
				Type:      types.StringType,
				Required:  true,
				Sensitive: true,
			},
		},
	}, nil
}

func (r ResourceSslType) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var plan SslCertificate

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	snis, err := ParseCert(plan.Certificate.Value, plan.PrivateKey.Value)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error detect snis from certificates",
			"Error detect snis from certificates, unexpected error: "+err.Error(),
		)
		return
	}

	kk := SSL{
		Certificate: plan.Certificate.Value,
		PrivateKey:  plan.PrivateKey.Value,
		SNIS:        snis,
	}

	result, err := r.client.CreateSsl(kk)

	if err != nil {
		resp.Diagnostics.AddError(
			"Can't create new ssl resource",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	var newState = SslCertificate{
		ID:          types.String{Value: result.ID},
		Certificate: types.String{Value: result.Certificate},
		PrivateKey:  types.String{Value: kk.PrivateKey},
	}

	diags = resp.State.Set(ctx, &newState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r ResourceSslType) Delete(ctx context.Context, request tfsdk.DeleteResourceRequest, response *tfsdk.DeleteResourceResponse) {
	var state SslCertificate

	diags := request.State.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteSsl(state.ID.Value)

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

func (r ResourceSslType) Read(ctx context.Context, request tfsdk.ReadResourceRequest, response *tfsdk.ReadResourceResponse) {
	var state SslCertificate

	diags := request.State.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	ssl, err := r.client.GetSsl(state.ID.Value)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't read ssl resource",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	//if ssl.Certificate == "" {
	//	state.ID = types.String{Value: ""}
	//}

	state.Certificate = types.String{Value: ssl.Certificate}

	diags = response.State.Set(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (r ResourceSslType) Update(ctx context.Context, request tfsdk.UpdateResourceRequest, response *tfsdk.UpdateResourceResponse) {
	var plan SslCertificate

	diags := request.State.Get(ctx, &plan)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	snis, err := ParseCert(plan.Certificate.Value, plan.PrivateKey.Value)

	if err != nil {
		response.Diagnostics.AddError(
			"Error detect snis from certificates",
			"Error detect snis from certificates, unexpected error: "+err.Error(),
		)
		return
	}

	kk := SSL{
		Certificate: plan.Certificate.Value,
		PrivateKey:  plan.PrivateKey.Value,
		SNIS:        snis,
	}

	result, err := r.client.UpdateSsl(plan.ID.Value, kk)

	if err != nil {
		response.Diagnostics.AddError(
			"Can't update ssl resource",
			"Unexpected error: "+err.Error(),
		)
		return
	}

	if result.Certificate != kk.Certificate {
		response.Diagnostics.AddError(
			"After update ssl resource: request and response cert is not equal",
			"Unexpected error",
		)
		return
	}

	var newState = SslCertificate{
		ID:          types.String{Value: result.ID},
		Certificate: types.String{Value: result.Certificate},
		PrivateKey:  types.String{Value: plan.PrivateKey.Value},
	}

	diags = response.State.Set(ctx, &newState)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (r ResourceSslType) ImportState(ctx context.Context, request tfsdk.ImportResourceStateRequest, response *tfsdk.ImportResourceStateResponse) {
	//TODO implement me
	panic("implement me")
}

func ParseCert(crt string, key string) ([]string, error) {
	certDERBlock, _ := pem.Decode([]byte(crt))
	if certDERBlock == nil {
		return []string{}, nil
	}

	_, err := tls.X509KeyPair([]byte(crt), []byte(key))
	if err != nil {
		return []string{}, err
	}

	x509Cert, err := x509.ParseCertificate(certDERBlock.Bytes)

	if err != nil {
		return []string{}, err
	}

	var snis []string
	if x509Cert.DNSNames != nil && len(x509Cert.DNSNames) > 0 {
		snis = x509Cert.DNSNames
	} else if x509Cert.IPAddresses != nil && len(x509Cert.IPAddresses) > 0 {
		for _, ip := range x509Cert.IPAddresses {
			snis = append(snis, ip.String())
		}
	} else {
		if x509Cert.Subject.Names != nil && len(x509Cert.Subject.Names) > 1 {
			var attributeTypeNames = map[string]string{
				"2.5.4.6":  "C",
				"2.5.4.10": "O",
				"2.5.4.11": "OU",
				"2.5.4.3":  "CN",
				"2.5.4.5":  "SERIALNUMBER",
				"2.5.4.7":  "L",
				"2.5.4.8":  "ST",
				"2.5.4.9":  "STREET",
				"2.5.4.17": "POSTALCODE",
			}
			for _, tv := range x509Cert.Subject.Names {
				oidString := tv.Type.String()
				typeName, ok := attributeTypeNames[oidString]
				if ok && typeName == "CN" {
					valueString := fmt.Sprint(tv.Value)
					snis = append(snis, valueString)
				}
			}
		}
	}

	return snis, nil
}
