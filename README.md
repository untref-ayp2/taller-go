# Taller de Go

Repositorio complementario de la sección **Taller de Go** de los apuntes de Algoritmos y Programación II.

## Estructura

```
01-introduccion/           # ← capítulo 2-1 del apunte
└── ejemplos/00-hola/

02-paquetes-y-modulos/     # ← capítulo 2-2
└── ejemplos/

03-elementos-basicos/      # ← capítulo 2-3 (tipos, variables, consts, condicionales, ciclos)
└── ejemplos/{01-tipodatos,02-variables,03-constantes,04-condicionales,05-ciclos}/

04-funciones/              # ← capítulo 2-4
└── ejemplos/{genericas,matematicas}/

05-arreglos-slices/        # ← capítulo 2-5
└── ejemplos/arreglos/

06-maps/                   # ← capítulo 2-6
└── ejemplos/

07-punteros/               # ← capítulo 2-7
└── ejemplos/punteros/

08-structs-interfaces/     # ← capítulo 2-8
└── ejemplos/

09-archivos/               # ← capítulo 2-9
└── ejemplos/{leer-completo,leer-lineas,escribir}/

10-errores/                # ← capítulo 2-10
└── ejemplos/

11-oop/                    # ← capítulo 2-11
└── ejemplos/figuras/
```

Cada tema tiene una carpeta `ejercicios/` con esqueletos para resolver y tests asociados.

## Cómo usar

```bash
git clone https://github.com/untref-ayp2/taller-go.git
cd taller-go
```

Para ejecutar un ejemplo:

```bash
go run ./01-introduccion/ejemplos/00-hola
```

Para ejecutar todos los tests:

```bash
go test ./...
```

## Ejercicios

Los esqueletos de ejercicios están en `NN-tema/ejercicios/` con `// TODO: implementar`.
Las soluciones están en la rama `soluciones`:

```bash
git checkout soluciones
```

## Requisitos

Go 1.20 o superior.
