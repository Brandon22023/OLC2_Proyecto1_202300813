# What Next? 
Seguramente te estas preguntando que debes hacer. 
En primera instancia tenemos que hacer nuestras expresiones, 
definimos una gramatica simple. 

```bash 
Escribe el codigo de Vlang: (Ctrl+D) para terminar
3 + 5 * (2 - 1)
Se termino el analisis sintactico
+-- programa
    |-- expresion
    |   |-- expresion
    |   |   +-- valores
    |   |       +-- valor
    |   |           +-- "3"
    |   |-- "+"
    |   +-- expresion
    |       |-- expresion
    |       |   +-- valores
    |       |       +-- valor
    |       |           +-- "5"
    |       |-- "*"
    |       +-- expresion
    |           |-- "("
    |           |-- expresion
    |           |   |-- expresion
    |           |   |   +-- valores
    |           |   |       +-- valor
    |           |   |           +-- "2"
    |           |   |-- "-"
    |           |   +-- expresion
    |           |       +-- valores
    |           |           +-- valor
    |           |               +-- "1"
    |           +-- ")"
    +-- "<EOF>"
```
Agregue un metodo para generar este arbol cada vez que probamos la estructura.
## Hagamos la declaraciones de las variables 
Nota: Si no le ponemos las labels a los hermanos, no podemos generar nada en la gramatica. 



## Se que existe el Visitor pero...¿Ahora qué? 

Lo primero es que debemos hacer un visitor por las operaciones que queremos hacer
en este caso es interpretar. Por lo que, a partir de esta gramatica, si nos vamos a la carpeta `/parser` notaremos que en `vlang_visitor.go` y en  `base_visitor.go` 
estaran los metodos que nos proponen la interfaz del Visitor y la clase abstracta. veamoslos 
