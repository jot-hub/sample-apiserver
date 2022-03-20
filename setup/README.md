# Kubernetes-style API server for everything else
The [Kubernetes API server](https://github.com/kubernetes/apiserver) is a component of the Kubernetes control plane that exposes the Kubernetes API. The API server is the front end for the Kubernetes control plane.

Kubernetes API server is serving Kubernetes resources and it has the ability to support custom resources(Custom Resource Definitions). Kubernetes API server accepts schema definitions in OpenAPIV3. A CRD(CustomResourceDefinition) corresponding to a businessobject (say, customer) can be setup like shown [here](https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/#create-a-customresourcedefinition).

It is worthwhile to explore the K8s API server (standalone, independent of other K8s components) as one of the options to build a general purpose API server for real-world/business/enterprise APIs.

This experiment could be divided into the following steps:

1. Running a "standalone" K8s API server (with none of the other K8s components) 
2. Setting up CRDs to represent business objects in that standalone K8s API server

## Standalone Kubernetes-style API server


comments:
CRD is a kind that is not available in sample-apiserver - only flunders and fischers are available
Try to include the apiextensions server in the sample-apiserver setup:
1. This will help setup any object based on CustomResourceDefinition handled by the apiextensions server

createServerChain - sets up an apiextensions server, builds the generic/core kube-apiserver with it and adds aggregated api server on top of it

postponed: setting up oidc instead of --client-ca-file









