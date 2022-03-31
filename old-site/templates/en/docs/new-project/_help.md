```bash
$ buffalo help new
Creates a new Buffalo application

Usage:
  buffalo new [name] [flags]

Flags:
      --api                  skip all front-end code and configure for an API server
      --ci-provider string   specify the type of ci file you would like buffalo to generate [none, travis, gitlab-ci, circleci] (default "none")
      --config string        config file (default is $HOME/.buffalo.yaml)
      --db-type string       specify the type of database you want to use [cockroach, mariadb, mysql, postgres] (default "postgres")
      --docker string        specify the type of Docker file to generate [none, multi, standard] (default "multi")
  -d, --dry-run              dry run
  -f, --force                delete and remake if the app already exists
  -h, --help                 help for new
      --module string        specify the root module (package) name. [defaults to 'automatic']
      --skip-config          skips using the config file
      --skip-pop             skips adding pop/soda to your app
      --skip-webpack         skips adding Webpack to your app
      --skip-yarn            use npm instead of yarn for frontend dependencies management
      --vcs string           specify the Version control system you would like to use [none, git, bzr] (default "git")
  -v, --verbose              verbosely print out the go get commands
```
