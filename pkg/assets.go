package pkg

import "embed"

//go:embed api/index.html
//go:embed openapi/api.yaml
var SwaggerUI embed.FS
