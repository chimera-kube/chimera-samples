package main

import (
	"log"
	"os"

	"github.com/chimera-kube/chimera/pkg/chimera"

	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
)

const (
	admissionPort = 8080
	admissionName = "accept-all.admission.rule"
)

var (
	admissionHost = os.Getenv("ADMISSION_CALLBACK_HOST")
)

func main() {
	err := chimera.StartServer(
		admissionName,
		admissionHost,
		admissionPort,
		[]chimera.Webhook{
			{
				Rules: []admissionregistrationv1.RuleWithOperations{
					{
						Operations: []admissionregistrationv1.OperationType{"*"},
						Rule: admissionregistrationv1.Rule{
							APIGroups:   []string{"*"},
							APIVersions: []string{"v1"},
							Resources:   []string{"*"},
						},
					},
				},
				Callback: chimera.AllowRequest,
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
