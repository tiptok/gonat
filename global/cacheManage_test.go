package global

import "testing"

func TestVehicleCache(t *testing.T) {
	//var TimerManage *comm.TimerWork
	cacheVeh := &MSVehiclesCacheLoader{}
	cacheVeh.Load(nil)
	// cacheVeh.NewCache("VehicleInfoCache", 60, MSVehiclesCacheLoader{}.Load)
	// cacheVeh.TimerTask.Work(nil)
	// TimerManage.RegistTask(cacheVeh.TimerTask)
	// TimerManage.Start()

	// time.Sleep(time.Second * 30)
	// fmt.Println("Cache Size:", len(VehiclesCache.CacheValue.DataStore))
	// //测B061171  测B061201 测B061231
	// c1 := cacheVeh.GetCache("测B061171")
	// if c1 == nil {
	// 	fmt.Println("Not Exists")
	// } else {
	// 	fmt.Println("取到缓存车辆信息:", c1)
	// }
}
