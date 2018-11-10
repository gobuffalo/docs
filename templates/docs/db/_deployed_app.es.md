<%= title("Migraciones una vez desplegadas", {name:"migrating-deployed-app"}) %>

<%= note() { %>
Esta sección es solo para usuarios de Buffalo..
<% } %>

Cuando construye su aplicación, las migraciones se almacenan dentro de su binario. Su binario tiene un comando oculto `migrate` que realiza las migraciones, al igual que cuando se usa` buffalo pop migrate`:

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
