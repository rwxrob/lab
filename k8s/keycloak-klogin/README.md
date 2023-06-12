# Desktop (containers) for developing klogin binary against Keycloak OIDC provider

## Goal

Create a safe (containerized) sandbox that emulates an enterprise in order to create and develop applications with a custom root certificate authority, intermediate CA, and traditional, domain-based server certificates.

## Steps

1. Register a domain name pointing to a LAN ip address
    * Example: `home.rwx.gg -> 192.168.1.200`
    * Add to `/etc/hosts` for tab completion
1. Create root CA
1. Create intermediate CA
1. Create a server certificate for Keycloak
1. Configure TLS in Keycloak
    1. Mount share into container
    1. Run container with volume mounts and TLS arguments
1. Add certificates to required "local" keyrings
1. Configure minikube to use TLS certs and OIDC
1. Code `klogin` to update `~/.kube/config`


Related:

* Setting up a development containerized workspace  
  <https://link.excalidraw.com/l/6rjSvoGlOkc/7nG4ZhdXCZA>
* Kind, Keycloak --- Securing Kubernetes api server with OIDC \| by Charles-Edouard Brétéché \| Medium  
  <https://medium.com/@charled.breteche/kind-keycloak-securing-kubernetes-api-server-with-oidc-371c5faef902>
* How to Secure Your Kubernetes Cluster with OpenID Connect and RBAC \| Okta Developer  
  <https://developer.okta.com/blog/2021/11/08/k8s-api-server-oidc>
* Subscribe to our RSS feed or Email newsletter  
  <https://www.redhat.com/sysadmin/podman-inside-container>
* minikube  
  <https://minikube.sigs.k8s.io/docs/tutorials/openid_connect_auth/>
* How to Secure Keycloak with HTTPS - Mastertheboss  
  <https://www.mastertheboss.com/keycloak/secure-keycloak-with-https/>
* Private CA Part 1: Building your own root and intermediate certificate authority - Artiom\'s Sketchpad  
  <https://www.flexlabs.org/2019/07/private-ca-1-building-root-and-intermediate-ca>
* Configuring TLS - Keycloak  
  <https://www.keycloak.org/server/enabletls>
* Generate self-signed certificate with a custom root CA - Azure Application Gateway \| Microsoft Learn  
  <https://learn.microsoft.com/en-us/azure/application-gateway/self-signed-certificates>
* <https://gist.github.com/fntlnz/cf14feb5a46b2eda428e000157447309>
* Generate self-signed certificate with a custom root CA - Azure Application Gateway \| Microsoft Learn  
  <https://learn.microsoft.com/en-us/azure/application-gateway/self-signed-certificates>

