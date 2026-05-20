# Ejercicios: Tipos parametrizables (genéricos)

## 01 - Búsqueda genérica

Completar las funciones `Contiene` y `Posicion` en `01-busqueda/busqueda.go`.
Ambas son genéricas y usan el constraint `comparable`.

→ `01-busqueda/`

## 02 - Ordenamiento genérico

Completar la función `OrdenarBurbuja` en `02-ordenamiento/ordenar.go`.
Usa una función comparadora para determinar el orden, lo que permite
ordenar cualquier tipo de dato. El tipo es `[T any]` porque la comparación
se delega a la función `menor`, no se hace directamente sobre `T`.

→ `02-ordenamiento/`

## 03 - Máximo y mínimo

Completar las funciones `Maximo` y `Minimo` en `03-max-min/maxmin.go`.
Usan un constraint personalizado `Ordenable` que agrupa tipos
con orden natural.

→ `03-max-min/`

## 04 - Filtrado genérico

Completar la función `Filtrar` en `04-filtrar/filtrar.go`.
Recibe un slice y una función condición, y devuelve un nuevo slice
con los elementos que la cumplen. Acá también se usa `[T any]` porque
la condición la evalúa la función pasada como parámetro, no el código
de `Filtrar`.

→ `04-filtrar/`
