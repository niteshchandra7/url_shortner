package repository

type DatabaseRepo interface {
	// Exists(link string) bool
	GetShortenLinkFromLink(link string) (string, bool)
	CreateAndInsertShortenLinkFromLink(link string) string
	GetLinkFromShortenLink(shorten_link string) (string, bool)
}
