// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	tfTypes "github.com/Styra/terraform-provider-styra/internal/provider/types"
	"github.com/Styra/terraform-provider-styra/internal/sdk"
	"github.com/Styra/terraform-provider-styra/internal/sdk/models/operations"
	"github.com/Styra/terraform-provider-styra/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SecretResource{}
var _ resource.ResourceWithImportState = &SecretResource{}

func NewSecretResource() resource.Resource {
	return &SecretResource{}
}

// SecretResource defines the resource implementation.
type SecretResource struct {
	client *sdk.StyraDas
}

// SecretResourceModel describes the resource data model.
type SecretResourceModel struct {
	Description types.String             `tfsdk:"description"`
	ID          types.String             `tfsdk:"id"`
	IfNoneMatch types.String             `tfsdk:"if_none_match"`
	Metadata    tfTypes.MetaV2ObjectMeta `tfsdk:"metadata"`
	Name        types.String             `tfsdk:"name"`
	RequestID   types.String             `tfsdk:"request_id"`
	Secret      types.String             `tfsdk:"secret"`
	SecretID    types.String             `tfsdk:"secret_id"`
}

func (r *SecretResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_secret"
}

func (r *SecretResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Secret Resource",
		Attributes: map[string]schema.Attribute{
			"description": schema.StringAttribute{
				Required: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"if_none_match": schema.StringAttribute{
				Optional:    true,
				Description: `if set to '*' then the request fill fail if the secret already exists`,
			},
			"metadata": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"created_at": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
					"created_by": schema.StringAttribute{
						Computed: true,
					},
					"created_through": schema.StringAttribute{
						Computed: true,
					},
					"last_modified_at": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
					"last_modified_by": schema.StringAttribute{
						Computed: true,
					},
					"last_modified_through": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"request_id": schema.StringAttribute{
				Computed: true,
			},
			"secret": schema.StringAttribute{
				Required: true,
			},
			"secret_id": schema.StringAttribute{
				Required:    true,
				Description: `secret ID`,
			},
		},
	}
}

func (r *SecretResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.StyraDas)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.StyraDas, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *SecretResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SecretResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	secretID := data.SecretID.ValueString()
	ifNoneMatch := new(string)
	if !data.IfNoneMatch.IsUnknown() && !data.IfNoneMatch.IsNull() {
		*ifNoneMatch = data.IfNoneMatch.ValueString()
	} else {
		ifNoneMatch = nil
	}
	secretsV1SecretsPutRequest := *data.ToSharedSecretsV1SecretsPutRequest()
	request := operations.CreateUpdateSecretRequest{
		SecretID:                   secretID,
		IfNoneMatch:                ifNoneMatch,
		SecretsV1SecretsPutRequest: secretsV1SecretsPutRequest,
	}
	res, err := r.client.Secrets.CreateUpdateSecret(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.SecretsV1SecretsPutResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedSecretsV1SecretsPutResponse(res.SecretsV1SecretsPutResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	secretId1 := data.SecretID.ValueString()
	request1 := operations.GetSecretRequest{
		SecretID: secretId1,
	}
	res1, err := r.client.Secrets.GetSecret(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if res1.SecretsV1SecretsGetResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedSecretsV1Secret(res1.SecretsV1SecretsGetResponse.Result)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SecretResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SecretResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	secretID := data.SecretID.ValueString()
	request := operations.GetSecretRequest{
		SecretID: secretID,
	}
	res, err := r.client.Secrets.GetSecret(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.SecretsV1SecretsGetResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedSecretsV1Secret(res.SecretsV1SecretsGetResponse.Result)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SecretResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SecretResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	secretID := data.SecretID.ValueString()
	ifNoneMatch := new(string)
	if !data.IfNoneMatch.IsUnknown() && !data.IfNoneMatch.IsNull() {
		*ifNoneMatch = data.IfNoneMatch.ValueString()
	} else {
		ifNoneMatch = nil
	}
	secretsV1SecretsPutRequest := *data.ToSharedSecretsV1SecretsPutRequest()
	request := operations.CreateUpdateSecretRequest{
		SecretID:                   secretID,
		IfNoneMatch:                ifNoneMatch,
		SecretsV1SecretsPutRequest: secretsV1SecretsPutRequest,
	}
	res, err := r.client.Secrets.CreateUpdateSecret(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.SecretsV1SecretsPutResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedSecretsV1SecretsPutResponse(res.SecretsV1SecretsPutResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	secretId1 := data.SecretID.ValueString()
	request1 := operations.GetSecretRequest{
		SecretID: secretId1,
	}
	res1, err := r.client.Secrets.GetSecret(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if res1.SecretsV1SecretsGetResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedSecretsV1Secret(res1.SecretsV1SecretsGetResponse.Result)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SecretResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SecretResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	secretID := data.SecretID.ValueString()
	request := operations.DeleteSecretRequest{
		SecretID: secretID,
	}
	res, err := r.client.Secrets.DeleteSecret(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *SecretResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("secret_id"), req.ID)...)
}
