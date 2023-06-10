// proof of concept: broadcast message to many goroutines
package main

import (
	"fmt"
	"time"
)

type chaincast struct {
	v     interface{}      //message payload.
	out   chan interface{} //message output channel
	spawn chan<- chaincast //send to this channel to spawn goroutine for chain element.
	next  chan chaincast   //channel for passing message to next
	prev  chan chaincast   //channel for receiving message from prev
}

func initchain() (chan interface{}, chaincast) {
	var head, tail chaincast
	in := make(chan interface{})
	head.next = make(chan chaincast)
	tail.prev = head.next
	spawnchan := make(chan chaincast)
	head.spawn = spawnchan
	tail.spawn = spawnchan

	go func() {
		for {
			fmt.Println("initchain(): input")
			head.next <- chaincast{<-in, nil, nil, nil, nil}
			fmt.Println("initchain(): input ok")
		}
	}()
	//	go cast(tail)
	go func(c <-chan chaincast) {
		fmt.Println("initchain(): func() spawner start.")
		for {
			go cast(<-c)
			fmt.Println("initchain(): func() spawner started new.")
		}
		fmt.Println("initchain(): func() spawner end")
	}(spawnchan)
	spawnchan <- tail
	return in, head
}
func addListener(head chaincast) chan interface{} {
	out := make(chan interface{})
	c := make(chan chaincast)
	n := chaincast{nil, out, head.spawn, c, nil}
	head.next <- n
	return out
}
func cast(self chaincast) {
	//	fmt.Println("cast called: ", self);
	for {
		e := <-(self.prev)
		//		fmt.Println("cast: received something.")
		if e.spawn == nil { //received a message
			//			fmt.Println("cast: received a message.")
			if self.out != nil && e.out == nil { //send message to output channel, ignoring head and tail
				//				self.out <- e.v
				select {
				case self.out <- e.v:
				default: //remove from chain and exit
					fmt.Println("cast(): about to exit.")
					self.next <- e
					self.next <- chaincast{nil, nil, self.spawn, nil, self.prev}
					fmt.Println("cast(): exit.")
					return
				}
			}
		} else if self.next == nil && e.prev == nil && e.next != nil { //insert into chain before tail node.
			e.prev = self.prev
			self.prev = e.next
			//			fmt.Println("tail: adding node: ", self)
			//			go cast(e)
			self.spawn <- e
		} else if e.out == nil && e.prev != nil && e.next == nil { //"remove" preceding elem
			fmt.Println("cast(): about to remove.")
			self.prev = e.prev
			continue
		}
		if self.next != nil {
			//fmt.Println("cast(): passing to next.: ", self.next)
			self.next <- e
		}
	}
	// fmt.Println("cast(): exit.")
}
func listener(desc string, o <-chan interface{}) {
	for i := 0; true; i++ {
		if i == 3 && desc == "routine3" {
			break
		}
		fmt.Println(desc+": ", <-o)
	}
}
func broadcaster(in chan<- interface{}) {
	for i := 0; i < 6; i++ {
		time.Sleep(time.Millisecond * 30)
		in <- "message " + string(int('1')+i)
	}
}
func main() {
	in, head := initchain()
	go listener("routine1", addListener(head))
	go listener("routine2", addListener(head))
	go listener("routine3", addListener(head))
	go listener("routine4", addListener(head))
	go listener("routine5", addListener(head))
	time.Sleep(time.Second / 10)
	go broadcaster(in)

	time.Sleep(time.Second * 3)
}
