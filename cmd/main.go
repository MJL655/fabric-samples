package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/studyzy/fabric-sdk-go/pkg/client/ledger"
	"github.com/studyzy/fabric-sdk-go/pkg/core/config"
	"github.com/studyzy/fabric-sdk-go/pkg/fabsdk"
	"github.com/studyzy/fabric-sdk-go/pkg/util/pathvar"
)
var (
	fabricSdk        *fabsdk.FabricSDK
	userName string="User1"
	orgName string=""
	configFabricPath="./config-fabric.yaml"
)
func main() {
	configProvider := config.FromFile(pathvar.Subst(configFabricPath))
	fabricSdk, _ = fabsdk.New(configProvider)
	queryChainInfo()
}


func queryChainInfo() error {

	channelContext := fabricSdk.ChannelContext("mychannel", fabsdk.WithUser(userName), fabsdk.WithOrg(orgName))
	client, err := ledger.New(channelContext)
	if err != nil {
		return err
	}
	result, err := client.QueryInfo()
	if err != nil {
		log.Fatal(err.Error())
		return err
	} else {
		height := result.BCI.Height
		blockHash := hex.EncodeToString(result.BCI.CurrentBlockHash)
		if err != nil {
			return err
		}
		fmt.Println("block_number", height)
		fmt.Println("current_block_hash", blockHash)
	}
	return nil
}
