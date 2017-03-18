package model

import (
	"reflect"
	"strings"
	"testing"
	"text/template"
)

var expectedRuntime = PivioRuntime{
	CPU:         "XL",
	RAM:         "S",
	Disk:        "S",
	HostType:    "Docker",
	NetworkZone: "DMZ",
}

var expectedService = PivioServices{
	Provides: []PivioServiceProvides{
		{
			Description:       "Web Display of the Announcement",
			ServiceName:       "web-announcement-service",
			Protocol:          "https",
			Port:              80,
			TransportProtocol: "tcp",
		},
		{
			Description:       "REST API for updating CfP data",
			ServiceName:       "rest-announcement-service",
			Protocol:          "https",
			Port:              9449,
			TransportProtocol: "tcp",
		},
	},
	DependsOn: PivioServiceDependsOn{
		Internal: []Service{{Name: "user-service"}, {Name: "email-announcement-service"}},
	},
}

var expectedRenderedTemplate = `
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

func TestPivioFromFile(t *testing.T) {
	p, err := PivioFromFile("../pivio.yaml")
	if err != nil {
		t.Fatal(err)
	}

	if len(p.Services.Provides) != 2 {
		t.Fatalf("Expected 2 provided services, found %d", len(p.Services.Provides))
	}

	// Does not catch differences in Environments field
	if !reflect.DeepEqual(p.Runtime, expectedRuntime) {
		t.Fatal("Runtimes do not match.")
	}

	if !reflect.DeepEqual(p.Services.Provides, expectedService.Provides) {
		t.Fatal("Provided services do not match.")
	}

	if !reflect.DeepEqual(p.Services.DependsOn.External, expectedService.DependsOn.External) {
		t.Fatal("External dependencies do not match.")
	}

	if !reflect.DeepEqual(p.Services.DependsOn.Internal, expectedService.DependsOn.Internal) {
		t.Fatal("Internal dependencies do not match.")
	}
}

func TestRender(t *testing.T) {
	p, err := PivioFromFile("../pivio.yaml")
	if err != nil {
		t.Fatal(err)
	}

	tpl, err := template.ParseFiles("../templates/network-dsl.tpl")
	if err != nil {
		t.Fatal(err)
	}

	rendered, err := p.Render(tpl)
	if err != nil {
		t.Fatal(err)
	}

	if strings.TrimSpace(rendered) != strings.TrimSpace(expectedRenderedTemplate) {
		t.Fatal("Rendered template does not match expectation, found", rendered)
	}
}
