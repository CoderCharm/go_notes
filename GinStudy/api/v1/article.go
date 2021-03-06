/*
* @Time    : 2020-11-17 11:26
* @Author  : CoderCharm
* @File    : article.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

**/

package v1

import (
	"gin_study/model/request"
	"gin_study/model/response"
	"gin_study/service"
	"gin_study/utils"
	"github.com/gin-gonic/gin"
)

// @Tags ArticleAPI
// @Summary 获取推荐文章
// @accept application/json
// @Produce application/json
// @Success 200 string {string}"{"code":200,"msg":"success","data":{}}"
// @Router /mini/api/article/get/recommend [get]
func GetRecommendArticle(c *gin.Context) {
	RecommendArticleList := service.FetchRecommendArticleList()

	//global.GIN_REDIS.Set()

	response.OkWithData(RecommendArticleList, c)
}

// @Tags ArticleAPI
// @Summary 获取文章列表
// @accept application/json
// @Produce application/json
// @Param cateInfo body request.ArticleCategory true "文章分类"
// @Success 200 string {string}"{"code":200,"msg":"success","data":{}}"
// @Router /mini/api/article/get/list [get]
func GetArticleList(c *gin.Context) {
	// 初始化请求参数
	cateInfo := request.ArticleCategory{CateId: 0,
		PageInfo: request.PageInfo{Page: 1, PageSize: 10},
	}

	_ = c.ShouldBindQuery(&cateInfo)

	if err := utils.Verify(cateInfo.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	articleList, total := service.FetchArticleIndexList(cateInfo)

	response.OkWithDetailed(response.PageResult{
		Data:     articleList,
		Total:    total,
		Page:     cateInfo.Page,
		PageSize: cateInfo.PageSize,
	}, "success", c)

}

// @Tags ArticleAPI
// @Summary 获取文章分类
// @accept application/json
// @Produce application/json
// @Success 200 string {string}"{"code":200,"msg":"success","data":{}}"
// @Router /mini/api/article/get/category [get]
func GetCategoryList(c *gin.Context) {
	CategoryList := service.FetchCategoryList()
	response.OkWithData(CategoryList, c)
}

// @Tags ArticleAPI
// @Summary 获取文章详情
// @accept application/json
// @Produce application/json
// @Param data body request.ArticleDetail true "文章href链接"
// @Success 200 string {string}"{"code":200,"msg":"success","data":{}}"
// @Router /mini/api/article/get/detail [get]
func GetArticleDetail(c *gin.Context) {
	href := request.ArticleDetail{}
	_ = c.ShouldBindQuery(&href)
	if err := utils.Verify(href, utils.ArticleDetailVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	DetailArticle := service.FetchArticleDetail(href)
	response.OkWithData(DetailArticle, c)
}
