package badgemaker

import (
	"embed"
	"infoplus/server/utils"
	"log"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

//go:embed font
var embedStatic embed.FS

func CreateBadgeTheme(c *gin.Context) *BadgeTheme {
	badgeStyleMap := map[string]string{
		"":              "flat.html",
		"plastic":       "plastic.html",
		"flat":          "flat.html",
		"for_the_badge": "for_the_badge.html"}

	style := utils.GetQueryOrDefault(c, "style", "flat")

	badgeStyle, ok := badgeStyleMap[style]
	if !ok {
		badgeStyle = "flat.html"
	}

	colorRegex := regexp.MustCompile(`^([0-9a-f]{6}|[0-9a-f]{3})$`)

	labelColor := utils.GetQueryOrDefault(c, "labelcolor", DefaultGrey)

	if !colorRegex.MatchString(labelColor) {
		labelColor = DefaultGrey
	}
	messageColor := utils.GetQueryOrDefault(c, "messagecolor", DefaultGreen)

	if !colorRegex.MatchString(messageColor) {
		messageColor = DefaultGreen
	}
	return &BadgeTheme{
		Style:        badgeStyle,
		LabelColor:   labelColor,
		MessageColor: messageColor}
}

type BadgeTheme struct {
	Style        string
	LabelColor   string
	MessageColor string
}

func (b *BadgeTheme) CreateBadgeWithTheme(label string, message string) *Badge {
	if b.Style == "for_the_badge.html" {
		label = strings.ToUpper(label)
		message = strings.ToUpper(message)
	}

	return CreateBadge(label, message)
}

type Badge struct {
	Label        string
	Message      string
	TotalWidth   float64
	LabelWidth   float64
	MessageWidth float64
	LabelX       float64
	MessageX     float64
}

func CreateBadge(label string, message string) *Badge {
	fontData, err := embedStatic.ReadFile("font/verdana.ttf")
	if err != nil {
		log.Println(err)
	}
	fontTTF, err := truetype.Parse(fontData)
	if err != nil {
		log.Println(err)
	}
	d := font.Drawer{
		Face: truetype.NewFace(fontTTF, &truetype.Options{
			Size:    float64(11),
			DPI:     72,
			Hinting: font.HintingFull,
		})}

	labelWidth := float64(d.MeasureString(label)>>6) + 13
	messageWidth := float64(d.MeasureString(message)>>6) + 13
	totalWidth := labelWidth + messageWidth
	labelX := labelWidth/2.0 + 1
	messageX := labelWidth + messageWidth/2.0 - 1

	return &Badge{
		Label:        label,
		Message:      message,
		TotalWidth:   totalWidth,
		LabelWidth:   labelWidth,
		MessageWidth: messageWidth,
		LabelX:       labelX,
		MessageX:     messageX}
}
