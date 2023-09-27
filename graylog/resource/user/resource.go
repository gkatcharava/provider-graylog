package user

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		Create: create,
		Read:   read,
		Update: update,
		Delete: destroy,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"user_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			// password is required to create but not required to update
			"password": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"full_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"permissions": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"first_name": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
			},
			// "preferences": {
			//   "updateUnfocussed": false,
			//   "enableSmartSearch": true
			// }
			"timezone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_timeout_ms": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3600000, //nolint:gomnd
			},
			"external": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			// startpage: null
			"roles": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"read_only": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"session_active": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"last_activity": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
