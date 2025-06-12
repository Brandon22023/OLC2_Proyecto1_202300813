package symbols

import (
    "fmt"
    "os"
)

type Symbol struct {
    ID        string
    SymType   string // "Variable" o "Función"
    DataType  string
    Scope     string
    Line      int
    Column    int
}

type SymbolTable struct {
    Symbols []Symbol
}

func NewSymbolTable() *SymbolTable {
    return &SymbolTable{Symbols: []Symbol{}}
}

func (st *SymbolTable) AddSymbol(sym Symbol) {
    // Evita duplicados en el mismo ámbito
    for _, s := range st.Symbols {
        if s.ID == sym.ID && s.Scope == sym.Scope {
            return // No agregar duplicado
        }
    }
    st.Symbols = append(st.Symbols, sym)
}

// Genera un archivo HTML con la tabla de símbolos
func (st *SymbolTable) ToHTML(filename string) error {
    html := `
<!DOCTYPE html>
<html lang="es">
<head>
<meta charset="UTF-8">
<title>Tabla de Símbolos</title>
<style>
    body {
        background: linear-gradient(120deg, #232526, #414345);
        font-family: 'Segoe UI', 'Roboto', Arial, sans-serif;
        color: #f5f6fa;
        margin: 0;
        padding: 40px;
    }
    h1 {
        text-align: center;
        color: #00b894;
        letter-spacing: 2px;
        margin-bottom: 30px;
        text-shadow: 1px 1px 8px #222;
    }
    table {
        border-collapse: collapse;
        margin: 0 auto;
        min-width: 900px;
        background: rgba(44, 62, 80, 0.98);
        border-radius: 18px;
        overflow: hidden;
        box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.37);
    }
    th, td {
        padding: 16px 22px;
        text-align: center;
    }
    th {
        background: linear-gradient(90deg, #00b894 60%, #0984e3 100%);
        color: #fff;
        font-size: 1.1em;
        letter-spacing: 1px;
        border-bottom: 3px solid #636e72;
    }
    tr {
        transition: background 0.2s;
    }
    tr:nth-child(even) {
        background: rgba(0, 184, 148, 0.08);
    }
    tr:hover {
        background: #636e72;
        color: #fff;
        cursor: pointer;
    }
    td {
        font-size: 1em;
        border-bottom: 1px solid #636e72;
    }
    .scope-global {
        color: #00cec9;
        font-weight: bold;
    }
    .scope-if, .scope-IF {
        color: #fdcb6e;
        font-weight: bold;
    }
    .scope-funcion, .scope-Función {
        color: #6c5ce7;
        font-weight: bold;
    }
    .scope-for {
        color: #e17055;
        font-weight: bold;
    }
</style>
</head>
<body>
<h1>Tabla de Símbolos</h1>
<table id="tablaSimbolos" style="display:none;">
<thead>
<tr>
<th>ID</th>
<th>Tipo símbolo</th>
<th>Tipo dato</th>
<th>Ámbito</th>
<th>Línea</th>
<th>Columna</th>
</tr>
</thead>
<tbody>
</tbody>
</table>
<script>
const rows = ` + "`"

    // Prepara las filas como un string para el efecto typing
    rows := ""
    for _, s := range st.Symbols {
        scopeClass := ""
        switch s.Scope {
        case "global", "GLOBAL":
            scopeClass = "scope-global"
        case "if-block", "IF":
            scopeClass = "scope-if"
        case "Función", "funcion":
            scopeClass = "scope-funcion"
        case "for-block", "FOR":
            scopeClass = "scope-for"
        }
        rows += fmt.Sprintf(
            "<tr><td>%s</td><td>%s</td><td>%s</td><td class='%s'>%s</td><td>%d</td><td>%d</td></tr>",
            s.ID, s.SymType, s.DataType, scopeClass, s.Scope, s.Line, s.Column)
    }

    html += rows + "`" + `;
const table = document.getElementById('tablaSimbolos');
const tbody = table.getElementsByTagName('tbody')[0];
let i = 0;
let tempRow = '';
table.style.display = '';
function typeRow() {
    if (i >= rows.length) return;
    tempRow += rows[i];
    tbody.innerHTML = tempRow;
    i++;
    if (rows[i-1] === '>') {
        setTimeout(typeRow, 60);
    } else {
        setTimeout(typeRow, 4);
    }
}
typeRow();
</script>
</body>
</html>
`
    return os.WriteFile(filename, []byte(html), 0644)
}