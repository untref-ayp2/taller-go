# Ejercicios: Maps

1. Escribí una función `ContarPalabras` que cuente las palabras en un
   string y devuelva un mapa que mapee las palabras a su número de
   ocurrencias. La función `Split` del paquete `strings` puede ser útil.
   → `contar-palabras/`

2. Escribí una función que compare dos mapas de cadenas y devuelva `true`
   si los mapas contienen las mismas claves y los mismos valores. Usá el
   siguiente prototipo: `func Igual(x, y map[string]int) bool`.
   → `igual/`

3. Los anagramas son palabras que tienen las mismas letras en diferente
   orden. Escribí una función `Anagramas` que tome dos strings y devuelva
   `true` si son anagramas. Usá el prototipo `func Anagramas(s1, s2 string) bool`.
   La complejidad debe ser $O(n)$, donde n es la longitud de los strings.
   → `anagramas/`
