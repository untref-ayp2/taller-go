# Ejercicios: Archivos

1. Escribí un programa que lea un archivo de texto y muestre por
   pantalla la cantidad de líneas que contiene.
   → `contar-lineas/`

2. Escribí una función `CopiarArchivo(origen, destino string) error`
   que copie el contenido de un archivo a otro. Usá `os.ReadFile` y
   `os.WriteFile`.
   → `copiar-archivo/`

3. Escribí un programa que lea un archivo `numeros.txt` donde cada línea
   contiene un número entero, los sume y muestre el resultado. Usá
   `bufio.Scanner` y `strconv.Atoi` para convertir cada línea a `int`.
   → `sumar-numeros/`

4. Escribí una función `AgregarLinea(archivo, linea string) error` que
   agregue una línea al final de un archivo existente. Usá `os.OpenFile`
   con la bandera `os.O_APPEND|os.O_WRONLY`.
   → `agregar-linea/`

5. Escribí un programa que dado un archivo de texto, genere otro archivo
   donde cada línea esté numerada (similar al ejemplo de `bufio.Scanner`).
   → `numerar-lineas/`
