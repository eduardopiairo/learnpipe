# CNOE: An Open-Source IDP Framework

Prerequisites:
- A Coder account for the PE Course instance
- Install [Coder CLI](https://coder.com/docs/install/cli)

## Setup
1 - Loggin to the PlatformEngineering.org Coder environment
```
coder login https://sandbox.platformengineering.org/
```

2- Forward a local port to the workspace

```
coder port-forward edpiairo-cnoe --tcp 8443:8443
```

3 - Get access to Coder instance itself. Start by running the following (in a second terminal): 
```
coder config-ssh
```

4 - Use the next command to actually login to the remote Coder-hosted machine
```
ssh coder.<CODER-WORKSPACE-NAME>
ssh edpiairo-cnoe.coder
```


The Coder setup has already set up the reference implementation for Coder and set up URLs to access each:
- Argo CD: https://cnoe.localtest.me:8443/argocd
- Argo Workflows: https://cnoe.localtest.me:8443/argo-workflows
- Backstage: https://cnoe.localtest.me:8443/
- Gitea: https://cnoe.localtest.me:8443/gitea
- Keycloak: https://cnoe.localtest.me:8443/keycloak/admin/master/console/


## Get the Secrets
```
idpbuilder get secrets
```

## Step 1: Simple Deployment with Backstage


## Step 2: Adding a New Resource

