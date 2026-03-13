package utils

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func GuardarImagen(imagen *multipart.FileHeader) (nombreImg *string, err error) {
	file, err := imagen.Open()
	if err != nil {
		return nil, errors.New("Error al abrir archivo")
	}
	defer file.Close()
	uploadDir := "./uploads"
	err = os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("Error al crear carpeta: %v", err)
	}
	ext := filepath.Ext(imagen.Filename)
	timestamp := time.Now().UnixNano() / 1000
	nombreUnico := fmt.Sprintf("%d%s", timestamp, ext)
	fmt.Println(nombreUnico)
	out, err := os.Create(filepath.Join(uploadDir, nombreUnico))
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Error al guardar archivo")
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return nil, errors.New("Error al copiar archivo")
	}

	return &nombreUnico, nil
}
