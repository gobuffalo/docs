---
name: Registros
seoDescription: Configure logging in Buffalo
seoKeywords: ["buffalo", "go", "golang", "configuration", "logs", "logging", "custom"]
aliases:
  - /docs/guides/logging
  - /en/docs/guides/logging
---

# Registros

Los registros de Buffalo se manejan usando la librería [logrus](https://github.com/sirupsen/logrus).

## Valores Predeterminados

El registrador por defecto genera regustros en un formato legible:

```bash
INFO[2020-02-21T07:42:34+01:00] /en/ content_type=text/html duration=26.189949ms human_size="21 kB" method=GET params="{\"lang\":[\"en\"]}" path=/en/ render=22.730816ms request_id=9b8d9260225fe99609a2-7cc679f4ae458b9925e3 size=21182 status=200
```

## Personalizar el Registrador

```go
// JSONLogger wraps a logrus JSON logger into a buffalo Logger
func JSONLogger(lvl logger.Level) logger.FieldLogger {
    l := logrus.New()
    l.Level = lvl
    l.SetFormatter(&logrus.JSONFormatter{})
    l.SetOutput(os.Stdout)
    return logger.Logrus{FieldLogger: l}
}

//...

app = buffalo.New(buffalo.Options{
// ...
    Logger:       JSONLogger(logger.DebugLevel),
}
```
