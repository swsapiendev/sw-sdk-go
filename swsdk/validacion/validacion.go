package validacion

import (
	"github.com/swsapiendev/sw-sdk-go/swsdk/autenticacion"
)

// ValidarCFDI valida un archivo XML de CFDI
func ValidarCFDI(xmlPath string) (*ValidacionResponse, error) {
	client := autenticacion.SetUser()
	return validarCFDIHelper(client, xmlPath)
}
