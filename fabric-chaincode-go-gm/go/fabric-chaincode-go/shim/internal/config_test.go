package internal

import (
	"fmt"
	"google.golang.org/grpc/keepalive"
	"io/ioutil"
	"testing"
	"time"
)

func TestSm2P(t *testing.T){

	loadConfig()
}


// LoadConfig loads the chaincode configuration
func loadConfig() (Config, error) {
	var err error


	conf := Config{
		ChaincodeName: "CORE_CHAINCODE_ID_NAME",
		// hardcode to match chaincode server
		KaOpts: keepalive.ClientParameters{
			Time:                1 * time.Minute,
			Timeout:             20 * time.Second,
			PermitWithoutStream: true,
		},
	}


	var key []byte

	key, err = ioutil.ReadFile("server.key")
	fmt.Printf("key---,%v",key)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read private key file: %s", err)
	}

	var cert []byte

	cert, err = ioutil.ReadFile("server.crt")
	fmt.Printf("cert---,%v",cert)

	if err != nil {
		return Config{}, fmt.Errorf("failed to read public key file: %s", err)
	}


	root, err := ioutil.ReadFile("ca.crt")
	fmt.Printf("root---,%v",root)

	if err != nil {
		return Config{}, fmt.Errorf("failed to read root cert file: %s", err)
	}

	tlscfg, err := LoadTLSConfig(false, key, cert, root)
	if err != nil {
		return Config{}, err
	}

	conf.TLS = tlscfg

	return conf, nil
}
