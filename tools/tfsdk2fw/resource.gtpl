// Code generated by tools/tfsdk2fw/main.go. Manual editing is required.

package {{ .PackageName }}

import (
	"context"
	{{if .HasTimeouts }}"time"{{- end}}

	{{if .HasTimeouts }}"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"{{- end}}
	{{range .FrameworkValidatorsPackages }}
	"github.com/hashicorp/terraform-plugin-framework-validators/{{ . }}"
	{{- end}}
	{{if .ImportFrameworkAttr }}"github.com/hashicorp/terraform-plugin-framework/attr"{{- end}}
	{{if .EmitResourceImportState }}"github.com/hashicorp/terraform-plugin-framework/path"{{- end}}
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	{{if gt (len .FrameworkPlanModifierPackages) 0 }}"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"{{- end}}
	{{- range .FrameworkPlanModifierPackages }}
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/{{ . }}"
	{{- end}}
	{{if gt (len .FrameworkValidatorsPackages) 0 }}"github.com/hashicorp/terraform-plugin-framework/schema/validator"{{- end}}
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/framework"
	{{if .ImportProviderFrameworkTypes }}fwtypes "github.com/hashicorp/terraform-provider-aws/internal/framework/types"{{- end}}
	{{ range .GoImports -}}
	{{ if .Alias }}{{ .Alias }} {{ end }}"{{ .Path }}"
	{{ end }}
)

// @FrameworkResource("{{ .TFTypeName }}")
func newResource{{ .Name }}(context.Context) (resource.ResourceWithConfigure, error) {
	r := &resource{{ .Name }}{}
{{- if gt .DefaultCreateTimeout 0 }}
	r.SetDefaultCreateTimeout({{ .DefaultCreateTimeout }} * time.Nanosecond) // TODO Convert to more human-friendly duration.
{{- end}}
{{- if gt .DefaultReadTimeout 0 }}
	r.SetDefaultReadTimeout({{ .DefaultReadTimeout }} * time.Nanosecond) // TODO Convert to more human-friendly duration.
{{- end}}
{{- if gt .DefaultUpdateTimeout 0 }}
	r.SetDefaultUpdateTimeout({{ .DefaultUpdateTimeout }} * time.Nanosecond) // TODO Convert to more human-friendly duration.
{{- end}}
{{- if gt .DefaultDeleteTimeout 0 }}
	r.SetDefaultDeleteTimeout({{ .DefaultDeleteTimeout }} * time.Nanosecond) // TODO Convert to more human-friendly duration.
{{- end}}

	return r, nil
}

type resource{{ .Name }} struct {
	framework.ResourceWithConfigure
{{- if .HasTimeouts }}
	framework.WithTimeouts
{{- end}}
}

// Metadata should return the full name of the resource, such as
// examplecloud_thing.
func (r *resource{{ .Name }}) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = "{{ .TFTypeName }}"
}

// Schema returns the schema for this resource.
func (r *resource{{ .Name }}) Schema(ctx context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	s := {{ .Schema }}
{{if .HasTimeouts }}
	if s.Blocks == nil {
		s.Blocks = make(map[string]schema.Block)
	}
	s.Blocks["timeouts"] = timeouts.Block(ctx, timeouts.Opts{
	{{- if gt .DefaultCreateTimeout 0 }}
		Create: true,
	{{- end}}
	{{- if gt .DefaultReadTimeout 0 }}
		Read: true,
	{{- end}}
	{{- if gt .DefaultUpdateTimeout 0 }}
		Update: true,
	{{- end}}
	{{- if gt .DefaultDeleteTimeout 0 }}
		Delete: true,
	{{- end}}
	})
{{- end}}

    response.Schema = s
}

// Create is called when the provider must create a new resource.
// Config and planned state values should be read from the CreateRequest and new state values set on the CreateResponse.
func (r *resource{{ .Name }}) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var data resource{{ .Name }}Data

	response.Diagnostics.Append(request.Plan.Get(ctx, &data)...)

	if response.Diagnostics.HasError() {
		return
	}

{{- if gt .DefaultCreateTimeout 0 }}
	createTimeout := r.CreateTimeout(ctx, data.Timeouts)
{{- end}}

	data.ID = types.StringValue("TODO")

    response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

// Read is called when the provider must read resource values in order to update state.
// Planned state values should be read from the ReadRequest and new state values set on the ReadResponse.
func (r *resource{{ .Name }}) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var data resource{{ .Name }}Data

	response.Diagnostics.Append(request.State.Get(ctx, &data)...)

	if response.Diagnostics.HasError() {
		return
	}

{{- if gt .DefaultReadTimeout 0 }}
	readTimeout := r.ReadTimeout(ctx, data.Timeouts)
{{- end}}

    response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

// Update is called to update the state of the resource.
// Config, planned state, and prior state values should be read from the UpdateRequest and new state values set on the UpdateResponse.
func (r *resource{{ .Name }}) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
{{if .EmitResourceUpdateSkeleton }}var old, new resource{{ .Name }}Data

	response.Diagnostics.Append(request.State.Get(ctx, &old)...)

	if response.Diagnostics.HasError() {
		return
	}

	response.Diagnostics.Append(request.Plan.Get(ctx, &new)...)

	if response.Diagnostics.HasError() {
		return
	}

{{- if gt .DefaultUpdateTimeout 0 }}
	updateTimeout := r.UpdateTimeout(ctx, new.Timeouts)
{{- end}}

    response.Diagnostics.Append(response.State.Set(ctx, &new)...){{- else}}// Noop.{{- end}}
}

// Delete is called when the provider must delete the resource.
// Config values may be read from the DeleteRequest.
//
// If execution completes without error, the framework will automatically call DeleteResponse.State.RemoveResource(),
// so it can be omitted from provider logic.
func (r *resource{{ .Name }}) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var data resource{{ .Name }}Data

	response.Diagnostics.Append(request.State.Get(ctx, &data)...)

	if response.Diagnostics.HasError() {
		return
	}

{{- if gt .DefaultDeleteTimeout 0 }}
	deleteTimeout := r.DeleteTimeout(ctx, data.Timeouts)
{{- end}}

	tflog.Debug(ctx, "deleting TODO", map[string]any{
		"id": data.ID.ValueString(),
	})
}

{{if .EmitResourceImportState }}
// ImportState is called when the provider must import the state of a resource instance.
// This method must return enough state so the Read method can properly refresh the full resource.
//
// If setting an attribute with the import identifier, it is recommended to use the ImportStatePassthroughID() call in this method.
func (r *resource{{ .Name }}) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
}
{{- end}}

{{if .EmitResourceModifyPlan }}
// ModifyPlan is called when the provider has an opportunity to modify
// the plan: once during the plan phase when Terraform is determining
// the diff that should be shown to the user for approval, and once
// during the apply phase with any unknown values from configuration
// filled in with their final values.
//
// The planned new state is represented by
// ModifyPlanResponse.Plan. It must meet the following
// constraints:
// 1. Any non-Computed attribute set in config must preserve the exact
// config value or return the corresponding attribute value from the
// prior state (ModifyPlanRequest.State).
// 2. Any attribute with a known value must not have its value changed
// in subsequent calls to ModifyPlan or Create/Read/Update.
// 3. Any attribute with an unknown value may either remain unknown
// or take on any value of the expected type.
//
// Any errors will prevent further resource-level plan modifications.
func (r *resource{{ .Name }}) ModifyPlan(ctx context.Context, request resource.ModifyPlanRequest, response *resource.ModifyPlanResponse) {
	r.SetTagsAll(ctx, request, response)
}
{{- end}}

type resource{{ .Name }}Data struct {
	{{ .Struct }}
	{{if .HasTimeouts }}Timeouts timeouts.Value `tfsdk:"timeouts"`{{- end}}
}
