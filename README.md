# Cloud Platform Project

This project provide Web based coding solution for small companies / labs. It orchestras WebIDE containers, shared volumes, etc. in a Kubernetes cluster. Here is the navigation of the project:

| Repo                                                                                  | Description |
| ------------------------------------------------------------------------------------- | ----------- |
| [cloud-platform-apiserver](https://github.com/davidliyutong/cloud-platform-apiserver) | API server  |
| [cloud-platform-frontend](https://github.com/davidliyutong/cloud-platform-frontend)   | Frontend    |

## Get-Started

### Prerequisites

The Cloud Platform Project need a fully functional Kubernetes cluster with the following components:

-   A Ingress controller, such as [Nginx Ingress Controller](https://kubernetes.github.io/ingress-nginx/deploy/), to handle traffic
-   A storage provisioner, such as [Longhorn](https://longhorn.io/), to provide storage to workloads
-   A configured `clusterIssuer`, [Cert-Manager](https://cert-manager.io/) to generate certificates for Ingress resources.

To deploy the project, the admin must also own a **valid domain** and properly configure its DNS resolution. Some hosts of this domain are occupied by this project:

-   `*.coder.domain.com` for WebIDE
-   `*.vnc.domain.com` for VNC connection
-   `*.ssh.domain.com` for ssh connection
-   `clpl.domain.com` for frontend
-   `apiserver.domain.com` for backend

### Deploy the Project

The project is deployed using [Kustomize](https://kustomize.io/). To deploy the project, the admin must first download the project's manifests:

```bash
git clone https://github.com/davidliyutong/cloud-platform-project.git
cd cloud-platform-project
```

### Prepare the Cluster

In case there is no existing Kubernetes cluster, we provide a list of Ansible playbooks to help the admin deploy a Kubernetes cluster with:

-   Rocky Linux 9 as OS
-   Rancher Kubernetes Engine 2 as K8S distribution
-   Longhorn as storage provisioner
-   Cert-Manager as certificate management system

| Name                                                                               | Description                        | Notice                                                                       |
| ---------------------------------------------------------------------------------- | ---------------------------------- | ---------------------------------------------------------------------------- |
| [init_base_server.yaml](manifests/ansible/playbooks/init_base_server.yaml)         | Install a base server suitable for | The playbook is based on Rocky Linux / AlmaLinux / CentOS, tested on Rocky 9 |
| [install_rke2_server.yaml](manifests/ansible/playbooks/install_rke2_server.yaml)   | Install the Server node            |                                                                              |
| [install_rke2_agent.yaml](manifests/ansible/playbooks/install_rke2_agent.yaml)     | Install the Agent node             | The server node should be installed and launched first.                      |
| [install_longhorn.yaml](manifests/ansible/playbooks/install_longhorn.yaml)         | Install Longhorn                   |                                                                              |
| [install_cert_manager.yaml](manifests/ansible/playbooks/install_cert_manager.yaml) | Install Cert-Manager               |                                                                              |

For example, you can launch a playbook via :

```shell
ansible-playbook -i <path_to_inventory> <path_to_playbook>
```

> Most of playbook will prompt for input and ask to select a target host.

### Helm Installation

Currently, this project only support installation via Helm. To install this project, make sure Helm is installed and Kubectl is properly configured.

First, create the namespace `clpl`, which is namespace for project

```shell
kubectl create ns clpl
```

Create CustomResourceDefinitions

```shell
kubectl apply -f helm/crds/oidcconfigs.yaml
kubectl apply -f helm/crds/siteconfigs.yaml
```

Create a `value.yaml` by copying the template:

```shell
cd cloud-platform-project
cp helm/values.yaml.template helm/values.yaml
```

Modify the `helm/values.yaml` to fit your needs. You can verify the installation by `helm install cloud-platform ./helm --dry-run --debug --namespace=clpl`

**Basic configuration:**

| Path                                  | Description                                    | Type    | Default Value        | Supported Values                                                   |
| ------------------------------------- | ---------------------------------------------- | ------- | -------------------- | ------------------------------------------------------------------ |
| `namespace`                           | The namespace where the resources are deployed | string  | `default`            | Any valid Kubernetes namespace                                     |
| `rbac.create`                         | Whether to create RBAC resources               | boolean | `true`               | `true`, `false`                                                    |
| `rbac.serviceAccount.name`            | The name of the service account to use         | string  | `"clpl-admin"`       | Any valid Kubernetes service account name                          |
| `tls.label`                           | Label for the TLS                              | string  | `"speit-site"`       | Any string                                                         |
| `tls.domain`                          | Domain for the TLS                             | string  | `"speit.site"`       | Any valid domain name                                              |
| `tls.issuer.name`                     | Name of the TLS issuer                         | string  | `letsencrypt-cf-dns` | Any valid issuer name                                              |
| `tls.issuer.kind`                     | Kind of the TLS issuer                         | string  | `ClusterIssuer`      | `ClusterIssuer`, `Issuer`                                          |
| `cloudflareIssuer.create`             | Whether to create Cloudflare issuer            | boolean | `false`              | `true`, `false`                                                    |
| `cloudflareIssuer.name`               | Name of the Cloudflare issuer                  | string  | `letsencrypt-cf-dns` | Any valid issuer name                                              |
| `cloudflareIssuer.apiToken`           | API token for the Cloudflare issuer            | string  | `"your-api-token"`   | Any valid API token, the token should be able to edit DNS settings |
| `cloudflareIssuer.apiTokenSecretName` | Secret name of the Cloudflare issuer API token | string  | `"cloudflare-api"`   | Any valid Kubernetes secret name                                   |
| `cloudflareIssuer.email`              | Email associated with the Cloudflare issuer    | string  | `null`               | Any valid email address                                            |
| `ingress.enabled`                     | Whether to enable ingress                      | boolean | `true`               | `true`, `false`                                                    |
| `ingress.className`                   | Class name for the ingress                     | string  | `"nginx"`            | Any valid class name                                               |
| `ingress.annotations`                 | Annotations for the ingress                    | map     | `{}`                 | Any valid annotations                                              |

**Database:**

| Path                                | Description                                  | Type    | Default Value             | Supported Values                  |
| ----------------------------------- | -------------------------------------------- | ------- | ------------------------- | --------------------------------- |
| `datastore.create`                  | Whether to create datastore                  | boolean | `true`                    | `true`, `false`                   |
| `datastore.name`                    | Name of the datastore                        | string  | `clpl-mongodb`            | Any valid datastore name          |
| `datastore.imageRef`                | Image reference for the datastore            | string  | `"docker.io/mongo:6.0.6"` | Any valid image reference         |
| `datastore.imagePullPolicy`         | Image pull policy for the datastore          | string  | `IfNotPresent`            | `Always`, `IfNotPresent`, `Never` |
| `datastore.resources.limits.cpu`    | CPU limit for the datastore                  | string  | `"1000m"`                 | Any valid CPU limit               |
| `datastore.resources.limits.memory` | Memory limit for the datastore               | string  | `"1024Mi"`                | Any valid memory limit            |
| `datastore.credentials.username`    | Username for the datastore                   | string  | `"clpl"`                  | Any valid username                |
| `datastore.credentials.password`    | Password for the datastore                   | string  | `"clpl"`                  | Any valid password                |
| `datastore.volume.create`           | Whether to create a volume for the datastore | boolean | `true`                    | `true`, `false`                   |
| `datastore.volume.pvc`              | Persistent volume claim for the datastore    | string  | `clpl-mongodb-data`       | Any valid PVC name                |
| `datastore.volume.storage`          | Storage size for the datastore volume        | string  | `"10Gi"`                  | Any valid storage size            |
| `datastore.volume.storageClassName` | Storage class for the datastore volume       | string  | `null`                    | Any valid storage class name      |
| `datastore.svc.name`                | Name of the datastore service                | string  | `clpl-mongodb`            | Any valid service name            |
| `datastore.svc.port`                | Port of the datastore service                | integer | `27017`                   | Any valid port number             |

**Apiserver:**

| Path                                | Description                                     | Type    | Default Value                              | Supported Values                  |
| ----------------------------------- | ----------------------------------------------- | ------- | ------------------------------------------ | --------------------------------- |
| `apiserver.name`                    | Name of the API server                          | string  | `clpl-apiserver`                           | Any valid server name             |
| `apiserver.imageRef`                | Image reference for the API server              | string  | `"docker.io/davidliyutong/clpl-apiserver"` | Any valid image reference         |
| `apiserver.imageTag`                | Image tag for the API server                    | string  | `latest`                                   | Any valid image tag               |
| `apiserver.imagePullPolicy`         | Image pull policy for the API server            | string  | `Always`                                   | `Always`, `IfNotPresent`, `Never` |
| `apiserver.resources.limits.cpu`    | CPU limit for the API server                    | string  | `"1000m"`                                  | Any valid CPU limit               |
| `apiserver.resources.limits.memory` | Memory limit for the API server                 | string  | `"1024Mi"`                                 | Any valid memory limit            |
| `apiserver.volume.create`           | Whether to create a volume for the API server   | boolean | `true`                                     | `true`, `false`                   |
| `apiserver.volume.pvc`              | Persistent volume claim for the API server      | string  | `clpl-apiserver-log`                       | Any valid PVC name                |
| `apiserver.volume.storage`          | Storage size for the API server volume          | string  | `"5Gi"`                                    | Any valid storage size            |
| `apiserver.volume.storageClassName` | Storage class for the API server volume         | string  | `null`                                     | Any valid storage class name      |
| `apiserver.svc.name`                | Name of the API server service                  | string  | `clpl-apiserver`                           | Any valid service name            |
| `apiserver.svc.port`                | Port of the API server service                  | integer | `8080`                                     | Any valid port number             |
| `apiserver.env.api.numWorkers`      | Number of workers for the API server            | integer | `4`                                        | Number of cores                   |
| `apiserver.env.api.debug`           | Whether to enable debug mode for the API server | boolean | `false`                                    | `true`, `false`                   |
| `apiserver.env.config.tokenSecret`  | Secret token for the API server configuration   | string  | `"top-secret"`                             | Any valid secret token            |

**OICD:**

| Path                                  | Description                            | Type    | Default Value                                  | Supported Values                                 |
| ------------------------------------- | -------------------------------------- | ------- | ---------------------------------------------- | ------------------------------------------------ |
| `apiserver.env.config.useOIDC`        | Whether to use OIDC for the API server | boolean | `false`                                        | `true`, `false`                                  |
| `apiserver.env.oidc.baseURL`          | The base URL for OIDC                  | string  | `https://authentik.example.com/application/o`  | OIDC provider URL                                |
| `apiserver.env.oidc.clientID`         | The client ID for OIDC                 | string  | `xxx`                                          | OIDC client ID                                   |
| `apiserver.env.oidc.clientSecret`     | The client secret for OIDC             | string  | `yyy`                                          | OIDC client secret                               |
| `apiserver.env.oidc.frontendLoginURL` | The frontend login URL for OIDC        | string  | `http://127.0.0.1:8081/login`                  | **Login URL of the frontend**                    |
| `apiserver.env.oidc.name`             | The name for OIDC                      | string  | `cloud-platform`                               | OIDC name                                        |
| `apiserver.env.oidc.redirectURL`      | The redirect URL for OIDC              | string  | `http://127.0.0.1:8081/v1/auth/oidc/authorize` | **Token process api of apiserver**               |
| `apiserver.env.oidc.scopes`           | The scopes for OIDC                    | string  | `openid+profile+email`                         | OIDC scopes, `+` as seperator                    |
| `apiserver.env.oidc.userFilter`       | The user filter for OIDC               | string  | `{}`                                           | Mongo-style user filter                          |
| `apiserver.env.oidc.userInfoPath`     | The user info path for OIDC            | string  | `$`                                            | Json-style user info path                        |
| `apiserver.env.oidc.userNamePath`     | The user name path for OIDC            | string  | `preferred_username`                           | Json-style user name path, relative to user info |
| `apiserver.env.oidc.emailPath`        | The email path for OIDC                | string  | `email`                                        | Json-style email path, relative to user info     |

**Frontend:**

| Path                               | Description                                  | Type    | Default Value                           | Supported Values                  |
| ---------------------------------- | -------------------------------------------- | ------- | --------------------------------------- | --------------------------------- |
| `frontend.name`                    | Name of the frontend                         | string  | `clpl-frontend`                         | Any valid service name            |
| `frontend.replicaCount`            | The number of replicas for the frontend      | integer | `1`                                     | Any positive integer              |
| `frontend.imageRef`                | Docker image reference for the frontend      | string  | `docker.io/davidliyutong/clpl-frontend` | Docker image to use               |
| `frontend.imageTag`                | The tag of the Docker image for the frontend | string  | `latest`                                | Any valid Docker image tag        |
| `frontend.imagePullPolicy`         | The pull policy for the Docker image         | string  | `Always`                                | `Always`, `IfNotPresent`, `Never` |
| `frontend.resources.limits.cpu`    | The CPU limit for the frontend               | string  | `1000m`                                 | Any valid CPU limit               |
| `frontend.resources.limits.memory` | The memory limit for the frontend            | string  | `1024Mi`                                | Any valid memory limit            |
| `frontend.svc.name`                | The name of the frontend service             | string  | `clpl-frontend`                         | Any valid service name            |
| `frontend.svc.port`                | The port of the frontend service             | integer | `80`                                    | Any valid port number             |
| `frontend.env.coreDNS`             | The CoreDNS configuration for the frontend   | string  | `10.43.0.10`                            | **Cluster CoreDNS service IP**    |

Once everything is settled, install the project:

```shell
helm install cloud-platform ./helm --namespace=clpl
```

### Helm Uninstallation

```shell
helm uninstall cloud-platform ./helm --namespace=clpl
```

### Helm Upgratioin

```shell
helm upgrade cloud-platform ./helm --namespace=clpl
```

## Mantenance tools

A maintenance tool is created under `tools/maintenance`. The tool is written in Go, and is designed to facilatate the maintenance of the project. It can be built with following commands:

```shell
cd tools/maintenance
go build main.go -o maintenance
```

Run `./maintenance --help` to get help about this tool

## Guides

### User Guide

See [docs/manuals](docs/manuals/)

### Develop Guide

See [docs/develop](docs/develop)
