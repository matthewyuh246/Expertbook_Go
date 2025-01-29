package main

// import "sync"

// type Group struct {
// 	cancel func()
// 	wg     sync.WaitGroup

// 	errOnce sync.Once
// 	err     error
// }

// func (g *Group) Go(f func() error) {
// 	g.wg.Add(1)

// 	go func() {
// 		defer g.wg.Done()

// 		if err := f(); err != nil {
// 			g.errOnce.Do(func() {
// 				g.err = err
// 				if g.cancel != nil {
// 					g.cancel()
// 				}
// 			})
// 		}
// 	}()
// }
