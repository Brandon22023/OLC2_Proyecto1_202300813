
# Manual Técnico - Intérprete VLangCherry

## Universidad de San Carlos de Guatemala  
**Facultad de Ingeniería**  
**Escuela de Ingeniería en Ciencias y Sistemas**  
**Curso:** Organización de Lenguajes y Compiladores 2  
**Proyecto 1 - Vaciones del Primer semestre 2025**  
**Lenguaje:** VLangCherry  
**Gupo:** 6

---

## 1. Introducción

Este manual técnico documenta la implementación de un intérprete para el lenguaje de programación VLangCherry. Este lenguaje, inspirado en Go, está diseñado para ilustrar conceptos fundamentales de compiladores, y ser ejecutado eficientemente en sistemas de bajos recursos.

---

## 2. Tecnologías Utilizadas

- **ANTLR**: Generador del analizador léxico y sintáctico.
- **Go (Golang)**: Lenguaje base para el desarrollo del intérprete.
- **Linux**: Plataforma objetivo.
- **Frameworks UI**: Fyne (opcional).

---

## 3. Arquitectura del Sistema

### 3.1 Flujo General

1. **Entrada**: Código fuente `.vch`
2. **Análisis Léxico/Sintáctico**: ANTLR
3. **AST y Análisis Semántico**: Implementado en Go
4. **Interpretación y Ejecución**
5. **Generación de Reportes**
6. **Salida en consola/interfaz**

### 3.2 Componentes Principales

- **Editor de código**
  - Crear, abrir y guardar archivos `.vch`
  - Múltiples pestañas y líneas visibles

- **Herramientas**
  - Botón "Ejecutar" que llama al intérprete

- **Reportes**
  - Errores léxicos/sintácticos/semánticos
  - Tabla de símbolos
  - AST visual

- **Consola**
  - Salida de ejecución
  - Mensajes de error y advertencias

---

## 4. Análisis y Procesamiento

### 4.1 ANTLR
- Se definió la gramática de VLangCherry
- Archivos generados: Lexer, Parser, ASTListener/Visitor

### 4.2 Árbol de Sintaxis Abstracta (AST)
- Estructura que representa el código
- Recorrido con patrones Visitor
- Base para análisis semántico y ejecución

### 4.3 Tabla de Símbolos
- Variables, funciones, structs, slices
- Tipos, ámbito, posición

### 4.4 Reportes
- Errores (tipo, ubicación, descripción)
- AST (visual o JSON)
- Símbolos

---

## 5. Características del Lenguaje VLangCherry

### 5.1 Tipado Estático
- No permite cambiar tipo tras la declaración

### 5.2 Tipos de Datos

| Tipo      | Descripción           | Valor por defecto |
|-----------|------------------------|-------------------|
| int       | Enteros                | 0                 |
| float64   | Flotantes 64 bits      | 0.0               |
| string    | Cadenas de texto       | ""                |
| bool      | Booleanos              | false             |
| rune      | Caracteres             | '\0'              |
| nil       | Ausencia de valor (referencias) |

### 5.3 Estructuras de Datos
- **Slices**: []int, []string, etc.
- **Slices multidimensionales**
- **Structs**: Tipos personalizados

### 5.4 Sentencias de Control

- `if`, `else if`, `else`
- `switch-case`
- `for`, `for-in`
- `break`, `continue`, `return`

### 5.5 Funciones

- Definición global
- Parámetros con tipo
- Retorno único
- Paso por valor/referencia

---

## 6. Funciones Embebidas

- `print(...)`
- `Atoi(string) → int`
- `parseFloat(string) → float64`
- `typeof(valor) → string`

---

## 7. Reglas Semánticas

- Verificación de tipos
- Identificadores únicos por ámbito
- No acceso fuera del ámbito
- No redefinición en mismo bloque
- No operaciones sobre `nil`
- Concatenación solo entre strings
- Verificación de tipos en operadores aritméticos, lógicos y de comparación

---

## 8. Reglas Léxicas

### Comentarios
- Una línea: `//`
- Multilínea: `/* ... */`

### Identificadores válidos
- Letras, números, guión bajo
- No comenzar con número
- No palabras reservadas

---

## 9. Interfaz Gráfica

- Implementada con:
  - Fyne (Go)
  - Electron + REST API (Go backend)

### Funcionalidades Mínimas

- Crear, abrir, guardar archivos `.vch`
- Ejecutar código
- Visualizar reportes y consola

---

## 10. Reportes Generados

### 10.1 Reporte de Errores

| No. | Descripción                        | Línea | Columna | Tipo      |
|-----|------------------------------------|-------|---------|-----------|
| 1   | El struct “Persona” no fue definido| 5     | 1       | Semántico |

### 10.2 Tabla de Símbolos

| ID       | Tipo símbolo | Tipo dato | Ámbito    | Línea | Columna |
|----------|---------------|-----------|-----------|--------|---------|
| x        | Variable       | int       | Global    | 2      | 5       |

### 10.3 AST
- Representación visual/jerárquica del código
- Se puede exportar como imagen o JSON

---

