package rdis

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/mangohow/cloud-ide-webserver/internal/model"
	"github.com/mangohow/cloud-ide-webserver/pkg/utils"
)

const (
	HostsHashKey = "hosts"
)

func AddRunningSpace(uid string, space *model.RunningSpace) error {
	data, err := json.Marshal(space)
	if err != nil {
		return err
	}
	str := utils.Bytes2String(data)
	res := client.HSet(context.Background(), HostsHashKey, uid, str, space.Sid, space.Host)
	return res.Err()
}

// CheckHasRunningSpace 检查用户是否有正在运行的工作空间
func CheckHasRunningSpace(uid string) (bool, error) {
	res := client.HGet(context.Background(), HostsHashKey, uid)
	if err := res.Err(); err != nil {
		if err == redis.Nil {
			return false, nil
		}
		return false, err
	}

	if res.Val() == "" {
		return false, nil
	}

	return true, nil
}

func DeleteRunningSpace(uid string) error {
	space, err := GetRunningSpace(uid)
	if err != nil {
		return err
	}

	r := client.HDel(context.Background(), HostsHashKey, uid, space.Sid)

	return r.Err()
}

func GetRunningSpace(uid string) (*model.RunningSpace, error) {
	res := client.HGet(context.Background(), HostsHashKey, uid)
	if err := res.Err(); err != nil {
		return nil, err
	}
	val := res.Val()
	var space model.RunningSpace
	err := json.Unmarshal(utils.String2Bytes(val), &space)
	if err != nil {
		return nil, err
	}

	return &space, nil
}

func CheckRunningSpaceAndDelete(uid string) (bool, error) {
	res := client.HGet(context.Background(), HostsHashKey, uid)
	if err := res.Err(); err != nil {
		if err == redis.Nil {
			return false, nil
		}
	}
	if res.Val() == "" {
		return false, nil
	}

	space := &model.RunningSpace{}
	err := json.Unmarshal(utils.String2Bytes(res.Val()), space)
	if err != nil {
		return true, err
	}
	sid := space.Sid
	if err = client.HDel(context.Background(), HostsHashKey, uid, sid).Err(); err != nil {
		return true, err
	}

	return true, nil
}

func CheckIsRunning(sid string) (bool, error) {
	res := client.HGet(context.Background(), HostsHashKey, sid)
	if err := res.Err(); err != nil {
		if err == redis.Nil {
			return false, nil
		}
		return false, err
	}

	if res.Val() == "" {
		return false, nil
	}

	return true, nil
}
