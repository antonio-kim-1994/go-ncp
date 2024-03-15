package service

//type SchedulingRequestData struct {
//	JobType     string `json:"jobType"`     // product, contract
//	RequestType string `json:"requestType"` // pause, resume
//	TimeCycle   string `json:"timeCycle"`   // minute
//}
//
//var (
//	c       *cron.Cron
//	entryID cron.EntryID
//)
//
//func (d SchedulingRequestData) startCronJob() {
//	var ()
//
//	if c == nil {
//		c = cron.New()
//	}
//	if entryID != 0 {
//		c.Remove(entryID)
//	}
//
//	switch d.JobType {
//	case "contract":
//		entryID, err = c.AddFunc(d.TimeCycle, checkContractDifference)
//	case "product":
//	}
//
//}
//
//func checkContractDifference() {
//	var (
//		now = time.Now().Format("2006-01-02")
//		db = repository.GetDB()
//		err error
//	)
//
//	dbDate
//	result, err := MergeContractData(now, )
//}
