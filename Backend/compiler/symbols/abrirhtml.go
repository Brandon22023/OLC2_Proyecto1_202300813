package symbols

import (
    "fmt"
    "os/exec"
    "runtime"
)

// OpenHTML abre un archivo HTML en el navegador predeterminado
func OpenHTML(path string) error {
    var cmd *exec.Cmd

    switch runtime.GOOS {
    case "linux":
        cmd = exec.Command("xdg-open", path)
    case "windows":
        cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", path)
    case "darwin":
        cmd = exec.Command("open", path)
    default:
        return fmt.Errorf("sistema operativo no soportado")
    }

    return cmd.Start()
}