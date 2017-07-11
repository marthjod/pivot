package model

import (
	"os"
	"strings"
	"testing"
	"text/template"
)

func TestPivio_RenderNetworkDsl(t *testing.T) {
	var (
		yaml                     = "../pivio.yaml"
		templateFile             = "../templates/network-dsl.tpl"
		expectedRenderedTemplate = `
service("web-announcement-service") {
    description    "Web Display of the Announcement"
    attach_to      "hg-CFPA"
    protocol       "tcp"
    port           80

    # WIP
    talks_to    "user-service"
    talks_to    "email-announcement-service"
    }

service("rest-announcement-service") {
    description    "REST API for updating CfP data"
    attach_to      "hg-CFPA"
    protocol       "tcp"
    port           9449

    # WIP
    talks_to    "user-service"
    talks_to    "email-announcement-service"
    }
`
	)
	f, err := os.Open(yaml)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer f.Close()
	p, err := Read(f)
	if err != nil {
		t.Fatal(err)
	}

	tpl, err := template.ParseFiles(templateFile)
	if err != nil {
		t.Fatal(err)
	}

	rendered, err := p.Render(tpl)
	if err != nil {
		t.Fatal(err)
	}

	if strings.TrimSpace(rendered) != strings.TrimSpace(expectedRenderedTemplate) {
		t.Fatalf("Rendered template does not match expectation, found\n%s", rendered)
	}
}
