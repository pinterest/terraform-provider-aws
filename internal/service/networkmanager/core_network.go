package networkmanager

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/service/networkmanager"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/flex"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

// This resource is explicitly NOT exported from the provider until design is finalized.
// Its Delete handler is used by sweepers.
func ResourceCoreNetwork() *schema.Resource {
	return &schema.Resource{
		CreateWithoutTimeout: resourceCoreNetworkCreate,
		ReadWithoutTimeout:   resourceCoreNetworkRead,
		UpdateWithoutTimeout: resourceCoreNetworkUpdate,
		DeleteWithoutTimeout: resourceCoreNetworkDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		CustomizeDiff: verify.SetTagsDiff,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 256),
			},
			"edges": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"asn": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"edge_location": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"inside_cidr_blocks": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"global_network_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringLenBetween(0, 50),
			},
			"policy_document": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.All(
					validation.StringLenBetween(0, 10000000),
					validation.StringIsJSON,
				),
				DiffSuppressFunc: verify.SuppressEquivalentJSONDiffs,
				StateFunc: func(v interface{}) string {
					json, _ := structure.NormalizeJsonString(v)
					return json
				},
			},
			"segments": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"edge_locations": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"shared_segments": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags":     tftags.TagsSchema(),
			"tags_all": tftags.TagsSchemaComputed(),
		},
	}
}

func resourceCoreNetworkCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).NetworkManagerConn
	defaultTagsConfig := meta.(*conns.AWSClient).DefaultTagsConfig
	tags := defaultTagsConfig.MergeTags(tftags.New(d.Get("tags").(map[string]interface{})))

	globalNetworkID := d.Get("global_network_id").(string)

	input := &networkmanager.CreateCoreNetworkInput{
		ClientToken:     aws.String(resource.UniqueId()),
		GlobalNetworkId: aws.String(globalNetworkID),
	}

	if v, ok := d.GetOk("description"); ok {
		input.Description = aws.String(v.(string))
	}

	if v, ok := d.GetOk("policy_document"); ok {
		input.PolicyDocument = aws.String(v.(string))
	}

	if len(tags) > 0 {
		input.Tags = Tags(tags.IgnoreAWS())
	}

	log.Printf("[DEBUG] Creating Network Manager Core Network: %s", input)
	output, err := conn.CreateCoreNetworkWithContext(ctx, input)

	if err != nil {
		return diag.Errorf("creating Core Network: %s", err)
	}

	d.SetId(aws.StringValue(output.CoreNetwork.CoreNetworkId))

	if _, err := waitCoreNetworkCreated(ctx, conn, d.Id(), d.Timeout(schema.TimeoutCreate)); err != nil {
		return diag.Errorf("waiting for Network Manager Core Network (%s) create: %s", d.Id(), err)
	}

	return resourceCoreNetworkRead(ctx, d, meta)
}

func resourceCoreNetworkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).NetworkManagerConn
	defaultTagsConfig := meta.(*conns.AWSClient).DefaultTagsConfig
	ignoreTagsConfig := meta.(*conns.AWSClient).IgnoreTagsConfig

	coreNetwork, err := FindCoreNetworkByID(ctx, conn, d.Id())

	if !d.IsNewResource() && tfresource.NotFound(err) {
		log.Printf("[WARN] Network Manager Core Network %s not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	if err != nil {
		return diag.Errorf("reading Network Manager Core Network (%s): %s", d.Id(), err)
	}

	d.Set("arn", coreNetwork.CoreNetworkArn)
	d.Set("description", coreNetwork.Description)
	d.Set("global_network_id", coreNetwork.GlobalNetworkId)
	d.Set("state", coreNetwork.State)

	if err := d.Set("created_at", aws.TimeValue(coreNetwork.CreatedAt).Format(time.RFC3339)); err != nil {
		return diag.Errorf("setting created_at: %s", err)
	}

	if err := d.Set("edges", flattenEdges(coreNetwork.Edges)); err != nil {
		return diag.Errorf("setting edges: %s", err)
	}

	if err := d.Set("segments", flattenSegments(coreNetwork.Segments)); err != nil {
		return diag.Errorf("setting segments: %s", err)
	}

	// getting the policy document uses a different API call
	// policy document is also optional
	resp, err := conn.GetCoreNetworkPolicyWithContext(ctx, &networkmanager.GetCoreNetworkPolicyInput{
		CoreNetworkId: aws.String(d.Id()),
	})

	// policy document is optional. API returns ResourceNotFoundException if there is no policy document
	if err != nil {
		if tfawserr.ErrCodeEquals(err, networkmanager.ErrCodeResourceNotFoundException) {
			log.Printf("[INFO] Network Manager Core Network %s no Policy Document", d.Id())
		} else {
			return diag.Errorf("reading Network Manager Core Network Policy Document (%s): %s", d.Id(), err)
		}
	}

	// policy document is optional
	if resp != nil && resp.CoreNetworkPolicy != nil {
		encodedPolicyDocument, err := protocol.EncodeJSONValue(resp.CoreNetworkPolicy.PolicyDocument, protocol.NoEscape)
		if err != nil {
			return diag.Errorf("reading Network Manager Core Network Policy Document encoding (%s): %s", d.Id(), err)
		}

		d.Set("policy_document", encodedPolicyDocument)
	}

	tags := KeyValueTags(coreNetwork.Tags).IgnoreAWS().IgnoreConfig(ignoreTagsConfig)

	//lintignore:AWSR002
	if err := d.Set("tags", tags.RemoveDefaultConfig(defaultTagsConfig).Map()); err != nil {
		return diag.Errorf("setting tags: %s", err)
	}

	if err := d.Set("tags_all", tags.Map()); err != nil {
		return diag.Errorf("setting tags_all: %s", err)
	}

	return nil
}

func resourceCoreNetworkUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).NetworkManagerConn

	if d.HasChange("tags_all") {
		o, n := d.GetChange("tags_all")

		if err := UpdateTagsWithContext(ctx, conn, d.Get("arn").(string), o, n); err != nil {
			return diag.Errorf("updating Network Manager Core Network (%s) tags: %s", d.Id(), err)
		}
	}

	return resourceCoreNetworkRead(ctx, d, meta)
}

func resourceCoreNetworkDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).NetworkManagerConn

	log.Printf("[DEBUG] Deleting Network Manager Core Network: %s", d.Id())
	_, err := conn.DeleteCoreNetworkWithContext(ctx, &networkmanager.DeleteCoreNetworkInput{
		CoreNetworkId: aws.String(d.Id()),
	})

	if tfawserr.ErrCodeEquals(err, networkmanager.ErrCodeResourceNotFoundException) {
		return nil
	}

	if err != nil {
		return diag.Errorf("deleting Network Manager Core Network (%s): %s", d.Id(), err)
	}

	if _, err := waitCoreNetworkDeleted(ctx, conn, d.Id(), d.Timeout(schema.TimeoutDelete)); err != nil {
		return diag.Errorf("waiting for Network Manager Core Network (%s) delete: %s", d.Id(), err)
	}

	return nil
}

func FindCoreNetworkByID(ctx context.Context, conn *networkmanager.NetworkManager, id string) (*networkmanager.CoreNetwork, error) {
	input := &networkmanager.GetCoreNetworkInput{
		CoreNetworkId: aws.String(id),
	}

	output, err := conn.GetCoreNetworkWithContext(ctx, input)

	if tfawserr.ErrCodeEquals(err, networkmanager.ErrCodeResourceNotFoundException) {
		return nil, &resource.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	if err != nil {
		return nil, err
	}

	if output == nil || output.CoreNetwork == nil {
		return nil, tfresource.NewEmptyResultError(input)
	}

	return output.CoreNetwork, nil
}

func StatusCoreNetworkState(ctx context.Context, conn *networkmanager.NetworkManager, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		output, err := FindCoreNetworkByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.State), nil
	}
}

func waitCoreNetworkCreated(ctx context.Context, conn *networkmanager.NetworkManager, id string, timeout time.Duration) (*networkmanager.CoreNetwork, error) {
	stateConf := &resource.StateChangeConf{
		// CoreNetwork is in PENDING state before AVAILABLE. No value for PENDING at the moment
		Pending: []string{networkmanager.CoreNetworkStateCreating, "PENDING"},
		Target:  []string{networkmanager.CoreNetworkStateAvailable},
		Timeout: timeout,
		Refresh: StatusCoreNetworkState(ctx, conn, id),
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*networkmanager.CoreNetwork); ok {
		return output, err
	}

	return nil, err
}

func waitCoreNetworkDeleted(ctx context.Context, conn *networkmanager.NetworkManager, id string, timeout time.Duration) (*networkmanager.CoreNetwork, error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{networkmanager.CoreNetworkStateDeleting},
		Target:  []string{},
		Timeout: timeout,
		Refresh: StatusCoreNetworkState(ctx, conn, id),
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*networkmanager.CoreNetwork); ok {
		return output, err
	}

	return nil, err
}

func flattenEdges(coreNetworkEdge []*networkmanager.CoreNetworkEdge) []interface{} {
	if coreNetworkEdge == nil {
		return []interface{}{}
	}

	coreNetworkEdges := []interface{}{}
	for _, edge := range coreNetworkEdge {
		values := map[string]interface{}{}
		values["asn"] = aws.Int64Value(edge.Asn)
		values["edge_location"] = aws.StringValue(edge.EdgeLocation)
		values["inside_cidr_blocks"] = flex.FlattenStringSet(edge.InsideCidrBlocks)
		coreNetworkEdges = append(coreNetworkEdges, values)
	}
	return coreNetworkEdges
}

func flattenSegments(coreNetworkSegment []*networkmanager.CoreNetworkSegment) []interface{} {
	if coreNetworkSegment == nil {
		return []interface{}{}
	}

	coreNetworkSegments := []interface{}{}
	for _, segment := range coreNetworkSegment {
		values := map[string]interface{}{}
		values["edge_locations"] = flex.FlattenStringSet(segment.EdgeLocations)
		values["name"] = aws.StringValue(segment.Name)
		values["shared_segments"] = flex.FlattenStringSet(segment.SharedSegments)
		coreNetworkSegments = append(coreNetworkSegments, values)
	}
	return coreNetworkSegments
}
