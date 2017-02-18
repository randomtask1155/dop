package main

import (
	"github.com/gorilla/websocket"
	"fmt"
	"crypto/tls"
	"net/http"
	"flag"
	"os"
	"os/user"
	"path"
	"io/ioutil"
	"encoding/json"
	"time"
	"github.com/cloudfoundry/dropsonde/dropsonde_unmarshaller"
	"github.com/cloudfoundry/sonde-go/events"
)

var (
	wssURL = flag.String("url", "", "Web socket address example: wss://doppler.system.domain:443/apps/41abc841-cbc8-4cab-854d-640a7c8b6a5f/stream")
	accessToken = flag.String("token", "", "Provide an access token used to authenticate with doppler endpoint. Defaults to ~/.cf/config.json")
)

// CFConfig json struct that matches params required from ~/.cf/config.json
type CFConfig struct {
	AccessToken string `json:"AccessToken"`
}

// Get access token from ~/.cf/config.json
func getAccessToken() (string, error) {
	if *accessToken != "" {
		return *accessToken, nil
	}
	
	usr, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("Could not get users home directory: %s", err)
	}
	config := path.Join(usr.HomeDir, ".cf/config.json")
	_,err = os.Stat(config)
	if err != nil {
		return "", fmt.Errorf("Unalbe to stat %s: %s", config, err)
	}
	
	b, err := ioutil.ReadFile(config)
	if err != nil {
		return "", fmt.Errorf("Reading Config %s failed: %s", config, err)
	}
	cfConfig := CFConfig{}
	err = json.Unmarshal(b, &cfConfig)
	if err != nil {
		return "", fmt.Errorf("Could not parse config %s: %s", config, err)
	}
	if cfConfig.AccessToken == "" {
		return "", fmt.Errorf("Invalid access token found in %s", config)
	}
	return cfConfig.AccessToken, nil
}

func createSocket() (*websocket.Conn, error) {
	badConn := new(websocket.Conn)
	token, err := getAccessToken()
	if err != nil {
		return badConn, err
	}
	dialer := websocket.DefaultDialer	
	dialer.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	h := http.Header{}
	h.Add("Authorization", token)
	conn, resp, err := dialer.Dial(*wssURL, h)
	if err != nil {
		var errString error 
		if resp != nil {
			errString =  fmt.Errorf("HTTP Response: %s\nWEBSOCKET Error: %s", resp.Status, err)
		} else {
			errString = fmt.Errorf("WEBSOCKET Error:%s", err)
		}
		return badConn, errString
	}
	return conn, nil
}

func quit(status int) {
	os.Exit(status)
}

func main() {
	flag.Parse()

	if *wssURL == "" {
		fmt.Println("-url must be specified")
		quit(1)
	}

	conn, err := createSocket() 
	if err != nil {
		fmt.Println(err)
		quit(2)
	}
	defer conn.Close()
	
	input := make(chan []byte)
	output := make(chan *events.Envelope)
	dn := dropsonde_unmarshaller.NewDropsondeUnmarshaller()
	
	fmt.Println("starting dropsnode unmarshaller...")
	go dn.Run(input, output)
	
	fmt.Println("starting output collector...")
	go func(output chan *events.Envelope){
		for{
			select {
			case e := <-output:
				fmt.Printf("%v\n", e)
			}
		}
	}(output)
	
	fmt.Println("starting read loop...")
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			time.Sleep(1 * time.Second)
		}
		input <- p
	}
}
