/**
* @file   : user
* @descrip: handler 处理不完的业务，或者比较大的业务交给这里的 user
* @author : ch-yk
* @create : 2018-09-05 下午3:14
* @email  : commonheart.yk@gmail.com
**/

package service

import (
	"fmt"
	"sync"

	"gogin/internal/util"
	"gogin/model"
)

//TODO: 其实如果没有 耗时操作，比如 GenShortId, 可以不用并发，直接倒腾数据结构即可
func ListUser(username string, offset, limit int) ([]*model.UserInfo, uint64, error) {
	infos := make([]*model.UserInfo, 0)
	//拿到的是 []*UserModel，即对应数据库查到了哪些记录 (model.ListUser() 就是查询数据库)
	users, count, err := model.ListUser(username, offset, limit)
	if err != nil {
		return nil, count, err
	}

	//下面开始把从数据库查到的记录，拼装成 handler 找我(service) 要的格式 `[]*model.UserInfo`

	//先给 slice 填充 所有记录的 id, 为了后续把填充完毕的结果集放入 infos
	ids := []uint64{}
	for _, user := range users {
		ids = append(ids, user.Id)
	}

	//并发处理: 从 []*UserModel 列表到 []*model.UserInfo 列表
	wg := sync.WaitGroup{}
	//并发填充临时结果集 userList,
	userList := model.UserList{
		Lock:  new(sync.Mutex), //锁对象 -- 因为下面操作 map 需要加锁 (非线程安全)
		IdMap: make(map[uint64]*model.UserInfo, len(users)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	//开个新任务去填充 userList，所以新协程中操作 userList 必须加锁
	// 遍历数据库记录 users []*UserModel
	for _, u := range users {
		wg.Add(1)

		//新协程
		go func(u *model.UserModel) {
			defer wg.Done()

			shortId, err := util.GenShortId()
			if err != nil {
				errChan <- err
				return
			}

			userList.Lock.Lock()
			defer userList.Lock.Unlock()

			userList.IdMap[u.Id] = &model.UserInfo{
				Id:        u.Id,
				Username:  u.Username,
				SayHello:  fmt.Sprintf("Hello %s", shortId),
				Password:  u.Password,
				CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
			}
		}(u)
	}

	//等待上面对于 userList 的填充 (整个 for 循环)完毕
	go func() {
		wg.Wait()
		close(finished) //然后回到主协程 相当于 finished <-
	}()

	select {
		case <-finished:
		case err := <-errChan:
			return nil, count, err
	}

	//从临时结果集 userList这个结构中的 map 倒腾到最终结果集

	for _, id := range ids {
		infos = append(infos, userList.IdMap[id])
	}

	return infos, count, nil
}
