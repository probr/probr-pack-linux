package ubuntu

import (
	"errors"
	"fmt"

	"github.com/probr/probr-pack-ubuntu/internal/config"
	"github.com/probr/probr-sdk/utils"
)

func (scenario *scenarioState) ensureUfwFirewallIsConfigured() error {
	// Standard auditing logic to ensures panics are also audited
	stepTrace, payload, err := utils.AuditPlaceholders()
	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = utils.ReformatError("[ERROR] Unexpected behavior occured: %s", panicErr)
		}
		scenario.audit.AuditScenarioStep(scenario.currentStep, stepTrace.String(), payload, err)
	}()

	stepTrace.WriteString("Validate that test is intended to pass; ")
	payload = struct {
		Ip string
	}{
		config.Vars.ServicePacks.Ubuntu.Ip,
	}
	session := GetSession()
	_, err1 := ConnectAndRunShellCmd("sudo ufw status | grep Status", *session)

	if err1 != nil {
		return err
	}
	session.Close()
	return errors.New("firewall is not able to connect to the VM")
}

func (scenario *scenarioState) ensureSSHRootLoginIsDisabled() error {
	// Standard auditing logic to ensures panics are also audited
	stepTrace, payload, err := utils.AuditPlaceholders()
	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = utils.ReformatError("[ERROR] Unexpected behavior occured: %s", panicErr)
		}
		scenario.audit.AuditScenarioStep(scenario.currentStep, stepTrace.String(), payload, err)
	}()

	stepTrace.WriteString("Validate that test is intended to pass; ")
	payload = struct {
		Ip string
	}{
		config.Vars.ServicePacks.Ubuntu.Ip,
	}
	session1 := GetSession()
	response, _ := ConnectAndRunShellCmd("dpkg -s sudo", *session1)

	fmt.Println("Response----------------------->", response)
	if response != "inactive" {
		return err
	}
	session1.Close()
	return errors.New("not able to connect to the VM")
}
