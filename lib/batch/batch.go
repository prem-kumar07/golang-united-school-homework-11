package batch

import (
	"context"
	//"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	g, _ := errgroup.WithContext(context.Background())
	res = make([]user,n)
    g.SetLimit(int(pool))
	for i := 0; i < int(n); i++ {
		i:=i
		g.Go(func() error {
				res[i] = getOne(int64(i))
			    return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil
	}
	//fmt.Printf("res %v",res)
	return res

}
