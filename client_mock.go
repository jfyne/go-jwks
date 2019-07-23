package jwks

import (
	jose "github.com/square/go-jose/v3"
)

type jWKSClientMock struct {
	secret string
}

func NewMockClient(secret string) JWKSClient {
	return &jWKSClientMock{
		secret: secret,
	}
}

func (c *jWKSClientMock) GetSignatureKey(keyId string) (*jose.JSONWebKey, error) {
	return mockKey(c.secret), nil
}

func (c *jWKSClientMock) GetEncryptionKey(keyId string) (*jose.JSONWebKey, error) {
	return mockKey(c.secret), nil
}

func (c *jWKSClientMock) GetKey(keyId string, use string) (*jose.JSONWebKey, error) {
	return mockKey(c.secret), nil
}

func mockKey(secret string) *jose.JSONWebKey {
	return &jose.JSONWebKey{
		KeyID:     "key1",
		Algorithm: "HS256",
		Key:       []byte(secret),
	}
}
