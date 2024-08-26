package gut

import "time"

type Stopwatch struct {
	start, stop time.Time
}

func (s Stopwatch) ElapsedTime() time.Duration {
	if s.IsRunning() {
		return time.Since(s.start)
	}
	return s.stop.Sub(s.start)
}

func (s Stopwatch) IsRunning() bool {
	return s.start.After(s.stop)
}

func (s *Stopwatch) Start() {
	if s.IsRunning() {
		return
	}
	s.Restart()
}

func (s *Stopwatch) Restart() {
	s.start = time.Now()
}

func (s *Stopwatch) Stop() {
	if s.IsRunning() {
		s.stop = time.Now()
	}
}
