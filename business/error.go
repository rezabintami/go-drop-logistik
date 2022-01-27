package business

import "errors"

var (
	ErrInternalServer = errors.New("something gone wrong, contact administrator")

	ErrNotFound = errors.New("data not found")

	ErrIDNotFound = errors.New("id not found")

	ErrArticleIDResource = errors.New("(ArticleID) not found or empty")

	ErrArticleTitleResource = errors.New("(ArticleTitle) not found or empty")

	ErrCategoryNotFound = errors.New("category not found")

	ErrDuplicateData = errors.New("duplicate data")

	ErrEmailPasswordNotFound = errors.New("(Email) or (Password) empty")
)
