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
└── ejercicios/{promedio,aplicar}/

05-arreglos-slices/          # ← capítulo 2-5
├── ejemplos/arreglos/
└── ejercicios/{invertir,rotar,eliminar,eliminar-duplicados}/

06-maps/                     # ← capítulo 2-6
├── ejemplos/
└── ejercicios/{contar-palabras,igual,anagramas}/

07-punteros/                 # ← capítulo 2-7
├── ejemplos/punteros/
└── ejercicios/{swap,sumar-punteros,dividir,inicializar-arreglo,maximo}/

08-structs-interfaces/       # ← capítulo 2-8
├── ejemplos/
└── ejercicios/notificaciones/

09-archivos/                 # ← capítulo 2-9
├── ejemplos/{leer-completo,leer-lineas,escribir}/
└── ejercicios/{contar-lineas,copiar-archivo,sumar-numeros,agregar-linea,numerar-lineas}/

10-errores/                  # ← capítulo 2-10
├── ejemplos/
└── ejercicios/{dividir,clasificar,buscar-producto,extraer,leer-config}/

11-oop/                      # ← capítulo 2-11
├── ejemplos/figuras/
└── ejercicios/sistema-figuras/
```

Cada directorio `ejercicios/` contiene un `README.md` con los enunciados y esqueletos con tests.

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
make test
```

O con más detalle:

```bash
make test-v
```

Para más información sobre cómo trabajar, ver [CONTRIBUTING.md](CONTRIBUTING.md).

## Requisitos

- Go 1.20 o superior
- Opcional: [golangci-lint](https://golangci-lint.run/) para linting local

## Licencia

MIT
