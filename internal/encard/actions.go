package encard

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v3"
)

func Run(ctx context.Context, cmd *cli.Command) error {
	opts, err := Setup(cmd)
	if err != nil {
		log.Fatalf("%v", err)
	}

	args := cmd.Args().Slice()

	cards, _ := LoadCards(args, opts.cfg.CardsDir)

	if len(cards) == 0 {
		return fmt.Errorf("no cards found")
	}

	session := NewSession(cards, opts)
	model := NewModel(session)

	program := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		return err
	}
	return nil
}

func Verify(ctx context.Context, cmd *cli.Command) error {
	opts, err := Setup(cmd)
	if err != nil {
		log.Fatalf("%v", err)
	}
	args := cmd.Args().Slice()
	cards, _ := LoadCards(args, opts.cfg.CardsDir)

	if len(cards) == 0 {
		return fmt.Errorf("no cards found")
	}

	fmt.Printf("loaded %d cards\n", len(cards))

	return nil
}

func ImageTest(ctx context.Context, cmd *cli.Command) error {
	RenderPNG("/home/sjsanc/work/encard/pkg/kitty/testdata/png-cat.png")
	return nil
}

func RenderPNG(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	encodedData := base64.StdEncoding.EncodeToString(data)
	chunkSize := 4096
	pos := 0
	var sb strings.Builder
	for pos < len(encodedData) {
		sb.WriteString("\x1b_G")
		if pos == 0 {
			sb.WriteString("a=T,f=100,")
		}
		endPos := pos + chunkSize
		if endPos > len(encodedData) {
			endPos = len(encodedData)
		}
		chunk := encodedData[pos:endPos]
		pos = endPos

		if pos < len(encodedData) {
			sb.WriteString("m=1")
		}
		if len(chunk) > 0 {
			sb.WriteString(";" + chunk)
		}

		sb.WriteString("\x1b\\")
	}
	fmt.Print(sb.String())
	return nil
}
