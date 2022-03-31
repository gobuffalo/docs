## Buffalo Plugin Cache Improvements

When running `buffalo fix` on an existing application the Buffalo plugins cache will be cleared to prevent issues with out of date caches.

The cache will also now only store plugins that can provide a successful `available` sub-command.

