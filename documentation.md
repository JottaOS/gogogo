## El camino de aprender Go..

## Módulos?

Crear módulos con `go mod init`. Ejemplo: `go mod init example/hello`

Para instalar un módulo, luego de importarlo en el código escribo `go mod tidy`

Para utilizar un módulo creado por mí, localmente, debo editar como se busca el módulo, Ej: `go mod edit -replace example.com/greetings=../greetings`. Luego de esto, puedo correr de vuelta `go mod tidy` para sincronizar.

Las funciones que comienzan con nombre en Mayúsculas pueden ser exportadas a otros módulos (public) y los que comienzan en minúscula solo pueden ser utilizados dentro del mismo paquete (private)