# Authentik experimentation

Update: Ruled out use of Authentik because it lack support for a simple, one-line podman startup variation as well as any support for `grant_type=password`

Goal of authentik is to get an OIDC/LDAP provider service up as quickly as possible in a container for development and testing purposes (not necessarily as a production authentication service, which it can obviously also do).

## Motivation

Kubernetes authentication with OIDC connect is pretty common in the enterprise for clusters that support multiple users developing disparate applications but simulating that authentication quickly and locally is not always easy. Often it includes setting up and supporting a full LDAP server which powers an OIDC provider with OAuth2 like Keycloak. Such services are bulky and hard to manage when attempting to experiment and monitoring what is going on while doing that development even harder. For example, ensuring that tokens and refresh expirations are working as desired is not obvious. Therefore, having a quick development environment that emulates these things seems to have value for people supporting and developing for such situations.

## Related

* <https://goauthentik.io>
