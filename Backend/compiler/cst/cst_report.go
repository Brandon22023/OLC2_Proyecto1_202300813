package cst

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

type CSTresponse struct {
	SVGTree string `json:"svgtree"`
}

func ReadFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, _ := io.ReadAll(file)
	return string(content)
}

func CstReport(input string) string {

	// get the content (relative to this path) ../compiler/TSwiftLanguage.g4

	parserContent := ""

	_, filename, _, _ := runtime.Caller(0)

	path := filepath.Dir(filename)

	// remove \cst from the path
	path = path[:len(path)-4]

	parser, err := json.Marshal(ReadFile(filepath.Join(path, "/parser/Vlang.g4")))

	if err != nil {
		fmt.Println(err)
	}
	parserContent = string(parser)

	lexerContent := ""
	lexer, err := json.Marshal(ReadFile(filepath.Join(path, "/parser/Vlang.g4")))

	if err != nil {
		fmt.Println(err)
	}

	lexerContent = string(lexer)

	jinput, _:= json.Marshal(input)
	finput := string(jinput)

	payload := []byte(
		fmt.Sprintf(
			`{
				"grammar": %s,
				"input": %s,
				"lexgrammar": %s,
				"start": "%s"
			}`,
			parserContent,
			finput,
			lexerContent,
			"programa",
		),
	)

	req, err := http.NewRequest("POST", "http://lab.antlr.org/parse/", bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return ""
	}

	// create a map to store the json
	var data map[string]interface{}
    
	fmt.Println("Respuesta ANTLR:", string(body))
	// unmarshal the json
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshalling json:", err)
		return ""
	}

	// Validación segura de campos
    result, ok := data["result"].(map[string]interface{})
    if !ok || result == nil {
        fmt.Println("No se encontró el campo 'result' en la respuesta:", string(body))
        return ""
    }

    svgtree, ok := result["svgtree"].(string)
    if !ok {
        fmt.Println("No se encontró el campo 'svgtree' en la respuesta:", result)
        return ""
    }

    return svgtree
}

func SaveCSTSVG(svgContent, filename string) error {
    return os.WriteFile(filename, []byte(svgContent), 0644)
}