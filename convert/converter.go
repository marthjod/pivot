package convert

type Converter interface {
	Render() (string, error)
}
