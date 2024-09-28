package renderer

import (
	"strings"
	"unicode/utf8"

	"github.com/cloudhonk/faras/khel"
)

type FarasFrameConfig struct {
	Width   int
	Height  int
	Padding int
}

type FarasFrame struct {
	frame [][]rune
}

type GetPlayersManager interface {
	GetPlayers() []*khel.Juwadey
}

type FarasFrameBuilder struct {
	PlayerManager GetPlayersManager
	FarasFrame
	FarasFrameConfig
}

func NewFarasFrameBuilder(config FarasFrameConfig, playerManager GetPlayersManager) *FarasFrameBuilder {
	return &FarasFrameBuilder{
		FarasFrameConfig: config,
		PlayerManager:    playerManager,
	}
}

func (ffb *FarasFrameBuilder) InitFrame() *FarasFrameBuilder {
	ffb.frame = make([][]rune, ffb.Height)
	for i := range ffb.frame {
		ffb.frame[i] = make([]rune, ffb.Width)
	}

	for i, line := range ffb.frame {
		for j := range line {
			ffb.frame[i][j] = []rune(" ")[0]
		}
	}

	return ffb
}

func (ffb *FarasFrameBuilder) AddTable() *FarasFrameBuilder {

	var sb strings.Builder
	sb.WriteString(strings.Repeat(" ", ffb.Padding))
	sb.WriteString(strings.Repeat("-", ffb.Width-2*ffb.Padding))
	sb.WriteString(strings.Repeat(" ", ffb.Padding))

	ffb.frame[ffb.Padding] = []rune(sb.String()) // Top border

	ffb.frame[ffb.Height-ffb.Padding-1] = []rune(sb.String()) // Bottom border

	for i := ffb.Padding; i < ffb.Height-ffb.Padding; i++ {
		ffb.frame[i][ffb.Padding] = []rune("|")[0]
		ffb.frame[i][ffb.Width-ffb.Padding-1] = []rune("|")[0]
	}
	return ffb
}

func (ffb *FarasFrameBuilder) AddPlayers(players []*khel.Juwadey) *FarasFrameBuilder {
	for i, player := range players {
		nameLen := len(player.Name)
		switch {

		case i == khel.BOTTOM:

			for i, ch := range player.Name {
				if ffb.Padding+i < ffb.Width {
					ffb.frame[ffb.Height-ffb.Padding-3][ffb.Width/2-nameLen/2+i] = ch
				}
			}

			for i, taas := range player.Haat {
				unicodeCharCount := utf8.RuneCountInString(taas.String())
				for j := range unicodeCharCount {
					ffb.frame[ffb.Height-ffb.Padding-2][ffb.Width/2-3+j+i*unicodeCharCount] = []rune(taas.String())[j]
				}
			}

		case i == khel.RIGHT:

			for i, ch := range player.Name {
				if ffb.Padding+i < ffb.Width {
					ffb.frame[ffb.Height/2-1][ffb.Width-ffb.Padding-1-nameLen+i] = ch
				}
			}

			for i, taas := range player.Haat {
				unicodeCharCount := utf8.RuneCountInString(taas.String())
				for j := range unicodeCharCount {
					ffb.frame[ffb.Height/2][ffb.Width-ffb.Padding-1-6+j+i*unicodeCharCount] = []rune(taas.String())[j]
				}
			}

		case i == khel.LEFT:
			for i, ch := range player.Name {
				if ffb.Padding+i < ffb.Width {
					ffb.frame[ffb.Height/2-1][ffb.Padding+1+i] = ch
				}
			}

			for i, taas := range player.Haat {
				unicodeCharCount := utf8.RuneCountInString(taas.String())
				for j := range unicodeCharCount {
					ffb.frame[ffb.Height/2][ffb.Padding+1+j+i*unicodeCharCount] = []rune(taas.String())[j]
				}
			}

		case i == khel.TOP:
			for i, ch := range player.Name {
				if ffb.Padding+i < ffb.Width {
					ffb.frame[ffb.Padding+1][ffb.Width/2-nameLen/2+i] = ch
				}
			}

			for i, taas := range player.Haat {
				unicodeCharCount := utf8.RuneCountInString(taas.String())
				for j := range unicodeCharCount {
					ffb.frame[ffb.Padding+2][ffb.Width/2-3+j+i*unicodeCharCount] = []rune(taas.String())[j]
				}
			}

		}
	}
	return ffb
}

func (ffb *FarasFrameBuilder) AddLogo() *FarasFrameBuilder {
	// ASCII art for "FARAS"
	logo := []string{
		" *****   *****  *****   *****  ***** ",
		" *       *   *  *   *   *   *  *     ",
		" *****   *****  *****   *****  ***** ",
		" *       *   *  *  *    *   *      * ",
		" *       *   *  *   *   *   *  ***** ",
	}

	// Calculate the starting position to center the logo in the table
	logoHeight := len(logo)
	logoWidth := len(logo[0])

	startY := (ffb.Height / 2) - (logoHeight / 2) // Y-coordinate center
	startX := (ffb.Width / 2) - (logoWidth / 2)   // X-coordinate center

	// Add the logo into the frame
	for i, line := range logo {
		for j, ch := range line {
			if startY+i >= 0 && startY+i < ffb.Height && startX+j >= 0 && startX+j < ffb.Width {
				ffb.frame[startY+i][startX+j] = rune(ch)
			}
		}
	}

	return ffb
}

func (ffb *FarasFrameBuilder) Build(playerRef int) {

	ffb.
		InitFrame().
		AddTable().
		AddLogo().
		AddPlayers(khel.RotatePlayers(ffb.PlayerManager.GetPlayers(), playerRef))
}

func (ffb *FarasFrameBuilder) Flush() []byte {

	clearScreen := "\033[H\033[2J"
	var sb strings.Builder
	var frameLines []string

	for _, line := range ffb.frame {
		frameLines = append(frameLines, string(line))
	}
	frame := strings.Join(frameLines, "\n")
	sb.WriteString(clearScreen)
	sb.WriteString(frame)
	return []byte(sb.String())
}

// func main() {
// 	config := FarasFrameConfig{
// 		Width:   80,
// 		Height:  25,
// 		Padding: 2,
// 	}
// 	ffb := NewFarasFrameBuilder(config)
// 	ffb.InitFrame().AddTable().AddLogo().Build()
// }
