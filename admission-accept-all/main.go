package main

import (
	"log"
	"os"

	"github.com/chimera-kube/chimera-admission-library/pkg/chimera"

	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
)

const (
	admissionPort = 8080
	admissionName = "accept-all.admission.rule"
)

var (
	admissionHost = os.Getenv("CHIMERA_SAMPLES_CALLBACK_HOST")
)

func main() {
	err := chimera.StartTLSServer(
		chimera.AdmissionConfig{
			Name:         admissionName,
			CallbackHost: admissionHost,
			CallbackPort: admissionPort,
			Webhooks: chimera.WebhookList{
				{
					Rules: []admissionregistrationv1.RuleWithOperations{
						{
							Operations: []admissionregistrationv1.OperationType{admissionregistrationv1.OperationAll},
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
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
