# PocketIDP: IDP with a Platform Orchestrator at the centre

Prerequisites:
- A Coder account for the PE Course instance
- Install [Coder CLI](https://coder.com/docs/install/cli)
- Install [Humanitec CLI](https://developer.humanitec.com/app-humanitec-io/docs/platform-orchestrator/cli/)

## What will be used
- Coder
- Humanitec
- Score
- Gitea / Backstage

## Setup
1 - Loggin to the PlatformEngineering.org Coder environment
```
coder login https://sandbox.platformengineering.org/
```

2- Forward a local port to the workspace

```
coder port-forward <CODER-WORKSPACE-NAME> --tcp 8443:443
coder port-forward edpiairo-pocketidp --tcp 8443:443
```

3 - Get access to Coder instance itself. Start by running the following (in a second terminal): 
```
coder config-ssh
```

4 - Use the next command to actually login to the remote Coder-hosted machine
```
ssh coder.edpiairo-pocketidp
```

5 - Run Humanitec CLI and follow the prompted instruction
```
humctl login
```

## Backstage Setup

After running `humctl login` we'll now be able to go to the Gitea instance through your browser. 

Go to: https://git.localhost:8443

Sign in to Gitea with the username and password. Both are pre-set as: 5minadmin

## Add New Database Resource

### Step 1 - Create the Resource Class

Run the following command inside of your terminal:
```
humctl api post "/orgs/${HUMANITEC_ORG}/resources/types/postgres/classes" \
  -d '{
  "id": "5min",
  "description": "Specifically to be used when using 5min IDP settings"
}'
```

### Step 2 - Create the Postgres Resource Definition

Clone the following repository inside the Coder instance:
```
git clone https://github.com/humanitec-architecture/resource-packs-in-cluster
```

Go to the folder: 
```
cd /resource-packs-in-cluster/examples/postgres
```

Now lets run Terraform inside this folder to generate the resource definitions:
```
terraform init
terraform plan
terraform apply
```

List of available resource definitions
``` 
humctl get resource-definitions
```

Lets set the Matching Criteria to make sure that we can use this Resource Definition inside our 5min-idp environment type and using a the resource class we created earlier.
``` 
humctl api PUT /orgs/${HUMANITEC_ORG}/resources/defs/<resource-definition-id>/criteria --data '[{"env_type": "5min-local","class":"5min"}]'

humctl api PUT /orgs/pe-course-ggsbscty/resources/defs/5min-idp-zffh-postgres/criteria --data '[{"env_type": "5min-local","class":"5min"}]'
``` 

### Step 3 - Request a Resource With a Deployment
Let’s go back to our local Gitea instance and find the repository for our scaffolded application. You can go to https://git.localhost:8443/

Find the `score.yaml`.

For our deployment, we know that we want to be able to request a Postgres database. Let’s make sure this is available to us in our current context!

Go back to your terminal, and run the following command:
``` 
humctl score available-resource-types
``` 

We can see we have two different class of Postgress available: default and 5min. For this particular exercise, let's use the 5min class!

Going back to the Git repository and the `score.yaml` file, we're going to add three lines under resources. Your `score.yaml` file should now look similar to this:
``` 
resources:
  dns: # We need a DNS record to point to the service 
    type: dns
  route:
    type: route
    params:
      host: ${resources.dns.host}
      path: /
      port: 80
  database: # Adding our Postgres database and associated class
    type: postgres
    class: 5min
```

Commit the change to the `main branch`, and go to the Actions tab to follow the deployment. You can also follow the deployment in the Humanitec UI!