package services

import (
    "go/ws/models"
    "go/ws/services/cache"
)

// 获取礼物信息（礼物如果更新 后台要刷新缓存，防止礼物无效）
func GetGiftInfo(id int) (GrabInfo *models.Grab, err error) {
    GrabInfo, err = cache.GetGrabInfo(id)

    // 缓存获取失败 从db 中查询
    if GrabInfo.GrabPrice == 0 {
        // model search
        ModelGrabInfo := GrabInfo.GetGrabInfo(id)

        cache.CacheGiftInfo(&ModelGrabInfo, id)
        return &ModelGrabInfo, nil
    }

    return GrabInfo, nil
}
