<%= title("Writing a Plugin") %>

First, you must understand [how Buffalo finds plugins](#how-does-buffalo-find-plugins), before you can successfully write one.

The `buffalo-plugins` plugin adds a new generator to `buffalo generate` to help you build a new plugin quickly

```bash
$ buffalo generate plugin -h

buffalo generate plugin github.com/foo/buffalo-bar

Usage:
  buffalo-plugins plugin [flags]

Flags:
  -a, --author string       author's name
  -d, --dry-run             run the generator without creating files or running commands
  -f, --force               will delete the target directory if it exists
  -h, --help                help for plugin
  -l, --license string      choose a license from: [agpl, isc, lgpl-v2.1, mozilla, no-license, artistic, bsd, eclipse, lgpl-v3, mit, apache, bsd-3-clause, unlicense, cc0, gpl-v2, gpl-v3] (default "mit")
  -s, --short-name string   a 'short' name for the package
      --with-gen            creates a generator plugin
```

<%= exampleDir("docs/plugins/_example/standard") %>
