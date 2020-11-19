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
	userName ="User1"
	orgName ="Org1"
	configFabricPath="./config-fabric.yaml"
)
func main() {
	err:=queryChainInfo()
	if err!=nil{
		fmt.Println(err.Error())
	}
}


func queryChainInfo() error {
	configProvider := config.FromFile(pathvar.Subst(configFabricPath))
	fabricSdk, err := fabsdk.New(configProvider)
	if err!=nil{
		return err
	}
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
