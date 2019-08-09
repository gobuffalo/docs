## Migrations en production

<%= note() { %>
Cette section ne concerne que les utilisateurs de Buffalo.
<% } %>

Lorsque vous compilez votre application, les migrations sont stockées dans le binaire résultant. Ce binaire inclut une commande cachée `migrate` qui applique les migrations, tout comme la commande `buffalo pop migrate` que vous utiliseriez en développement :

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
