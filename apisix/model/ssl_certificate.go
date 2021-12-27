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
)

type SslCertificateType struct {
	ID          types.String `tfsdk:"id"`
	Certificate types.String `tfsdk:"certificate"`
	PrivateKey  types.String `tfsdk:"private_key"`
	Snis        types.List   `tfsdk:"snis"`
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

		"snis": {
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Computed:    true,
			Description: "a non-empty arrays of https SNI",
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultFunction(func(ctx context.Context, request tfsdk.ModifyAttributePlanRequest, response *tfsdk.ModifyAttributePlanResponse) (attr.Value, error) {
					var state SslCertificateType
					request.Config.Get(ctx, &state)
					snis, err := ParseCert(state.Certificate.Value, state.PrivateKey.Value)
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
	},
}

func SslCertificateTypeMapToState(data map[string]interface{}) (*SslCertificateType, error) {
	newState := SslCertificateType{}

	if v := data["id"]; v != nil {
		newState.ID = types.String{Value: v.(string)}
	}

	newState.Certificate = types.String{Value: data["cert"].(string)}

	if v := data["snis"]; v != nil {
		var values []attr.Value
		for _, value := range v.([]interface{}) {
			values = append(values, types.String{Value: value.(string)})
		}

		newState.Snis = types.List{ElemType: types.StringType, Elems: values}
	}

	return &newState, nil
}

func SslCertificateTypeStateToMap(state SslCertificateType) (map[string]interface{}, error) {

	requestObject := make(map[string]interface{})

	utils.StringTypeValueToMap(state.Certificate, requestObject, "cert", true)
	utils.StringTypeValueToMap(state.PrivateKey, requestObject, "key", true)
	utils.ListTypeValueToMap(state.Snis, requestObject, "snis", true)

	return requestObject, nil
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
