package analyzer

import (
	"testing"

	"github.com/nearform/gammaray/nodepackage"
	"github.com/nearform/gammaray/pathrunner"
	log "github.com/sirupsen/logrus"

	"github.com/google/go-cmp/cmp"
)

func TestHelloWorld(t *testing.T) {
	walkers := []nodepackage.Walker{
		pathrunner.PathRunner{},
	}
	vulns, err := Analyze("../test_data/hello-world", "", walkers...)
	if err != nil {
		panic(err)
	}
	numVulns := 0
	for provider, vulnList := range vulns {
		numVulns += len(vulnList)
		log.Print(provider, "> ", len(vulnList), " vulnerabilities:\n", vulnList)
	}
	if diff := cmp.Diff(numVulns, 0); diff != "" {
		t.Errorf("TestHelloWorld: vulnerabilities : (-got +want)\n%s", diff)
	}
}

func TestInsecureProject(t *testing.T) {
	walkers := []nodepackage.Walker{
		pathrunner.PathRunner{},
	}
	vulns, err := Analyze("../test_data/insecure-project", "", walkers...)
	if err != nil {
		panic(err)
	}

	for provider, vulnList := range vulns {
		providerVulns := len(vulnList)
		log.Print(provider, "> ", providerVulns, " vulnerabilities:\n", vulnList)
		// both OSSIndex and NodeSWG report bassmaster-1.0.0 and its dep hoek-1.5.2
		if diff := cmp.Diff(providerVulns, 2); diff != "" {
			t.Errorf("TestInsecureProject: %s vulnerabilities : (-got +want)\n%s", provider, diff)
		}
	}

}

func TestNotExistingProject(t *testing.T) {
	walkers := []nodepackage.Walker{
		pathrunner.PathRunner{},
	}
	_, err := Analyze("./does-not-exist", "", walkers...)
	if err == nil {
		t.Errorf("TestNotExistingProject: ./does-not-exist does not exist, it should not be analyzed !")
	} else if diff := cmp.Diff(err.Error(), "could not find any dependencies and all strategies to find them failed"); diff != "" {
		t.Errorf("TestNotExistingProject: err : (-got +want)\n%s", diff)
	}
}
