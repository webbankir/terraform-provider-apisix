package model

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"math/big"
)

type SslCertificateType struct {
	ID            types.String `tfsdk:"id"`
	IsEnabled     types.Bool   `tfsdk:"is_enabled"`
	Certificate   types.String `tfsdk:"certificate"`
	PrivateKey    types.String `tfsdk:"private_key"`
	Snis          types.List   `tfsdk:"snis"`
	ValidityEnd   types.Number `tfsdk:"validity_end"`
	ValidityStart types.Number `tfsdk:"validity_start"`
	Labels        types.Map    `tfsdk:"labels"`
}

var SslCertificateSchema = tfsdk.Schema{
	Attributes: map[string]tfsdk.Attribute{
		"id": {
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"certificate": {
			Type:        types.StringType,
			Required:    true,
			Description: "https certificate",
		},
		"private_key": {
			Type:        types.StringType,
			Required:    true,
			Sensitive:   true,
			Description: "https private key",
		},

		"validity_end": {
			Type:        types.NumberType,
			Optional:    true,
			Computed:    true,
			Description: "NotAfter",
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultFunction(func(ctx context.Context, request tfsdk.ModifyAttributePlanRequest, response *tfsdk.ModifyAttributePlanResponse) (attr.Value, error) {
					var state SslCertificateType
					request.Config.Get(ctx, &state)
					notAfter, err := CertNotAfter(state.Certificate.Value)
					if err != nil {
						return types.Number{Null: true}, err
					}
					return types.Number{Value: big.NewFloat(float64(notAfter))}, nil
				}),
			},
		},

		"validity_start": {
			Type:        types.NumberType,
			Optional:    true,
			Computed:    true,
			Description: "NotBefore",
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultFunction(func(ctx context.Context, request tfsdk.ModifyAttributePlanRequest, response *tfsdk.ModifyAttributePlanResponse) (attr.Value, error) {
					var state SslCertificateType
					request.Config.Get(ctx, &state)
					notBefore, err := CertNotBefore(state.Certificate.Value)
					if err != nil {
						return types.Number{Null: true}, err
					}
					return types.Number{Value: big.NewFloat(float64(notBefore))}, nil
				}),
			},
		},

		"snis": {
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Computed:    true,
			Description: "a non-empty arrays of https SNI",
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultFunction(func(ctx context.Context, request tfsdk.ModifyAttributePlanRequest, response *tfsdk.ModifyAttributePlanResponse) (attr.Value, error) {
					var state SslCertificateType
					request.Config.Get(ctx, &state)
					snis, err := CertSNIS(state.Certificate.Value, state.PrivateKey.Value)
					if err != nil {
						return types.List{Null: true}, err
					}

					var values []attr.Value
					for _, value := range snis {
						values = append(values, types.String{Value: value})
					}

					return types.List{ElemType: types.StringType, Elems: values}, nil
				}),
			},
		},
		"labels": {
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
		},
		"is_enabled": {
			Type:     types.BoolType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultBool(true),
			},
		},
	},
}

func SslCertificateTypeMapToState(jsonMap map[string]interface{}) (*SslCertificateType, error) {
	newState := SslCertificateType{}

	utils.MapValueToStringTypeValue(jsonMap, "id", &newState.ID)
	utils.MapValueToStringTypeValue(jsonMap, "cert", &newState.Certificate)
	utils.MapValueToListTypeValue(jsonMap, "snis", &newState.Snis)
	utils.MapValueToNumberTypeValue(jsonMap, "validity_start", &newState.ValidityStart)
	utils.MapValueToNumberTypeValue(jsonMap, "validity_end", &newState.ValidityEnd)
	utils.MapValueToMapTypeValue(jsonMap, "labels", &newState.Labels)

	if v := jsonMap["status"]; v != nil {
		if v.(float64) == 1 {
			newState.IsEnabled = types.Bool{Value: true}
		} else {
			newState.IsEnabled = types.Bool{Value: false}
		}
	} else {
		newState.IsEnabled = types.Bool{Null: true}
	}
	return &newState, nil
}

func SslCertificateTypeStateToMap(state SslCertificateType) (map[string]interface{}, error) {

	requestObject := make(map[string]interface{})

	utils.StringTypeValueToMap(state.Certificate, requestObject, "cert")
	utils.StringTypeValueToMap(state.PrivateKey, requestObject, "key")
	utils.ListTypeValueToMap(state.Snis, requestObject, "snis")
	utils.NumberTypeValueToMap(state.ValidityStart, requestObject, "validity_start")
	utils.NumberTypeValueToMap(state.ValidityEnd, requestObject, "validity_end")
	utils.MapTypeValueToMap(state.Labels, requestObject, "labels")
	if !state.IsEnabled.Null {
		if state.IsEnabled.Value {
			requestObject["status"] = 1
		} else {
			requestObject["status"] = 0
		}
	}
	return requestObject, nil
}

func CertSNIS(crt string, key string) ([]string, error) {
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

func CertNotAfter(crt string) (int64, error) {
	certDERBlock, _ := pem.Decode([]byte(crt))
	if certDERBlock == nil {
		return 0, nil
	}

	x509Cert, err := x509.ParseCertificate(certDERBlock.Bytes)
	return x509Cert.NotAfter.Unix(), err

}
func CertNotBefore(crt string) (int64, error) {
	certDERBlock, _ := pem.Decode([]byte(crt))
	if certDERBlock == nil {
		return 0, nil
	}

	x509Cert, err := x509.ParseCertificate(certDERBlock.Bytes)
	return x509Cert.NotBefore.Unix(), err

}
