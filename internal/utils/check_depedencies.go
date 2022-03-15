package utils

import (
	"os/exec"

	"github.com/NethermindEth/1Click/configs"
	log "github.com/sirupsen/logrus"
)

/*
CheckDependencies :
This function is responsible for checking if on-premise setup dependencies are installed on host machine

params :-
a. dependencies []string
List of dependencies to be checked

returns :-
a. []string
List of dependencies that are not installed
*/
func CheckDependencies(dependencies []string) (pending []string) {
	for _, dependency := range dependencies {
		_, err := exec.LookPath(dependency)
		if err != nil {
			log.Errorf(configs.DependencyNotInstalledError, dependency)
			pending = append(pending, dependency)
		}
	}
	return
}
