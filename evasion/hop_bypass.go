package luvevasion

import (
	luvutils "LuvFramework/utils"
	"fmt"
	"time"
)

var i int64 = 0

//The HopBypass Evasion function bypasses anti-virus software by allocating fake memory and hopping between multiple memory locations using function calls.
func HopBypass() ([]byte, error) {
	var mem []byte
	var err error

	mem, err = luvutils.AllocMem(100000000)
	if err != nil {
		return nil, err
	}
	keepMem(mem, 10)
	hop()
	return mem, nil
}

func keepMem(mem []byte, duration int) {
	//Make a semaphore channel to keep the memory alive
	sem := make(chan bool, 1)

	go func() {
		for i := 0; i < duration; i++ {
			time.Sleep(time.Second * 1)
			//Keep accessing the memory to keep it alive
			mem = append(mem, []byte("\x00")...)
			//Send a message to the semaphore channel
		}
		sem <- true
	}()
	//Wait for the semaphore channel to be released
	signal := <-sem
	fmt.Println("Signal:", signal)
}

//Hop Function

func hop() {
	i++
	if i == 100000000 {
		i = 0
	}
	hop2()
}

func hop2() {
	i++
	if i == 100000000 {
		i = 0
	}
	hop3()
}

func hop3() {
	i++
	if i == 100000000 {
		i = 0
	}
	hop4()
}

func hop4() {
	i++
	if i == 100000000 {
		i = 0
	}
	hop5()
}

func hop5() {
	i++
	if i == 100000000 {
		i = 0
	}
	hop6()
}

func hop6() {
	i++
	if i == 100000000 {
		i = 0
	}
	hop7()
}

func hop7() {
	i++
	if i == 100000000 {
		i = 0
	}
	hop8()
}

func hop8() {
	i++
	if i == 100000000 {
		i = 0
	}
	hop9()
}

func hop9() {
	i++
	if i == 100000000 {
		i = 0
	}

}
