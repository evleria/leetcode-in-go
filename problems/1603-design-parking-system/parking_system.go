package _603_design_parking_system

type ParkingSystem struct {
	left [3]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{[3]int{big, medium, small}}
}

func (ps *ParkingSystem) AddCar(carType int) bool {
	if ps.left[carType-1] > 0 {
		ps.left[carType-1]--
		return true
	}
	return false
}
