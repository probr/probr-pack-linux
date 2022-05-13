package access

import (
	"errors"
	"fmt"
	"strings"

	"github.com/probr/probr-pack-ubuntu/internal/config"
	"github.com/probr/probr-sdk/utils"
)

func (scenario *scenarioState) givenAnUbuntuVMMustBeUp() error {
	// Standard auditing logic to ensures panics are also audited
	_, _, err := utils.AuditPlaceholders()
	return err
}

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
	response, err := ConnectAndRunShellCmd("sudo ufw status | grep Status", *session)
	fmt.Println("Response====>", response)
	if response == "inactive\n" {
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
	response, _ := ConnectAndRunShellCmd("dpkg -s sudo | grep Status", *session1)

	fmt.Println("Response----------------------->", response)
	if response == "Status: install ok installed\n" {
		return err
	}
	session1.Close()
	return errors.New("not able to connect to the VM")
}

func (scenario *scenarioState) runCommandAndVerify(command, output string) error {
	// Supported values for 'response':
	//	'welcomed'
	//	'rejected'

	// Standard auditing logic to ensures panics are also audited
	stepTrace, payload, err := utils.AuditPlaceholders()
	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = utils.ReformatError("[ERROR] Unexpected behavior occured: %s", panicErr)
		}
		scenario.audit.AuditScenarioStep(scenario.currentStep, stepTrace.String(), payload, err)
	}()

	session1 := GetSession()
	response, _ := ConnectAndRunShellCmd(command, *session1)

	fmt.Println("Response----------------------->", response)
	if !strings.Contains(response, output) {
		return errors.New("not compliant")
	}

	fmt.Println(command, "-", output)

	// Validate input values

	return err
}
