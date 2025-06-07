#!/bin/bash

echo "[INFO] Ejecutando ANTLR4 directamente con el .jar..."

JAR="antlr-4.13.2-complete.jar"

# Verifica que el archivo exista
if [ ! -f "$JAR" ]; then
  echo "[ERROR] No se encuentra $JAR en el directorio actual"
  exit 1
fi

# Ejecuta ANTLR con tus gramáticas
java -Xmx500M -cp "$JAR" org.antlr.v4.Tool \
  -Dlanguage=Go \
  -visitor \
  -package compiler \
  -o . \
  parser/Vlang.g4

echo "[INFO] Generación completada."
