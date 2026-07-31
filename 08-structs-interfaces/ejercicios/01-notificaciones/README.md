# Ejercicio: Notificaciones

Implementá un sistema de notificaciones simple:

1. Definí un *struct* `Email` con campos `Direccion string` y
   `Mensaje string`. Implementá `NewEmail(direccion, mensaje string) Email`.

2. Definí un método `Enviar() string` sobre `Email` que devuelva
   `"Enviando email a <direccion>: <mensaje>"`.

3. Definí un *struct* `SMS` con campos `Numero string` y
   `Mensaje string`. Implementá `NewSMS(numero, mensaje string) SMS`.

4. Definí un método `Enviar() string` sobre `SMS` que devuelva
   `"Enviando SMS a <numero>: <mensaje>"`.

5. Definí una interfaz `Notificable` con el método `Enviar() string`.
   Escribí una función `EnviarNotificaciones(ns []Notificable)` que itere
   e imprima el resultado de `Enviar()` para cada uno.

6. En el `main`, creá un `Email` y un `SMS`, almacenalos en
   `[]Notificable`, y llamá a `EnviarNotificaciones`.
