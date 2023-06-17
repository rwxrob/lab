# Desktop (containers) for developing klogin binary against Keycloak OIDC provider

## Goal

Create a safe (containerized) sandbox that emulates an closed enterprise network in order to create and develop applications with a custom root certificate authority, intermediate CA, and traditional, domain-based server certificates.

## Target containers

* Simulated "offline" linux host to create and manage root and interm CA. (`ca`)
* Keycloak configured to use server cert (with keyring updates as well). (`keycloak`)
* Simulated Go development box with keyring updated with CA certs. (`dev`)

## Steps

1. Register a domain name pointing to a LAN ip address
    * Example: `home.rwx.gg -> 192.168.1.200`
    * Add to `/etc/hosts` for tab completion
1. Create root CA
1. Create intermediate CA
1. Create a server certificate for Keycloak
1. Configure TLS in Keycloak
    1. Mount share into container
    1. Run container with volume mounts and TLS arguments ([`startkeycloak`](startkeycloak))
1. Add root and intermediate CA to Chrome 
1. Add certificates to required "local" keyrings
1. Configure minikube to use TLS certs and OIDC
1. Code `klogin` to update `~/.kube/config`

## Add root and intermediate CA to Chrome

1. Copy the `crt` and `key` files onto the machine with Chrome on Windows.
1. Click on top right button and select **Settings** from menu.
1. Click on **Privacy and Security** on list on the left.
1. Click on **Security** in the middle of page.
1. Scroll down to **Manage Device Certificates** and click.
1. A system's dialog window will open.
1. Click on **Trusted Root CA** tab.
1. Click on **Import** and click **Next* on the Welcome screen.
1. Click **Browse** to find `crt` file.
1. Validate new CA is in the list.

Repeat the same steps but do it for **Intermediate CA** tab.

After finishing, confirm that the domain now shows as "Secure" when visiting pages served from it.

## Add `home.rwx.gg.crt` to root CA certs on `localhost`

On Ubuntu:

```sh
sudo cp /shared/home.rwx.gg.crt /usr/local/share/ca-certificates
sudo update-ca-certificates
```

After that the `--cacert` argument to `curl` is no longer needed.

Also see [`gettoken`](gettoken) for the logic to detect if root CA is loaded onto the system or not and if not use a simulated embedded file.

## Creating a root CA TLS HTTP client request in Go

Go will use the system TLS certs if there is no TLS CertPool added.

* <https://gist.github.com/michaljemala/d6f4e01c4834bf47a9c4>

```golang
  const caCert = `...`
  caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	transport := &http.Transport{TLSClientConfig: &tls.Config{RootCAs: caCertPool}}
	client := &http.Client{Transport: transport}
```

## Start minikube using own root CA with OIDC connect authentication

1. TODO 

## Related

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
* `certgen` script  
  <https://github.com/aurae-runtime/aurae/blob/main/hack/certgen>
* <https://www.digitalocean.com/community/tutorials/how-to-set-up-and-configure-a-certificate-authority-ca-on-ubuntu-20-04#prerequisites>
* Minimal CA + Cert script generator  
  <https://github.com/dexidp/dex/blob/master/examples/k8s/gencert.sh>
* OIDC Login to Kubernetes and Kubectl with Keycloak  
  <https://www.talkingquickly.co.uk/setting-up-oidc-login-kubernetes-kubectl-with-keycloak>
* <https://kubernetes.io/docs/reference/access-authn-authz/authentication/#openid-connect-tokens>
* <https://ahmet.im/blog/mastering-kubeconfig/>
* <https://kubernetes.io/docs/reference/config-api/kubeconfig.v1/>
