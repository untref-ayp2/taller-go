# Cómo trabajar en este repositorio

## Ejecutar los tests

Ejecutá todos los tests:

    make test

Ejecutar los tests con detalle de cada caso:

    make test-v

Ejecutar los tests de un ejercicio específico:

    go test -v ./03-elementos-basicos/ejercicios/01-temperatura/...

Para filtrar por nombre de test:

    go test -run TestNombre ./...

## Feedback automático (CI)

Cada vez que hacés `git push`, GitHub Actions ejecuta los tests automáticamente.
Vas a ver el resultado en la pestaña **Actions** de tu repositorio.

Si los tests pasan: se muestra un tick verde.
Si algún test falla: se muestra una cruz roja — revisá el error, corregí el código y volvé a pushear.

## Tu repositorio tiene un PR "Feedback"

Cuando empezaste la assignment, GitHub Classroom creó automáticamente un Pull Request
llamado **"Feedback"** en tu repositorio. Este PR es tu canal de comunicación con los docentes.

**No cierres este PR. Si lo cerrás, perdés el canal de ayuda.**

### Cómo funcionan los commits en el PR "Feedback"

Cada vez que hacés `git push`, tus commits aparecen automáticamente en el PR "Feedback".
No necesitás hacer nada especial — el PR se actualiza solo.

### Cómo pedir ayuda

1. **Commitear y pushear** tu intento (aunque no esté completo):

       git add .
       git commit -m "intento ejercicio 01-temperatura"
       git push

2. **Ir al PR "Feedback"** — en la pestaña Pull Requests de tu repositorio,
   click en el PR llamado "Feedback".

3. **Agregar un comentario** en el PR explicando:
   - Qué ejercicio estás haciendo.
   - Qué intentaste.
   - Qué error estás viendo.

4. **Mencionar al docente** — Para mencionar a un docente podes arrobarlo
   con el nombre de usuario o el correo electrónico, en el comentario para
   que reciba una notificación:

       @nombredeusuario ayuda con el ejercicio de temperatura

### Normas para pedir ayuda

- **Un solo PR por repositorio** — no necesitás abrir más PRs; el PR "Feedback"
  ya existe y es el canal correcto.
- **Antes de pedir ayuda** — intentá resolver el ejercicio por tu cuenta al
  menos 3 veces. En el comentario explicá qué intentaste.
- **No pedir que te lo resuelvan** — el objetivo es que aprendas, no que te
  den la solución.
- **No cerrar el PR** — el PR "Feedback" debe mantenerse abierto para que los
  docentes puedan verte y responderte.

## Convenciones de código

- **Ejercicios**: son `package main` con `main.go` + `main_test.go`.
- **Ejemplos**: son packages con nombre significativo (`condicionales`, `ciclos`, `genericas`).
- **Indentación**: tabs (ya configurado en `.editorconfig`).

## Estructura de capítulos

```
01-introduccion/             # capítulo 2-1
02-paquetes-y-modulos/       # capítulo 2-2
03-elementos-basicos/        # capítulo 2-3 (tipos, variables, consts, condicionales, ciclos)
04-funciones/                # capítulo 2-4
05-arreglos-slices/          # capítulo 2-5
06-maps/                     # capítulo 2-6
07-punteros/                 # capítulo 2-7
08-structs-interfaces/       # capítulo 2-8
09-archivos/                 # capítulo 2-9
10-errores/                  # capítulo 2-10
11-oop/                      # capítulo 2-11
12-genericos/                # capítulo 2-12
```

Cada capítulo tiene `ejemplos/` (código resuelto) y `ejercicios/` (esqueletos con `// TODO`).

## Comandos útiles

    make fmt    # formatear todo el código
    make lint   # verificar estilo con linter
    make build  # compilar todo sin ejecutar
    make clean  # limpiar archivos generados

    go run ./01-introduccion/ejemplos/00-hola  # ejecutar un ejemplo

## Requisitos

- Go 1.20 o superior.
- Opcional: golangci-lint (https://golangci-lint.run/) para verificar estilo
  localmente.
