package helpers

import "sync"

func Parallelize(functions ...func()) {
    var waitGroup sync.WaitGroup
    waitGroup.Add(len(functions))

	ch := make(chan struct{}, len(functions))

    for _, function := range functions {
		ch <- struct{}{}
		go func(copyFunc func()) {
			defer func() {
				<-ch
				waitGroup.Done()
			}()
			copyFunc()
		}(function)
	}
	waitGroup.Wait()
}
