## Unified Helpers Package

Previously [`github.com/gobuffalo/plush`](https://godoc.org/github.com/gobuffalo/plush) helper functions were scattered in multiple locations. These functions were hard to find, often non-exported, and poorly documented.

The new [`github.com/gobuffalo/helpers`](https://godoc.org/github.com/gobuffalo/helpers) has been introduced to be a single source for these helpers.

This new package is broken into categorized sub-packages with the appropriate helpers

```text
├── content
├── debug
├── encoders
├── env
├── escapes
├── forms
│   └── bootstrap
├── hctx
├── helpers
├── helptest
├── inflections
├── iterators
├── meta
├── paths
├── tags
└── text
```
