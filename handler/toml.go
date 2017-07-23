package handler

import (
	"fmt"
	"net/http"
	"os/exec"

	"bytes"
	"resumebuilder/model"
	"resumebuilder/resumebuilder"
)

// ParseTOML --
func (h *Handler) ParseTOML(w http.ResponseWriter, req *http.Request) {
	var config model.Config
	var buf bytes.Buffer
	if err := resumebuilder.ReadTOML(req.Body, &config); err != nil {
		fmt.Fprintf(w, "failed to read [%s]", err)
		return
	}
	defer req.Body.Close()
	tmpl := resumebuilder.Make(config)
	cmd := exec.Command("docker", "run", "-i", "latex")
	cmd.Stdin = bytes.NewBufferString(tmpl)
	cmd.Stdout = &buf
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(w, "command failed [%s]", err)
		return
	}
	fmt.Fprintf(w, "%s", buf.Bytes())
}

// GetTexFromTOML --
func (h *Handler) GetTexFromTOML(w http.ResponseWriter, req *http.Request) {
	var config model.Config
	if err := resumebuilder.ReadTOML(req.Body, &config); err != nil {
		fmt.Fprintf(w, "failed to read [%s]", err)
		return
	}
	defer req.Body.Close()
	tmpl := resumebuilder.Make(config)
	fmt.Fprintf(w, "%s", tmpl)
}
