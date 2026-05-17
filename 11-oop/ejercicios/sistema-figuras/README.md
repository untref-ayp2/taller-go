# Ejercicio: Sistema de figuras geométricas

Implementá un programa que modele figuras con los siguientes elementos:

a. Definí una interfaz `Figura` con los métodos `Area() float64` y
   `Perimetro() float64`.

b. Definí una interfaz `Exportable` con el método
   `Exportar(ruta string) error` que guarde la figura en un archivo de
   texto con sus datos (usando `os.WriteFile` o `fmt.Fprintf`).

c. Implementá un `Rectangulo` con los campos `Base` y `Altura`.
   El constructor `NewRectangulo(base, altura float64) (*Rectangulo, error)`
   debe validar que ambos valores sean mayores a cero, devolviendo un
   error en caso contrario.

d. Implementá un `Cuadrado` que componga a `Rectangulo` mediante un
   campo embebido. El constructor `NewCuadrado(lado float64) (*Cuadrado, error)`
   debe crear un cuadrado reutilizando `NewRectangulo`.

e. Tanto `Rectangulo` como `Cuadrado` deben implementar las interfaces
   `Figura` y `Exportable`. El método `Exportar` debe escribir en el
   archivo los datos de la figura (tipo, dimensiones, área y perímetro).

f. En el `main`, creá un slice de `Figura` con un rectángulo y un
   cuadrado, calculá sus áreas y perímetros, y exportá cada figura a
   un archivo separado.
