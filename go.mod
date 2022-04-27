module github.com/probr/probr-pack-ubuntu

go 1.14

require (
	github.com/cucumber/godog v0.11.0
	github.com/hashicorp/go-hclog v0.15.0 // indirect
	github.com/markbates/pkger v0.17.1
	github.com/probr/probr-sdk v0.1.6
	//home//hemantruhela//go//src//probr-sdk v0.1.6
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
)

// For Development Only
// replace github.com/probr/probr-sdk => ../probr-sdk

//replace github.com/probr/probr-sdk => ../probr-sdk
