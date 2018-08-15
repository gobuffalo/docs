module github.com/gobuffalo/gobuffalo

require (
	dmitri.shuralyov.com/text/kebabcase v0.0.0-20180217051803-40e40b42552a
	github.com/RoaringBitmap/roaring v0.4.9
	github.com/Smerity/govarint v0.0.0-20150407073650-7265e41f48f1
	github.com/ajg/form v0.0.0-20160802194845-cc2954064ec9
	github.com/blevesearch/bleve v0.7.0
	github.com/blevesearch/go-porterstemmer v1.0.1
	github.com/blevesearch/segment v0.0.0-20160915185041-762005e7a34f
	github.com/boltdb/bolt v1.3.1
	github.com/cockroachdb/cockroach-go v0.0.0-20180212155653-59c0560478b7 // indirect
	github.com/couchbase/vellum v0.0.0-20180427141700-eb6ae3743b3f
	github.com/dustin/go-humanize v0.0.0-20180421182945-02af3965c54e
	github.com/edsrzf/mmap-go v0.0.0-20170320065105-0bce6a688712
	github.com/fatih/color v1.7.0
	github.com/fatih/structs v1.0.0 // indirect
	github.com/fsnotify/fsnotify v1.4.7
	github.com/glycerine/go-unsnap-stream v0.0.0-20180323001048-9f0cb55181dd
	github.com/go-sql-driver/mysql v1.3.0
	github.com/gobuffalo/buffalo v0.0.0-20180530133015-7c0b5fe40f04
	github.com/gobuffalo/envy v1.6.3
	github.com/gobuffalo/github_flavored_markdown v1.0.0
	github.com/gobuffalo/makr v1.1.1 // indirect
	github.com/gobuffalo/mw-forcessl v0.0.0-20180802152810-73921ae7a130
	github.com/gobuffalo/mw-i18n v0.0.0-20180802152014-e3060b7e13d6
	github.com/gobuffalo/mw-paramlogger v0.0.0-20180807082017-6b90b69a724a
	github.com/gobuffalo/packr v1.11.0
	github.com/gobuffalo/plush v0.0.0-20180503143032-5d52aed81c89
	github.com/gobuffalo/pop v0.0.0-20180521154020-9985012f32f5
	github.com/gobuffalo/tags v2.0.6+incompatible
	github.com/gobuffalo/uuid v2.0.0+incompatible // indirect
	github.com/gobuffalo/validate v2.0.0+incompatible // indirect
	github.com/gobuffalo/x v0.0.0-20180117215853-11ca13c05abd
	github.com/golang/protobuf v1.1.0
	github.com/golang/snappy v0.0.0-20180518054509-2e65f85255db
	github.com/gorilla/context v1.1.1
	github.com/gorilla/mux v1.6.2
	github.com/gorilla/securecookie v1.1.1
	github.com/gorilla/sessions v0.0.0-20160922145804-ca9ada445741
	github.com/gorilla/websocket v1.2.0
	github.com/grokify/html-strip-tags-go v0.0.0-20180530080503-3f8856873ce5
	github.com/jmoiron/sqlx v0.0.0-20180406164412-2aeb6a910c2b
	github.com/lib/pq v0.0.0-20180523175426-90697d60dd84
	github.com/markbates/going v1.0.1
	github.com/markbates/grift v0.0.0-20180319170132-76f93617a788
	github.com/markbates/hmax v0.0.0-20170213234856-800e180dcd16
	github.com/markbates/inflect v1.0.0
	github.com/markbates/refresh v1.4.0
	github.com/markbates/sigtx v1.0.0
	github.com/markbates/willie v0.0.0-20180320154129-67934945178e
	github.com/mattn/anko v0.0.4
	github.com/mattn/go-colorable v0.0.9 // indirect
	github.com/mattn/go-isatty v0.0.3 // indirect
	github.com/mattn/go-sqlite3 v1.7.0
	github.com/microcosm-cc/bluemonday v1.0.0
	github.com/mitchellh/go-homedir v0.0.0-20180523094522-3864e76763d9
	github.com/monoculum/formam v0.0.0-20170619223434-99ca9dcbaca6
	github.com/mschoch/smat v0.0.0-20160514031455-90eadee771ae
	github.com/nicksnyder/go-i18n v1.10.0
	github.com/pelletier/go-toml v1.1.0
	github.com/philhofer/fwd v1.0.0
	github.com/pkg/errors v0.8.0
	github.com/russross/blackfriday v0.0.0-20180428102519-11635eb403ff
	github.com/serenize/snaker v0.0.0-20171204205717-a683aaf2d516 // indirect
	github.com/sergi/go-diff v1.0.0
	github.com/shurcooL/github_flavored_markdown v0.0.0-20171120162553-28433ea3fc83
	github.com/shurcooL/go v0.0.0-20180423040247-9e1955d9fb6e
	github.com/shurcooL/go-goon v0.0.0-20170922171312-37c2f522c041
	github.com/shurcooL/graphql v0.0.0-20180514000029-62c9ce094e75
	github.com/shurcooL/octiconssvg v0.0.0-20180509202950-f4482000187f
	github.com/shurcooL/sanitized_anchor_name v0.0.0-20170918181015-86672fcb3f95
	github.com/sirupsen/logrus v1.0.5
	github.com/steveyen/gtreap v0.0.0-20150807155958-0abe01ef9be2 // indirect
	github.com/stvp/slug v0.0.0-20150928221549-5ab8191bb1fe
	github.com/tinylib/msgp v0.0.0-20170627230615-b2b6a672cf1e
	github.com/unrolled/secure v0.0.0-20180416205222-a1cf62cc2159
	github.com/willf/bitset v1.1.3
	golang.org/x/crypto v0.0.0-20180527072434-ab813273cd59
	golang.org/x/net v0.0.0-20180801234040-f4c29de78a2a
	golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f
	golang.org/x/sys v0.0.0-20180525142821-c11f84a56e43
	golang.org/x/text v0.3.0 // indirect
	gopkg.in/yaml.v2 v2.2.1
)
