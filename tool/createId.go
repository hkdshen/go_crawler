package tool
import(
	"math"
	"sync"
)

type CreateID interface {
	GetID() uint32
}

type CreateId struct {
	id uint32 //当前ID
	idMutex sync.Mutex //操作id的锁
}

func (ci *CreateId) GetID() uint32 {
	ci.idMutex.Lock()
	defer ci.idMutex.Unlock()
	if ci.id < math.MaxUint32 {
		ci.id ++
	}else{
		ci.id = 0
	}
	return ci.id
}

func NewCreateID() CreateID {
	return &CreateId{}
}