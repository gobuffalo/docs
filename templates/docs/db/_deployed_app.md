## Migrations Once Deployed

<%= note() { %>
This section is only for Buffalo users.
<% } %>

When you build your app, the migrations are stored inside your binary. Your binary has a hidden `migrate` command baked in that performs the migrations, just like it does when you use `buffalo pop migrate`:

```bash
$ ./myapp migrate
DEBU[2018-01-12T06:14:20Z] select count(*) as row_count from (SELECT schema_migration.* FROM schema_migration AS schema_migration WHERE version = ?) a $1=20171213171622
DEBU[2018-01-12T06:14:20Z] select count(*) as row_count from (SELECT schema_migration.* FROM schema_migration AS schema_migration WHERE version = ?) a $1=20171213172104
DEBU[2018-01-12T06:14:20Z] select count(*) as row_count from (SELECT schema_migration.* FROM schema_migration AS schema_migration WHERE version = ?) a $1=20171213172249
DEBU[2018-01-12T06:14:20Z] select count(*) as row_count from (SELECT schema_migration.* FROM schema_migration AS schema_migration WHERE version = ?) a $1=20171213173148
DEBU[2018-01-12T06:14:20Z] select count(*) as row_count from (SELECT schema_migration.* FROM schema_migration AS schema_migration WHERE version = ?) a $1=20171219070903
DEBU[2018-01-12T06:14:20Z] select count(*) as row_count from (SELECT schema_migration.* FROM schema_migration AS schema_migration WHERE version = ?) a $1=20171219071524

0.0010 seconds
```
