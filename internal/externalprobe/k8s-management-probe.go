package externalprobe

import (
	"errors"
	"fmt"
	"github.com/conplementAG/k8s-semantic-detective/pkg/common/commands"
	"github.com/conplementAG/k8s-semantic-detective/pkg/common/fileprocessing"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"
)

func Probe(namespace string) {
	startTime := time.Now()
	testTemplate := renderTestConfigMap(time.Now().Format("20060102150405"), namespace)
	result := apply(testTemplate)

	if getApplyResult(result) == Failed {
		panic(errors.New(result))
	}

	fmt.Printf("management_success_latency_internal %f\n", float64(time.Since(startTime).Milliseconds()))
}

func renderTestConfigMap(timestamp string, namespace string) string {
	result := strings.Replace(TestConfigMap, "{{ timestamp }}", timestamp, -1)
	return strings.Replace(result, "{{ namespace }}", namespace, -1)
}

func getApplyResult(result string) ApplyResults {
	var pattern = regexp.MustCompile(`\bcreated|unchanged|configured\b`)
	found := pattern.FindString(result)

	if len(found) == 0 {
		return Failed
	} else {
		switch found {
		case "created":
			return Created
		case "unchanged":
			return Unchanged
		case "configured":
			return Configured
		default:
			panic("Unknown result found, probe status could not be determined. Result: " + result)
		}
	}
}

func apply(content string) string {
	temporaryFile := fileprocessing.WriteStringToTemporaryFile(content, "resource.yaml")
	defer fileprocessing.DeletePath(temporaryFile)

	command := fmt.Sprintf("kubectl apply -f %s --server=%s --token=%s --certificate-authority=%s", temporaryFile, getServer(), getToken(), getCert())

	// we dont need the console out debug messages, so we use the overload method here to prevent any stdout output
	return commands.ExecuteCommandWithSecretContents(commands.Create(command))
}

func getToken() string {
	token, _ := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/token")
	return string(token)
}

func getCert() string {
	return "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
}

func getServer() string {
	return fmt.Sprintf("https://%s:%s", os.Getenv("KUBERNETES_SERVICE_HOST"), os.Getenv("KUBERNETES_SERVICE_PORT_HTTPS"))
}
