package result

/*
错误码设计
第一位表示错误级别, 1 为系统错误, 2 为普通错误
第二三位表示服务模块代码
第四五位表示具体错误代码
*/
// 定义错误码
type Errno struct {
	Code      int
	Message   string
	EnMessage string
}

func (err Errno) Error() string {
	return err.Message
}

var (
	SUCCESS = &Errno{Code: 200, Message: "success", EnMessage: "success"}

	// 系统通用错误, 前缀为 100
	ErrInternalServer       = &Errno{Code: 10001, Message: "内部服务器错误", EnMessage: "Internal server error"}
	ErrBind                 = &Errno{Code: 10002, Message: "请求参数错误", EnMessage: "Request parameter error"}
	ErrUploadFile           = &Errno{Code: 10003, Message: "上传文件失败", EnMessage: "Failed to upload file"}
	ErrPageParam            = &Errno{Code: 10004, Message: "分页参数有误", EnMessage: "Wrong paging parameters"}
	ErrContextParam         = &Errno{Code: 10005, Message: "错误的上下文参数", EnMessage: "Wrong context parameters"}
	ErrRecordNotExist       = &Errno{Code: 10006, Message: "数据记录不存在", EnMessage: "Data record does not exist"}
	ErrDataPermissionDenied = &Errno{Code: 20326, Message: "数据权限拒绝", EnMessage: "Data Permission Denial"}

	// 认证错误, 前缀是 202
	ErrRefreshTokenInvalid = &Errno{Code: 20201, Message: "refresh token 无效", EnMessage: "refresh token invalid"}
	ErrTokenInvalid        = &Errno{Code: 20202, Message: "token无效", EnMessage: "invalid token"}
	ErrSignInvalid         = &Errno{Code: 20203, Message: "签名错误", EnMessage: "Wrong signature"}
	ErrTimeInvalid         = &Errno{Code: 20204, Message: "请求时间无效", EnMessage: "Invalid request time"}
	ErrAppKeyNotExist      = &Errno{Code: 20205, Message: "appKey不存在", EnMessage: "appKey does not exist"}
	ErrUserForbidden       = &Errno{Code: 20206, Message: "用户权限不足", EnMessage: "Insufficient user rights"}
	ErrSignParam           = &Errno{Code: 20207, Message: "app_key,sn,ts为headers必选参数", EnMessage: "app_key,sn,ts are required parameters for headers"}

	// 用户错误, 前缀为 203
	ErrUserNotFound        = &Errno{Code: 20301, Message: "账号不存在", EnMessage: "Account does not exist"}
	ErrPasswordIncorrect   = &Errno{Code: 20302, Message: "密码错误", EnMessage: "Password error"}
	ErrPasswordRetryLimit  = &Errno{Code: 20303, Message: "密码错误次数过多", EnMessage: "Password error too many times"}
	ErrAuthCodeSend        = &Errno{Code: 20304, Message: "验证码发送失败", EnMessage: "Failed to send captcha"}
	ErrAuthCode            = &Errno{Code: 20305, Message: "验证码校验失败", EnMessage: "captcha verification failure"}
	ErrResetPasswordSelect = &Errno{Code: 20306, Message: "auth_code与password_old参数二选一", EnMessage: "auth_code and password_old parameters choose one"}
	ErrImageFormat         = &Errno{Code: 20307, Message: "图片格式不支持", EnMessage: "Image format not supported"}
	ErrImageUploadStorage  = &Errno{Code: 20308, Message: "图片上传云存储错误", EnMessage: "Image upload cloud storage error"}
	ErrModifyParameter     = &Errno{Code: 20309, Message: "用户信息修改头像和昵称不能同时为空", EnMessage: "User information modification avatar and nickname cannot be empty at the same time"}
	ErrMobileOSInput       = &Errno{Code: 20310, Message: "mobile_os 只支持1:android 2:ios", EnMessage: "mobile_os only supports 1:android 2:ios"}
	ErrUserAlreadyExist    = &Errno{Code: 20311, Message: "用户已经存在", EnMessage: "User already exists"}
	ErrNicknameInput       = &Errno{Code: 20312, Message: "nickname格式有误", EnMessage: "nickname format is wrong"}
	ErrUserHouseInit       = &Errno{Code: 20313, Message: "用户房子房间初始化失败", EnMessage: "User house room initialization failed"}
	ErrUserShareYourself   = &Errno{Code: 20314, Message: "不能分享给自己", EnMessage: "Cannot share to myself"}
	ErrIotUserRegistered   = &Errno{Code: 20315, Message: "向iot注册用户失败", EnMessage: "Failed to register user with iot"}
	ErrUserClearToken      = &Errno{Code: 20316, Message: "清理历史token失败", EnMessage: "Failed to clear historical token"}
	ErrOwnerHouseNotExist  = &Errno{Code: 20317, Message: "房子owner查询失败", EnMessage: "House owner query failed"}
	ErrAuthCodeRetryOften  = &Errno{Code: 20318, Message: "验证码校验频繁", EnMessage: "Captcha check frequently"}
	ErrLoginRetryOften     = &Errno{Code: 20319, Message: "用户登录频繁", EnMessage: "User login frequently"}
	ErrNotAdminUser        = &Errno{Code: 20320, Message: "非管理员用户", EnMessage: "Non-admin user"}
	ErrSQL                 = &Errno{Code: 20321, Message: "查询失败", EnMessage: "Query failure"}
	ErrCreate              = &Errno{Code: 20322, Message: "创建失败", EnMessage: "Failed to create"}
	ErrNoSubUser           = &Errno{Code: 20323, Message: "不能操作非子用户", EnMessage: "Cannot operate on non-sub-users"}
	ErrUserDisabled        = &Errno{Code: 20324, Message: "用户已被禁用", EnMessage: "User has been disabled"}
	ErrRoleNotFound        = &Errno{Code: 20325, Message: "角色不存在", EnMessage: "Role does not exist"}
	ErrPermissionDenied    = &Errno{Code: 20326, Message: "权限校验失败", EnMessage: "Permission check failed"}
	ErrHasDevice           = &Errno{Code: 20327, Message: "请先移除所有设备", EnMessage: "Please remove all devices first"}
	ErrHasSharedHouse      = &Errno{Code: 20328, Message: "请先剔除家庭里其他的成员", EnMessage: "Please remove other members of the family first"}
	ErrHasBeSharedHouse    = &Errno{Code: 20329, Message: "请先退出别人的家庭", EnMessage: "Please exit someone else's family first"}
	ErrHasDeviceGroup      = &Errno{Code: 20330, Message: "请先删除所有群组", EnMessage: "Please remove all groups first"}
	ErrListRoleAuthority   = &Errno{Code: 20331, Message: "获取角色的权限列表失败", EnMessage: "Failed to get the permission list of the role"}
	ErrListAuthority       = &Errno{Code: 20332, Message: "获取权限列表失败", EnMessage: "Failed to get permission list"}
	ErrAuthority           = &Errno{Code: 20333, Message: "权限不存在", EnMessage: "Permission does not exist"}
	ErrCreateRole          = &Errno{Code: 20334, Message: "创建角色失败", EnMessage: "Failed to create role"}
	ErrListRole            = &Errno{Code: 20335, Message: "获取角色列表失败", EnMessage: "Failed to get the role list"}
	ErrUpdateRole          = &Errno{Code: 20336, Message: "更新角色失败", EnMessage: "Failed to update role"}
	ErrRoleInUse           = &Errno{Code: 20337, Message: "角色已被使用", EnMessage: "Role has been used"}
	ErrDeleteRole          = &Errno{Code: 20338, Message: "删除角色失败", EnMessage: "Failed to delete role"}
	ErrHasProductKey       = &Errno{Code: 20339, Message: "产品不存在", EnMessage: "Product does not exist"}
	ErrHasAppId            = &Errno{Code: 20340, Message: "App不存在", EnMessage: "App does not exist"}
	ErrCompanyNameInput    = &Errno{Code: 20341, Message: "company_name格式有误", EnMessage: "company_name format is wrong"}
	ErrPasswordInput       = &Errno{Code: 20342, Message: "password格式有误", EnMessage: "password format is wrong"}
	ErrCaptcha             = &Errno{Code: 20343, Message: "行为验证失败", EnMessage: "behavior verification failed"}
	ErrUserLocked          = &Errno{Code: 20344, Message: "您的密码输入错误次数过多,请5分钟后再尝试", EnMessage: "your password has been entered incorrectly too many times, please try again in 5 minutes"}
	ErrAuthCodeFrequency   = &Errno{Code: 20345, Message: "验证码获取太频繁", EnMessage: "get auth code too frequently"}
	ErrDomain              = &Errno{Code: 20346, Message: "仅支持在中国域注册或找回密码", EnMessage: "Only support in China domain registration or password retrieval"}

	//产品/设备, 前缀为 204
	ErrDeviceParingToken           = &Errno{Code: 20401, Message: "配网token获取失败", EnMessage: "Failed to get the token of the distribution network"}
	ErrDeviceOrder                 = &Errno{Code: 20402, Message: "设备指令发送失败", EnMessage: "Failed to send device command"}
	InvalidParingToken             = &Errno{Code: 20403, Message: "无效的配网token", EnMessage: "Invalid paring token"}
	ErrDeviceNotExist              = &Errno{Code: 20404, Message: "设备不存在", EnMessage: "Device does not exist"}
	ErrUserDeviceAuthFailure       = &Errno{Code: 20405, Message: "设备权限异常", EnMessage: "Device permission abnormal"}
	ErrProductNotExist             = &Errno{Code: 20406, Message: "产品不存在", EnMessage: "Product does not exist"}
	ErrProductIconGet              = &Errno{Code: 20407, Message: "获取产品icon信息失败", EnMessage: "Failed to get product icon information"}
	ErrDeviceHasBeenShared         = &Errno{Code: 20408, Message: "设备已经分享给用户", EnMessage: "The device has been shared to the user"}
	ErrProductAuthFailure          = &Errno{Code: 20409, Message: "产品权限异常", EnMessage: "Product permission abnormal"}
	ErrProductCategoryNotExist     = &Errno{Code: 20410, Message: "产品分类不存在", EnMessage: "Product category does not exist"}
	ErrProductTypeNotExist         = &Errno{Code: 20411, Message: "产品类型不存在", EnMessage: "Product type does not exist"}
	ErrProductNameInput            = &Errno{Code: 20412, Message: "product_name 长度需要大于2，小于255", EnMessage: "product_name length must be greater than 2 and less than 255"}
	ErrProductUpdateInputNil       = &Errno{Code: 20413, Message: "product_name 和 product_image 不能同时为空", EnMessage: "product_name and product_image cannot be empty at the same time"}
	ErrProductNameInputExist       = &Errno{Code: 20414, Message: "product_name 已存在", EnMessage: "product_name already exists"}
	ErrProductDescInputNil         = &Errno{Code: 20415, Message: "product_desc 不能为空", EnMessage: "product_desc cannot be empty"}
	ErrProtocolInput               = &Errno{Code: 20416, Message: "protocol 协议类型不正确，现支持Zigbee、WIFI、BLE、BLE+Wi-Fi四种", EnMessage: "protocol protocol type is incorrect, now support Zigbee, WIFI, BLE, BLE+Wi-Fi"}
	ErrProtocolInputNil            = &Errno{Code: 20417, Message: "protocol 不能为空", EnMessage: "protocol cannot be empty"}
	ErrParingTypeInput             = &Errno{Code: 20418, Message: "paring_type 配网类型不正确，现支持BLE、AP、EZ、ZB、ZBN、NO", EnMessage: "The paring_type is incorrect, now supports BLE, AP, EZ, ZB, ZBN, NO"}
	ErrParingTypeInputNil          = &Errno{Code: 20419, Message: "paring_type 不能为空", EnMessage: "paring_type cannot be empty"}
	ErrProductParams               = &Errno{Code: 20420, Message: "product 参数处理错误", EnMessage: "Product parameter processing error"}
	ErrUserDeviceNameRepeatBYHouse = &Errno{Code: 20421, Message: "该房间下设备昵称重复", EnMessage: "The nickname of the device in this room is duplicated"}
	ErrDeviceAlreadyBind           = &Errno{Code: 20422, Message: "该设备已被绑定，请解绑后重试。", EnMessage: "The device has been bound, please unbind it and try again."}
	ErrDeviceGroupList             = &Errno{Code: 20423, Message: "获取设备分组列表失败", EnMessage: "Failed to get device group list"}
	ErrDeviceGroupUpdate           = &Errno{Code: 20424, Message: "设备分组更新设备失败", EnMessage: "Failed to update devices in device group"}
	ErrDeviceGroupCreate           = &Errno{Code: 20425, Message: "设备分组创建失败", EnMessage: "Failed to create device group"}
	ErrDeviceGroupNotExist         = &Errno{Code: 20426, Message: "设备分组不存在", EnMessage: "Device group does not exist"}
	ErrDeviceGroupAuthFailure      = &Errno{Code: 20427, Message: "设备分组权限异常", EnMessage: "Abnormal device grouping permission"}
	ErrDeviceGroupDevice           = &Errno{Code: 20428, Message: "设备分组不存在该设备", EnMessage: "The device does not exist in the device group"}
	ErrDeviceGroupDeviceList       = &Errno{Code: 20429, Message: "获取设备分组的设备列表失败", EnMessage: "Failed to get device list of device group"}
	ErrDeviceGet                   = &Errno{Code: 20430, Message: "获取设备信息失败", EnMessage: "Failed to get device information"}
	ErrRoomGet                     = &Errno{Code: 20431, Message: "获取房间信息失败", EnMessage: "Failed to get room information"}
	ErrProductGet                  = &Errno{Code: 20432, Message: "获取产品信息失败", EnMessage: "Failed to get product information"}
	ErrDeviceCmdIsEmpty            = &Errno{Code: 20433, Message: "设备指令为空", EnMessage: "Device command is empty"}
	ErrDeviceQuitDeviceGroup       = &Errno{Code: 20434, Message: "设备退出设备分组失败", EnMessage: "Failed to exit device grouping"}
	ErrDeleteDeviceGroup           = &Errno{Code: 20435, Message: "设备分组删除失败", EnMessage: "Failed to delete device group"}
	ErrCleanDeviceGroup            = &Errno{Code: 20436, Message: "清理群组失败", EnMessage: "Failed to clear the group"}
	ErrProductGroupGet             = &Errno{Code: 20437, Message: "获取productKey所在的分组失败", EnMessage: "Failed to get the group in which the productKey is located"}
	ErrGroupProductsGet            = &Errno{Code: 20438, Message: "获取分组的productKey列表失败", EnMessage: "Failed to get the productKey list of the group"}
	ErrProductNameInputRepeat      = &Errno{Code: 20439, Message: "存在重复产品名称", EnMessage: "Duplicate product names exist"}
	ErrOrderProduct                = &Errno{Code: 20440, Message: "产品排序失败", EnMessage: "Failed to sort products"}
	ErrProductModelNotExist        = &Errno{Code: 20441, Message: "产品模版不存在", EnMessage: "Product template does not exist"}
	ErrProductStatusNotInDevelop   = &Errno{Code: 20442, Message: "只有开发中状态的产品允许修改", EnMessage: "Only products in development status are allowed to be modified"}
	ErrDeviceParingType            = &Errno{Code: 20443, Message: "配网模式输入错误", EnMessage: "Wrong input for distribution mode"}
	ErrDeviceStatusAndVersion      = &Errno{Code: 20444, Message: "获取设备在线状态跟版本失败", EnMessage: "Failed to get the online status and version of the device"}
	ErrMeshDeviceShare             = &Errno{Code: 20445, Message: "Mesh设备暂不支持分享", EnMessage: "Mesh device does not support sharing at the moment"}
	//设备定时
	ErrDeleteThingTime       = &Errno{Code: 20446, Message: "删除定时失败", EnMessage: "Failed to delete timing"}
	ErrUpdateThingTime       = &Errno{Code: 20447, Message: "修改定时失败", EnMessage: "Modify timing failed"}
	ErrAddThingTime          = &Errno{Code: 20448, Message: "添加定时失败", EnMessage: "Add Timing failed"}
	ErrThingTime             = &Errno{Code: 20449, Message: "定时不存在", EnMessage: "Timer does not exist"}
	ErrThingTimeNum          = &Errno{Code: 20450, Message: "最多只能设置10个定时", EnMessage: "Only 10 timings can be set at most"}
	ErrCountdownTimeNum      = &Errno{Code: 20451, Message: "最多只能设置1个倒计时", EnMessage: "Only 1 countdown timing can be set at most"}
	ErrPushThingTimeToDevice = &Errno{Code: 20452, Message: "定时下发到设备失败", EnMessage: "Failed to send timer to device"}

	ErrGetMeshAddress       = &Errno{Code: 20453, Message: "获取Mesh地址失败", EnMessage: "Failed to get mesh address"}
	ErrMeshAddress          = &Errno{Code: 20454, Message: "mesh_addr无效", EnMessage: "invalid mesh addr"}
	ErrProductOnOffShelf    = &Errno{Code: 20455, Message: "产品上/下架失败", EnMessage: "Product listing/removal failed"}
	ErrProductMultiLanguage = &Errno{Code: 20456, Message: "修改产品多语言失败", EnMessage: "Failed to modify product multi-language"}
	ErrCopyProduct          = &Errno{Code: 20457, Message: "复制产品失败", EnMessage: "Failed to copy product"}

	//同步产品
	ErrSyncProduct            = &Errno{Code: 20458, Message: "同步产品失败", EnMessage: "Failed to sync product"}
	ErrDestRegion             = &Errno{Code: 20459, Message: "错误的目标域", EnMessage: "Wrong target region"}
	ErrProductList            = &Errno{Code: 20460, Message: "错误的产品列表", EnMessage: "Wrong product list"}
	ErrProductSyncNotExist    = &Errno{Code: 20461, Message: "产品同步不存在", EnMessage: "ProductSync not exist"}
	ErrProductSyncAuthFailure = &Errno{Code: 20462, Message: "产品同步权限异常", EnMessage: "ProductSync permission exception"}

	ErrUpdateSecurityLevel      = &Errno{Code: 20463, Message: "修改产品安全级别失败", EnMessage: "Failed to modify product security level"}
	ErrGetProductPanelShareCode = &Errno{Code: 20464, Message: "获取面板分享码失败", EnMessage: "Failed to get panel sharing code"}
	ErrCreateDevice             = &Errno{Code: 20465, Message: "添加设备失败", EnMessage: "Failed to add device"}
	ErrCreateDeviceThan100      = &Errno{Code: 20466, Message: "产品最多只允许添加100个设备", EnMessage: "The product only allows up to 100 devices to be added"}

	//设备基础配置
	ErrCreateThingBaseConfig = &Errno{Code: 20467, Message: "创建配置失败", EnMessage: "Create config failed"}
	ErrNoThingBaseConfig     = &Errno{Code: 20468, Message: "配置不存在", EnMessage: "Config does not exist"}
	ErrUpdateThingBaseConfig = &Errno{Code: 20469, Message: "更新配置失败", EnMessage: "Update config failed"}
	ErrDeleteThingBaseConfig = &Errno{Code: 20470, Message: "删除配置失败", EnMessage: "Delete config failed"}
	ErrGetThingBaseConfig    = &Errno{Code: 20471, Message: "获取配置失败", EnMessage: "Get config failed"}
	ErrLimitThingBaseConfig  = &Errno{Code: 20472, Message: "配置数目已达最大限制", EnMessage: "Maximum config limit reached"}

	//房子房间, 前缀为 205
	ErrRoomNameExist                 = &Errno{Code: 20501, Message: "成员名称重复", EnMessage: "Duplicate room name"}
	ErrRoomNotExist                  = &Errno{Code: 20502, Message: "房间不存在", EnMessage: "Room does not exist"}
	ErrRoomAuthFailure               = &Errno{Code: 20503, Message: "房间权限异常", EnMessage: "Room permission abnormal"}
	ErrRoomDeleteFailure             = &Errno{Code: 20504, Message: "房间下有设备，无法删除", EnMessage: "There are devices under the room and cannot be deleted"}
	ErrHouseNotExist                 = &Errno{Code: 20505, Message: "房子不存在", EnMessage: "House does not exist"}
	ErrHouseAuthFailure              = &Errno{Code: 20506, Message: "房子权限异常", EnMessage: "House permission abnormal"}
	ErrHouseHasBeenShared            = &Errno{Code: 20507, Message: "房子已存在", EnMessage: "House already exists"}
	ErrHouseNameExist                = &Errno{Code: 20508, Message: "房子名称重复", EnMessage: "House name is duplicated"}
	ErrHouseDeleteFailureRoomExist   = &Errno{Code: 20509, Message: "房子下有房间，无法删除", EnMessage: "There are rooms under the house, can't delete"}
	ErrHouseDeleteFailureIsDefault   = &Errno{Code: 20510, Message: "默认房子，无法删除", EnMessage: "Default house, can't delete"}
	ErrDefaultHouseNotExist          = &Errno{Code: 20511, Message: "默认房子，查询失败", EnMessage: "Default house, query failed"}
	ErrRoomOrderUpdateInput          = &Errno{Code: 20512, Message: "更新房间排序不能为空", EnMessage: "Update room sort cannot be empty"}
	ErrHouseDeleteFailureDeviceExist = &Errno{Code: 20513, Message: "房子下有设备，无法删除", EnMessage: "There are devices under the house, can't delete"}
	ErrHouseCannotChangeOwner        = &Errno{Code: 20514, Message: "该房子不支持转让", EnMessage: "This house cannot change owner"}
	ErrHouseNewOwner                 = &Errno{Code: 20515, Message: "用户不是房子的成员", EnMessage: "The user is not a member of this house"}

	// RN 文件, 前缀为 206
	ErrGetRNFile    = &Errno{Code: 20601, Message: "获取RN文件信息失败", EnMessage: "Failed to get RN file information"}
	ErrUploadRNFile = &Errno{Code: 20602, Message: "上传RN文件失败", EnMessage: "Failed to upload RN file"}
	ErrRNNotExist   = &Errno{Code: 20603, Message: "RN信息不存在", EnMessage: "RN information does not exist"}

	// 意见反馈, 前缀为 207
	ErrGetFeedback  = &Errno{Code: 20701, Message: "获取意见反馈失败", EnMessage: "Failed to get feedback"}
	ErrPostFeedback = &Errno{Code: 20702, Message: "创建意见反馈失败", EnMessage: "Failed to create feedback"}

	// 移动 app, 前缀为 208
	ErrGetMobileApp                = &Errno{Code: 20801, Message: "获取app版本信息失败", EnMessage: "Failed to get app version information"}
	ErrAppNotExist                 = &Errno{Code: 20802, Message: "app不存在", EnMessage: "App does not exist"}
	ErrAppAuthFailure              = &Errno{Code: 20803, Message: "app权限异常", EnMessage: "App permission exception"}
	ErrGetAppIcon                  = &Errno{Code: 20804, Message: "获取app的icon失败", EnMessage: "Failed to get the app icon"}
	ErrAppVersionCode              = &Errno{Code: 20805, Message: "请输入高版本的版本号", EnMessage: "Please enter the version number of the higher version"}
	ErrCreateRelatedProduct        = &Errno{Code: 20806, Message: "添加关联产品失败", EnMessage: "Failed to add associated product"}
	ErrRelatedProductNotExist      = &Errno{Code: 20807, Message: "关联产品不存在", EnMessage: "Related product does not exist"}
	ErrRelatedProductNotDelete     = &Errno{Code: 20808, Message: "上架产品不允许删除", EnMessage: "Products listed are not allowed to be deleted"}
	ErrDeleteRelatedProduct        = &Errno{Code: 20809, Message: "关联产品删除失败", EnMessage: "Failed to delete associated product"}
	ErrAddOrDropRelatedProduct     = &Errno{Code: 20810, Message: "上/下架失败", EnMessage: "Failed to add/drop"}
	ErrWSClient                    = &Errno{Code: 20811, Message: "ws连接不存在", EnMessage: "ws does not exist"}
	ErrRelatedProductSort          = &Errno{Code: 20812, Message: "关联产品排序失败", EnMessage: "Failed to sort related products"}
	ErrGetParentUser               = &Errno{Code: 20813, Message: "查询父账号失败", EnMessage: "Failed to query parent account"}
	ErrCreateApp                   = &Errno{Code: 20814, Message: "创建app失败", EnMessage: "Failed to create app"}
	ErrAuthorizeAppSubUser         = &Errno{Code: 20815, Message: "授权app给子用户失败", EnMessage: "Failed to authorize app to sub user"}
	ErrGetNewAppId                 = &Errno{Code: 20816, Message: "获取最新app_id失败", EnMessage: "Failed to get the latest app_id"}
	ErrShowAppCategory             = &Errno{Code: 20817, Message: "切换app自定义分类是否展示失败", EnMessage: "Whether switching the app custom category fails to display"}
	ErrDeleteApp                   = &Errno{Code: 20818, Message: "删除app失败", EnMessage: "Failed to delete app"}
	ErrRelatedProductAlreadyExists = &Errno{Code: 20819, Message: "该app下此产品已存在", EnMessage: "This product already exists under this app"}
	ErrAppPolicyNotExist           = &Errno{Code: 20820, Message: "app policy不存在", EnMessage: "App policy does not exist"}
	ErrUpdateApp                   = &Errno{Code: 20821, Message: "编辑app失败", EnMessage: "Failed to update app"}

	// 产品引导页, 前缀为 209
	ErrGetProductGuide           = &Errno{Code: 20901, Message: "获取产品引导页失败", EnMessage: "Failed to get product guide page"}
	ErrCreateProductGuide        = &Errno{Code: 20902, Message: "创建产品引导页失败", EnMessage: "Failed to create product guide page"}
	ErrGetProductGuideLanguage   = &Errno{Code: 20903, Message: "获取产品引导页失败,语言查询失败", EnMessage: "Failed to get product guide page, language query failed"}
	ErrCreateProductInput        = &Errno{Code: 20904, Message: "创建产品引导页productKey或paringType必须一致", EnMessage: "Create product guide page productKey or paringType must be the same"}
	ErrProductGuideMultiLanguage = &Errno{Code: 20905, Message: "修改产品引导页多语言失败", EnMessage: "Failed to modify the multilingual product guide"}

	// 产品规格, 前缀为 210
	ErrGetProductSpec        = &Errno{Code: 21001, Message: "获取产品规格失败", EnMessage: "Failed to get product specification"}
	ErrCreateProductSpec     = &Errno{Code: 21002, Message: "创建产品规格失败", EnMessage: "Failed to create product specification"}
	ErrUpdateProductSpec     = &Errno{Code: 21003, Message: "更新产品规格失败", EnMessage: "Update product specification failed"}
	ErrGetProductProperty    = &Errno{Code: 21004, Message: "获取产品属性失败", EnMessage: "Failed to get product properties"}
	ErrCreateProductProperty = &Errno{Code: 21005, Message: "创建产品属性失败", EnMessage: "Failed to create product attributes"}
	ErrUpdateProductProperty = &Errno{Code: 21006, Message: "修改产品属性失败", EnMessage: "Modify Product Attributes Failure"}
	ErrDeleteProductProperty = &Errno{Code: 21007, Message: "删除产品属性失败", EnMessage: "Delete product attributes failure"}

	// 场景联动, 前缀为 211
	ErrSceneNotExist                = &Errno{Code: 21101, Message: "场景不存在", EnMessage: "Scenario does not exist"}
	ErrSceneCreate                  = &Errno{Code: 21102, Message: "场景创建失败", EnMessage: "Scene creation failure"}
	ErrSceneConditionNotExist       = &Errno{Code: 21103, Message: "自动场景至少有一个触发条件", EnMessage: "Automatic scenario has at least one trigger condition"}
	ErrSceneConditionExist          = &Errno{Code: 21104, Message: "手动场景不允许有触发条件", EnMessage: "Manual scenes are not allowed to have trigger conditions"}
	ErrSceneActionNotExist          = &Errno{Code: 21105, Message: "场景至少有一个执行动作", EnMessage: "Scene has at least one execution action"}
	ErrSceneConditionDevice         = &Errno{Code: 21106, Message: "条件类型为设备时Name,DeviceID,PropertyName,PropertyValue,MatchRule 不能为空", EnMessage: "Name,DeviceID,PropertyName,PropertyValue,MatchRule cannot be empty when the condition type is Device"}
	ErrSceneActionDevice            = &Errno{Code: 21107, Message: "动作类型为设备控制时Name,DeviceID,PropertyName,PropertyValue 不能为空", EnMessage: "Name,DeviceID,PropertyName,PropertyValue cannot be empty when the action type is Device Control"}
	ErrSceneConditionDeviceNotExist = &Errno{Code: 21108, Message: "条件类型为设备时检查到设备不存在", EnMessage: "If the condition type is Device, the device does not exist."}
	ErrSceneActionDeviceNotExist    = &Errno{Code: 21109, Message: "动作类型为设备控制时检查到设备不存在", EnMessage: "If the action type is device control, the device does not exist."}
	ErrSceneType                    = &Errno{Code: 21110, Message: "场景类型只支持 1.自动场景 2.手动场景", EnMessage: "Scene type only supports 1.Auto Scene 2.Manual Scene"}
	ErrSceneConditionType           = &Errno{Code: 21111, Message: "条件类型目前只支持 1.设备 2.定时 3.天气", EnMessage: "The condition type only supports 1.Device 2.Timing 3.Weather"}
	ErrSceneActionType              = &Errno{Code: 21112, Message: "动作类型目前只支持 1.设备控制 2.语音推送 3.短信推送 4.移动推送", EnMessage: "Action type currently only supports 1.Device control 2.Voice push 3.SMS push 4.Mobile push"}
	ErrSceneEnableType              = &Errno{Code: 21113, Message: "启用状态只支持 1.启用 2.不启用", EnMessage: "Enable status is only supported 1. Enable 2. Not enabled"}
	ErrActionNum                    = &Errno{Code: 21114, Message: "自动场景必须设置动作触发数量 action_num", EnMessage: "Automatic scene must set the number of action triggers action_num"}
	ErrCronFormat                   = &Errno{Code: 21115, Message: "Cron 格式错误", EnMessage: "Cron format error"}
	ErrSceneCycle                   = &Errno{Code: 21116, Message: "禁止场景循环调用", EnMessage: "Prohibit scenario cycle call"}
	ErrSceneName                    = &Errno{Code: 21117, Message: "用户下场景名称重复", EnMessage: "Duplicate scenario name under user"}
	ErrHouseIdNotEmpty              = &Errno{Code: 21118, Message: "house_id不能为空", EnMessage: "house_id cannot be empty"}
	ErrActionGet                    = &Errno{Code: 21119, Message: "获取动作信息失败", EnMessage: "Failed to get action information"}

	// 产品语义化，前缀为  212
	ErrGetProductSemantic    = &Errno{Code: 21201, Message: "获取产品功能语义化失败", EnMessage: "Failed to get product function semantics"}
	ErrCreateProductSemantic = &Errno{Code: 21202, Message: "创建产品功能语义化失败", EnMessage: "Failed to create product function semantics"}
	ErrDeleteProductSemantic = &Errno{Code: 21203, Message: "删除产品功能语义化失败", EnMessage: "Failed to delete product function semantics"}

	//hi link相关功能，前缀为  213
	ErrHiLinkDownload             = &Errno{Code: 21301, Message: "下载文件出错", EnMessage: "Error downloading file"}
	ErrHiLinkLicenseToken         = &Errno{Code: 21302, Message: "token无效", EnMessage: "token is invalid"}
	ErrHiLinkLicenseGet           = &Errno{Code: 21303, Message: "固件已经烧录到设备成功，不允许生成License", EnMessage: "Firmware has been burned to the device successfully, but license generation is not allowed"}
	ErrHiLinkLicenseStatus        = &Errno{Code: 21304, Message: "license状态只允许1或者2", EnMessage: "License status only allows 1 or 2"}
	ErrHiLinkLicenseNumber        = &Errno{Code: 21305, Message: "token预置license数量不足", EnMessage: "The number of token pre-set licenses is not enough"}
	ErrHiLinkLicenseExpiredToken  = &Errno{Code: 21306, Message: "token已过期", EnMessage: "token has expired"}
	ErrHiLinkMacNotExist          = &Errno{Code: 21307, Message: "mac地址不存在", EnMessage: "mac address does not exist"}
	ErrHiLinkMacExist             = &Errno{Code: 21308, Message: "mac地址已存在,且烧录成功", EnMessage: "The mac address already exists, and the burn-in is successful."}
	ErrHiLinkMacRetryCountLimit   = &Errno{Code: 21309, Message: "相同mac地址重试获取license达到上限", EnMessage: "Retry to get license with the same mac address reached the limit"}
	ErrHiLinkLicenseTokenStatus   = &Errno{Code: 21310, Message: "正在同步license，请稍后再试试", EnMessage: "Synchronizing licenses, please try again later"}
	ErrGetHiLicenseApplyRecord    = &Errno{Code: 21311, Message: "获取license申请记录失败", EnMessage: "Failed to obtain the license application record list"}
	ErrCreateHiLicenseApplyRecord = &Errno{Code: 21312, Message: "创建一条license申请记录失败", EnMessage: "Failed to create a license application record"}
	ErrGetHiLicenseApplyDevice    = &Errno{Code: 21313, Message: "获取license失败", EnMessage: "Failed to obtain license"}
	ErrGetBatchNotExist           = &Errno{Code: 21314, Message: "该批次不存在", EnMessage: "The batch does not exist"}
	ErrLicenseApplyFix            = &Errno{Code: 21315, Message: "该批次申请无异常，不需要补救", EnMessage: "There is no abnormality in the batch application and no remediation is required"}
	ErrGetConsumedNum             = &Errno{Code: 21316, Message: "获取license列表已消耗/未消耗数量失败", EnMessage: "Failed to obtain the consumed/unconsumed quantity of the license list"}
	ErrGetLicense                 = &Errno{Code: 21317, Message: "该token没有可用license", EnMessage: "No license is available for this token"}

	// 产品分类，前缀为 214
	ErrGetProductCategory            = &Errno{Code: 21401, Message: "产品一级分类查询失败", EnMessage: "Product first category query failed"}
	ErrCreateProductCategory         = &Errno{Code: 21402, Message: "产品一级分类创建失败", EnMessage: "Failed to create product first category"}
	ErrUpdateProductCategory         = &Errno{Code: 21403, Message: "产品一级分类修改失败", EnMessage: "Product first category modification failed"}
	ErrUpdateProductCategoryNotExist = &Errno{Code: 21404, Message: "产品一级分类不存在", EnMessage: "Product first category does not exist"}
	ErrHasSecondaryCategory          = &Errno{Code: 21405, Message: "该一级分类下有二级分类，不允许删除", EnMessage: "There is a secondary category under this first category, and deletion is not allowed."}
	ErrDeleteProductCategory         = &Errno{Code: 21406, Message: "产品一级分类删除失败", EnMessage: "Product first category deletion failure"}
	ErrOrderProductCategory          = &Errno{Code: 21407, Message: "产品一级/二级分类排序失败", EnMessage: "Product first-category/sub-category sorting failed"}
	ErrLanguageProductCategory       = &Errno{Code: 21408, Message: "产品一级/二级分类多语言设置失败", EnMessage: "Product first-category/sub-category multi-language setting failed"}
	ErrAlreadyProductCategory        = &Errno{Code: 21409, Message: "产品一级分类名称重复", EnMessage: "Duplicate product first category name"}

	// 分享模版，前缀为 216
	ErrShareUserNotExist         = &Errno{Code: 21601, Message: "共享用户不存在", EnMessage: "Shared user does not exist"}
	ErrShareHouseNotExist        = &Errno{Code: 21602, Message: "共享用户房子不存在", EnMessage: "Shared user house does not exist"}
	ErrShareDeviceNotExist       = &Errno{Code: 21603, Message: "共享用户设备不存在", EnMessage: "Shared user device does not exist"}
	ErrNotShareUser              = &Errno{Code: 21604, Message: "非共享账号", EnMessage: "Non-shared account"}
	ErrDeviceUnShareFailure      = &Errno{Code: 21605, Message: "设备取消分享失败", EnMessage: "Device unshare failed"}
	ErrHouseUnShareFailure       = &Errno{Code: 21606, Message: "房子取消分享失败", EnMessage: "House unshare failed"}
	ErrShareTokenInvalid         = &Errno{Code: 21607, Message: "分享token无效", EnMessage: "Share token is invalid"}
	ErrAlreadyShareHouse         = &Errno{Code: 21608, Message: "账号重复添加", EnMessage: "Duplicate account added"}
	ErrNoShareHouse              = &Errno{Code: 21609, Message: "房子没有分享给该用户", EnMessage: "House is not shared to this user"}
	ErrAlreadyShareDeviceGroup   = &Errno{Code: 21610, Message: "群组已经分享给该用户", EnMessage: "Group has been shared to this user"}
	ErrDeviceGroupShareFailure   = &Errno{Code: 21611, Message: "群组分享失败", EnMessage: "Group sharing failed"}
	ErrShareDeviceGroupNotExist  = &Errno{Code: 21612, Message: "共享设备群组不存在", EnMessage: "Shared device group does not exist"}
	ErrDeviceGroupUnShareFailure = &Errno{Code: 21613, Message: "设备群组取消分享失败", EnMessage: "Device group unshare failed"}
	ErrShareAlreadyExist         = &Errno{Code: 21614, Message: "已存在的分享", EnMessage: "Share already exist"}
	ErrShareNoExist              = &Errno{Code: 21615, Message: "不存在的分享", EnMessage: "Share not exist"}
	ErrShareThingNoExist         = &Errno{Code: 21616, Message: "分享事物不存在", EnMessage: "Share thing not exist"}
	ErrSharePermission           = &Errno{Code: 21617, Message: "分享权限失效", EnMessage: "illegal share permission"}

	//三方接口模块 ，前缀为 217
	ErrIotDeviceBind      = &Errno{Code: 21701, Message: "iot绑定设备接口异常", EnMessage: "iot bind device interface exception"}
	ErrIotDeviceShare     = &Errno{Code: 21702, Message: "iot分享设备接口异常", EnMessage: "iot sharing device interface abnormal"}
	ErrIotDeviceInfo      = &Errno{Code: 21703, Message: "iot查询设备信息失败", EnMessage: "iot query device information failed"}
	ErrIotGroupShare      = &Errno{Code: 21704, Message: "iot组分享接口异常", EnMessage: "iot group sharing interface abnormal"}
	ErrPlatformShadow     = &Errno{Code: 21705, Message: "平台设备影子获取失败", EnMessage: "Platform device shadow acquisition failure"}
	ErrPlatformDevice     = &Errno{Code: 21706, Message: "平台设备信息获取失败", EnMessage: "Platform device information acquisition failure"}
	ErrCreateDeviceShadow = &Errno{Code: 21707, Message: "创建设备影子失败", EnMessage: "Failed to create device shadow"}

	// 配置信息模块, 前缀为 218
	ErrGetConfigMap = &Errno{Code: 21801, Message: "获取配置信息失败", EnMessage: "Failed to get configuration information"}
	ErrFmtConfigMap = &Errno{Code: 21802, Message: "解析配置信息失败", EnMessage: "Failed to parse configuration information"}

	// 品类功能点， 前缀为 219
	ErrGetModelProperty    = &Errno{Code: 21901, Message: "获取品类属性失败", EnMessage: "Failed to get category attributes"}
	ErrCreateModelProperty = &Errno{Code: 21902, Message: "创建品类属性失败", EnMessage: "Failed to create category attributes"}
	ErrUpdateModelProperty = &Errno{Code: 21903, Message: "修改品类属性失败", EnMessage: "Failed to modify category attributes"}
	ErrDeleteModelProperty = &Errno{Code: 21904, Message: "删除品类属性失败", EnMessage: "Failed to delete category attributes"}
	ErrGetModelSpec        = &Errno{Code: 21905, Message: "获取品类规格失败", EnMessage: "Failed to get category specification"}
	ErrCreateModelSpec     = &Errno{Code: 21906, Message: "创建品类规格失败", EnMessage: "Failed to create category specification"}
	ErrUpdateModelSpec     = &Errno{Code: 21907, Message: "更新品类规格失败", EnMessage: "Update Category Specification Failure"}

	// 固件管理，前缀为 220
	ErrGetFirmware                  = &Errno{Code: 22001, Message: "获取固件失败", EnMessage: "Failed to get firmware"}
	ErrCreateFirmware               = &Errno{Code: 22002, Message: "创建固件失败", EnMessage: "Failed to create firmware"}
	ErrUpdateFirmware               = &Errno{Code: 22003, Message: "修改固件失败", EnMessage: "Modify firmware failure"}
	ErrGetFirmwareMD5               = &Errno{Code: 22004, Message: "获取固件的MD5值失败", EnMessage: "Failed to get MD5 value of firmware"}
	ErrUploadFirmwareToIot          = &Errno{Code: 22005, Message: "提交固件到iot失败", EnMessage: "Submit firmware to iot failed"}
	ErrGetVerifyDevice              = &Errno{Code: 22006, Message: "获取验证设备失败", EnMessage: "Failed to get authentication device"}
	ErrCreateVerifyDevice           = &Errno{Code: 22007, Message: "创建固件验证设备失败", EnMessage: "Failed to create firmware authentication device"}
	ErrEnableVerifyDevice           = &Errno{Code: 22008, Message: "禁用低版本验证设备失败", EnMessage: "Failed to disable low version verification device"}
	ErrIsNotNewVerifyDevice         = &Errno{Code: 22009, Message: "该固件不是最新版本，不允许校验", EnMessage: "The firmware is not the latest version and verification is not allowed"}
	ErrUpdateTimeoutVerifyDevice    = &Errno{Code: 22010, Message: "更新超时时间失败", EnMessage: "Failed to update timeout time"}
	ErrGetReleaseDevice             = &Errno{Code: 22011, Message: "获取发布详情列表失败", EnMessage: "Failed to get the release details list"}
	ErrCreateReleaseDevice          = &Errno{Code: 22012, Message: "固件发布失败", EnMessage: "Firmware release failed"}
	ErrCreateReleaseDeviceExist     = &Errno{Code: 22013, Message: "该固件已发布，不能重复发布", EnMessage: "The firmware has been released and cannot be released repeatedly"}
	ErrRevokeReleaseDevice          = &Errno{Code: 22014, Message: "撤销发布任务失败", EnMessage: "Unpublished task failed"}
	ErrCreateVerifyDeviceNotNew     = &Errno{Code: 22015, Message: "只能添加最新的固件版本验证/发布", EnMessage: "Only the latest firmware version verification can be added"}
	ErrGetDevice                    = &Errno{Code: 22016, Message: "设备不存在", EnMessage: "Device does not exist"}
	ErrDeviceExist                  = &Errno{Code: 22017, Message: "设备重复添加验证", EnMessage: "Repeat device addition verification"}
	ErrSelectUpgradeDevicesRequired = &Errno{Code: 22018, Message: "定向升级必须要传devices", EnMessage: "Targeted upgrade must upload devices"}
	ErrUpgradeTimeRequired          = &Errno{Code: 22019, Message: "定时升级必须要传时间段", EnMessage: "Time period must be transmitted for scheduled upgrade"}
	ErrVerifyDeviceCannotExceed20   = &Errno{Code: 22020, Message: "固件验证总数不能超过20个设备", EnMessage: "The total number of firmware verification cannot exceed 20 devices"}
	ErrNoVerifyNotRelease           = &Errno{Code: 22021, Message: "没有验证通过的固件不允许发布", EnMessage: "Firmware that has not been verified is not allowed to be released"}
	ErrGetFirmwareNotExist          = &Errno{Code: 22022, Message: "固件不存在", EnMessage: "Firmware does not exist"}
	ErrGetDeviceForProduct          = &Errno{Code: 22023, Message: "该产品下没有查询到设备", EnMessage: "No device is found under this product"}
	ErrUpdateVerifyStatus           = &Errno{Code: 22024, Message: "更新固件验证状态失败", EnMessage: "Failed to update firmware verification status"}
	ErrGetNewFirmware               = &Errno{Code: 22025, Message: "获取最新版本固件失败", EnMessage: "Failed to get the latest firmware"}
	ErrCreateFirmwareNotNew         = &Errno{Code: 22026, Message: "不能创建低版本固件", EnMessage: "Cannot create lower version firmware"}
	ErrReleaseDeviceCannotExceed20  = &Errno{Code: 22027, Message: "定向升级不能超过20个设备", EnMessage: "Targeted upgrade cannot exceed 20 devices"}
	ErrVerifyDeviceIsNew            = &Errno{Code: 22028, Message: "已经是最新版本，无需验证", EnMessage: "Already the latest version, no verification required"}
	ErrEnableFirmware               = &Errno{Code: 22029, Message: "禁用低版本固件的发布任务失败", EnMessage: "Failed to disable the release task of lower version firmware"}
	ErrAlreadyUpgradeFirmware       = &Errno{Code: 22030, Message: "已经是最新版本，无需升级", EnMessage: "Already the latest version, no need to upgrade"}
	ErrUpdatePublishedFirmware      = &Errno{Code: 22031, Message: "只有开发中/已试产状态的产品，可以编辑已发布固件", EnMessage: "Only products in development/trial production status can edit the released firmware"}
	ErrDoesNotVerify                = &Errno{Code: 22032, Message: "您选择的产品暂不支持验证", EnMessage: "The product you selected does not support verification"}
	ErrDoesNotSelectUpgrade         = &Errno{Code: 22033, Message: "不支持定向升级", EnMessage: "Does not support targeted upgrade"}
	ErrDeleteFirmware               = &Errno{Code: 22034, Message: "删除固件失败", EnMessage: "Failed to delete firmware"}
	ErrOtaUpdateCommandErr          = &Errno{Code: 22035, Message: "下发固件升级指令失败", EnMessage: "Failed to issue firmware upgrade command"}
	ErrUpdateBurnFirmware           = &Errno{Code: 22036, Message: "只有开发中/已试产状态的产品，可以编辑烧录固件", EnMessage: "Only products in development/trial production status can be edited and burned firmware"}

	// 产品二级分类，前缀为 221
	ErrGetProductSecondaryCategory      = &Errno{Code: 22101, Message: "产品二级分类查询失败", EnMessage: "Product secondary category query failed"}
	ErrCreateProductSecondaryCategory   = &Errno{Code: 22102, Message: "产品二级分类创建失败", EnMessage: "Failed to create product secondary category"}
	ErrUpdateProductSecondaryCategory   = &Errno{Code: 22103, Message: "产品二级分类修改失败", EnMessage: "Product secondary category modification failed"}
	ErrProductSecondaryCategoryNotExist = &Errno{Code: 22104, Message: "产品二级分类不存在", EnMessage: "Product secondary category does not exist"}
	ErrDeleteProductSecondaryCategory   = &Errno{Code: 22105, Message: "产品二级分类删除失败", EnMessage: "Product secondary category failed to delete"}
	ErrHasProductTemplate               = &Errno{Code: 22106, Message: "该二级分类下有产品模板，不允许删除", EnMessage: "There is a product template under this secondary category, and deletion is not allowed"}
	ErrAlreadyProductSecondaryCategory  = &Errno{Code: 22107, Message: "产品二级分类名称重复", EnMessage: "Product secondary category name duplicate"}
	ErrHasProduct                       = &Errno{Code: 22108, Message: "该二级分类下有产品，不允许删除", EnMessage: "There are products under this secondary category, deletion is not allowed"}

	//产品标准功能池， 前缀为 222
	ErrProductPropertyModelNotExist = &Errno{Code: 22201, Message: "产品标准功能池查询失败", EnMessage: "Product standard function pool query failed"}
	ErrCreateProductPropertyModel   = &Errno{Code: 22202, Message: "产品标准功能池创建失败", EnMessage: "Product standard function pool creation failed"}
	ErrUpdateProductPropertyModel   = &Errno{Code: 22203, Message: "产品标准功能池编辑失败", EnMessage: "Product standard function pool edit failure"}
	ErrDeleteProductPropertyModel   = &Errno{Code: 22204, Message: "产品标准功能池删除失败", EnMessage: "Product standard function pool deletion failure"}
	ErrPropertyModelMultiLanguage   = &Errno{Code: 22205, Message: "修改标准功能多语言失败", EnMessage: "Failed to modify standard function multi-language"}

	//产品模板， 前缀为 223
	ErrGetProductTemplate         = &Errno{Code: 22301, Message: "产品模板查询失败", EnMessage: "Product template query failed"}
	ErrProductTemplateNotExist    = &Errno{Code: 22302, Message: "产品模板不存在", EnMessage: "Product template does not exist"}
	ErrCreateProductTemplate      = &Errno{Code: 22303, Message: "产品模板创建失败", EnMessage: "Product template creation failed"}
	ErrUpdateProductTemplate      = &Errno{Code: 22304, Message: "更新产品模板失败", EnMessage: "Update product template failed"}
	ErrProductTemplateAuthFailure = &Errno{Code: 22305, Message: "产品模板权限异常", EnMessage: "Product template permission exception"}
	ErrOrderProductTemplate       = &Errno{Code: 22306, Message: "产品模板排序失败", EnMessage: "Product template sorting failure"}
	ErrProductTemplateParam       = &Errno{Code: 22307, Message: "免开发模版，必须要传通信方式、配网方式", EnMessage: "No development template, communication mode and network distribution mode must be transmitted"}
	ErrDeleteProductTemplate      = &Errno{Code: 22308, Message: "产品模板下有关联产品，不允许删除", EnMessage: "There are related products under the product template and it is not allowed to delete"}
	//产品模版 - 面板设置
	ErrGetProductTemplatePanel      = &Errno{Code: 22310, Message: "获取面板失败", EnMessage: "Failed to get panel"}
	ErrUpdateProductTemplatePanel   = &Errno{Code: 22311, Message: "修改面板失败", EnMessage: "Failed to modify panel"}
	ErrCreateProductTemplatePanel   = &Errno{Code: 22312, Message: "创建面板失败", EnMessage: "Failed to create panel"}
	ErrProductTemplatePanelNotExist = &Errno{Code: 22313, Message: "面板不存在", EnMessage: "Panel does not exist"}
	ErrDeleteProductTemplatePanel   = &Errno{Code: 22314, Message: "删除面板失败", EnMessage: "Failed to delete panel"}
	ErrProductTemplatePanelExist    = &Errno{Code: 22315, Message: "版本号不允许重复", EnMessage: "Panel version number is not allowed to be repeated"}
	//产品模版 - 语音技能
	ErrVoicePlatformNotExist       = &Errno{Code: 22318, Message: "语音平台不存在", EnMessage: "Voice platform does not exist"}
	ErrEnableProductTemplateVoice  = &Errno{Code: 22319, Message: "启用/禁用语音平台失败", EnMessage: "Failed to enable/disable the voice platform"}
	ErrProductTemplateVoiceDisable = &Errno{Code: 22320, Message: "产品模版未启用该语音平台，不能开通", EnMessage: "The voice platform is not enabled in the product template and cannot be activated"}
	ErrOpenProductVoicePrice       = &Errno{Code: 22321, Message: "服务费用错误，开通语音服务失败", EnMessage: "Wrong service fee, failed to activate voice service"}
	ErrCloseProductVoice           = &Errno{Code: 22322, Message: "关闭语音服务失败", EnMessage: "Failed to turn off the voice service"}
	ErrOpenProductVoice            = &Errno{Code: 22323, Message: "开通语音服务失败", EnMessage: "Failed to activate voice service"}
	ErrDisableProductTemplateVoice = &Errno{Code: 22324, Message: "该模版下已开通生效语音产品，不能禁用", EnMessage: "Voice products have been activated under this template and cannot be disabled"}
	ErrCreateVoiceExample          = &Errno{Code: 22325, Message: "创建语音示例失败", EnMessage: "Failed to create voice sample"}
	ErrVoiceExampleNotExist        = &Errno{Code: 22326, Message: "语音示例不存在", EnMessage: "Voice sample does not exist"}
	ErrUpdateVoiceExample          = &Errno{Code: 22327, Message: "修改语音示例失败", EnMessage: "Failed to update voice sample"}
	ErrDeleteVoiceExample          = &Errno{Code: 22328, Message: "删除语音示例失败", EnMessage: "Failed to delete voice sample"}

	//模组管理， 前缀为 224
	ErrGetModule              = &Errno{Code: 22401, Message: "模组查询失败", EnMessage: "Module query failure"}
	ErrCreateModule           = &Errno{Code: 22402, Message: "模组创建失败", EnMessage: "Module creation failure"}
	ErrUpdateModule           = &Errno{Code: 22403, Message: "模组编辑失败", EnMessage: "Module edit failure"}
	ErrDeleteModule           = &Errno{Code: 22404, Message: "模组删除失败", EnMessage: "Module deletion failure"}
	ErrModuleNotExist         = &Errno{Code: 22405, Message: "模组不存在", EnMessage: "Module does not exist"}
	ErrDeleteModuleHasProduct = &Errno{Code: 22406, Message: "有产品关联该模组，不允许删除", EnMessage: "There is a product associated with the module and deletion is not allowed"}

	//虚拟设备， 前缀为 225
	ErrVirtualDeviceNotExist   = &Errno{Code: 22501, Message: "虚拟设备不存在", EnMessage: "Virtual device does not exist"}
	ErrDeleteVirtualDevice     = &Errno{Code: 22502, Message: "虚拟设备删除失败", EnMessage: "Virtual device deletion failed"}
	ErrGetVirtualDevice        = &Errno{Code: 22503, Message: "虚拟设备查询失败", EnMessage: "Virtual device query failed"}
	ErrCreateVirtualDevice     = &Errno{Code: 22504, Message: "虚拟设备创建失败", EnMessage: "Virtual device creation failed"}
	ErrCreateVirtualDeviceCode = &Errno{Code: 22505, Message: "缓存虚拟设备分享码失败", EnMessage: "Failed to cache virtual device sharing code"}
	ErrDeleteVirtualDeviceUser = &Errno{Code: 22506, Message: "用户绑定的虚拟设备删除失败", EnMessage: "User-bound virtual device deletion failure"}
	ErrRealDeviceOnline        = &Errno{Code: 22507, Message: "真实设备暂时不支持上下线操作", EnMessage: "Real devices do not support online and offline operations temporarily"}

	//数据统计，前缀为 226
	ErrSTParseInt             = &Errno{Code: 22601, Message: "ParseInt失败", EnMessage: "ParseInt failure"}
	ErrSTGetUserProductKey    = &Errno{Code: 22602, Message: "获取用户的ProductKey失败", EnMessage: "Failed to get the user's ProductKey"}
	ErrSTActivationCount      = &Errno{Code: 22603, Message: "获取激活数据失败", EnMessage: "Failed to get activation data"}
	ErrSTActivationTrendCount = &Errno{Code: 22604, Message: "获取激活趋势失败", EnMessage: "Failed to get activation trend"}
	ErrSTActiveCount          = &Errno{Code: 22605, Message: "获取活跃数据失败", EnMessage: "Failed to get active data"}
	ErrSTActiveTrendCount     = &Errno{Code: 22606, Message: "获取活跃趋势失败", EnMessage: "Failed to get active trend"}
	ErrSTRegionCount          = &Errno{Code: 22607, Message: "获取地区数据失败", EnMessage: "Failed to get region data"}
	ErrSTProductDeviceCount   = &Errno{Code: 22608, Message: "获取产品设备数失败", EnMessage: "Failed to get product device count"}
	ErrSTOverview             = &Errno{Code: 22609, Message: "获取数据概览失败", EnMessage: "Failed to get data overview"}
	ErrSTFaultCount           = &Errno{Code: 22610, Message: "获取故障数量失败", EnMessage: "Failed to get data fault count"}
	ErrSTFaultTrend           = &Errno{Code: 22611, Message: "获取故障趋势失败", EnMessage: "Failed to get data fault trend"}

	//移动设备推送
	ErrMobileDeviceSync = &Errno{Code: 22701, Message: "移动设备同步基础信息失败", EnMessage: "Mobile device synchronization base info failed"}

	//面板管理, 前缀为 228
	ErrGetPanelUpdateRecord                 = &Errno{Code: 22801, Message: "获取面板更新记录列表失败", EnMessage: "Failed to get the panel update record list"}
	ErrFileUrlNotExist                      = &Errno{Code: 22802, Message: "ios/android文件地址，至少存在一个", EnMessage: "ios/android file address, at least one exists"}
	ErrUpdatePanelCannotExceed100           = &Errno{Code: 22803, Message: "一次最多更新100个产品", EnMessage: "Up to 100 products can be updated at a time"}
	ErrCreatePanelUpdateRecord              = &Errno{Code: 22804, Message: "创建面板更新记录失败", EnMessage: "Failed to create panel update record"}
	ErrUpdatePanelUpdateRecord              = &Errno{Code: 22805, Message: "修改面板更新记录状态失败", EnMessage: "Failed to update the record status in the modification panel"}
	ErrProductNotInDevelopOrTrialProduction = &Errno{Code: 22806, Message: "只有开发中/试产中的产品允许修改面板", EnMessage: "Only products in development/trial production are allowed to modify the panel"}
	ErrPanelUpdateRecordNotExist            = &Errno{Code: 22807, Message: "面板更新记录不存在", EnMessage: "Panel update record does not exist"}

	// 消息订阅
	ErrMessageSubscribeNotAllowed    = &Errno{Code: 23000, Message: "没有权限操作", EnMessage: "message subscribe not allowed "}
	ErrMessageSubscribeNotExist      = &Errno{Code: 23000, Message: "消息订阅数据不存在", EnMessage: "message subscribe does not exist"}
	ErrDataSourceNotExist            = &Errno{Code: 23001, Message: "数据源数据不存在", EnMessage: "data source does not exist"}
	ErrSubscribeActionNotExist       = &Errno{Code: 23002, Message: "行为操作数据不存在", EnMessage: "subscribe action does not exist"}
	ErrDataSourceAlreadyExist        = &Errno{Code: 23002, Message: "数据源已存在", EnMessage: "data source already exist"}
	ErrMessageSubscribeMsgTypeFormat = &Errno{Code: 23003, Message: "订阅消息类型格式错误", EnMessage: "message type format is incorrect"}
	ErrMessageSubscribeForeignPolicy = &Errno{Code: 23004, Message: "已启动的消息订阅无法进行编辑", EnMessage: "message subscribe foreign policy"}

	//语音技能，前缀为 229
	ErrUpdateDeviceMap   = &Errno{Code: 22901, Message: "更新设备映射表失败", EnMessage: "Failed to update device mapping list"}
	ErrGetDeviceStatus   = &Errno{Code: 22902, Message: "查询设备功能点失败", EnMessage: "Failed to search device status"}
	ErrGetDeviceModel    = &Errno{Code: 22903, Message: "查询设备模板失败", EnMessage: "Failed to search model"}
	ErrControlDevice     = &Errno{Code: 22904, Message: "设备控制失败", EnMessage: "Failed to control device"}
	ErrLinkSkillSearch   = &Errno{Code: 22905, Message: "未找到用户记录", EnMessage: "Failed to find user-platform record"}
	ErrLinkSkillUpdate   = &Errno{Code: 22906, Message: "更新用户记录失败", EnMessage: "Failed to update user-platform record"}
	ErrCheckVerifyFailed = &Errno{Code: 22907, Message: "校验验证码失败", EnMessage: "Check verify code failed"}
	ErrInvalidVerifyCode = &Errno{Code: 22908, Message: "非法验证码", EnMessage: "Invalid verify code"}
	ErrExpiredVerifyCode = &Errno{Code: 22909, Message: "未找到此验证码或验证码已过期", EnMessage: "Verify code is expired"}

	// 云配置，前缀为 231
	ErrCreateCloudConfigAlreadyExist = &Errno{Code: 23101, Message: "该token已存在云配置，不能重复创建", EnMessage: "The token already exists in the cloud configuration and cannot be created repeatedly"}
	ErrCreateCloudConfig             = &Errno{Code: 23102, Message: "创建云配置失败", EnMessage: "Failed to create cloud configuration"}
	ErrGetCloudConfig                = &Errno{Code: 23103, Message: "获取云配置信息失败", EnMessage: "Failed to obtain cloud configuration information"}
	ErrGetCloudConfigNotExist        = &Errno{Code: 23104, Message: "云配置不存在", EnMessage: "Cloud configuration does not exist"}
	ErrUpdateCloudConfig             = &Errno{Code: 23105, Message: "修改云配置失败", EnMessage: "Failed to modify cloud configuration"}

	// 产品标签，前缀为 232
	ErrCreateProductLabelAlreadyExist = &Errno{Code: 23201, Message: "不能创建重复标签", EnMessage: "Cannot create duplicate labels"}
	ErrCreateProductLabelMost20       = &Errno{Code: 23202, Message: "一次最多创建10条，总数不能超过20条", EnMessage: "A maximum of 10 entries can be created at a time, and the total number cannot exceed 20 entries"}
	ErrCreateProductLabel             = &Errno{Code: 23203, Message: "创建产品标签失败", EnMessage: "Failed to create product label"}
	ErrGetProductLabelNotExist        = &Errno{Code: 23204, Message: "产品标签不存在", EnMessage: "Product label does not exist"}
	ErrDeleteProductLabel             = &Errno{Code: 23205, Message: "删除产品标签失败", EnMessage: "Failed to delete product label"}
	ErrUpdateProductLabel             = &Errno{Code: 23206, Message: "修改产品标签失败", EnMessage: "Failed to modify product label"}

	//群组管理， 前缀为 233
	ErrCreateGroupsMost100         = &Errno{Code: 23301, Message: "一个群组最多添加100个产品", EnMessage: "A group can add up to 100 products"}
	ErrCreateGroupsAlreadyExist    = &Errno{Code: 23302, Message: "一个产品只能存在一个群组中", EnMessage: "A product can only exist in one group"}
	ErrCreateGroups                = &Errno{Code: 23303, Message: "创建群组失败", EnMessage: "Failed to create group"}
	ErrGetGroupsNotExist           = &Errno{Code: 23304, Message: "群组不存在", EnMessage: "Group does not exist"}
	ErrUpdateGroups                = &Errno{Code: 23305, Message: "修改群组失败", EnMessage: "Failed to modify group"}
	ErrDeleteGroupsHasProduct      = &Errno{Code: 23306, Message: "此群组下存在产品，不能删除", EnMessage: "There are products in this group and cannot be deleted"}
	ErrDeleteGroups                = &Errno{Code: 23307, Message: "删除群组失败", EnMessage: "Failed to delete group"}
	ErrDeleteGroupProduct          = &Errno{Code: 23308, Message: "删除群组下的产品失败", EnMessage: "Failed to delete products under the group"}
	ErrUpdateProductGroupConfig    = &Errno{Code: 23309, Message: "修改产品是否支持群组配置失败", EnMessage: "Failed to modify whether the product supports group configuration"}
	ErrHasProductGroup             = &Errno{Code: 23310, Message: "该产品已关联群组管理，请先删除此群组", EnMessage: "This product has been associated with group management, please delete this group first"}
	ErrHasProductGroupAlreadyExist = &Errno{Code: 23311, Message: "一个群组下不能存在重复产品", EnMessage: "Duplicate products cannot exist in a group"}
	ErrCreateProductGroup          = &Errno{Code: 23312, Message: "创建群组下产品失败", EnMessage: "Failed to create product under group"}
	ErrNotSupportGroupProduct      = &Errno{Code: 23313, Message: "产品不存在或不支持群组", EnMessage: "The product does not exist or does not support groups"}

	//租户管理， 前缀为 234
	ErrGetTenantNotExist       = &Errno{Code: 23401, Message: "租户不存在", EnMessage: "Tenant does not exist"}
	ErrGetTenantParentNotExist = &Errno{Code: 23402, Message: "租户所属父账号不存在", EnMessage: "The parent account of the tenant does not exist"}

	//OEM App， 前缀为 235
	ErrCreateOemApp                 = &Errno{Code: 23501, Message: "创建OEM App失败", EnMessage: "Failed to create OEM App"}
	ErrUpdateOemApp                 = &Errno{Code: 23502, Message: "编辑OEM App失败", EnMessage: "Failed to edit OEM App"}
	ErrUpdateOemAppIcon             = &Errno{Code: 23503, Message: "编辑OEM App图标失败", EnMessage: "Failed to edit OEM App icon"}
	ErrOemAppConfigNotExist         = &Errno{Code: 23504, Message: "OEM App配置不存在", EnMessage: "OEM App configuration does not exist"}
	ErrUpdateOemAppThemeColor       = &Errno{Code: 23505, Message: "编辑OEM App主题色失败", EnMessage: "Failed to edit OEM App theme color"}
	ErrUpdateOemAppActionBarIcon    = &Errno{Code: 23506, Message: "编辑OEM App底部操作栏图标失败", EnMessage: "Failed to edit the icon in the operation bar at the bottom of the OEM App"}
	ErrUpdateOemAppIosCertificate   = &Errno{Code: 23507, Message: "编辑OEM App Ios证书失败", EnMessage: "Failed to edit OEM App Ios certificate"}
	ErrOemAppHistoryVersionNotExist = &Errno{Code: 23508, Message: "OEM App历史版本不存在", EnMessage: "OEM App history version does not exist"}
	ErrOemAppBuildFailure           = &Errno{Code: 23509, Message: "OEM App构建失败", EnMessage: "OEM App failed to build"}
	ErrOemAppBuildNotExist          = &Errno{Code: 23510, Message: "OEM App构建记录不存在", EnMessage: "OEM App build record does not exist"}
)
