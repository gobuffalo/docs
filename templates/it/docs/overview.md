<% seoDescription("Buffalo framework overview") %>
<% seoKeywords(["buffalo", "framework", "overview", "go", "golang", "mux", "bootstrap", "jquery"]) %>

<%= h1("Overview") %>

Benvenuto a bordo!

Anche se Buffalo può considerarsi un framework, è di fatto un ecosistema di librerie Go e Javascript pensate per lavorare insieme.

In questo capitolo, faremo un riassunto in ciò che troverai nella tua app Buffalo.

## Librerie backend

### buffalo

Buffalo è la "colla" tra tutti i componenti. Include le librerie e amministra il funzionamento.

### gorilla/mux

[gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) è uno dei più usati router di Go. Anche se alcuni router sono più veloci (come [httprouter](https://github.com/julienschmidt/httprouter)), gorilla/mux è quello che offre più funzionalità restando abbastanza veloce.

### pop

[pop](https://github.com/gobuffalo/pop) è l'ORM di default di Buffalo. Fornisce il toolbox `soda` che ti supporta con le necessità del tuo database e supporta molti tipi di database, come PostgreSQL, MySQL e SQLite.

### plush

[plush](https://github.com/gobuffalo/plush) è il template engine di default di Buffalo. La sua sintassi è simile a quella dei template ERB (di Ruby).

### packr

[packr](https://github.com/gobuffalo/packr) è il software Go per incorporare i tuoi asset statici (template, immagini e così via). Packr mira a produrre un file binario finale con tutto al suo interno.

## Librerie frontend

### Bootstrap

[Bootstrap](https://getbootstrap.com/) è tra le librerie frontend più famose. Aiuta a scrivere interfacce responsive usando componenti semplici come tabelle, carousel e layout a griglia.

### jQuery

[jQuery](https://jquery.com/) è una corposa libreria che mira a semplificare la manipolazione del DOM e le query AJAX. Anche se ora è meno usata, molti progetti la usano come strategia per supportare tutti i browser.

### Webpack

[Webpack](https://webpack.js.org/) è un famoso gestore di pacchetti per Javascript. Si prenderà cura dei tuoi Javascript, CSS, immagini e asset statici.

Webpack è configurato di default per versionare e minimizzare i tuoi asset.

## Prossimi passi

* [Installazione](/it/docs/getting-started/installation) - Installa Buffalo!
