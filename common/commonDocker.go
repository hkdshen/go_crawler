package common
import (
	"errors"
)
//channel容器
type ChanParams struct {
	reqChanLength uint32
	respChanLength uint32
	itemChanLength uint32
	errorChanLength uint32
	chanDescription string
}
//创建channel容器
func NewChanParams(reqChanLen uint32,respChanLen uint32,itemChanLen uint32,errorChanLen uint32) *ChanParams{
	return &ChanParams{reqChanLen,respChanLen,
		              itemChanLen,errorChanLen,""}
}

func (cp *ChanParams) Check() error {
	var err string
	if cp.reqChanLength >0 && cp.respChanLength >0 && cp.itemChanLength >0 && cp.errorChanLength >0 {
		err = ""
	}else{
		err = "The"
		if cp.reqChanLength == 0 {
			err += "reqChanLength,"
		}else if cp.respChanLength == 0 {
			err += "respChanLength,"
		}else if cp.itemChanLength == 0 {
			err += "itemChanLength,"
		}else if cp.errorChanLength == 0 {
			err += "errorChanLength"
		}
		err += "need greater than zero !\n"
	}

	if err == "" {
		return nil
	}else{
		return errors.New(err)
	}
}

func (cp *ChanParams) GetReqChanLength() uint32 {
	return cp.reqChanLength
}

func (cp *ChanParams) GetRespChanLength() uint32 {
	return cp.respChanLength
}

func (cp *ChanParams) GetItemChanLength() uint32 {
	return cp.itemChanLength
}

func (cp *ChanParams) GetErrorChanLength() uint32 {
	return cp.errorChanLength
}

//管理池尺寸
type PoolCommonSize struct{
	downloaderPoolSize uint32
	analyzerPoolSize uint32
}

func NewPoolCommon(downloadPoolSize uint32,analyzePoolSize uint32) *PoolCommonSize {
	return &PoolCommonSize{downloadPoolSize,analyzePoolSize}
}

func (pcs *PoolCommonSize) GetDownloadPoolSize() uint32 {
	return pcs.downloaderPoolSize
}

func (pcs *PoolCommonSize) GetAnalyzerPoolSize() uint32 {
	return pcs.analyzerPoolSize
}