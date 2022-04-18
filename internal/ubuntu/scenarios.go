package ubuntu

import (
	"errors"
	"fmt"

	"github.com/probr/probr-pack-ubuntu/internal/config"
	"github.com/probr/probr-sdk/utils"
)

func (scenario *scenarioState) gNomeDisplayManager() error {
	// Standard auditing logic to ensures panics are also audited
	stepTrace, payload, err := utils.AuditPlaceholders()
	fmt.Println("GNOEM display manageer =======================================", scenario.currentStep)
	defer func() {
		fmt.Println("audit=================>", scenario.audit)
		if panicErr := recover(); panicErr != nil {
			err = utils.ReformatError("[ERROR] Unexpected behavior occured: %s", panicErr)
		}
		scenario.audit.AuditScenarioStep(scenario.currentStep, stepTrace.String(), payload, err)
	}()

	stepTrace.WriteString("Validate that test is intended to pass; ")

	payload = struct {
		UserName string
		Password string
		Ip       string
	}{
		config.Vars.ServicePacks.Ubuntu.Username,
		config.Vars.ServicePacks.Ubuntu.Password,
		config.Vars.ServicePacks.Ubuntu.Ip,
	}
	fmt.Println("config.Vars.ServicePacks.Ubuntu.Username===============>", config.Vars.ServicePacks.Ubuntu.Username)
	//	response := ConnectAndRunShellCmd("ls -arlt", config.Vars.ServicePacks.Ubuntu.Username,
	//		config.Vars.ServicePacks.Ubuntu.Password,
	//		config.Vars.ServicePacks.Ubuntu.Ip,
	//		"22")
	response := ConnectAndRunShellCmd("ls -arlt", "ansibleuser", "password@123", "172.17.60.136", "22")
	fmt.Println("Response----------------------->", response)
	if len(response) == 0 {
		err = errors.New("not able to connect to the VM")
	}
	return err

}
