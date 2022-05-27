package v1

import (
	"gin-blog/pkg/app"
	"gin-blog/pkg/e"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/util"
	"gin-blog/service/tag_service"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summary Update article tag
// @Produce  json
// @Param id path int true "ID"
// @Param name body string true "Name"
// @Param state body int false "State"
// @Param modified_by body string true "ModifiedBy"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags/{id} [put]
func GetTags(c *gin.Context) {
    appG := app.Gin{C: c}
	name := c.Query("name")

    var state int = -1
    if arg := c.Query("state"); arg != "" {
        state = com.StrTo(arg).MustInt()
    }

    tagService := tag_service.Tag{
        Name: name,
        State: state,
        PageNum: util.GetPage(c),
        PageSize: setting.AppSetting.PageSize,
    }

    tags, err := tagService.GetAll()
    if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_TAGS_FAIL, nil)
		return
	}
    count,err := tagService.Count()
    if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_TAG_FAIL, nil)
		return
	}
    // code := e.SUCCESS

    // data["lists"] = models.GetTags(util.GetPage(c), setting.AppSetting.PageSize, maps)
    // data["total"] = models.GetTagTotal(maps)

    // c.JSON(http.StatusOK, gin.H{
    //     "code" : code,
    //     "msg" : e.GetMsg(code),
    //     "data" : data,
    // })
    appG.Response(http.StatusOK,e.SUCCESS, map[string]interface{}{
		"lists": tags,
		"total": count,
	})
}
type AddTagForm struct {
	Name      string `form:"name" valid:"Required;MaxSize(100)"`
	CreatedBy string `form:"created_by" valid:"Required;MaxSize(100)"`
	State     int    `form:"state" valid:"Range(0,1)"`
}

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} string  "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func AddTag(c *gin.Context) {
	var (
        appG = app.Gin{C: c}
        form  AddTagForm
    )
    httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
    tagService := tag_service.Tag{
        Name: form.Name,
        CreatedBy: form.CreatedBy,
        State: form.State,
    }
    exists, err := tagService.ExistByName()
    if err != nil {
        appG.Response(http.StatusInternalServerError,e.ERROR_ADD_TAG_FAIL,nil)
        return
    }
    if exists {
        appG.Response(http.StatusOK,e.ERROR_EXIST_TAG,nil)
        return
    }
    err = tagService.Add()
    if err!=nil {
        appG.Response(http.StatusInternalServerError,e.ERROR_ADD_TAG_FAIL,nil)
        return
    }

    appG.Response(http.StatusOK,e.SUCCESS,nil)
}

type EditTagForm struct{
    ID int  `form:"id" valid:"Required;Min(1)"`
    Name       string `form:"name" valid:"Required;MaxSize(100)"`
	ModifiedBy string `form:"modified_by" valid:"Required;MaxSize(100)"`
	State      int    `form:"state" valid:"Range(0,1)"`

}
// @Summary 修改文章标签
// @Produce  json
// @Param id path int true "ID"
// @Param name query string true "ID"
// @Param state query int false "State"
// @Param modified_by query string true "ModifiedBy"
// @Success 200 {string} string  "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags/{id} [put]
func EditTag(c *gin.Context) {
    var (
        appG = app.Gin{C: c}
        form EditTagForm
    )
    httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
    tagService := tag_service.Tag{
        ID: form.ID,
        Name: form.Name,
        ModifiedBy: form.ModifiedBy,
        State: form.State,
    }
    
    exists, err := tagService.ExistByID()
    if err!=nil {
        appG.Response(http.StatusInternalServerError,e.ERROR_EXIST_TAG_FAIL,nil)
        return
    }
    if !exists {
        appG.Response(http.StatusOK,e.ERROR_NOT_EXIST_TAG,nil)
        return
    }

    err = tagService.Edit()
    if err!=nil {
        appG.Response(http.StatusInternalServerError,e.ERROR_DELETE_TAG_FAIL,nil)
        return 
    }

    appG.Response(http.StatusOK,e.SUCCESS,nil)
    



    
}

// @Summary Delete article tag
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags/{id} [delete]
func DeleteTag(c *gin.Context) {
    appG := app.Gin{C: c}
    valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
    valid.Min(id, 1, "id").Message("ID必须大于0")

    if  valid.HasErrors() {
        app.MarkErrors(valid.Errors)
        appG.Response(http.StatusBadRequest,e.INVALID_PARAMS,nil)
        return
    }
    tagService := tag_service.Tag{
        ID: id,
    }
    exists, err := tagService.ExistByID()
    if err!=nil {
        appG.Response(http.StatusInternalServerError,e.ERROR_DELETE_TAG_FAIL,nil)
        return
    }
    if !exists {
        appG.Response(http.StatusOK,e.ERROR_NOT_EXIST_TAG,nil)
        return
    }
    err  = tagService.Delete()
    if err!=nil {
        appG.Response(http.StatusInternalServerError,e.ERROR_DELETE_TAG_FAIL,nil)
        return
    }

    appG.Response(http.StatusOK,e.SUCCESS,nil)
   
   
}
