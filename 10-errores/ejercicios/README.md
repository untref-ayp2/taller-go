# Ejercicios: Errores

1. **División segura.** Escribí una función `dividir(a, b int) (int, error)`
   que retorne el resultado de a/b y un error si b es cero. Usá `errors.New`.
   → `dividir/`

2. **Clasificar nota.** Escribí una función `clasificar(nota int) (string, error)`
   que devuelva `"aprobado"` si nota >= 4, `"desaprobado"` si nota < 4,
   o un error con `fmt.Errorf` si la nota está fuera del rango 0-10
   (incluyendo el valor inválido en el mensaje).
   → `clasificar/`

3. **Buscar producto.** Declará un error centinela `ErrProductoNoEncontrado`.
   Escribí una función `buscarProducto(codigos []string, target string) (string, error)`
   que recorra el slice y devuelva el código si existe, o
   `ErrProductoNoEncontrado` en caso contrario.
   → `buscar-producto/`

4. **Saldo insuficiente.** Definí un struct `SaldoInsuficienteError` con
   campos `Saldo` y `Monto` (ambos `float64`) que implemente `error`.
   Escribí una función `extraer(saldo, monto float64) (float64, error)`
   que devuelva el saldo restante o un `SaldoInsuficienteError` si el
   monto supera al saldo.
   → `extraer/`

5. **Leer configuración.** Escribí una función `leerConfig(ruta string) (string, error)`
   que llame a `os.ReadFile(ruta)`. Si hay error, agregale contexto con
   `fmt.Errorf` y `%w` indicando que falló al leer la configuración.
   En el `main`, usá `errors.Is` para detectar el error original.
   → `leer-config/`
