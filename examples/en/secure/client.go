package secure

import (
	"encoding/base64"
	"fmt"
	"os"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/converter"
	"go.temporal.io/sdk/temporal"
)

const TaskQueue = "secure"

// NewCodecs returns the codec chain. On encoding the codecs are applied from the end:
// compression first, then encryption; on decoding, in the reverse order.
func NewCodecs(keyID string, key []byte) []converter.PayloadCodec {
	return []converter.PayloadCodec{
		&Codec{KeyID: keyID, Keys: map[string][]byte{keyID: key}},
		converter.NewZlibCodec(converter.ZlibCodecOptions{}),
	}
}

// Dial creates a client that encrypts payloads and error text.
// Workers created from this client use the same settings.
func Dial(keyID string, key []byte) (client.Client, error) {
	dc := converter.NewCodecDataConverter(converter.GetDefaultDataConverter(), NewCodecs(keyID, key)...)
	return client.Dial(client.Options{
		DataConverter: dc,
		FailureConverter: temporal.NewDefaultFailureConverter(temporal.DefaultFailureConverterOptions{
			DataConverter:          dc,
			EncodeCommonAttributes: true, // encrypt error messages and stack traces too
		}),
	})
}

// KeyFromEnv reads a base64 AES-256 key from an environment variable.
func KeyFromEnv(name string) ([]byte, error) {
	key, err := base64.StdEncoding.DecodeString(os.Getenv(name))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	if len(key) != 32 {
		return nil, fmt.Errorf("%s: need a 32-byte key, got %d bytes", name, len(key))
	}
	return key, nil
}
