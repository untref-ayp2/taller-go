# Ejercicios: Errores

## 01 - División segura

Completar la función `Dividir` en `01-dividir/dividir.go`.
Recibe dos enteros y devuelve el cociente junto con un error si el divisor
es cero. El error se crea con `errors.New`.

→ `01-dividir/`

## 02 - Clasificar edad

Completar la función `ClasificarEdad` en `02-clasificar/clasificar.go`.
Recibe una edad y devuelve `"nino"`, `"adolescente"`, `"adulto"` o
`"adulto mayor"` según corresponda. Si la edad es negativa, devuelve un error.

→ `02-clasificar/`

## 03 - Buscar producto

Completar la función `BuscarProducto` en `03-buscar-producto/buscar_producto.go`.
Recibe un slice de `Producto` y un código, y devuelve el producto encontrado.
Si no existe, devuelve el error centinela `ErrProductoNoEncontrado`.

→ `03-buscar-producto/`

## 04 - Saldo insuficiente

Completar la función `Extraer` en `04-extraer/extraer.go`.
Recibe un saldo y un monto, y devuelve el saldo restante. Si el monto supera
al saldo, devuelve un `SaldoInsuficienteError` (struct con campos `Saldo`
y `Monto` que implementa `error`).

→ `04-extraer/`

## 05 - Leer archivo con contexto

Completar la función `LeerArchivo` en `05-leer-archivo/leer_archivo.go`.
Llama a `os.ReadFile` y, si hay error, lo envuelve con `fmt.Errorf`
usando `%w` para agregar contexto. Los tests verifican el error con
`errors.Is` contra `os.ErrNotExist`.

→ `05-leer-archivo/`
