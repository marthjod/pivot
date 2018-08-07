package model

import (
	"os"
	"reflect"
	"testing"
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

func TestPivio_Read(t *testing.T) {
	f, err := os.Open("../examples/filebackend/pivio.yaml")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer f.Close()
	p, err := Read(f)
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

func TestPivio_ReadJSONMultiple(t *testing.T) {
	f, err := os.Open("../examples/apibackend/pivio.json")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer f.Close()
	pivios, err := ReadJSONMultiple(f)
	if err != nil {
		t.Fatal(err)
	}

	if len(pivios) == 0 {
		t.Fatalf("expected at least 1 service in unmarshaled struct")
	}

	p := pivios[0]

	if len(p.Services.DependsOn.Internal) != 2 {
		t.Fatalf("Expected 2 provided services, found %d", len(p.Services.DependsOn.Internal))
	}

	// Does not catch differences in Environments field
	if !reflect.DeepEqual(p.Runtime, expectedRuntime) {
		t.Fatalf("Runtimes do not match, expected %+v, got %+v", expectedRuntime, p.Runtime)
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
