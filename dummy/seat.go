package dummy

import "san_dong/model"

var Seat map[int]bool = map[int]bool{}

func InitSeat() {
	for i := 0; i < 20; i++ {
		Seat[i] = false // mantap array
	}
}

func CheckSeatStatus(seat *model.Seat) bool {
	return Seat[seat.Pos]
}
