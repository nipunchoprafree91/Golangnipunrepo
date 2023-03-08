package Docker

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func CreateDocker(dockerfile, tag string) error {
	fmt.Println("Deleting docker ,Image from base Image", dockerfile)
	cmd := exec.Command("docker", "image", "rm", "-f", tag)
	fmt.Printf("command Executed is %v\n\n", cmd)

	stdoutStderr, err := cmd.CombinedOutput()
	fmt.Printf("command output %s", stdoutStderr)

	if err != nil {

		log.Fatalf("Aborting docker creation with Error:%v", err)
		return err
	}

	fmt.Println("Creating docker ,Image from base Image", dockerfile)
	cmd = exec.Command("docker", "build", "-f", dockerfile, "-t", tag, ".")
	fmt.Printf("command Executed is %v\n\n", cmd)

	stdoutStderr, err = cmd.CombinedOutput()
	fmt.Printf("command output %s", stdoutStderr)

	if err != nil {

		log.Fatalf("Aborting docker creation with Error:%v", err)
		return err
	}

	tagaftersplit := strings.Split(tag, ":")
	nameaftersplit := tagaftersplit[0]
	label := "reference=" + nameaftersplit

	cmd = exec.Command("docker", "image", "ls", "--filter", label)
	fmt.Printf("command Executed is %v\n\n", cmd)
	stdoutStderr, err = cmd.CombinedOutput()

	fmt.Printf("command output %s\n\n", stdoutStderr)
	if err != nil {
		log.Fatalf("Error %v while printing docker Image list", err)
		return err
	}

	return nil
}
