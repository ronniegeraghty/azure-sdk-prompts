/* ============================================================
   VARIATION B (developed) — ANCHORED
   Reusable primitives + four artboards covering the full
   lifecycle: plan → running → failure detail → final.
   ============================================================ */

/* ---------- primitives ---------- */

function HeaderB({ label, trailing }) {
  const rule = "─".repeat(Math.max(10, 60 - label.length - (trailing ? String(trailing).length + 4 : 0)));
  return (
    <div>
      <span className="c-dimmer">──</span>{" "}
      <span className="b">{label}</span>
      {trailing && <> <span className="c-dimmer">·</span> <span className="c-dim">{trailing}</span></>}
      {"  "}
      <span className="c-dimmer">{rule}</span>
    </div>
  );
}

function Kv({ k, v, color = "c-fg" }) {
  return (
    <span>
      <span className="c-dimmer">{k}</span>{" "}
      <span className={color}>{v}</span>
    </span>
  );
}

function PhaseSym({ phase, spin }) {
  if (phase === "pass")       return <span className="c-green">✓</span>;
  if (phase === "fail")       return <span className="c-red">✗</span>;
  if (phase === "generating") return <span className="c-yellow">{spin}</span>;
  if (phase === "grading")    return <span className="c-yellow">{spin}</span>;
  return <span className="c-dimmer">·</span>;
}

function PhaseWord({ phase }) {
  if (phase === "generating") return <span className="c-yellow">generating</span>;
  if (phase === "grading")    return <span className="c-yellow">grading</span>;
  if (phase === "pass")       return <span className="c-green">pass</span>;
  if (phase === "fail")       return <span className="c-red">fail</span>;
  return <span className="c-dimmer">queued</span>;
}

/* ---------- collapsed (history) block: one line ---------- */
function CollapsedBlock({ label, phase, gen, passed, total }) {
  const col = phase === "pass" ? "c-green" : "c-red";
  return (
    <div>
      <PhaseSym phase={phase} />{"  "}
      <span className="c-fg">{label.padEnd(20)}</span>{" "}
      <span className={col}>{phase}</span>{"   "}
      <span className="c-dimmer">{passed}/{total} criteria</span>{"   "}
      <span className="c-dimmer">{fmtElapsed(gen.elapsed)}</span>{"   "}
      <span className="c-dimmer">{gen.turns}t · {(gen.tokensIn/1000).toFixed(0)}k tok</span>
      {gen.files && <>{"   "}<span className="c-cyan">{gen.files}</span></>}
    </div>
  );
}

/* ---------- grader row ---------- */
function GraderRowB({ name, state, passed, total }) {
  const spin = useSpinner();
  const sym =
    state === "pass"    ? <span className="c-green">✓</span> :
    state === "fail"    ? <span className="c-red">✗</span> :
    state === "running" ? <span className="c-yellow">{spin}</span> :
                          <span className="c-dimmer">·</span>;

  const bar = total > 0 ? (
    <>{"["}
      <span className={state === "fail" ? "c-red" : "c-green"}>{"▰".repeat(passed || 0)}</span>
      <span className="c-dimmer">{"▱".repeat(Math.max(0, total - (passed || 0)))}</span>
    {"]"}</>
  ) : null;

  const trailing =
    state === "queued"  ? <span className="c-dimmer">queued</span> :
    state === "running" ? <span className="c-yellow">grading…</span> :
                          <span className={state === "fail" ? "c-red" : "c-dimmer"}>{(passed ?? 0)}/{total}</span>;

  return (
    <div>
      {"     "}{sym}{"  "}
      <span className={state === "queued" ? "c-dimmer" : "c-fg"}>{name.padEnd(18)}</span>{" "}
      {bar}{"  "}
      {trailing}
    </div>
  );
}

/* ---------- full block ---------- */
function FullBlockB({ label, worker, phase, gen, graders, failure }) {
  const spin = useSpinner();
  const live = useElapsed(phase === "generating" || phase === "grading", gen?.elapsed ?? 0);
  const t = phase === "pass" || phase === "fail" ? gen.elapsed : live;

  return (
    <div>
      <div>
        <PhaseSym phase={phase} spin={spin} />{"  "}
        <span className="b">{label}</span>
        {worker && <><span className="c-dimmer">   ·  </span><span className="c-dim">w{worker}</span></>}
        <span className="c-dimmer">   │  </span><PhaseWord phase={phase} />
        {phase !== "queued" && <><span className="c-dimmer">   ·   </span><Kv k="t" v={fmtElapsed(t)} /></>}
        {gen?.files && <><span className="c-dimmer">   ·   </span><Kv k="files" v={gen.files} color="c-cyan" /></>}
      </div>

      {phase !== "queued" && (
        <div className="c-dimmer">
          {"     "}
          <Kv k="turns" v={gen.turns} />{"  "}
          <Kv k="calls" v={gen.calls} />{"  "}
          <Kv k="tok" v={`${(gen.tokensIn/1000).toFixed(1)}k↓ ${gen.tokensOut.toLocaleString()}↑`} />{"  "}
          <Kv k="tools" v={<>azure<span className="c-dimmer">(mcp 1/1)</span>, azure-sdk-python<span className="c-dimmer">(40/40)</span></>} />
        </div>
      )}

      {graders && graders.map((g, i) => <GraderRowB key={i} {...g} />)}

      {failure && (
        <div className="fail-block">
          <span className="b c-red">↳ {failure.name}</span>{"  "}<span className="c-dimmer">{failure.rules.length} rules failed</span>
          {failure.rules.map((r, i) => (
            <div key={i}>
              {"  "}<span className="c-red">✗</span>{" "}<span className="c-fg">{r.name.padEnd(22)}</span>{" "}<span className="c-dim">{r.reason}</span>
            </div>
          ))}
          {failure.hint && <div className="c-dimmer">  hint: {failure.hint}</div>}
        </div>
      )}
    </div>
  );
}

/* ---------- progress footer ---------- */
function ProgressFooterB({ done, running, total, elapsed, eta, workers }) {
  const width = 36;
  const passW = Math.round((done / total) * width);
  const runW  = Math.round((running / total) * width);
  const restW = Math.max(0, width - passW - runW);
  return (
    <div>
      {"  "}
      <span className="c-green">{"█".repeat(passW)}</span>
      <span className="c-yellow">{"█".repeat(runW)}</span>
      <span className="c-dimmer">{"░".repeat(restW)}</span>{"  "}
      <span className="c-fg">{done}</span><span className="c-dim">/{total}</span>{"   "}
      <span className="c-dimmer">·</span>{"   "}
      <Kv k="workers" v={workers} color="c-yellow" />{"   "}
      <Kv k="elapsed" v={elapsed} />{"   "}
      <Kv k="eta" v={eta} />
    </div>
  );
}

Object.assign(window, {
  HeaderB, Kv, PhaseSym, PhaseWord,
  CollapsedBlock, GraderRowB, FullBlockB, ProgressFooterB,
});
