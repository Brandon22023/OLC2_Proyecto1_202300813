package errors

import (
    "fmt"
    "os"
    "compiler/repl"
)

func SaveErrorsHTML(errorTable *repl.ErrorTable, ruta string) error {
    html := `
<!DOCTYPE html>
<html lang="es">
<head>
<meta charset="UTF-8">
<title>Reporte de Errores</title>
<style>
    body {
        background: linear-gradient(120deg, #232526, #ff7675 120%);
        font-family: 'Segoe UI', 'Roboto', Arial, sans-serif;
        color: #fff;
        margin: 0;
        padding: 40px;
    }
    h1 {
        text-align: center;
        color: #d63031;
        letter-spacing: 2px;
        margin-bottom: 30px;
        text-shadow: 1px 1px 8px #222;
        font-size: 2.5em;
    }
    .icon {
        font-size: 1.3em;
        vertical-align: middle;
        margin-right: 8px;
    }
    table {
        border-collapse: collapse;
        margin: 0 auto;
        min-width: 950px;
        background: rgba(44, 62, 80, 0.98);
        border-radius: 18px;
        overflow: hidden;
        box-shadow: 0 8px 32px 0 rgba(214, 48, 49, 0.37);
        animation: fadeIn 1.2s;
    }
    @keyframes fadeIn {
        from { opacity: 0; transform: translateY(40px);}
        to { opacity: 1; transform: translateY(0);}
    }
    th, td {
        padding: 16px 22px;
        text-align: center;
    }
    th {
        background: linear-gradient(90deg, #d63031 60%, #fdcb6e 100%);
        color: #fff;
        font-size: 1.1em;
        letter-spacing: 1px;
        border-bottom: 3px solid #636e72;
    }
    tr {
        transition: background 0.2s;
    }
    tr:nth-child(even) {
        background: rgba(214, 48, 49, 0.08);
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
    .tipo-lexico {
        color: #00b894;
        font-weight: bold;
        letter-spacing: 1px;
    }
    .tipo-sintactico {
        color: #0984e3;
        font-weight: bold;
        letter-spacing: 1px;
    }
    .tipo-semantico {
        color: #fdcb6e;
        font-weight: bold;
        letter-spacing: 1px;
    }
    .icon-lexico::before {
        content: "📝";
    }
    .icon-sintactico::before {
        content: "⚠️";
    }
    .icon-semantico::before {
        content: "❌";
    }
    .desc {
        text-align: left;
        font-size: 1.05em;
        font-family: 'Segoe UI', 'Roboto', Arial, sans-serif;
    }
</style>
</head>
<body>
<h1>Reporte de Errores</h1>
<table id="tablaErrores" style="display:none;">
<thead>
<tr>
    <th>No.</th>
    <th>Descripción</th>
    <th>Línea</th>
    <th>Columna</th>
    <th>Tipo</th>
</tr>
</thead>
<tbody>
</tbody>
</table>
<script>
const rows = ` + "`"

    // Prepara las filas como un string para el efecto typing
    rows := ""
    for i, err := range errorTable.Errors {
        tipoClass := ""
        iconClass := ""
        switch err.Type {
        case "Error léxico":
            tipoClass = "tipo-lexico"
            iconClass = "icon icon-lexico"
        case "Error sintáctico":
            tipoClass = "tipo-sintactico"
            iconClass = "icon icon-sintactico"
        case "Error semántico":
            tipoClass = "tipo-semantico"
            iconClass = "icon icon-semantico"
        default:
            tipoClass = ""
            iconClass = ""
        }
        rows += fmt.Sprintf(
            "<tr><td>%d</td><td class='desc'>%s</td><td>%d</td><td>%d</td><td class='%s'><span class='%s'></span>%s</td></tr>",
            i+1, err.Msg, err.Line, err.Column, tipoClass, iconClass, err.Type)
    }

    html += rows + "`" + `;
const table = document.getElementById('tablaErrores');
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
    return os.WriteFile(ruta, []byte(html), 0644)
}

func RemoveDuplicateErrors(errors []repl.Error) []repl.Error {
    unique := make([]repl.Error, 0, len(errors))
    seen := make(map[string]struct{})
    for _, err := range errors {
        key := fmt.Sprintf("%s|%d|%d", err.Msg, err.Line, err.Column)
        if _, exists := seen[key]; !exists {
            seen[key] = struct{}{}
            unique = append(unique, err)
        }
    }
    return unique
}