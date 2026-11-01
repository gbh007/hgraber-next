package legacyadapter

import (
	"context"

	"github.com/gbh007/hgraber-next/external/parserdto"
	"github.com/gbh007/hgraber-next/internal/domain/core"
)

func ParseBookAttr(ctx context.Context, p legacyParser, attr string) ([]string, error) {
	switch attr {
	case core.AttributeCodeAuthor:
		return p.Authors(ctx)

	case core.AttributeCodeCategory:
		return p.Categories(ctx)

	case core.AttributeCodeCharacter:
		return p.Characters(ctx)

	case core.AttributeCodeGroup:
		return p.Groups(ctx)

	case core.AttributeCodeLanguage:
		return p.Languages(ctx)

	case core.AttributeCodeParody:
		return p.Parodies(ctx)

	case core.AttributeCodeTag:
		return p.Tags(ctx)

	default:
		return []string{}, parserdto.UnknownAttributeError
	}
}
