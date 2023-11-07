package do

import "sync"

type LoginReq struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

type LoginRsp struct {
	Token    string `json:"token"`
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
}

type RegisterReq struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

type RegisterRsp struct {
}

type ObjectPool struct {
	pool sync.Pool
}

func NewObjectPool() *ObjectPool {
	return &ObjectPool{
		pool: sync.Pool{
			New: func() interface{} {
				return &LoginReq{} // 创建新的对象
			},
		},
	}
}

func (p *ObjectPool) AcquireObject() *LoginReq {
	obj := p.pool.Get() // 从对象池获取对象
	if obj == nil {
		return &LoginReq{} // 如果对象池为空，创建新的对象
	}
	return obj.(*LoginReq)
}

func (p *ObjectPool) ReleaseObject(obj *LoginReq) {
	p.pool.Put(obj) // 将对象放回对象池
}
