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

// NewCodecs возвращает цепочку codecs. При кодировании они применяются с конца:
// сначала сжатие, потом шифрование; при декодировании — в обратном порядке.
func NewCodecs(keyID string, key []byte) []converter.PayloadCodec {
	return []converter.PayloadCodec{
		&Codec{KeyID: keyID, Keys: map[string][]byte{keyID: key}},
		converter.NewZlibCodec(converter.ZlibCodecOptions{}),
	}
}

// Dial создаёт клиент, который шифрует payloads и текст ошибок.
// Воркеры, созданные из этого клиента, используют те же настройки.
func Dial(keyID string, key []byte) (client.Client, error) {
	dc := converter.NewCodecDataConverter(converter.GetDefaultDataConverter(), NewCodecs(keyID, key)...)
	return client.Dial(client.Options{
		DataConverter: dc,
		FailureConverter: temporal.NewDefaultFailureConverter(temporal.DefaultFailureConverterOptions{
			DataConverter:          dc,
			EncodeCommonAttributes: true, // шифровать и сообщения об ошибках со стеками
		}),
	})
}

// KeyFromEnv читает ключ AES-256 в base64 из переменной окружения.
func KeyFromEnv(name string) ([]byte, error) {
	key, err := base64.StdEncoding.DecodeString(os.Getenv(name))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	if len(key) != 32 {
		return nil, fmt.Errorf("%s: нужен ключ длиной 32 байта, получено %d", name, len(key))
	}
	return key, nil
}
