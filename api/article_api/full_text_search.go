package article_api

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (ArticleApi) FullTextSearchView(c *gin.Context) {
	var cr models.PageInfo
	_ = c.ShouldBindQuery(&cr)

	boolQuery := elastic.NewBoolQuery()

	if cr.Key != "" {
		boolQuery.Must(elastic.NewMultiMatchQuery(cr.Key, "title", "body"))
	}
	if cr.Page == 0 {
		cr.Page = 1
	}
	if cr.Limit == 0 {
		cr.Limit = 10
	}
	from := (cr.Page - 1) * cr.Limit

	result, err := global.ESClient.
		Search(models.FullTextModel{}.Index()).
		Query(boolQuery).
		Highlight(elastic.NewHighlight().Field("body")).
		From(from).
		Size(cr.Limit).
		Do(context.Background())
	if err != nil {
		return
	}
	count := result.Hits.TotalHits.Value //搜索到结果总条数
	fullTextList := make([]models.FullTextModel, 0)
	for _, hit := range result.Hits.Hits {
		var model models.FullTextModel
		json.Unmarshal(hit.Source, &model)

		body, ok := hit.Highlight["body"]
		if ok {
			model.Body = body[0]
		}

		fullTextList = append(fullTextList, model)
	}

	res.OkWithList(fullTextList, count, c)
}
