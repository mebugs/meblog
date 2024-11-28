package blogService

import (
	"fmt"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/common/mysql/blogDB"
	"time"
)

// AddCategory 创建文章分类
func AddCategory(traceID string, req *blogModel.CategoryAddReq) *baseModel.ResBody {
	// 先处理资源
	if len(req.Source) < 1 {
		return baseModel.Fail(constant.CategoryAddNoScNG)
	}
	dbSourceReq, errNum, err := thirdAddSource(traceID, &blogModel.SourceAddReq{
		FileType: 1, // ICON
		SourceDoReq: blogModel.SourceDoReq{
			Name:         "Category:" + req.Title,
			FileStrArray: req.Source,
		},
	})
	if err != nil {
		switch errNum {
		case 1:
			// 解析数据库错误
			return checkSourceDBErr(err)
		case 2:
			return baseModel.Fail(constant.SourceFileUpNg)
		}
	}
	// 创建对象初始化
	dbReq := req.ToDbReq()
	dbReq.SourceId = dbSourceReq.Id
	err = blogDB.CategoryTable.InsertOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "AddCategory Fail . Err Is : %v", err)
		// 解析数据库错误
		return checkCategoryDBErr(err)
	}
	// 分类无需独立的资源关联
	return baseModel.Success(constant.CategoryAddSS, true)
}

// PageCategory 查询文章分类分页
func PageCategory(traceID string, req *blogModel.CategoryPageReq) *baseModel.ResBody {
	// 查询分页
	total, list, err := blogDB.CategoryTable.Page(categoryPageQuery(req))
	if err != nil {
		log.ErrorTF(traceID, "PageCategory Fail . Err Is : %v", err)
		return baseModel.Fail(constant.CategoryGetNG)
	}
	// 循环List获取资源地址
	sourceIds := make([]uint64, len(list))
	for i, r := range list {
		sourceIds[i] = r.SourceId
	}
	sources, err := blogDB.SourceTable.GetByIds(sourceIds)
	if err != nil {
		log.ErrorTF(traceID, "PageCategory GetSource Fail . Err Is : %v", err)
		return baseModel.Fail(constant.CategoryGetNG)
	}

	// 读取资源地址
	sourceMap := make(map[uint64]string, len(sources))
	for _, r := range sources {
		sourceMap[r.Id] = fmt.Sprintf(constant.SourceFileUrl, r.FilePath, fmt.Sprintf("%d", r.Id), r.BackEnd, time.Now().Unix())
	}
	return baseModel.SuccessUnPop(baseModel.SetPageRes(blogModel.ToCategoryPageRes(list, sourceMap), total))
}

// GetCategory 文章分类详情
func GetCategory(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	res, err := blogDB.CategoryTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetCategory Fail . Err Is : %v", err)
		return baseModel.Fail(constant.CategoryGetNG)
	}
	source, err := blogDB.SourceTable.GetOneById(res.SourceId)
	if err != nil {
		log.ErrorTF(traceID, "GetCategory GetSource Fail . Err Is : %v", err)
		return baseModel.Fail(constant.CategoryGetNG)
	}
	return baseModel.SuccessUnPop(blogModel.ToCategoryGetRes(&res, fmt.Sprintf(constant.SourceFileUrl, source.FilePath, fmt.Sprintf("%d", source.Id), source.BackEnd, time.Now().Unix())))
}

// EditCategory 编辑文章分类
func EditCategory(traceID string, req *blogModel.CategoryEditReq) *baseModel.ResBody {
	dbReq, err := blogDB.CategoryTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetCategory Fail . Err Is : %v", err)
		return baseModel.Fail(constant.CategoryGetNG)
	}
	// 如果存在图片资源，更新图片
	if len(req.Source) > 1 {
		return baseModel.Fail(constant.CategoryAddNoScNG)
	}
	errNum, err := thirdEditSource(traceID, &blogModel.SourceEditReq{
		Id: dbReq.SourceId,
		SourceDoReq: blogModel.SourceDoReq{
			Name:         "Category:" + dbReq.Title,
			FileStrArray: req.Source,
		},
	})
	if err != nil {
		switch errNum {
		case 1:
			return baseModel.Fail(constant.SourceGetNG)
		case 2:
			return baseModel.Fail(constant.SourceFileUpNg)
		case 3:
			return checkSourceDBErr(err)
		}
	}
	// 对象更新
	req.ToDbReq(&dbReq)
	err = blogDB.CategoryTable.UpdateOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "EditCategory %d Fail . Err Is : %v", dbReq.Id, err)
		// 解析数据库错误
		return checkCategoryDBErr(err)
	}
	return baseModel.Success(constant.CategoryEditSS, true)
}

// DelCategory 文章分类移除
func DelCategory(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	dbReq, err := blogDB.CategoryTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetCategory Fail . Err Is : %v", err)
		return baseModel.Fail(constant.CategoryGetNG)
	}
	// 如果存在数据禁止删除
	count, err := blogDB.PostTable.CountByObject(blogDB.Post{CategoryId: req.Id})
	if err != nil {
		log.ErrorTF(traceID, "CountCategory %d Fail . Err Is : %v", req.Id, err)
		return baseModel.Fail(constant.CategoryGetNG)
	}
	if count > 0 {
		log.ErrorTF(traceID, "DelCategory %d Fail . Had Post", req.Id)
		return baseModel.Fail(constant.CategoryDelNoNG)
	}
	// 物理删除
	err = blogDB.CategoryTable.DeleteOne(dbReq.Id)
	if err != nil {
		log.ErrorTF(traceID, "DelCategory %d Fail . Err Is : %v", dbReq.Id, err)
		// 硬删除直接报错
		return baseModel.Fail(constant.CategoryDelNG)
	}
	return baseModel.Success(constant.CategoryDelSS, true)
}

// MergeCategory 分组迁移
func MergeCategory(traceID string, req *blogModel.CategoryToReq) *baseModel.ResBody {
	_, err := blogDB.CategoryTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetCategory %d Fail . Err Is : %v", req.Id, err)
		return baseModel.Fail(constant.CategoryGetNG)
	}
	_, err = blogDB.CategoryTable.GetOneById(req.ToId)
	if err != nil {
		log.ErrorTF(traceID, "GetToCategory %d Fail . Err Is : %v", req.ToId, err)
		return baseModel.Fail(constant.CategoryGetNG)
	}
	// 批量更新分组
	err = blogDB.PostTable.Executor().ToCategory(req.Id, req.ToId)
	if err != nil {
		log.ErrorTF(traceID, "MergeCategory %d To %d Fail . Err Is : %v", req.Id, req.ToId, err)
		return baseModel.Fail(constant.CategoryMergeNG)
	}
	return baseModel.Success(constant.CategoryMergeSS, true)
}

// ListCategory 分组列表
func ListCategory(traceID string) *baseModel.ResBody {
	resList, err := blogDB.CategoryTable.GetAll()
	if err != nil {
		log.ErrorTF(traceID, "ListCategory Fail . Err Is : %v", err)
		return baseModel.Fail(constant.CategoryGetNG)
	}
	resp := make([]baseModel.SelectNumRes, 0)
	for _, item := range resList {
		resp = append(resp, baseModel.SelectNumRes{
			Label: item.Title,
			Value: item.Id,
		})
	}
	return baseModel.SuccessUnPop(resp)
}
