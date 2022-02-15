package service

import "time"

func repeat(cb func() error, interval time.Duration, repeat int) {
	for i := 0; i < repeat; i++ {
		err := cb()

		if err != nil {
			time.Sleep(interval)
			continue
		}

		return
	}
}
