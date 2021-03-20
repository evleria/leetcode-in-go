package _396_design_underground_system

type Entry struct {
	Station string
	Time    int
}

type UndergroundSystem struct {
	Stats   map[[2]string][2]int // key: [startStation, endStation], value: [timeSpent, trips]
	Enroute map[int]Entry        // key: id, value: {startStation, startTime}
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		Stats:   map[[2]string][2]int{},
		Enroute: map[int]Entry{},
	}
}

func (us *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	us.Enroute[id] = Entry{stationName, t}
}

func (us *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	entry := us.Enroute[id]
	delete(us.Enroute, id)
	key := [2]string{entry.Station, stationName}
	stats := us.Stats[key]
	us.Stats[key] = [2]int{stats[0] + t - entry.Time, stats[1] + 1}
}

func (us *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	if stats := us.Stats[[2]string{startStation, endStation}]; stats[1] > 0 {
		return float64(stats[0]) / float64(stats[1])
	}
	return 0
}
