package types

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Southclaws/supervillain"
)

type Api struct {
	Meta                           Meta
	CreateShortcodeRequestPayload  CreateShortcodeRequestPayload
	CreateShortcodeResponsePayload CreateShortcodeResponsePayload
	GetShortcodesResponsePayload   GetShortcodesResponsePayload
	GetShortcodesMeta              GetShortcodesMeta
}

func TypesInitialization() {
	var withTypes bool
	flag.BoolVar(&withTypes, "types", false, "generate types directory")
	flag.Parse()
	if !withTypes {
		return
	}
	fmt.Println("Generating types...")

	file := []string{"import { z } from \"zod\";"}
	file = append(
		file, supervillain.StructToZodSchema(Api{}),
	)
	indexFilePath, err := filepath.Abs("../../frontend/@types/api.ts")
	if err != nil {
		panic(errors.Join(errors.New("failed to create absolute path for zod shema"), err))
	}
	fmt.Println("Write index file ", indexFilePath)
	err = os.WriteFile(
		indexFilePath,
		[]byte(strings.Join(file, "\n\n")),
		0600,
	)
	if err != nil {
		panic(errors.Join(errors.New("failed to write tozod shema file"), err))
	}

}
