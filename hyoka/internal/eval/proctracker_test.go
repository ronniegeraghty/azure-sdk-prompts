package eval

import (
	"os/exec"
	"sync"
	"testing"
	"time"
)

func TestProcessTracker_RegisterDeregister(t *testing.T) {
	pt := &ProcessTracker{}

	// Start a real process to get a valid PID
	cmd := exec.Command("sleep", "30")
	if err := cmd.Start(); err != nil {
		t.Fatalf("failed to start test process: %v", err)
	}
	defer cmd.Process.Kill()

	pid := cmd.Process.Pid
	pt.Register(pid)

	pt.mu.Lock()
	if _, ok := pt.procs[pid]; !ok {
		t.Errorf("expected pid %d to be registered", pid)
	}
	pt.mu.Unlock()

	pt.Deregister(pid)

	pt.mu.Lock()
	if _, ok := pt.procs[pid]; ok {
		t.Errorf("expected pid %d to be deregistered", pid)
	}
	pt.mu.Unlock()
}

func TestProcessTracker_RegisterInitializesMap(t *testing.T) {
	pt := &ProcessTracker{}
	if pt.procs != nil {
		t.Fatal("expected procs map to be nil initially")
	}

	cmd := exec.Command("sleep", "30")
	if err := cmd.Start(); err != nil {
		t.Fatalf("failed to start test process: %v", err)
	}
	defer cmd.Process.Kill()

	pt.Register(cmd.Process.Pid)
	if pt.procs == nil {
		t.Fatal("expected procs map to be initialized after Register")
	}
}

func TestProcessTracker_DeregisterNonexistent(t *testing.T) {
	pt := &ProcessTracker{}
	// Should not panic on nil map
	pt.Deregister(99999)
}

func TestProcessTracker_TerminateAll_Empty(t *testing.T) {
	pt := &ProcessTracker{}
	errs := pt.TerminateAll(1 * time.Second)
	if len(errs) != 0 {
		t.Errorf("expected 0 errors for empty tracker, got %d", len(errs))
	}
}

func TestProcessTracker_TerminateAll_GracefulExit(t *testing.T) {
	pt := &ProcessTracker{}

	// Start a process that responds to SIGTERM
	cmd := exec.Command("sleep", "60")
	if err := cmd.Start(); err != nil {
		t.Fatalf("failed to start test process: %v", err)
	}

	pt.Register(cmd.Process.Pid)

	// TerminateAll sends SIGTERM, waits up to timeout, then SIGKILL stragglers.
	// The process should be killed within the timeout+SIGKILL cycle.
	errs := pt.TerminateAll(500 * time.Millisecond)

	// Verify the tracker is cleared
	pt.mu.Lock()
	if pt.procs != nil {
		t.Errorf("expected procs to be nil after TerminateAll, got %v", pt.procs)
	}
	pt.mu.Unlock()

	// Errors from Kill are acceptable (process may already be gone)
	_ = errs
	cmd.Wait()
}

func TestProcessTracker_TerminateAll_MultiplProcesses(t *testing.T) {
	pt := &ProcessTracker{}

	cmds := make([]*exec.Cmd, 3)
	for i := 0; i < 3; i++ {
		cmds[i] = exec.Command("sleep", "60")
		if err := cmds[i].Start(); err != nil {
			t.Fatalf("failed to start process %d: %v", i, err)
		}
		pt.Register(cmds[i].Process.Pid)
	}

	errs := pt.TerminateAll(500 * time.Millisecond)
	if len(errs) != 0 {
		t.Errorf("expected 0 errors, got %d: %v", len(errs), errs)
	}

	// Wait for all processes to exit
	for _, cmd := range cmds {
		cmd.Wait()
	}
}

func TestProcessTracker_TerminateAll_AlreadyExitedProcess(t *testing.T) {
	pt := &ProcessTracker{}

	// Start and immediately stop a process
	cmd := exec.Command("true")
	if err := cmd.Start(); err != nil {
		t.Fatalf("failed to start test process: %v", err)
	}
	cmd.Wait() // Wait for it to exit

	pt.Register(cmd.Process.Pid)

	// Should not error — already-exited processes are removed during SIGTERM
	errs := pt.TerminateAll(1 * time.Second)
	if len(errs) != 0 {
		t.Errorf("expected 0 errors for already-exited process, got %d: %v", len(errs), errs)
	}
}

func TestProcessTracker_ConcurrentAccess(t *testing.T) {
	pt := &ProcessTracker{}
	var wg sync.WaitGroup

	// Concurrent Register/Deregister should not race
	cmds := make([]*exec.Cmd, 10)
	for i := 0; i < 10; i++ {
		cmds[i] = exec.Command("sleep", "30")
		if err := cmds[i].Start(); err != nil {
			t.Fatalf("failed to start process %d: %v", i, err)
		}
	}
	defer func() {
		for _, cmd := range cmds {
			cmd.Process.Kill()
			cmd.Wait()
		}
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(pid int) {
			defer wg.Done()
			pt.Register(pid)
		}(cmds[i].Process.Pid)
	}
	wg.Wait()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(pid int) {
			defer wg.Done()
			pt.Deregister(pid)
		}(cmds[i].Process.Pid)
	}
	wg.Wait()
}

func TestProcessTracker_TerminateAll_ClearsMap(t *testing.T) {
	pt := &ProcessTracker{}

	cmd := exec.Command("sleep", "60")
	if err := cmd.Start(); err != nil {
		t.Fatalf("failed to start test process: %v", err)
	}

	pt.Register(cmd.Process.Pid)
	pt.TerminateAll(500 * time.Millisecond)

	// Map should be nil after TerminateAll
	pt.mu.Lock()
	defer pt.mu.Unlock()
	if pt.procs != nil {
		t.Error("expected procs map to be nil after TerminateAll")
	}

	cmd.Wait()
}
