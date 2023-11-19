package testdata

//
//func producer(ch chan int, wg *sync.WaitGroup) {
//	defer wg.Done()
//	for i := 0; i < 5; i++ {
//		ch <- i
//	}
//	close(ch)
//}
//
//func fanIn(input1, input2 <-chan int) <-chan int {
//	ch := make(chan int)
//	var wg sync.WaitGroup
//	wg.Add(2)
//
//	go func() {
//		defer wg.Done()
//		for v := range input1 {
//			ch <- v
//		}
//	}()
//
//	go func() {
//		defer wg.Done()
//		for v := range input2 {
//			ch <- v
//		}
//	}()
//
//	go func() {
//		wg.Wait()
//		close(ch)
//	}()
//
//	return ch
//}
//
//func TestFanIn(t *testing.T) {
//
//	ch1 := make(chan int)
//	ch2 := make(chan int)
//
//	var wg sync.WaitGroup
//	wg.Add(1)
//
//	go producer(ch1, &wg)
//	go producer(ch2, &wg)
//
//	result := fanIn(ch1, ch2)
//	//<-result
//	for val := range result {
//		fmt.Println(val)
//	}
//	//for {
//	//	select {
//	//	case data := <-result:
//	//		fmt.Println(data)
//	//	}
//	//}
//}
