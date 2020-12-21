package main

import (
	"github.com/lvxin0315/gapi/core/generate"
	"github.com/lvxin0315/gapi/models"
)

/**
 * @Author lvxin0315@163.com
 * @Description model生成对应service
 * @Date 11:29 上午 2020/12/11
 * @Param
 * @return
 **/

const ServiceDir = "services"

func main() {
	genService := generate.GenService{
		ServiceDir: ServiceDir,
	}
	genService.AutoService(
		models.EbAgreementModel{},
		models.EbArticleModel{},
		models.EbArticleCategoryModel{},
		models.EbArticleContentModel{},
		models.EbAuxiliaryModel{},
		models.EbCacheModel{},
		models.EbCategoryModel{},
		models.EbDeliveryServiceModel{},
		models.EbDiyModel{},
		models.EbExpressModel{},
		models.EbLiveAnchorModel{},
		models.EbLiveGoodsModel{},
		models.EbLiveRoomModel{},
		models.EbLiveRoomGoodsModel{},
		models.EbMemberCardModel{},
		models.EbMemberCardBatchModel{},
		models.EbMemberRightModel{},
		models.EbMemberShipModel{},
		models.EbOtherOrderModel{},
		models.EbOtherOrderStatusModel{},
		models.EbQrcodeModel{},
		models.EbShippingTemplatesModel{},
		models.EbShippingTemplatesFreeModel{},
		models.EbShippingTemplatesRegionModel{},
		models.EbSmsRecordModel{},
		models.EbStoreBargainModel{},
		models.EbStoreBargainUserModel{},
		models.EbStoreBargainUserHelpModel{},
		models.EbStoreCartModel{},
		models.EbStoreCategoryModel{},
		models.EbStoreCombinationModel{},
		models.EbStoreCouponIssueModel{},
		models.EbStoreCouponIssueUserModel{},
		models.EbStoreCouponProductModel{},
		models.EbStoreCouponUserModel{},
		models.EbStoreOrderModel{},
		models.EbStoreOrderCartInfoModel{},
		models.EbStoreOrderEconomizeModel{},
		models.EbStoreOrderInvoiceModel{},
		models.EbStoreOrderStatusModel{},
		models.EbStorePinkModel{},
		models.EbStoreProductModel{},
		models.EbStoreProductAttrModel{},
		models.EbStoreProductAttrResultModel{},
		models.EbStoreProductAttrValueModel{},
		models.EbStoreProductCateModel{},
		models.EbStoreProductCouponModel{},
		models.EbStoreProductDescriptionModel{},
		models.EbStoreProductLogModel{},
		models.EbStoreProductRelationModel{},
		models.EbStoreProductReplyModel{},
		models.EbStoreProductRuleModel{},
		models.EbStoreSeckillModel{},
		models.EbStoreSeckillTimeModel{},
		models.EbStoreServiceModel{},
		models.EbStoreServiceFeedbackModel{},
		models.EbStoreServiceLogModel{},
		models.EbStoreServiceRecordModel{},
		models.EbStoreServiceSpeechcraftModel{},
		models.EbStoreVisitModel{},
		models.EbSystemAdminModel{},
		models.EbSystemAttachmentModel{},
		models.EbSystemAttachmentCategoryModel{},
		models.EbSystemCityModel{},
		models.EbSystemConfigModel{},
		models.EbSystemConfigTabModel{},
		models.EbSystemFileModel{},
		models.EbSystemGroupModel{},
		models.EbSystemGroupDataModel{},
		models.EbSystemLogModel{},
		models.EbSystemMenusModel{},
		models.EbSystemNoticeModel{},
		models.EbSystemNoticeAdminModel{},
		models.EbSystemRoleModel{},
		models.EbSystemStoreModel{},
		models.EbSystemStoreStaffModel{},
		models.EbSystemUserLevelModel{},
		models.EbTemplateMessageModel{},
		models.EbUserModel{},
		models.EbUserAddressModel{},
		models.EbUserBillModel{},
		models.EbUserBrokerageFrozenModel{},
		models.EbUserEnterModel{},
		models.EbUserExtractModel{},
		models.EbUserFriendsModel{},
		models.EbUserGroupModel{},
		models.EbUserInvoiceModel{},
		models.EbUserLabelModel{},
		models.EbUserLabelRelationModel{},
		models.EbUserLevelModel{},
		models.EbUserNoticeModel{},
		models.EbUserNoticeSeeModel{},
		models.EbUserRechargeModel{},
		models.EbUserSearchModel{},
		models.EbUserSignModel{},
		models.EbUserVisitModel{},
		models.EbWechatKeyModel{},
		models.EbWechatMediaModel{},
		models.EbWechatMessageModel{},
		models.EbWechatNewsCategoryModel{},
		models.EbWechatReplyModel{},
		models.EbWechatUserModel{},
	)
}