package config

import (
	"encoding/json"
	"fmt"
	"log"

	sdkConfig "github.com/probr/probr-sdk/config"
	"github.com/probr/probr-sdk/config/setter"
	"github.com/probr/probr-sdk/utils"
)

// Vars is a stateful object containing the variables required to execute this pack
var Vars varOptions

// Init will set values with the content retrieved from a filepath, env vars, or defaults
func (ctx *varOptions) Init() (err error) {
	fmt.Println("ctx.Varsfile==============>", ctx.VarsFile)
	if ctx.VarsFile != "" {
		ctx.decode()
		if err != nil {
			log.Printf("[ERROR] %v", err)
			return
		}
	} else {
		log.Printf("[DEBUG] No vars file provided, unexpected behavior may occur")
	}
	sdkConfig.GlobalConfig.VarsFile = ctx.VarsFile
	sdkConfig.GlobalConfig.IP = ctx.ServicePacks.Ubuntu.Ip
	fmt.Println("IP from args------>", sdkConfig.GlobalConfig.IP)
	sdkConfig.GlobalConfig.Init()
	sdkConfig.GlobalConfig.PrepareOutputDirectory(sdkConfig.GlobalConfig.IP+"/audit", sdkConfig.GlobalConfig.IP+"/cucumber")
	log.Printf("[DEBUG] Config initialized by %s", utils.CallerName(1))
	return
}

// decode uses an SDK helper to create a YAML file decoder,
// parse the file to an object, then extracts the values from
// ServicePacks.ubuntu into this context
func (ctx *varOptions) decode() (err error) {
	configDecoder, file, err := sdkConfig.NewConfigDecoder(ctx.VarsFile)
	if err != nil {
		return
	}
	err = configDecoder.Decode(&ctx)
	file.Close()
	return err
}

// LogConfigState will write the config file to the write directory
func (ctx *varOptions) LogConfigState() {
	json, _ := json.MarshalIndent(ctx, "", "  ")
	log.Printf("[INFO] Config State: %s", json)
}

func (ctx *varOptions) Tags() string {
	return sdkConfig.ParseTags(ctx.ServicePacks.Ubuntu.TagInclusions, ctx.ServicePacks.Ubuntu.TagExclusions)
}

func (ctx *ubuntu) setEnvAndDefaults() {
	setter.SetVar(&ctx.Ip, "PROBR_UBUNTU_IP", "172.17.60.136")
	setter.SetVar(&ctx.Password, "PROBR_UBUNTU_PASSWORD", "password@123")
	setter.SetVar(&ctx.Username, "PROBR_UBUNTU_USERNAME", "ansibleuser")
	setter.SetVar(ctx.Port, "PROBR_UBUNTU_PORT", "22")
}
