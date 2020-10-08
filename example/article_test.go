package example

import (
	"testing"

	lorem "github.com/drhodes/golorem"
)

func TestArticleManager(t *testing.T) {
	var err error

	am, err := NewArticleManager("am.dat", "", "amsecret")
	if err != nil {
		t.Errorf("NewArticleManager return err: %s", err)
	}

	for i := 1; i <= 10; i++ {
		articleName := lorem.Sentence(8, 12)
		articleContent := lorem.Paragraph(3, 6)
		tag1 := lorem.Word(3, 8)
		tag2 := lorem.Word(4, 9)
		tag3 := lorem.Word(5, 10)

		article, err := am.NewArticle(articleName, articleContent, articleContent, []string{tag1, tag2, tag3})
		if err != nil {
			t.Errorf("NewArticle return err: %s", err)
		}

		if article.ID == "" {
			t.Errorf("NewArticle Id is nil")
		}

		retRecord, err := am.GetByID(article.ID)
		if err != nil {
			t.Errorf("GetByID return err: %s", err)
		}

		if article.Name != retRecord.Name || article.Content != retRecord.Content {
			t.Errorf("GetByID returned values are not equal")
		}

	}

	return
}
