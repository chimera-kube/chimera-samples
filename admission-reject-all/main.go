package main

import (
	"log"
	"net/http"
	"os"

	"github.com/chimera-kube/chimera/pkg/chimera"

	admissionv1 "k8s.io/api/admission/v1"
	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
)

const (
	admissionPort = 8080
	admissionName = "reject-all.admission.rule"
)

var (
	admissionHost = os.Getenv("AW_CALLBACK_HOST")
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
						Operations: []admissionregistrationv1.OperationType{admissionregistrationv1.OperationAll},
						Rule: admissionregistrationv1.Rule{
							APIGroups:   []string{"*"},
							APIVersions: []string{"v1"},
							Resources:   []string{"*"},
						},
					},
				},
				Callback: func(*admissionv1.AdmissionRequest) (chimera.WebhookResponse, error) {
					return chimera.NewRejectRequest().WithCode(http.StatusTeapot).WithMessage("Very small, but still -- a teapot"), nil
				},
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
