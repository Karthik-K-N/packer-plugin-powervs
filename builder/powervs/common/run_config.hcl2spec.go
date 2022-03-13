// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package common

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatCOS is an auto-generated flat version of COS.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatCOS struct {
	Bucket *string `mapstructure:"bucket" required:"true" cty:"bucket" hcl:"bucket"`
	Object *string `mapstructure:"object" required:"true" cty:"object" hcl:"object"`
	Region *string `mapstructure:"region" required:"true" cty:"region" hcl:"region"`
}

// FlatMapstructure returns a new FlatCOS.
// FlatCOS is an auto-generated flat version of COS.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*COS) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatCOS)
}

// HCL2Spec returns the hcl spec of a COS.
// This spec is used by HCL to read the fields of COS.
// The decoded values from this spec will then be applied to a FlatCOS.
func (*FlatCOS) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"bucket": &hcldec.AttrSpec{Name: "bucket", Type: cty.String, Required: false},
		"object": &hcldec.AttrSpec{Name: "object", Type: cty.String, Required: false},
		"region": &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
	}
	return s
}

// FlatCapture is an auto-generated flat version of Capture.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatCapture struct {
	Name *string         `mapstructure:"name" required:"true" cty:"name" hcl:"name"`
	COS  *FlatCaptureCOS `mapstructure:"cos" required:"false" cty:"cos" hcl:"cos"`
}

// FlatMapstructure returns a new FlatCapture.
// FlatCapture is an auto-generated flat version of Capture.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Capture) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatCapture)
}

// HCL2Spec returns the hcl spec of a Capture.
// This spec is used by HCL to read the fields of Capture.
// The decoded values from this spec will then be applied to a FlatCapture.
func (*FlatCapture) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"name": &hcldec.AttrSpec{Name: "name", Type: cty.String, Required: false},
		"cos":  &hcldec.BlockSpec{TypeName: "cos", Nested: hcldec.ObjectSpec((*FlatCaptureCOS)(nil).HCL2Spec())},
	}
	return s
}

// FlatCaptureCOS is an auto-generated flat version of CaptureCOS.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatCaptureCOS struct {
	Bucket    *string `mapstructure:"bucket" required:"true" cty:"bucket" hcl:"bucket"`
	Region    *string `mapstructure:"region" required:"true" cty:"region" hcl:"region"`
	AccessKey *string `mapstructure:"access_key" required:"true" cty:"access_key" hcl:"access_key"`
	SecretKey *string `mapstructure:"secret_key" required:"true" cty:"secret_key" hcl:"secret_key"`
}

// FlatMapstructure returns a new FlatCaptureCOS.
// FlatCaptureCOS is an auto-generated flat version of CaptureCOS.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*CaptureCOS) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatCaptureCOS)
}

// HCL2Spec returns the hcl spec of a CaptureCOS.
// This spec is used by HCL to read the fields of CaptureCOS.
// The decoded values from this spec will then be applied to a FlatCaptureCOS.
func (*FlatCaptureCOS) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"bucket":     &hcldec.AttrSpec{Name: "bucket", Type: cty.String, Required: false},
		"region":     &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"access_key": &hcldec.AttrSpec{Name: "access_key", Type: cty.String, Required: false},
		"secret_key": &hcldec.AttrSpec{Name: "secret_key", Type: cty.String, Required: false},
	}
	return s
}

// FlatSource is an auto-generated flat version of Source.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatSource struct {
	Name       *string         `mapstructure:"name" required:"true" cty:"name" hcl:"name"`
	Import     *bool           `mapstructure:"import" required:"false" cty:"import" hcl:"import"`
	COS        *FlatCOS        `mapstructure:"cos" required:"false" cty:"cos" hcl:"cos"`
	StockImage *FlatStockImage `mapstructure:"stock_image" required:"false" cty:"stock_image" hcl:"stock_image"`
}

// FlatMapstructure returns a new FlatSource.
// FlatSource is an auto-generated flat version of Source.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Source) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatSource)
}

// HCL2Spec returns the hcl spec of a Source.
// This spec is used by HCL to read the fields of Source.
// The decoded values from this spec will then be applied to a FlatSource.
func (*FlatSource) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"name":        &hcldec.AttrSpec{Name: "name", Type: cty.String, Required: false},
		"import":      &hcldec.AttrSpec{Name: "import", Type: cty.Bool, Required: false},
		"cos":         &hcldec.BlockSpec{TypeName: "cos", Nested: hcldec.ObjectSpec((*FlatCOS)(nil).HCL2Spec())},
		"stock_image": &hcldec.BlockSpec{TypeName: "stock_image", Nested: hcldec.ObjectSpec((*FlatStockImage)(nil).HCL2Spec())},
	}
	return s
}

// FlatStockImage is an auto-generated flat version of StockImage.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatStockImage struct {
	Name *string `mapstructure:"name" required:"true" cty:"name" hcl:"name"`
}

// FlatMapstructure returns a new FlatStockImage.
// FlatStockImage is an auto-generated flat version of StockImage.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*StockImage) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatStockImage)
}

// HCL2Spec returns the hcl spec of a StockImage.
// This spec is used by HCL to read the fields of StockImage.
// The decoded values from this spec will then be applied to a FlatStockImage.
func (*FlatStockImage) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"name": &hcldec.AttrSpec{Name: "name", Type: cty.String, Required: false},
	}
	return s
}