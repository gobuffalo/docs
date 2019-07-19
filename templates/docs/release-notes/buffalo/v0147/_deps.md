## Dependency Diet

```diff
--- v0.14.0-go.mod	2019-07-19 16:14:14.000000000 -0400
+++ v0.14.7-go.mod	2019-07-19 16:13:40.000000000 -0400
@@ -3,30 +3,23 @@
 go 1.12

 require (
+	github.com/cockroachdb/apd v1.1.0 // indirect
+	github.com/cockroachdb/cockroach-go v0.0.0-20181001143604-e0a95dfd547c // indirect
 	github.com/codegangsta/negroni v1.0.0 // indirect
-	github.com/gobuffalo/buffalo v0.14.0
-	github.com/gobuffalo/buffalo-pop v1.9.0
+	github.com/gobuffalo/buffalo v0.14.7-beta.2
+	github.com/gobuffalo/buffalo-pop v1.16.0
 	github.com/gobuffalo/envy v1.7.0
-	github.com/gobuffalo/fizz v1.9.2 // indirect
-	github.com/gobuffalo/httptest v1.4.0 // indirect
-	github.com/gobuffalo/logger v1.0.1 // indirect
-	github.com/gobuffalo/makr v1.2.0 // indirect
 	github.com/gobuffalo/mw-csrf v0.0.0-20190129204204-25460a055517
 	github.com/gobuffalo/mw-forcessl v0.0.0-20190224202501-6d1ef7ffb276
 	github.com/gobuffalo/mw-i18n v0.0.0-20190224203426-337de00e4c33
 	github.com/gobuffalo/mw-paramlogger v0.0.0-20190224201358-0d45762ab655
-	github.com/gobuffalo/nulls v0.1.0 // indirect
 	github.com/gobuffalo/packr v1.30.1
 	github.com/gobuffalo/packr/v2 v2.5.2
-	github.com/gobuffalo/plush v3.8.3+incompatible // indirect
 	github.com/gobuffalo/pop v4.11.2+incompatible
 	github.com/gobuffalo/suite v2.8.1+incompatible
-	github.com/gorilla/mux v1.7.3 // indirect
-	github.com/gorilla/sessions v1.2.0 // indirect
+	github.com/jackc/fake v0.0.0-20150926172116-812a484cc733 // indirect
 	github.com/markbates/grift v1.1.0
-	github.com/markbates/oncer v1.0.0 // indirect
-	github.com/markbates/refresh v1.8.0 // indirect
-	github.com/monoculum/formam v0.0.0-20190307031628-bc555adff0cd // indirect
+	github.com/satori/go.uuid v1.2.0 // indirect
+	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24 // indirect
 	github.com/unrolled/secure v1.0.0
-	golang.org/x/net v0.0.0-20190620200207-3b0461eec859 // indirect
 )
```
