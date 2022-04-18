// Package welcome provides the implementation required to execute the BDD tests described in container_registry_access.feature file
package ubuntu

import (
	"bytes"
	"fmt"
	"net"
	"strings"

	"github.com/cucumber/godog"
	"github.com/probr/probr-pack-ubuntu/internal/summary"
	audit "github.com/probr/probr-sdk/audit"
	"github.com/probr/probr-sdk/probeengine"
	"golang.org/x/crypto/ssh"
)

type probeStruct struct{}

// scenarioState holds the steps and state for any scenario in this probe
type scenarioState struct {
	name        string
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
	return "ubuntu"
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
		scenario.name = s.Name
		scenario.probe = summary.State.GetProbeLog(probe.Name())
		scenario.audit = summary.State.GetProbeLog(probe.Name()).InitializeAuditor(s.Name, s.Tags)
		probeengine.LogScenarioStart(s)
	})

	ctx.BeforeStep(func(st *godog.Step) {
		scenario.currentStep = st.Text
	})

	// Background
	//ctx.Step(`^GNOME Display Manager is disabled$`, scenario.gNomeDisplayManager)
	ctx.Step(`^an Ubuntu VM must be up$`, scenario.gNomeDisplayManager())
	ctx.Step(`^gnome display Manager is disabled$`, scenario.gNomeDisplayManager())

	ctx.AfterStep(func(st *godog.Step, err error) {
		scenario.currentStep = ""
	})

	ctx.AfterScenario(func(s *godog.Scenario, err error) {
		probeengine.LogScenarioEnd(s)
	})

}

func ConnectAndRunShellCmd(command, username, password, hostname, port string) string {

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	fmt.Println("\nConnecting to ", hostname, port)

	hostaddress := strings.Join([]string{hostname, port}, ":")
	client, err := ssh.Dial("tcp", hostaddress, config)
	if err != nil {
		panic(err.Error())
	}
	session, err := client.NewSession()
	if err != nil {
		panic(err.Error())
	}
	defer session.Close()
	fmt.Println("To exit this program, hit Control-C")
	fmt.Printf("Enter command to execute on %s : ", hostname)
	fmt.Println("Executing command ", command)
	var buff bytes.Buffer
	session.Stdout = &buff
	if err := session.Run(command); err != nil {
		panic(err.Error())
	}

	return buff.String()

}
