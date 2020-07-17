package tool
import (
	"errors"
	"fmt"
	"sync"

	common "go_crawler/common"
)

type ChannelManagerStatus uint8

const (
	CHANMANAGERSTATUSUNINIT ChannelManagerStatus = 0 //未初始化
	CHANMANAGERSTATUSINIT   ChannelManagerStatus = 1 //已初始化
	CHANMANAGERSTATUSSTOP   ChannelManagerStatus = 2 //暂停
	CHANMANAGERSTATUSCLOSE  ChannelManagerStatus = 3 //关闭
)

// 映射关系字典。
var statusNameMap = map[ChannelManagerStatus]string{
	CHANMANAGERSTATUSUNINIT: "unInitialize",
	CHANMANAGERSTATUSINIT:   "initialize",
	CHANMANAGERSTATUSSTOP:   "stop",
	CHANMANAGERSTATUSCLOSE:  "close",
}

// 通道管理接口类型。
type ChannelManager interface {
	// 初始化通道管理器。
	Init(channelParams common.ChanParams, reset bool) bool
	// 关闭通道管理器。
	Close() bool
	// 获取请求传输通道。
	GetReqChan() (chan common.HttpRequest, error)
	// 获取响应传输通道。
	GetRespChan() (chan common.HttpResponse, error)
	// 获取条目传输通道。
	GetItemChan() (chan common.Item, error)
	// 获取错误传输通道。
	GetErrorChan() (chan error, error)
	// 获取通道管理器的状态。
	GetStatus() ChannelManagerStatus
}

//通道管理实例
type ChanManager struct{
	chanParams  common.ChanParams         // 通道容器。
	reqCh       chan common.HttpRequest   // 请求通道。
	respCh      chan common.HttpResponse  // 响应通道。
	itemCh      chan common.Item          // 信息通道。
	errorCh     chan error                // 错误通道。
	status      ChannelManagerStatus      // 通道管理器的状态。
	rwMutex     sync.RWMutex              // 读写锁。
}

func (cm *ChanManager) Init(chanParams common.ChanParams,reset bool) bool {
	if err := chanParams.Check(); err != nil {
		panic(err)
	}
	cm.rwMutex.Lock()
	defer cm.rwMutex.Unlock()
	if cm.status == CHANMANAGERSTATUSINIT && !reset {
		return false
	}
	cm.chanParams = chanParams
	cm.reqCh = make(chan common.HttpRequest, chanParams.GetReqChanLength())
	cm.respCh = make(chan common.HttpResponse, chanParams.GetRespChanLength())
	cm.itemCh = make(chan common.Item, chanParams.GetItemChanLength())
	cm.errorCh = make(chan error, chanParams.GetErrorChanLength())
	cm.status = CHANMANAGERSTATUSINIT
	return true
}

func (cm *ChanManager) Close() bool {
	cm.rwMutex.Lock()
	defer cm.rwMutex.Unlock()
	if cm.status != CHANMANAGERSTATUSINIT {
		return false
	}
	close(cm.reqCh)
	close(cm.respCh)
	close(cm.itemCh)
	close(cm.errorCh)
	cm.status = CHANMANAGERSTATUSCLOSE
	return true
}

func (cm *ChanManager) checkStatus() error {
	if cm.status == CHANMANAGERSTATUSINIT {
		return nil
	}
	statusName, ok := statusNameMap[cm.status]
	if !ok {
		statusName = fmt.Sprintf("%d", cm.status)
	}
	errMsg :=
		fmt.Sprintf("Channel manager in an abnormal status: %s!\n",
			statusName)
	return errors.New(errMsg)
}

func (cm *ChanManager) GetReqChan() (chan common.HttpRequest, error) {
	cm.rwMutex.RLock()
	defer cm.rwMutex.RUnlock()
	if err := cm.checkStatus(); err != nil {
		return nil, err
	}
	return cm.reqCh, nil
}

func (cm *ChanManager) GetRespChan() (chan common.HttpResponse, error) {
	cm.rwMutex.RLock()
	defer cm.rwMutex.RUnlock()
	if err := cm.checkStatus(); err !=nil {
		return nil , err
	}
	return cm.respCh, nil
}

func (cm *ChanManager) GetItemChan() (chan common.Item, error) {
	cm.rwMutex.RLock()
	defer cm.rwMutex.RUnlock()
	if err := cm.checkStatus(); err != nil {
		return nil, err
	}
	return cm.itemCh, nil
}

func (cm *ChanManager) GetErrorChan() (chan error, error) {
	cm.rwMutex.RLock()
	defer cm.rwMutex.RUnlock()
	if err := cm.checkStatus(); err != nil {
		return nil, err
	}
	return cm.errorCh, nil
}

func (cm *ChanManager) GetStatus() ChannelManagerStatus {
	return cm.status
}

func NewChanManager (chanParams common.ChanParams) ChannelManager {
	NewCM := &ChanManager{}
	NewCM.Init(chanParams,true)
	return NewCM
}

