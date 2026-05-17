package web

import (
	"io"
	"regexp"
	"strings"

	"github.com/nguyenthenguyen/docx"
)

func ReadTXT(reader io.Reader) string {

	buf := new(strings.Builder)

	io.Copy(buf, reader)

	return buf.String()
}

func ReadDOCX(path string) (string, error) {

	doc, err := docx.ReadDocxFile(path)
	if err != nil {
		return "", err
	}

	defer doc.Close()

	// Extract raw content
	raw := doc.Editable().GetContent()

	// Remove XML tags
	re := regexp.MustCompile(`<[^>]*>`)

	clean := re.ReplaceAllString(raw, " ")

	// Normalize spacing
	clean = strings.Join(strings.Fields(clean), " ")

	return clean, nil
}