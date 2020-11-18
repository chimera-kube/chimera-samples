# chimera-samples

`chimera-samples` is a collection of examples using the [`chimera`
library](https://github.com/chimera-kube/chimera) to write admission
webhooks.

The following examples can be found in this project:

- `admission-accept-all`: an admission controller that accepts every
  request it has been registered against.

- `admission-reject-all`: an admission controller that rejects every
  request it has been registered against, returning a 418 error and a
  custom message as the admission review object.

As validation webhooks they are not very useful, but it shows how easy
it is to create a main program that can use `chimera` as a
library. The code that validates requests and objects can be shared
among your real webhook and this one, so it's really easy to perform
end to end tests of your webhooks integrated with a Kubernetes
cluster.
