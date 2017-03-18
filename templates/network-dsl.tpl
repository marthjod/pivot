{{range .Services.Provides}}
service("{{.ServiceName}}") {
    description    "{{.Description}}"
    attach_to      "hg-{{$.ShortName}}"
    protocol       "{{.TransportProtocol}}"
    port           {{.Port}}

    # WIP
    {{range $.Services.DependsOn.Internal -}}
    talks_to    "{{.Name}}"
    {{end -}}
}
{{end}}