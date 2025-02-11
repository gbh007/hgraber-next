package parsing

import (
	"net/url"
	"strings"

	"github.com/google/uuid"
)

type URLMirror struct {
	ID          uuid.UUID
	Name        string
	Prefixes    []string
	Description string
}

type UrlCloner struct {
	mirrors map[string][]string
}

func NewUrlCloner(mirrors []URLMirror) UrlCloner {
	uc := UrlCloner{
		mirrors: make(map[string][]string),
	}

	for _, mirror := range mirrors {
		if len(mirror.Prefixes) > 1 { // Если префикс только 1, то зеркал быть не может.
			for i1, p1 := range mirror.Prefixes {
				values := make([]string, 0, len(mirror.Prefixes)-1)

				for i2, p2 := range mirror.Prefixes {
					if i1 == i2 {
						continue
					}

					values = append(values, p2)
				}

				uc.mirrors[p1] = values
			}
		}
	}

	return uc
}

func (uc UrlCloner) GetClones(u url.URL) ([]url.URL, error) {
	out := make([]url.URL, 0)

	s := u.String()

	for prefix, replacements := range uc.mirrors {
		if strings.HasPrefix(s, prefix) {
			for _, v := range replacements {
				mirror, err := url.Parse(strings.Replace(s, prefix, v, 1))
				if err != nil {
					return nil, err
				}

				out = append(out, *mirror)
			}
		}
	}

	return out, nil
}
