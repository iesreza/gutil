package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"github.com/iesreza/gutil/log"
	"io/ioutil"
)

type PublicKey rsa.PublicKey
type PrivateKey rsa.PrivateKey

// GenerateKeyPair generates a new key pair
func GenerateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey) {
	privkey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		log.Error("Unable to generate RSA %s", err)
	}
	return privkey, &privkey.PublicKey
}

// PrivateKeyToBytes private key to bytes
func PrivateKeyToBytes(priv *rsa.PrivateKey) []byte {
	privBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(priv),
		},
	)

	return privBytes
}

// PublicKeyToBytes public key to bytes
func PublicKeyToBytes(pub *rsa.PublicKey) []byte {
	pubASN1, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		log.Error("Unable to cast public key to bytes %s", err)
	}

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})

	return pubBytes
}

// BytesToPrivateKey bytes to private key
func BytesToPrivateKey(priv []byte) *PrivateKey {
	block, _ := pem.Decode(priv)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		log.Info("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			log.Error("Unable to decrypt pem block %s", err)
		}
	}
	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		log.Error("Unable to parse private key %s", err)
	}
	return (*PrivateKey)(key)
}

// BytesToPublicKey bytes to public key
func BytesToPublicKey(pub []byte) *PublicKey {
	block, _ := pem.Decode(pub)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		log.Info("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			log.Error("Unable to decrypt pem block %s", err)
		}
	}
	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		log.Error("Unable to parse public key %s", err)
	}
	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		log.Error("not ok")
	}

	return (*PublicKey)(key)
}

// encryptWithPublicKey encrypts data with public key
func EncryptWithPublicKey(msg []byte, pub *rsa.PublicKey) []byte {
	hash := sha256.New()

	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pub, msg, nil)
	if err != nil {
		log.Error("Unable to encrypt with public key %s", err)
	}
	return ciphertext
}

// DecryptWithPrivateKey decrypts data with private key
func DecryptWithPrivateKey(ciphertext []byte, priv *rsa.PrivateKey) []byte {

	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
	if err != nil {
		log.Error("Unable to decrypt with privatekey %s", err)
	}
	return plaintext
}

func OpenPublicKey(path string) *PublicKey {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error("Unable to open public key file %s", err)
		return nil
	}
	return BytesToPublicKey(b)
}

func ParsePublicKey(key string) *PublicKey {
	return BytesToPublicKey([]byte(key))
}

func ParsePrivateKey(key string) *PrivateKey {
	return BytesToPrivateKey([]byte(key))
}

func (key *PrivateKey) GetPublicKey() *PublicKey {
	return (*PublicKey)(&((*rsa.PrivateKey)(key)).PublicKey)
}

func (key *PrivateKey) Decrypt(cipher interface{}) []byte {
	switch cipher.(type) {
	case []byte:
		return DecryptWithPrivateKey(cipher.([]byte), (*rsa.PrivateKey)(key))
	case string:
		return DecryptWithPrivateKey(cipher.([]byte), (*rsa.PrivateKey)(key))
	default:
		log.ErrorF("Invalid cipher type")
	}
	return nil
}

func (key *PublicKey) Encrypt(obj interface{}) []byte {
	switch obj.(type) {
	case []byte:
		return EncryptWithPublicKey(obj.([]byte), (*rsa.PublicKey)(key))
	case string:
		return EncryptWithPublicKey([]byte(obj.(string)), (*rsa.PublicKey)(key))
	default:
		b, err := json.Marshal(obj)
		if err != nil {
			log.ErrorF("Unable to serialize object to json")
		}
		return EncryptWithPublicKey(b, (*rsa.PublicKey)(key))
	}
}

func (key *PrivateKey) Encrypt(obj interface{}) []byte {
	return key.GetPublicKey().Encrypt(obj)
}

func (key *PublicKey) Bytes() []byte {
	return PublicKeyToBytes((*rsa.PublicKey)(key))
}

func (key *PrivateKey) Bytes() []byte {
	return PrivateKeyToBytes((*rsa.PrivateKey)(key))
}

func (key *PublicKey) String() string {
	return string(key.Bytes())
}

func (key *PrivateKey) String() string {
	return string(key.Bytes())
}
