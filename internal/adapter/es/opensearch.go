package es

import (
	"context"
	"crypto/tls"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/opensearch-project/opensearch-go/v2"
	"github.com/opensearch-project/opensearch-go/v2/signer/awsv2"

	conf "grapefruit/config"
	"net/http"
)

var client *opensearch.Client

func MustInit() {
	env := conf.GetLabelVal(context.Background())
	if env == "prod" {
		prod()
		return
	}
	dev()
	return
}

func dev() {
	esConfig := conf.GetESConfig()
	var err error
	client, err = opensearch.NewClient(opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Addresses: []string{esConfig.Endpoint},
		Username:  "", //
		Password:  "",
	})
	if err != nil {
		panic(err)
	}
}

func prod() {
	ctx := context.Background()
	esConfig := conf.GetESConfig()
	awsCfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("cn-north-1"),
		config.WithCredentialsProvider(
			aws.NewCredentialsCache(
				credentials.NewStaticCredentialsProvider(
					esConfig.CredentialKey,
					esConfig.CredentialSecret, "",
				)),
		),
	)
	if err != nil {
		panic(err)
	}

	signer, err := awsv2.NewSignerWithService(awsCfg, "es")
	if err != nil {
		panic(err)
	}

	client, err = opensearch.NewClient(opensearch.Config{
		Addresses: []string{esConfig.Endpoint},
		Signer:    signer,
	})
	if err != nil {
		panic(err)
	}

}

func GetESClient() *opensearch.Client {
	return client
}
