package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aaronland/go-aws/v3/auth"
	requestsigner "github.com/opensearch-project/opensearch-go/v4/signer/awsv2"
)

func main() {

	var method string
	var uri string

	var credentials_uri string

	flag.StringVar(&method, "method", "GET", "")
	flag.StringVar(&uri, "uri", "", "")
	flag.StringVar(&credentials_uri, "credentials-uri", "", "")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Signs an OpenSearch HTTP request using AWS credentials.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options] path(N) path(N)\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	body_r := strings.NewReader(strings.Join(flag.Args(), " "))

	ctx := context.Background()

	aws_cfg, err := auth.NewConfig(ctx, credentials_uri)

	if err != nil {
		log.Fatalf("Failed to create new AWS config, %v", err)
	}

	signer, err := requestsigner.NewSignerWithService(aws_cfg, "es")

	if err != nil {
		log.Fatalf("Failed to create request signer, %v", err)
	}

	req, err := http.NewRequest(method, uri, body_r)

	if err != nil {
		log.Fatalf("Failed to create new HTTP request, %v", err)
	}

	err = signer.SignRequest(req)

	if err != nil {
		log.Fatalf("Failed to sign request, %v", err)
	}

	err = req.Write(os.Stdout)

	if err != nil {
		log.Fatalf("Failed to write request, %v", err)
	}
}
