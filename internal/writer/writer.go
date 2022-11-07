package writer

import (
	"fmt"
	"os"

	"github.com/svartvalp/paoias2/internal/model"
)

type Writer struct {
}

func (w *Writer) Write(file string, res []model.ArgFunc) error {
	f, err := os.Create(file)
	if err != nil {
		return nil
	}
	defer func() { _ = f.Close() }()
	for _, r := range res {
		_, err := f.WriteString(fmt.Sprintf("%v;%v;\n", r.T, r.Y))
		if err != nil {
			return err
		}
	}

	return nil
}

func New() *Writer {
	return &Writer{}
}
