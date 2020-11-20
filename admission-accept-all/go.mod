module github.com/chimera-kube/chimera-samples/admission-accept-all

go 1.15

require (
	github.com/chimera-kube/chimera-admission-library v0.0.0-00010101000000-000000000000
	k8s.io/api v0.18.6
)

// TODO: remove when chimera-admission-library is open to the world
replace github.com/chimera-kube/chimera-admission-library => ../../chimera-admission-library
