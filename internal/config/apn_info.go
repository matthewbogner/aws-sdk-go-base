package config

import (
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Builds the user-agent string for APN
func (apn APNInfo) BuildUserAgentString() string {
	builder := smithyhttp.NewUserAgentBuilder()
	builder.AddKeyValue("APN", "1.0")
	builder.AddKeyValue(apn.PartnerName, "1.0")
	for _, p := range apn.Products {
		p.buildUserAgentPart(builder)
	}
	return builder.Build()
}
