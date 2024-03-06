package cron_ser

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/redis_ser"
)

func SyncArticleData() {
	//查询es所有数据,为后面数据更新做准备
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Size(10000).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		return
	}
	//拿到redis中的缓存数据
	diggInfo := redis_ser.NewDigg().GetInfo()
	lookInfo := redis_ser.NewArticleLook().GetInfo()
	commentInfo := redis_ser.NewCommentCount().GetInfo()

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Log.Error(err)
			continue
		}

		digg := diggInfo[hit.Id]
		look := lookInfo[hit.Id]
		comment := commentInfo[hit.Id]

		//新的数据=之前数据+缓存数据
		newDigg := article.DiggCount + digg
		newLook := article.LookCount + look
		newComment := article.CommentCount + comment

		if digg == 0 && look == 0 && comment == 0 {
			global.Log.Infof("%s无变化", article.Title)
			continue
		}

		//更新
		_, err := global.ESClient.
			Update().
			Index(models.ArticleModel{}.Index()).
			Id(hit.Id).
			Doc(map[string]int{
				"digg_count":    newDigg,
				"look_count":    newLook,
				"comment_count": newComment,
			}).
			Do(context.Background())
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		global.Log.Infof("%s, 点赞数据同步成功， 点赞数: %d 浏览数: %d 评论数: %d", article.Title, newDigg, newLook, newComment)
	}

	//清楚数据
	redis_ser.NewDigg().Clear()
	redis_ser.NewArticleLook().Clear()
	redis_ser.NewCommentCount().Clear()
}
