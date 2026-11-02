package security

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"go.temporal.io/sdk/client"
)

// DialMTLS подключается с клиентским сертификатом: self-hosted с mTLS
// или Temporal Cloud по сертификату namespace.
func DialMTLS(hostPort, namespace, certFile, keyFile, caFile string) (client.Client, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("сертификат клиента: %w", err)
	}
	tlsCfg := &tls.Config{Certificates: []tls.Certificate{cert}, MinVersion: tls.VersionTLS12}

	if caFile != "" { // свой CA сервера для self-hosted; Temporal Cloud он не нужен
		pem, err := os.ReadFile(caFile)
		if err != nil {
			return nil, err
		}
		pool := x509.NewCertPool()
		if !pool.AppendCertsFromPEM(pem) {
			return nil, fmt.Errorf("не удалось прочитать CA из %s", caFile)
		}
		tlsCfg.RootCAs = pool
	}

	return client.Dial(client.Options{
		HostPort:          hostPort,
		Namespace:         namespace,
		ConnectionOptions: client.ConnectionOptions{TLS: tlsCfg},
	})
}

// DialCloudAPIKey подключается к Temporal Cloud по API-ключу. hostPort — региональный
// эндпоинт вида us-east-1.aws.api.temporal.io:7233, namespace — вида orders.a1b2c.
func DialCloudAPIKey(hostPort, namespace, apiKey string) (client.Client, error) {
	return client.Dial(client.Options{
		HostPort:          hostPort,
		Namespace:         namespace,
		ConnectionOptions: client.ConnectionOptions{TLS: &tls.Config{}},
		Credentials:       client.NewAPIKeyStaticCredentials(apiKey),
	})
}
