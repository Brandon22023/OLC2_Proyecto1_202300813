package main

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"math"
	"net/http"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// --- Botón personalizado sin fondo azul ---
type CustomButton struct {
    widget.BaseWidget
    Label       string
    BgColor     color.Color
    BorderColor color.Color
    OnTapped    func()
}

func NewCustomButton(label string, bg color.Color, border color.Color, tapped func()) *CustomButton {
    btn := &CustomButton{
        Label:       label,
        BgColor:     bg,
        BorderColor: border,
        OnTapped:    tapped,
    }
    btn.ExtendBaseWidget(btn)
    return btn
}
func gradientRect(width, height int) *canvas.Image {
    cornerRadius := 9.0 // Debe coincidir con el CornerRadius del borde
    img := image.NewNRGBA(image.Rect(0, 0, width, height))
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            // Cálculo para esquinas redondeadas
            dx := math.Min(float64(x), float64(width-1-x))
            dy := math.Min(float64(y), float64(height-1-y))
            if dx < cornerRadius && dy < cornerRadius {
                dist := math.Hypot(cornerRadius-dx, cornerRadius-dy)
                if dist > cornerRadius {
                    img.Set(x, y, color.NRGBA{0, 0, 0, 0}) // Transparente
                    continue
                }
            }
            // Degradado de azul marino oscuro a azul más claro
            r := uint8(20 + (30*y)/height)
            g := uint8(30 + (40*y)/height)
            b := uint8(60 + (100*y)/height)
            img.Set(x, y, color.NRGBA{R: r, G: g, B: b, A: 255})
        }
    }
    return canvas.NewImageFromImage(img)
}

func (b *CustomButton) CreateRenderer() fyne.WidgetRenderer {
    // Fondo degradado
    grad := gradientRect(170, 44)
    grad.FillMode = canvas.ImageFillStretch
    grad.SetMinSize(fyne.NewSize(170, 44))

    border := canvas.NewRectangle(b.BorderColor)
    border.CornerRadius = 22 
    border.SetMinSize(fyne.NewSize(174, 48))

    text := canvas.NewText(b.Label, color.White)
    text.Alignment = fyne.TextAlignCenter
    text.TextStyle = fyne.TextStyle{Bold: true}
    text.TextSize = 16

    cont := container.NewStack(
        border,
        grad,
        container.NewCenter(text),
    )
    return widget.NewSimpleRenderer(cont)
}

func (b *CustomButton) Tapped(_ *fyne.PointEvent) {
    if b.OnTapped != nil {
        b.OnTapped()
    }
    b.Refresh()
}

// Reemplaza styledButton por este:
func styledButton(label string, bg color.Color, border color.Color, tapped func()) fyne.CanvasObject {
    return NewCustomButton(label, bg, border, tapped)
}

type TimesTheme struct{}

func (t *TimesTheme) Font(s fyne.TextStyle) fyne.Resource {
    res, err := fyne.LoadResourceFromPath("fonts/Times New Roman.ttf")
    if err != nil {
        return theme.DefaultTheme().Font(s)
    }
    return res
}
func (t *TimesTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
    if n == theme.ColorNameForeground {
        return color.White // Texto blanco
    }
    if n == theme.ColorNameBackground {
        return color.RGBA{36, 41, 56, 255} // Fondo oscuro
    }
    return theme.DefaultTheme().Color(n, v)
}
func (t *TimesTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
    return theme.DefaultTheme().Icon(n)
}
func (t *TimesTheme) Size(n fyne.ThemeSizeName) float32 {
    if n == theme.SizeNameText {
        return 20 // Tamaño de letra grande
    }
    return theme.DefaultTheme().Size(n)
}
func main() {
    a := app.NewWithID("olc2.proyecto1.202300813")
    a.Settings().SetTheme(&TimesTheme{}) 
    w := a.NewWindow("OLC2 Proyecto 1 - IDE")

    // Colores modernos y vibrantes
    bgColor := color.RGBA{R: 36, G: 41, B: 56, A: 255}
    panelColor := color.RGBA{R: 36, G: 41, B: 56, A: 255}

    // Fondo de la ventana
    background := canvas.NewRectangle(bgColor)
    var archivoActual fyne.URI
    // --- Editor sin números de línea ---
    entrada := widget.NewMultiLineEntry()
    entrada.SetText("// Escribe tu código aquí...")
    entrada.Wrapping = fyne.TextWrapWord
    // Botones con colores personalizados
    
    // --- Consola ---
    salida := widget.NewMultiLineEntry()
    salida.SetText("// Salida del programa...")
    salida.Wrapping = fyne.TextWrapWord
    //salida.Disable()
    var btnReportes fyne.CanvasObject
    buttons := container.NewHBox(
    layout.NewSpacer(),
    styledButton("Crear archivo",
    color.RGBA{R: 70, G: 130, B: 180, A: 255},    // Azul acero apagado
    color.RGBA{R: 120, G: 144, B: 156, A: 120},   // Gris azulado suave
    func() {
        dialog := dialog.NewFileSave(
            func(writer fyne.URIWriteCloser, err error) {
                if err != nil || writer == nil {
                    return
                }
                defer writer.Close()
                // Archivo vacío, sin contenido
                _, err = writer.Write([]byte(""))
                if err != nil {
                    dialog.ShowError(err, w)
                    return
                }
                archivoActual = writer.URI() // Guarda la ruta para futuros guardados
                entrada.SetText("") // Limpia el editor
                dialog.ShowInformation("Archivo creado", "¡Archivo .vch creado correctamente!", w)
            }, w)
        dialog.SetFileName("nuevo.vch")
        dialog.Show()
    }),
    styledButton("Abrir archivo",
            color.RGBA{R: 120, G: 144, B: 156, A: 255},
            color.RGBA{R: 176, G: 190, B: 197, A: 120},
            func() {
                dialog := dialog.NewFileOpen(
                    func(reader fyne.URIReadCloser, err error) {
                        if err != nil || reader == nil {
                            return
                        }
                        defer reader.Close()
                        data, err := io.ReadAll(reader)
                        if err == nil {
                            entrada.SetText(string(data))
                            archivoActual = reader.URI() // Guarda la ruta del archivo abierto
                        }
                    }, w)
                dialog.SetFilter(nil)
                dialog.Show()
            }),
    styledButton("Guardar Archivo",
    color.RGBA{R: 139, G: 195, B: 74, A: 255},
    color.RGBA{R: 205, G: 220, B: 57, A: 120},
    func() {
        if archivoActual != nil {
            writer, err := storage.Writer(archivoActual)
            if err != nil {
                dialog.ShowError(err, w)
                return
            }
            defer writer.Close()
            _, err = writer.Write([]byte(entrada.Text))
            if err != nil {
                dialog.ShowError(err, w)
            } else {
                dialog.ShowInformation("Guardado", "¡Archivo guardado correctamente!", w)
            }
        } else {
            // Si no hay archivo abierto, muestra "Guardar como..."
            dialog := dialog.NewFileSave(
                func(writer fyne.URIWriteCloser, err error) {
                    if err != nil || writer == nil {
                        return
                    }
                    defer writer.Close()
                    _, err = writer.Write([]byte(entrada.Text))
                    if err != nil {
                        dialog.ShowError(err, w)
                        return
                    }
                    archivoActual = writer.URI() // Guarda la ruta para futuros guardados
                    dialog.ShowInformation("Guardado", "¡Archivo guardado correctamente!", w)
                }, w)
            dialog.SetFileName("nuevo.vch")
            dialog.Show()
        }
    }),
    styledButton("Ejecutar",
        color.RGBA{R: 149, G: 117, B: 205, A: 255},   // Morado pastel
        color.RGBA{R: 100, G: 181, B: 246, A: 120},   // Azul claro pastel
        func() {

            go func() {
                code := entrada.Text
                resp, err := http.Post("http://localhost:3000/analizar", "text/plain", strings.NewReader(code))
                if err != nil {
                    salida.SetText("Error al conectar con el backend: " + err.Error())
                    return
                }
                defer resp.Body.Close()
                body, _ := io.ReadAll(resp.Body)
                fmt.Println("Respuesta del backend:", string(body)) // <-- Esto imprime en la terminal
                salida.SetText(string(body))
            }()
        }),
    // Aquí guardamos la referencia
        func() fyne.CanvasObject {
            btnReportes = styledButton(
                "Reportes",
                color.RGBA{R: 255, G: 202, B: 40, A: 255},
                color.RGBA{R: 255, G: 224, B: 130, A: 120},
                func() {
                    menu := fyne.NewMenu("Reportes",
                        fyne.NewMenuItem("Reporte de Errores", func() {
                            go func() {
                                resp, err := http.Get("http://localhost:3000/reporte-errores")
                                if err != nil {
                                    dialog.ShowError(err, w)
                                    return
                                }
                                defer resp.Body.Close()
                                ruta, _ := io.ReadAll(resp.Body)
                                dialog.ShowInformation("Ruta del reporte de errores", string(ruta), w)
                                // Si quieres abrirlo automáticamente en el navegador (Linux):
                                // exec.Command("xdg-open", string(ruta)).Start()
                            }()
                        }),
                        fyne.NewMenuItem("Reporte de Tabla de Símbolos", func() {
                            go func() {
                            resp, err := http.Get("http://localhost:3000/reporte-simbolos")
                            if err != nil {
                                dialog.ShowError(err, w)
                                return
                            }
                            defer resp.Body.Close()
                            ruta, _ := io.ReadAll(resp.Body)
                            // Puedes mostrar la ruta o abrir el HTML con el navegador del sistema
                            dialog.ShowInformation("Ruta del reporte", string(ruta), w)
                            // Si quieres abrirlo automáticamente en el navegador (Linux):
                            // exec.Command("xdg-open", string(ruta)).Start()
                        }()
                        }),
                        fyne.NewMenuItem("Reporte CST", func() {
                            go func() {
                                resp, err := http.Get("http://localhost:3000/reporte-cst")
                                if err != nil {
                                    dialog.ShowError(err, w)
                                    return
                                }
                                defer resp.Body.Close()
                                ruta, _ := io.ReadAll(resp.Body)
                                dialog.ShowInformation("Ruta del CST", string(ruta), w)
                                // Si quieres abrirlo automáticamente en el navegador (Linux):
                                // exec.Command("xdg-open", string(ruta)).Start()
                            }()
                        }),
                    )
                    // Calcula la posición absoluta del botón
                    pos := fyne.CurrentApp().Driver().AbsolutePositionForObject(btnReportes)
                    // Ajusta la posición para que el menú salga justo debajo
                    pos = pos.Add(fyne.NewPos(0, btnReportes.Size().Height))
                    widget.ShowPopUpMenuAtPosition(menu, w.Canvas(), pos)
                },
            )
            return btnReportes
        }(),
    layout.NewSpacer(),
)

    

    editorWithLines := container.NewVScroll(entrada)

    entradaBG := container.NewBorder(
        widget.NewLabelWithStyle("EDITOR", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
        nil, nil, nil,
        editorWithLines,
    )
    entradaBG = container.NewStack(
        canvas.NewRectangle(panelColor),
        entradaBG,
    )

    

    salidaBG := container.NewBorder(
        widget.NewLabelWithStyle("CONSOLA", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
        nil, nil, nil,
        salida,
    )
    salidaBG = container.NewStack(
        canvas.NewRectangle(panelColor),
        salidaBG,
    )

    // Paneles principales lado a lado (horizontal split)
    mainPanels := container.NewHSplit(
        entradaBG,
        salidaBG,
    )
    mainPanels.Offset = 0.5

    content := container.NewBorder(
        buttons,
        nil,
        nil,
        nil,
        mainPanels,
    )

    // Fondo y contenido: el content ocupa todo el espacio central
    w.SetContent(container.NewStack(
        background,
        content,
    ))

    w.Resize(fyne.NewSize(1400, 900)) // ventana más grande
    w.CenterOnScreen()
    w.ShowAndRun()
}