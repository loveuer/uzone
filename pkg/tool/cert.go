package tool

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/youmark/pkcs8"
	"log"
	"os"
)

func LoadTlsConfig(caCertFilePath, certFilePath, keyFilePath string, keyPassword ...string) (*tls.Config, error) {
	cert, err := LoadCertificate(certFilePath, keyFilePath, keyPassword...)
	if err != nil {
		return nil, err
	}

	var (
		caBytes []byte
	)

	if caBytes, err = os.ReadFile(caCertFilePath); err != nil {
		return nil, err
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caBytes)

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}

	return config, nil
}

func LoadCertificate(certFilePath, keyFilePath string, keyPassword ...string) (tls.Certificate, error) {
	var (
		err  error
		cert tls.Certificate
	)

	if len(keyPassword) == 0 || keyPassword[0] == "" {
		return tls.LoadX509KeyPair(certFilePath, keyFilePath)
	}

	var (
		crtBytes []byte
		keyBytes []byte
	)

	if crtBytes, err = os.ReadFile(certFilePath); err != nil {
		return cert, err
	}

	if keyBytes, err = os.ReadFile(keyFilePath); err != nil {
		return cert, err
	}

	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return cert, fmt.Errorf("pem decode key bytes failed")
	}

	data, param, err := pkcs8.ParsePrivateKey(block.Bytes, []byte(keyPassword[0]))
	if err != nil {
		return cert, err
	}

	kb, err := x509.MarshalPKCS8PrivateKey(data)
	if err != nil {
		log.Fatalf("[4] err: %v", err)
	}

	keyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: kb,
	})

	c, err := tls.X509KeyPair(crtBytes, keyPem)
	//c, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("[5] err: %s", err.Error())
	}

	_ = param
	return c, err
}
