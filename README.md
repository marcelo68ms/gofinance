# Instalação

```go 
go get github.com/lucasmdomingues/gofinance
```

# Exemplos

### Obter somente cotações de moedas e bolsas.

```go
import (
	"encoding/json"
	"fmt"
	"github.com/lucasmdomingues/gofinance"
)

func main() {

	quotations, err := gofinance.GetQuotations("KEY")
	if err != nil {
		fmt.Println(err)
		return
	}

	json, err := json.Marshal(quotations)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", json)
}
```

### Obter somente taxas CDI e SELIC.

```go
import (
	"encoding/json"
	"fmt"
	"gofinance"
)

func main() {

	taxes, err := gofinance.GetTaxes("KEY")
	if err != nil {
		fmt.Println(err)
		return
	}

	json, err := json.Marshal(taxes)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", json)
}

```
### HG Finance
https://hgbrasil.com/status/finance/

### TO DO

* Obtendo dados históricos
* Personalizando a resposta
