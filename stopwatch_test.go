package gut

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type StopwatchTestSuite struct {
	suite.Suite
	Stopwatch Stopwatch
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (s *StopwatchTestSuite) SetupTest() {
	s.Stopwatch = Stopwatch{}
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (s *StopwatchTestSuite) TestInitialState() {
	s.False(s.Stopwatch.IsRunning(), "Stopwatch should not be running.")
	s.Zero(s.Stopwatch.ElapsedTime(), "Elapsed time should be zero.")
}

func (s *StopwatchTestSuite) TestStart() {
	s.Stopwatch.Start()
	s.True(s.Stopwatch.IsRunning(), "Stopwatch should be running.")
	time.Sleep(time.Nanosecond)
	s.True(s.Stopwatch.ElapsedTime() > 0, "Elapsed time should be greater than zero.")

	time.Sleep(time.Microsecond)
	alreadyElapsedTime := s.Stopwatch.ElapsedTime()
	s.Stopwatch.Start()
	s.True(s.Stopwatch.IsRunning(), "Start() should be idempotent: stopwatch should be running.")
	time.Sleep(time.Nanosecond)
	s.True(s.Stopwatch.ElapsedTime() > alreadyElapsedTime, "Start() should be idempotent: elapsed time should be increasing.")
}

func (s *StopwatchTestSuite) TestRestart() {
	s.Stopwatch.Restart()
	s.True(s.Stopwatch.IsRunning(), "Stopwatch should be running.")
	time.Sleep(time.Nanosecond)
	s.True(s.Stopwatch.ElapsedTime() > 0, "Elapsed time should be greater than zero.")
}

func (s *StopwatchTestSuite) TestStartThenRestart() {
	s.Stopwatch.Start()
	time.Sleep(time.Microsecond)
	alreadyElapsedTime := s.Stopwatch.ElapsedTime()
	s.Stopwatch.Restart()
	s.True(s.Stopwatch.IsRunning(), "Stopwatch should be running.")
	time.Sleep(time.Nanosecond)
	s.True(s.Stopwatch.ElapsedTime() < alreadyElapsedTime, "Elapsed time should be reset.")
}

func (s *StopwatchTestSuite) TestStop() {
	s.Stopwatch.Stop()
	s.False(s.Stopwatch.IsRunning(), "Stopwatch should not be running.")
	time.Sleep(time.Nanosecond)
	s.Zero(s.Stopwatch.ElapsedTime(), "Elapsed time should be zero.")
}

func (s *StopwatchTestSuite) TestStartThenStop() {
	s.Stopwatch.Start()
	time.Sleep(time.Nanosecond)
	s.Stopwatch.Stop()
	alreadyElapsedTime := s.Stopwatch.ElapsedTime()
	s.False(s.Stopwatch.IsRunning(), "Stopwatch should not be running.")
	time.Sleep(time.Nanosecond)
	s.Equal(alreadyElapsedTime, s.Stopwatch.ElapsedTime(), "Elapsed time should be unchanged.")

	s.Stopwatch.Stop()
	s.False(s.Stopwatch.IsRunning(), "Stop() should be idempotent: stopwatch should not be running.")
	time.Sleep(time.Nanosecond)
	s.Equal(alreadyElapsedTime, s.Stopwatch.ElapsedTime(), "Stop() should be idempotent: elapsed time should be unchanged.")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestStopwatchTestSuite(t *testing.T) {
	suite.Run(t, new(StopwatchTestSuite))
}
