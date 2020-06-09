package cluster

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/mumoshu/terraform-provider-eksctl/pkg/awsclicompat"
)

type Configuration struct {
	AWSSession *session.Session
}

func configureFunc() func(*schema.ResourceData) (interface{}, error) {
	return func(d *schema.ResourceData) (interface{}, error) {
		s := awsclicompat.NewSession(d.Get("region").(string))

		return &Configuration{
			AWSSession: s,
		}, nil
	}
}
