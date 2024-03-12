
# Infra

## Sumãrio

1. [Logger](#logger)
   1. [Funcionalidade](#funcionalidade)
   2. [Como usar](#como-usar)

### Logger

path: `infra/config/logger.go`

##### Funcionalidade

Ferramenta para registrar mensagens durante a execução do programa.
Utilidades:

1. Registro de eventos
2. Depuração
3. Monitoramento

##### Como usar

Para instanciar o logger em algum lugar da aplicação é necessário usar o package `config`. Ao topo do arquivo, declare o logger com o escopo do arquivo, vamos considerar um controller:

```go
//

package controllers

import (
 "pet-dex-backend/v2/infra/config"
 ... //imports
)

var logger = config.GetLogger("pet-controller") // Instânciando o logger

... // controller

func (pc *PetController) Create(w http.ResponseWriter, r *http.Request){
 ...
 logger.Error(err) // Utilizando
 logger.Info("Olha lá")
}
```

> É nessário passar o contexto o logger vai ser executado pelo parametro da função

O logger tem alguns métodos:

- Debug - Debugf
- Info - Infof
- Warn - Warnf
- Error - Errorf
