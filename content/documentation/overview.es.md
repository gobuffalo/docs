---
name: Vistazo General
icon: "images/overview.svg"
seoDescription: Buffalo framework overview
seoKeywords: ["buffalo", "framework", "overview", "go", "golang", "mux", "bootstrap", "jquery"]
aliases:
  - /es/docs/overview
---

# Vistazo General

Bienvenido a bordo!

Si bien Buffalo puede considerarse como un marco, es principalmente un ecosistema de bibliotecas Go y Javascript seleccionadas para encajar juntas. La mayoría de estos componentes se pueden cambiar por otros, pero solo brindaremos soporte para esta combinación predeterminada.


En este capítulo, haremos un recorrido por los componentes predeterminados que se envían con su aplicación Buffalo.

## Bibliotecas de back-end

- ### Buffalo

  Buffalo es el *"pegamento"* entre todos los componentes proporcionados. Envuelve las bibliotecas y administra el flujo de trabajo.

- ### gorilla/mux

  [gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) es uno de los enrutadores más utilizados en Go. Si bien algunos enrutadores son más rápidos (como [httprouter](https://github.com/julienschmidt/httprouter)), gorilla/mux es el que brinda la mayor cantidad de funciones y es lo suficientemente rápido.

- ### pop

  [pop](https://github.com/gobuffalo/pop) es el ORM predeterminado para Buffalo. Proporciona la caja de herramientas `soda` para ayudarte con tus necesidades de base de datos y es compatible con varias bases de 
  datos, como PostgreSQL, MySQL y SQLite.

- ### plush

  [plush](https://github.com/gobuffalo/plush) es el motor de plantillas predeterminado de Buffalo. Su sintaxis es cercana a las plantillas ERB (en Ruby).

## Bibliotecas de front-end

- ### bootstrap
  [bootstrap](https://getbootstrap.com/) es una de las bibliotecas de herramientas frontend más famosas. Ayuda a crear interfaces responsivas utilizando componentes comunes como tablas, carruseles o diseños de cuadrícula.

- ### jQuery

  [jQuery](https://jquery.com/) es una rica biblioteca que tiene como objetivo simplificar la manipulación DOM y las consultas AJAX. Si bien ahora se usa menos, muchos proyectos todavía lo tienen como complemento para ayudar a admitir todos los navegadores.

- ### webpack
  [webpack](https://webpack.js.org/) es un conocido paquete de activos de Javascript. Se encargará de sus archivos Javascript, CSS, imágenes y activos estáticos.

  {{<note>}}
  **webpack** está configurado de forma predeterminada para aplicar hash y minificar sus activos.
  {{</note>}}

## Siguientes pasos

* [Instalación](/es/docs/getting-started/installation) - Instala Buffalo!
