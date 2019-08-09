# Vue d'ensemble

Bienvenue à bord ! Bien que Buffalo puisse être considéré comme un framework, on parlera plutôt d'un écosystème de biliothèques Go et Javascript testées pour travailler ensemble. La plupart de ces composants peut être remplacé par un autre, mais nous ne fournirons un support que pour les composants par défaut.

Dans ce chapitre, nous allons faire un petit tour des briques par défaut proposées dans votre application Buffalo.

## Bibliothèques backend

### buffalo

Buffalo est le « liant » entre les différents composants fournis. Il s'assure que les biliothèques interagissent bien ensemble et gère leur ordre d'utilisation.

### gorilla/mux

[gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) est l'un des routeurs les plus utilisés en Go. Bien que d'autres routeurs soient plus rapides (comme [httprouter](https://github.com/julienschmidt/httprouter)), gorilla/mux est le routeur qui propose le plus de fonctionnalités, tout en étant suffisement rapide.

### pop

[pop](https://github.com/gobuffalo/pop) est l'ORM par défaut de Buffalo. Il fournit la boîte à outils `soda` pour vous aider à gérer votre base de données, et supporte tant PostgreSQL et MySQL que SQLite.

### plush

[plush](https://github.com/gobuffalo/plush) est le moteur de templates par défaut de Buffalo. Sa syntaxe est proche de celle des templates ERB (en Ruby).

### packr

[packr](https://github.com/gobuffalo/packr) est un empaqueteur en Go pour vos assets statiques (templates, images, etc.). Packr a pour but de produire un binaire final avec tout ce qu'il faut à l'intérieur. 

## Bibilothèques frontend

### Bootstrap

[Bootstrap](https://getbootstrap.com/) est l'une des bibliothèques frontend les plus connues en matière de design Web. Elle aide à construire des interfaces adaptables à tous les écrans en utilisant des composants standards tels que des tableaux, des carousels ou des grilles de mise en page.

### jQuery

[jQuery](https://jquery.com/) est une bibliothèque très fournie qui a pour but de manipuler l'arbre DOM et de rendre les requêtes AJAX plus simples. Bien que moins utilisée maintenant, de nombreux projets continuent de l'employer comme compagnon pour les aider à supporter tous les navigateurs.

### Webpack

[Webpack](https://webpack.js.org/) est un empaqueteur Javascript renommé pour vos assets. Il s'occupera de vos fichiers Javascript, CSS, images et tout autres assets statiques.

Par défaut, Webpack est configuré pour minifier et hasher vos assets.

## Prochaines étapes

* [Installation](/fr/docs/getting-started/installation) - Installez Buffalo !