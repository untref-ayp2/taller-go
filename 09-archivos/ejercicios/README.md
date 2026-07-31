# Ejercicios: Archivos

## 01 - Contar líneas

Completar la función `ContarLineas` en `01-contar-lineas/contar_lineas.go`.
Lee un archivo de texto y devuelve la cantidad de líneas que contiene,
usando `bufio.Scanner`.

→ `01-contar-lineas/`

## 02 - Copiar archivo

Completar la función `CopiarArchivo` en `02-copiar-archivo/copiar_archivo.go`.
Recibe las rutas de origen y destino, y copia el contenido usando
`os.ReadFile` y `os.WriteFile`.

→ `02-copiar-archivo/`

## 03 - Sumar números

Completar la función `SumarNumeros` en `03-sumar-numeros/sumar_numeros.go`.
Lee un archivo donde cada línea contiene un número entero, los suma
usando `bufio.Scanner` y `strconv.Atoi`, y devuelve el resultado.

→ `03-sumar-numeros/`

## 04 - Agregar línea

Completar la función `AgregarLinea` en `04-agregar-linea/agregar_linea.go`.
Agrega una línea al final de un archivo existente, leyendo el contenido
actual con `os.ReadFile` y escribiéndolo de vuelta con `os.WriteFile`.

→ `04-agregar-linea/`

## 05 - Numerar líneas

Completar la función `NumerarLineas` en `05-numerar-lineas/numerar_lineas.go`.
Genera un archivo donde cada línea del original está numerada, usando
`fmt.Fprintf` para escribir las líneas formateadas.

→ `05-numerar-lineas/`
