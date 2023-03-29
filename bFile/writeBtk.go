package File

import "os"

type BTFile struct {
	file *os.File
}

func NewBtkFile(name string) *BTFile {
	file, er := os.Create(name)
	if er != nil {
		return nil
	}
	b := new(BTFile)
	b.file = file
	return b
}

func (b *BTFile) Write(line string) {
	b.file.WriteString(line + "\r\n")
}

func (b *BTFile) WrteNewLine() {
	b.Write("\r\n")
}

func (b *BTFile) WriteTable(line string) {
	b.Write("    " + line)
}

func (b *BTFile) WriteTwoTable(line string) {
	b.Write("        " + line)
}

func (b *BTFile) Close() {
	b.file.Close()
}
