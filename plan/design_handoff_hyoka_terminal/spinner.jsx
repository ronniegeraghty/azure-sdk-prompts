// shared spinner hook
const SPIN_FRAMES = ["⠋","⠙","⠹","⠸","⠼","⠴","⠦","⠧","⠇","⠏"];

function useSpinner(intervalMs = 90) {
  const [i, setI] = React.useState(0);
  React.useEffect(() => {
    const id = setInterval(() => setI(x => (x + 1) % SPIN_FRAMES.length), intervalMs);
    return () => clearInterval(id);
  }, [intervalMs]);
  return SPIN_FRAMES[i];
}

// Fake "elapsed time" that ticks up, for verisimilitude
function useElapsed(running, startSeconds = 0) {
  const [t, setT] = React.useState(startSeconds);
  React.useEffect(() => {
    if (!running) return;
    const id = setInterval(() => setT(x => x + 0.1), 100);
    return () => clearInterval(id);
  }, [running]);
  return t;
}

function fmtElapsed(s) {
  if (s < 60) return `${s.toFixed(1)}s`;
  const m = Math.floor(s / 60);
  const sec = Math.floor(s % 60);
  return `${m}:${String(sec).padStart(2, "0")}`;
}

Object.assign(window, { SPIN_FRAMES, useSpinner, useElapsed, fmtElapsed });
