---
name: Configuración
seoDescription: "Maneja la configuracion con Buffalo"
seoKeywords:
  - "buffalo"
  - "go"
  - "golang"
  - "configuration"
  - "env"
  - "framework"
  - "web"
weight: 4
aliases:
  - /documentation/getting_started/config-vars
  - /docs/getting-started/config-vars
  - /en/docs/getting-started/config-vars
  - /es/docs/getting-started/config-vars
---

# Configuracion

En este capitulo, tu vas a aprender como manejar la configuracion de tu proyecto con Buffalo.

Las variables de entorno son un buen camino para separar entornos con diferentes valores o secretos, de la base del codigo de tu aplicación. ([como describe en la aplicación de los Doce factores](https://12factor.net/config)). Eso puede ayudar a definir el comportamiento quese basa en el contexto de la aplicación (Como requerir SSL en producción) y aislas secretos de tu codigo. De esta manera, todos los desarroladores no tienen que conocer las llaves de producción para servicios sensibles como la API de un Banco y pueden utilizar las llaves de la API en sandbox.

## Variables de entorno Permitidas

Las siguientes variables son usadas por Buffalo:

| Variable                 | Por defecto                  | Uso                                                                |
| ---                      | ---                          | ---                                                                |
| `GO_ENV`                 | `development`                | El entorno (dev, qa, production etc.) que Buffalo esté corriendo.  |
| `GO_BIN`                 | `go`                         | El compilador de Go para usar.                                     |
| `BUFFALO_PLUGIN_PATH`    | `$PATH`                      | Donde Buffalo va a buscar sus plugins.                             |
| `BUFFALO_PLUGIN_TIMEOUT` | `1s`                         | Cuanto tiempo espera Buffalo para que un plugin Responda           |
| `ADDR`                   | `127.0.0.1` or `0.0.0.0`     | La direccion que se debe usar en el servidor.                      |
| `PORT`                   | `3000`                       | El puerto que se debe configurar para el server.                   |
| `HOST`                   | `http://127.0.0.1:$PORT`     | La "URL" de la aplicación (i.e. Lo que escriben los usuarios).     |
| `SESSION_SECRET`         | `""`                         | Utilizada para asegurar las sesiones.                              |

## Configuracion Personalizada

Tu puedes proporcionar tus propias variables y recuperarlas desde su aplicación. El paquete [envy](https://github.com/gobuffalo/envy) hace esto facil!

```go
import "github.com/gobuffalo/envy"

// [...]

// Get MYSECRET env variable, default to empty string if it's not set
var MYSECRET = envy.Get("MYSECRET", "")

// Get REQUIREDSECRET env variable, return an error if it's not set
REQUIREDSECRET, err := envy.MustGet("REQUIREDSECRET")
```

## Soporte para archivos .env

<!--%= sinceVersion("0.10.3") %-->

Buffalo entrega soporte con `.env` (**desde buffalo >= 0.10.3**), lo que significa que buffalo cargará los archivos `.env` dentro de la variable del entorno una vez que la aplicacion esté corriendo. Para hacer eso Buffalo usa [`envy.Load`](https://github.com/gobuffalo/envy/blob/e613c80275b86293880eddeb27417c9a7c670ff3/envy.go#L53) el cual buscará el archivo `.env` en la raiz de tu aplicación.

Si tu no estas familiarizado como luce un archivo `.env`, aqui tienes un ejemplo:

```text
SENDGRID_API_KEY=ABCCOQ7GFRVCW0ODHPFQ3FTP5SLL1Q
SENDGRID_EMAIL=email@myapp.com

APP_DEBUG=true
APP_LOG_LEVEL=debug
APP_URL=https://myapp.com
```

Las aplicaciones generadas (**con buffalo >= 0.10.3**) van a crear un archivo `.env` por defecto en la raiz de su aplicación. Este archivo va será mirado por Buffalo para los cambios pero va a ser ignorado por git (Añadido en el `.gitignore`). Esto es una buena manera de evitar que los desarrolladores suban las credenciales por error.

## Siguiente Paso

* [Integracion de Herramientas](/es/documentation/getting_started/integrations) - Trabajar con Buffalo, usando herramientas existentes.
