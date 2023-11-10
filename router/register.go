package router

import (
	"github.com/gin-gonic/gin"
	"nursing_work/handler/carer"
	"nursing_work/handler/chat"
	"nursing_work/handler/collect"
	"nursing_work/handler/hire"
	"nursing_work/handler/like"
	"nursing_work/handler/mid"
	"nursing_work/handler/post"
	"nursing_work/handler/user"
	"nursing_work/handler/view"
)

func Register(r *gin.Engine) {
	r.POST("/api/v1/login", user.AppWeChatLogin)
	//r.POST("/api/v1/email-login", user.t)
	v1 := r.Group("/api/v1", mid.TokenMiddleWare)
	p := v1.Group("/post")
	{
		p.GET("/recommendation", post.Recmt) //推荐的帖子
		p.GET("/comment", post.Cmt)          //评论
		p.GET("/comment/reply", post.Rep)    //评论的回复

		p.POST("", post.PtCt)                    //发布帖子
		p.DELETE("", post.PtDt)                  //删除帖子
		p.POST("/comment", post.CmtCt)           //发布评论
		p.DELETE("/comment", post.CmtDt)         //删除评论
		p.POST("/comment/reply", post.ReplyCt)   //发布回复
		p.DELETE("/comment/reply", post.ReplyDt) //删除回复
	}

	l := v1.Group("/like")
	{
		l.POST("", like.Create)   //点赞
		l.DELETE("", like.Delete) //取消点赞
	}

	c := v1.Group("/collect")
	{
		c.POST("", collect.Create)   //收藏
		c.DELETE("", collect.Delete) //取消收藏
	}

	u := v1.Group("/user")
	{
		u.GET("/:id", user.Base)                  //获取用户基本信息
		u.POST("/subscribe", user.Subscribe)      //关注
		u.DELETE("/subscribe", user.DisSubscribe) //取关
		u.GET("/:id/subscribe", user.Following)   //关注的人
		u.GET("/:id/fans", user.Fans)             //粉丝
		u.GET("/:id/post", user.Post)             //已发布的帖子
		u.GET("/:id/collection", user.Collection) //收藏的帖子
		u.GET("/:id/like", user.Like)             //点赞的帖子
		u.PUT("/avatar", user.Avatar)             //更新头像
		u.PUT("/info", user.InfoUpt)              //更新基本信息
		u.PUT("/identity", user.Identity)         //更新身份
	}
	//----------------------------------------------------------------
	h := v1.Group("/carer")
	{
		h.GET("/recommendation", carer.Recmt) //获取推荐护工
		h.GET("/search", carer.Search)        //搜索护工
		h.GET("", carer.Type)                 //根据类型返回 1
		h.GET("/view", carer.View)            // 获得护工评价 2
		//h.GET("/is_hire", carer.IsHire)       //验证是否雇佣
		//h.POST("/view", carer.ViewCt)         //发表评价
		//h.DELETE("/view", carer.ViewDt)       //删除评价
	}

	v := v1.Group("/view")
	{
		v.POST("", view.Create)   //评价指定护工 3
		v.DELETE("", view.Delete) //删除评价  4

		v.GET("/:id/detail", view.View) //我的评价 5
	}

	{
		v1.POST("/hire", hire.Create)   //建立雇佣 6
		v1.DELETE("/hire", hire.Delete) //取消雇佣 7
		v1.GET("/hire", hire.Verify)    //检测是否雇佣 8
	}

	// 聊天
	ch := v1.Group("/chat")
	{
		ch.GET("", chat.Chat)
		ch.GET("/history-one", chat.GetHistory)
		ch.GET("/history-list", chat.GetHistoryUser)
	}
}
