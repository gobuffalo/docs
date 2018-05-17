```bash
$ buffalo db g model --help

Generates a model for your database

Usage:
  buffalo db generate model [name] [flags]

Aliases:
  model, m


Flags:
  -s, --skip-migration   Skip creating a new fizz migration for this model.

Global Flags:
  -c, --config string   The configuration file you would like to use.
  -d, --debug           Use debug/verbose mode
  -e, --env string      The environment you want to run migrations against. Will use $GO_ENV if set. (default "development")
  -p, --path string     Path to the migrations folder (default "./migrations")
```

You can remove generated model by running:

```bash
$ buffalo db destroy model [name]
```

Or in short form:

```bash
$ buffalo db d m [name]
```
