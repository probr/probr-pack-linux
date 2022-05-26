// Package welcome provides the implementation required to execute the BDD tests described in container_registry_access.feature file
package group

import (
	"bytes"
	"fmt"
	"net"
	"strings"

	"github.com/cucumber/godog"
	conf "github.com/probr/probr-pack-ubuntu/internal/config"
	"github.com/probr/probr-pack-ubuntu/internal/summary"
	audit "github.com/probr/probr-sdk/audit"
	"github.com/probr/probr-sdk/probeengine"
	"golang.org/x/crypto/ssh"
)

type probeStruct struct{}

// scenarioState holds the steps and state for any scenario in this probe
type scenarioState struct {
	name string
	//Session     *ssh.Session
	currentStep string
	audit       *audit.Scenario
	probe       *audit.Probe
}

// Probe meets the service pack interface for adding the logic from this file
var Probe probeStruct
var scenario scenarioState

// Name presents the name of this probe for external reference
func (probe probeStruct) Name() string {
	// The return value for `Name` should match the probe directory
	// and it's feature file, so each may be properly addressed for
	// packing and opening the files.
	return "group"
}

// Path presents the path of these feature files for external reference
func (probe probeStruct) Path() string {
	// this should reference the probe parent directory (usually `internal/<probe-name>`)
	return probeengine.GetFeaturePath("internal", probe.Name())
}

// ProbeInitialize handles any overall Test Suite initialisation steps.  This is registered with the
// test handler as part of the init() function.
func (probe probeStruct) ProbeInitialize(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {

	})

	ctx.AfterSuite(func() {
	})
}

// ScenarioInitialize provides initialization logic before each scenario is executed
func (probe probeStruct) ScenarioInitialize(ctx *godog.ScenarioContext) {
	ctx.BeforeScenario(func(s *godog.Scenario) {

		beforeScenario(&scenario, probe.Name(), s)
	})

	ctx.BeforeStep(func(st *godog.Step) {
		scenario.currentStep = st.Text
	})

	// Background
	ctx.Step("^an Ubuntu VM must be up$", scenario.givenAnUbuntuVMMustBeUp)
	ctx.Step(`^run "([^"]*)" and then output is  "([^"]*)"$`, scenario.runCommandAndVerify)

	//^run "([^"]*)" and then output is  "([^"]*)"$

	ctx.AfterStep(func(st *godog.Step, err error) {
		scenario.currentStep = ""
	})

	ctx.AfterScenario(func(s *godog.Scenario, err error) {
		//scenario.Session.Close()
		probeengine.LogScenarioEnd(s)
	})

}

func ConnectAndRunShellCmd(command string, session ssh.Session) (string, error) {
	fmt.Println("Executing command ", command)
	var buff bytes.Buffer
	session.Stdout = &buff
	err := session.Run(command)
	if err != nil {
		fmt.Println("Error=======>", err)
		return "", err
	}

	return buff.String(), nil
}

func beforeScenario(s *scenarioState, probeName string, gs *godog.Scenario) {
	s.name = gs.Name
	//s.Session = GetSession()
	s.probe = summary.State.GetProbeLog(probeName)
	s.audit = summary.State.GetProbeLog(probeName).InitializeAuditor(s.name, gs.Tags)
	probeengine.LogScenarioStart(gs)
}

func GetSession() *ssh.Session {
	config := &ssh.ClientConfig{
		User: conf.Vars.ServicePacks.Ubuntu.Username,
		Auth: []ssh.AuthMethod{ssh.Password(conf.Vars.ServicePacks.Ubuntu.Password)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	hostaddress := strings.Join([]string{conf.Vars.ServicePacks.Ubuntu.Ip, conf.Vars.ServicePacks.Ubuntu.Port}, ":")
	client, err := ssh.Dial("tcp", hostaddress, config)
	if err != nil {
		fmt.Println("Error---------------->", err)
		panic(err.Error())
	}
	session, err := client.NewSession()
	if err != nil {
		panic(err.Error())
	}
	return session

}
