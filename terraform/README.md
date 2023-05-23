# Terraform

## Workflow

1. init
2. validate
3. plan
4. apply
5. destroy

## Configuration files

Terraform configuration files (aka terraform manifests) are files with .tf extension.

## HCL

HCL HashiCorp Language:
- Blocks
- Arguments
- Identifiers
- Comments

### Blocks

```
# Template
<BLOCK TYPE> "<BLOCK LABEL>" "<BLOCK LABEL>"   {
  # Block body
  <IDENTIFIER> = <EXPRESSION> # Argument
}
```

```
# AWS Example
resource "aws_instance" "ec2demo" { # BLOCK
  ami           = "ami-04d29b6f966df1537" # Argument
  instance_type = var.instance_type # Argument with value as expression (Variable value replaced from varibales.tf
}
```

### Comments

```
# Line comment

// Line comment

/* Multi
line
comments*/
```

## Arguments
An argument assigns a value to a particular name.

## Meta-arguments
The meta-arguments within resource blocks allow to define the behavior of a resource behavior.


## Top-level blocks

- Fundamental blocks
    - Terraform (settings) block
    - Provider block
    - Resource block
- Variable blocks
    - Input variables block
    - Output values block
    - Local values block
- Calling / Referencing blocks
    - Data sources block
    - Modules block