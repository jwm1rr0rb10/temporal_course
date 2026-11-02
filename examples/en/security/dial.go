package security

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"go.temporal.io/sdk/client"
)

// DialMTLS connects with a client certificate: self-hosted with mTLS
// or Temporal Cloud with a namespace certificate.
func DialMTLS(hostPort, namespace, certFile, keyFile, caFile string) (client.Client, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("client certificate: %w", err)
	}
	tlsCfg := &tls.Config{Certificates: []tls.Certificate{cert}, MinVersion: tls.VersionTLS12}

	if caFile != "" { // your own server CA for self-hosted; Temporal Cloud doesn't need it
		pem, err := os.ReadFile(caFile)
		if err != nil {
			return nil, err
		}
		pool := x509.NewCertPool()
		if !pool.AppendCertsFromPEM(pem) {
			return nil, fmt.Errorf("could not parse the CA from %s", caFile)
		}
		tlsCfg.RootCAs = pool
	}

	return client.Dial(client.Options{
		HostPort:          hostPort,
		Namespace:         namespace,
		ConnectionOptions: client.ConnectionOptions{TLS: tlsCfg},
	})
}

// DialCloudAPIKey connects to Temporal Cloud with an API key. hostPort is a regional
// endpoint such as us-east-1.aws.api.temporal.io:7233, namespace looks like orders.a1b2c.
func DialCloudAPIKey(hostPort, namespace, apiKey string) (client.Client, error) {
	return client.Dial(client.Options{
		HostPort:          hostPort,
		Namespace:         namespace,
		ConnectionOptions: client.ConnectionOptions{TLS: &tls.Config{}},
		Credentials:       client.NewAPIKeyStaticCredentials(apiKey),
	})
}
