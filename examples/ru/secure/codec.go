package secure

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"

	commonpb "go.temporal.io/api/common/v1"
	"google.golang.org/protobuf/proto"
)

const (
	metadataEncoding = "encoding"
	encodingName     = "binary/encrypted"
	metadataKeyID    = "encryption-key-id"
)

// Codec шифрует payloads алгоритмом AES-256-GCM. ID ключа пишется в метаданные,
// поэтому ротация ключей не ломает расшифровку старых историй.
type Codec struct {
	KeyID string            // каким ключом шифровать новые данные
	Keys  map[string][]byte // все ключи, которыми могут быть зашифрованы старые данные
}

func (c *Codec) Encode(payloads []*commonpb.Payload) ([]*commonpb.Payload, error) {
	key, ok := c.Keys[c.KeyID]
	if !ok {
		return nil, fmt.Errorf("нет ключа %q", c.KeyID)
	}
	out := make([]*commonpb.Payload, len(payloads))
	for i, p := range payloads {
		plain, err := proto.Marshal(p) // шифруем payload целиком, вместе с метаданными
		if err != nil {
			return nil, err
		}
		sealed, err := seal(key, plain)
		if err != nil {
			return nil, err
		}
		out[i] = &commonpb.Payload{
			Metadata: map[string][]byte{
				metadataEncoding: []byte(encodingName),
				metadataKeyID:    []byte(c.KeyID),
			},
			Data: sealed,
		}
	}
	return out, nil
}

func (c *Codec) Decode(payloads []*commonpb.Payload) ([]*commonpb.Payload, error) {
	out := make([]*commonpb.Payload, len(payloads))
	for i, p := range payloads {
		if string(p.GetMetadata()[metadataEncoding]) != encodingName {
			out[i] = p // не наш формат: например, данные, записанные до включения шифрования
			continue
		}
		keyID := string(p.GetMetadata()[metadataKeyID])
		key, ok := c.Keys[keyID]
		if !ok {
			return nil, fmt.Errorf("нет ключа %q", keyID)
		}
		plain, err := open(key, p.GetData())
		if err != nil {
			return nil, err
		}
		var orig commonpb.Payload
		if err := proto.Unmarshal(plain, &orig); err != nil {
			return nil, err
		}
		out[i] = &orig
	}
	return out, nil
}

func seal(key, plain []byte) ([]byte, error) {
	gcm, err := newGCM(key)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}
	return gcm.Seal(nonce, nonce, plain, nil), nil
}

func open(key, sealed []byte) ([]byte, error) {
	gcm, err := newGCM(key)
	if err != nil {
		return nil, err
	}
	if len(sealed) < gcm.NonceSize() {
		return nil, errors.New("зашифрованные данные короче nonce")
	}
	nonce, ciphertext := sealed[:gcm.NonceSize()], sealed[gcm.NonceSize():]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

func newGCM(key []byte) (cipher.AEAD, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewGCM(block)
}
