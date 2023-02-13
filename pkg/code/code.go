package code

const (
	QuerySuccess uint32 = iota
	QueryFailed

	LoginSuccess
	LoginFailed
	LoginUserDeleted

	SpaceCreateSuccess
	SpaceCreateFailed
	SpaceCreateNameDuplicate
	SpaceCreateReachMaxCount

	SpaceStartSuccess
	SpaceStartFailed

	SpaceDeleteSuccess
	SpaceDeleteFailed
	SpaceDeleteIsRunning

	SpaceStopSuccess
	SpaceStopFailed
	SpaceStopIsNotRunning

	UserNameAvailable
	UserNameUnavailable
)

type UserStatus uint32

const (
	StatusNormal UserStatus = iota
	StatusDeleted
)

var messageForCode = map[uint32]string{
	QuerySuccess:             "查询成功",
	QueryFailed:              "查询失败",
	LoginSuccess:             "登录成功",
	LoginFailed:              "登录失败",
	LoginUserDeleted:         "用户已注销",
	SpaceCreateSuccess:       "创建成功",
	SpaceCreateFailed:        "创建失败",
	SpaceCreateNameDuplicate: "不能和已有工作空间名称重复",
	SpaceCreateReachMaxCount: "达到最大工作空间创建上限,请删除其它工作空间后重试",
	SpaceStartSuccess:        "工作空间启动成功",
	SpaceStartFailed:         "工作空间启动失败,请重试",
	SpaceDeleteSuccess:       "删除工作空间成功",
	SpaceDeleteFailed:        "删除工作空间失败",
	SpaceDeleteIsRunning:     "无法删除正在运行的工作空间,请先停止运行",
	SpaceStopSuccess:         "停止工作空间成功",
	SpaceStopFailed:          "停止工作空间失败",
	SpaceStopIsNotRunning:    "工作空间未运行",
	UserNameAvailable:        "用户名可用",
	UserNameUnavailable:      "用户名不可用",
}

func GetMessage(code uint32) string {
	return messageForCode[code]
}
