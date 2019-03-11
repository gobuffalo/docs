```bash
$ buffalo generate resource -h

Generate a new actions/resource file

Usage:
  buffalo generate resource [name] [flags]

Aliases:
  resource, r

Examples:
$ buffalo g resource users
Generates:

- actions/users.go
- actions/users_test.go
- models/user.go
- models/user_test.go
- migrations/2016020216301234_create_users.up.fizz
- migrations/2016020216301234_create_users.down.fizz

$ buffalo g resource users --skip-migration
Generates:

- actions/users.go
- actions/users_test.go
- models/user.go
- models/user_test.go

$ buffalo g resource users --skip-model
Generates:

- actions/users.go
- actions/users_test.go

$ buffalo g resource users --use-model users
Generates:

- actions/users.go
- actions/users_test.go

Flags:
  -d, --dry-run            dry run
  -h, --help               help for resource
  -n, --name string        allows to define a different model name for the resource being generated.
  -s, --skip-migration     tells resource generator not-to add model migration
      --skip-model         tells resource generator not to generate model nor migrations
      --skip-templates     tells resource generator not to generate templates for the resource
      --use-model string   tells resource generator to reference an existing model in generated code
  -v, --verbose            verbosely print out the go get commands
```
