# Taller de Go

Repositorio complementario de la sección **Taller de Go** de los apuntes de Algoritmos y Programación II.

## Estructura

```
01-introduccion/             # ← capítulo 2-1
├── ejemplos/00-hola/
└── ejercicios/saludo-personalizado/

02-paquetes-y-modulos/       # ← capítulo 2-2
└── ejemplos/{01-crear-modulo,02-importar-paquete,03-dependencias}/

03-elementos-basicos/        # ← capítulo 2-3 (tipos, variables, consts, condicionales, ciclos)
├── ejemplos/{01-tipodatos,02-variables,03-constantes,04-condicionales,05-ciclos}/
└── ejercicios/{01-temperatura,02-positivo-negativo,03-dia-semana,04-suma-1-100,05-maximo}/

04-funciones/                # ← capítulo 2-4
├── ejemplos/{genericas,matematicas}/
└── ejercicios/{01-promedio,02-aplicar}/

05-arreglos-slices/          # ← capítulo 2-5
├── ejemplos/arreglos/
└── ejercicios/{01-invertir,02-rotar,03-eliminar,04-eliminar-duplicados}/

06-maps/                     # ← capítulo 2-6
├── ejemplos/
└── ejercicios/{01-contar-palabras,02-igual,03-anagramas}/

07-punteros/                 # ← capítulo 2-7
├── ejemplos/punteros/
└── ejercicios/{01-swap,02-sumar-punteros,03-dividir,04-inicializar-arreglo,05-concatenar}/

08-structs-interfaces/       # ← capítulo 2-8
├── ejemplos/
└── ejercicios/01-notificaciones/

09-archivos/                 # ← capítulo 2-9
├── ejemplos/{leer-completo,leer-lineas,escribir}/
└── ejercicios/{01-contar-lineas,02-copiar-archivo,03-sumar-numeros,04-agregar-linea,05-numerar-lineas}/

10-errores/                  # ← capítulo 2-10
├── ejemplos/
└── ejercicios/{01-dividir,02-clasificar,03-buscar-producto,04-extraer,05-leer-config}/

11-oop/                      # ← capítulo 2-11
├── ejemplos/figuras/
└── ejercicios/01-sistema-empleados/

12-genericos/                # ← capítulo 2-12
├── ejemplos/
└── ejercicios/{01-busqueda,02-ordenamiento,03-max-min,04-filtrar}/
```

Cada directorio `ejercicios/` contiene un `README.md` con los enunciados y esqueletos con tests.

## Cómo usar

```bash
# Ejecutar todos los tests
make test

# Con detalle de cada caso
make test-v

# Tests de un ejercicio específico
go test -v ./03-elementos-basicos/ejercicios/01-temperatura/...
```

Para ejecutar un ejemplo:

```bash
go run ./01-introduccion/ejemplos/00-hola
```

Para más información, ver [CONTRIBUTING.md](CONTRIBUTING.md).

## Requisitos

- Go 1.22 o superior
- Opcional: [golangci-lint](https://golangci-lint.run/) para linting local

## Licencia

MIT
