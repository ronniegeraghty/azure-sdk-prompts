/* ============================================================
   VARIATION B — Lifecycle scenes (v2)
   - Completed blocks stay fully expanded
   - Collapsed one-liners appear ONLY at the end as the summary
   - Adds: larger-queue scene (2 workers, 5-eval queue) + verbose
   ============================================================ */

/* ---------- Scene 1: startup / plan ---------- */
function TerminalB_Plan() {
  return (
    <div className="term">
      <div>
        <span className="c-dimmer">$</span>{" "}
        <span className="c-fg">hyoka run</span>{" "}
        <span className="c-dim">--prompt-id key-vault-dp-python-crud --config python-pairwise --workers 3</span>
      </div>
      <div>{"\n"}</div>

      <HeaderB label="plan" />
      <div>
        {"  "}<Kv k="prompts" v="1" />{"   "}
        <Kv k="configs" v="3" />{"   "}
        <Kv k="evals" v="3" color="b" />{"   "}
        <Kv k="workers" v="3" color="c-yellow" />
      </div>
      <div>
        {"  "}<span className="c-dimmer">prompt  </span><span className="c-fg">key-vault-dp-python-crud</span>{"  "}
        <span className="c-dimmer">python · CRUD · azure-keyvault</span>
      </div>
      <div>{"\n"}</div>

      <HeaderB label="queue" />
      <div>{"  "}<span className="c-dimmer">·</span>{"  "}<span className="c-fg">gpt-5.3-codex</span>{"       "}<span className="c-dimmer">queued</span></div>
      <div>{"  "}<span className="c-dimmer">·</span>{"  "}<span className="c-fg">claude-opus-4.6</span>{"     "}<span className="c-dimmer">queued</span></div>
      <div>{"  "}<span className="c-dimmer">·</span>{"  "}<span className="c-fg">claude-sonnet-4.5</span>{"   "}<span className="c-dimmer">queued</span></div>
      <div>{"\n"}</div>

      <div className="c-dim">starting 3 workers…</div>
    </div>
  );
}

/* ---------- Scene 2: mid-flight — completed stays expanded ---------- */
function TerminalB_Mid() {
  return (
    <div className="term">
      <HeaderB label="running" trailing="1 done · 2 active · 0 queued" />
      <div>{"\n"}</div>

      <FullBlockB
        label="claude-opus-4.6"
        worker={1}
        phase="pass"
        gen={{ elapsed: 118.5, turns: 4, calls: 4, tokensIn: 179339, tokensOut: 1460, files: "kv_secrets_crud.py" }}
        graders={[
          { name: "output_files",    state: "pass", passed: 2, total: 2 },
          { name: "prompt_criteria", state: "pass", passed: 8, total: 8 },
        ]}
      />
      <div>{"\n"}</div>

      <FullBlockB
        label="gpt-5.3-codex"
        worker={2}
        phase="grading"
        gen={{ elapsed: 72.3, turns: 3, calls: 3, tokensIn: 102653, tokensOut: 1886, files: null }}
        graders={[
          { name: "output_files",    state: "fail",    passed: 0, total: 2 },
          { name: "prompt_criteria", state: "running" },
        ]}
      />
      <div>{"\n"}</div>

      <FullBlockB
        label="claude-sonnet-4.5"
        worker={3}
        phase="generating"
        gen={{ elapsed: 42.1, turns: 2, calls: 2, tokensIn: 88410, tokensOut: 910 }}
        graders={[
          { name: "output_files",    state: "queued" },
          { name: "prompt_criteria", state: "queued" },
        ]}
      />
      <div>{"\n"}</div>

      <ProgressFooterB done={1} running={2} total={3} elapsed="2:10" eta="~1:20" workers="3" />
    </div>
  );
}

/* ---------- Scene 3: 2 workers, LONGER queue — 5 evals remaining ---------- */
function TerminalB_Queue() {
  return (
    <div className="term">
      <div>
        <span className="c-dimmer">$</span>{" "}
        <span className="c-fg">hyoka run</span>{" "}
        <span className="c-dim">--prompt-set azure-python --workers 2</span>
      </div>
      <div>{"\n"}</div>

      <HeaderB label="running" trailing="2 done · 2 active · 3 queued" />
      <div>{"\n"}</div>

      {/* Completed — still fully expanded */}
      <FullBlockB
        label="key-vault-crud · claude-opus-4.6"
        worker={1}
        phase="pass"
        gen={{ elapsed: 118.5, turns: 4, calls: 4, tokensIn: 179339, tokensOut: 1460, files: "kv_secrets_crud.py" }}
        graders={[
          { name: "output_files",    state: "pass", passed: 2, total: 2 },
          { name: "prompt_criteria", state: "pass", passed: 8, total: 8 },
        ]}
      />
      <div>{"\n"}</div>

      <FullBlockB
        label="key-vault-crud · gpt-5.3-codex"
        worker={2}
        phase="fail"
        gen={{ elapsed: 72.3, turns: 3, calls: 3, tokensIn: 102653, tokensOut: 1886, files: null }}
        graders={[
          { name: "output_files",    state: "fail", passed: 0, total: 2 },
          { name: "prompt_criteria", state: "pass", passed: 8, total: 8 },
        ]}
        failure={{
          name: "output_files",
          rules: [
            { name: "min_files",          reason: "need ≥ 1 produced file, found 0" },
            { name: "min_bytes_per_file", reason: "no produced files to check (≥ 1 required)" },
          ],
          hint: "agent completed 3 turns but never called `create`",
        }}
      />
      <div>{"\n"}</div>

      {/* Active */}
      <FullBlockB
        label="key-vault-crud · claude-sonnet-4.5"
        worker={1}
        phase="grading"
        gen={{ elapsed: 95.4, turns: 3, calls: 3, tokensIn: 135783, tokensOut: 3876, files: "README.md, keyvault_crud.py, requirements.txt" }}
        graders={[
          { name: "output_files",    state: "pass",    passed: 2, total: 2 },
          { name: "prompt_criteria", state: "running" },
        ]}
      />
      <div>{"\n"}</div>

      <FullBlockB
        label="blob-upload · claude-opus-4.6"
        worker={2}
        phase="generating"
        gen={{ elapsed: 28.0, turns: 1, calls: 1, tokensIn: 42110, tokensOut: 320 }}
        graders={[
          { name: "output_files",    state: "queued" },
          { name: "prompt_criteria", state: "queued" },
        ]}
      />
      <div>{"\n"}</div>

      {/* Queue tail */}
      <HeaderB label="queued" trailing="3 remaining" />
      <div>{"  "}<span className="c-dimmer">·</span>{"  "}<span className="c-fg">blob-upload</span>{"             "}<span className="c-dim">gpt-5.3-codex</span>{"       "}<span className="c-dimmer">queued</span></div>
      <div>{"  "}<span className="c-dimmer">·</span>{"  "}<span className="c-fg">blob-upload</span>{"             "}<span className="c-dim">claude-sonnet-4.5</span>{"   "}<span className="c-dimmer">queued</span></div>
      <div>{"  "}<span className="c-dimmer">·</span>{"  "}<span className="c-fg">cosmos-query</span>{"            "}<span className="c-dim">claude-opus-4.6</span>{"     "}<span className="c-dimmer">queued</span></div>
      <div>{"\n"}</div>

      <ProgressFooterB done={2} running={2} total={7} elapsed="4:22" eta="~6:10" workers="2" />
    </div>
  );
}

/* ---------- Scene 4: final — collapsed summary at the very end ---------- */
function TerminalB_Final() {
  return (
    <div className="term">
      <div className="c-dim">… (expanded blocks above, from the live run) …</div>
      <div>{"\n"}</div>

      <HeaderB label="summary" trailing="run 20260424-034027" />
      <div>{"\n"}</div>

      <CollapsedBlock label="claude-opus-4.6"   phase="pass" gen={{ elapsed: 118.5, turns: 4, tokensIn: 179339, files: "kv_secrets_crud.py" }}                passed={10} total={10} />
      <CollapsedBlock label="claude-sonnet-4.5" phase="pass" gen={{ elapsed:  95.4, turns: 3, tokensIn: 135783, files: "README.md, keyvault_crud.py, …" }} passed={12} total={12} />
      <CollapsedBlock label="gpt-5.3-codex"     phase="fail" gen={{ elapsed:  72.3, turns: 3, tokensIn: 102653, files: null }}                              passed={8}  total={10} />

      <div>{"\n"}</div>
      <div>
        {"  "}
        <span className="c-green">{"█".repeat(24)}</span>
        <span className="c-red">{"█".repeat(12)}</span>
        {"   "}
        <span className="c-green">2</span><span className="c-dim">/3 passed</span>{"   "}
        <span className="c-dimmer">·</span>{"   "}
        <span className="c-dim">elapsed</span> <span className="c-fg">4:46</span>
      </div>
      <div>{"\n"}</div>

      <HeaderB label="next" />
      <div>
        {"  "}<span className="c-dim">view  </span>{"  "}<span className="c-cyan">hyoka view 20260424-034027</span>{"   "}<span className="c-dimmer">↗ open report in browser</span>
      </div>
      <div>
        {"  "}<span className="c-dim">trend </span>{"  "}<span className="c-cyan">hyoka trend --last 10</span>{"         "}<span className="c-dimmer">AI analysis across recent runs</span>
      </div>
      <div>
        {"  "}<span className="c-dim">retry </span>{"  "}<span className="c-cyan">hyoka run --rerun-failed 20260424-034027</span>{"  "}<span className="c-dimmer">rerun only the failed eval</span>
      </div>
      <div>{"\n"}</div>
      <div className="c-dimmer">  report  reports/20260424-034027/</div>
    </div>
  );
}

/* ---------- Scene 5: verbose mode ---------- */
function VerboseToolsB() {
  return (
    <div>
      <div className="c-dim">     tools</div>
      <div>{"       "}<span className="c-dimmer">├─</span>{" "}<span className="c-fg">azure</span>{"                 "}<span className="c-dimmer">mcp</span>{"           "}<span className="c-green">✓ loaded</span></div>
      <div>{"       "}<span className="c-dimmer">├─</span>{" "}<span className="c-fg">generator-skills</span>{"      "}<span className="c-dimmer">skills dir</span>{"    "}<span className="c-green">1/1 skills</span></div>
      <div>{"       "}<span className="c-dimmer">│  └─</span>{" "}<span className="c-dim">azure-sdk-for-rust-bestpractices</span></div>
      <div>{"       "}<span className="c-dimmer">└─</span>{" "}<span className="c-fg">azure-sdk-python</span>{"      "}<span className="c-dimmer">plugin</span>{"        "}<span className="c-green">40/40 skills</span>{"  "}<span className="c-dimmer">·</span>{"  "}<span className="c-green">1/1 mcp</span></div>
      <div>{"       "}{"   "}<span className="c-dimmer">├─ azure-keyvault-py              ├─ azure-identity-py</span></div>
      <div>{"       "}{"   "}<span className="c-dimmer">├─ azure-storage-blob-py          ├─ azure-cosmos-py</span></div>
      <div>{"       "}{"   "}<span className="c-dimmer">├─ azure-servicebus-py            ├─ azure-eventhub-py</span></div>
      <div>{"       "}{"   "}<span className="c-dimmer">├─ fastapi-router-py              ├─ pydantic-models-py</span></div>
      <div>{"       "}{"   "}<span className="c-dimmer">└─ … 32 more</span>{"  "}<span className="c-dim">(hyoka show tools --expand)</span></div>
    </div>
  );
}

function TerminalB_Verbose() {
  const spin = useSpinner();
  return (
    <div className="term">
      <div>
        <span className="c-dimmer">$</span>{" "}
        <span className="c-fg">hyoka run</span>{" "}
        <span className="c-dim">--prompt-id key-vault-dp-python-crud --config python-pairwise/claude-opus-4.6 --verbose</span>
      </div>
      <div>{"\n"}</div>

      <HeaderB label="eval 2/3" trailing="claude-opus-4.6 · key-vault-dp-python-crud" />
      <div>{"\n"}</div>

      <VerboseToolsB />
      <div>{"\n"}</div>

      <div>
        <span className="c-yellow">{spin}</span>{"  "}<span className="b">generator</span>
        <span className="c-dimmer">   │  </span><span className="c-yellow">generating</span>
        <span className="c-dimmer">   ·   </span><Kv k="t" v="1:58" />
        <span className="c-dimmer">   ·   </span><Kv k="session" v="a7f3b2" color="c-dim" />
      </div>
      <div className="c-dimmer">
        {"     "}<Kv k="turns" v="3" />{"  "}<Kv k="calls" v="3" />{"  "}<Kv k="tok" v="179.3k↓ 1,460↑" />{"  "}<Kv k="ctx" v="52%" />
      </div>
      <div>{"\n"}</div>

      <div>
        {"     "}<span className="c-dimmer">▸</span>{" "}<span className="c-dim">turn 1  </span>{" "}
        <span className="c-cyan">skill       </span>{" "}<span className="c-fg">azure-keyvault-py</span>{"              "}<span className="c-green">ok</span>{"   "}<span className="c-dimmer">6.2s</span>
      </div>
      <div>
        {"     "}<span className="c-dimmer">▸</span>{" "}<span className="c-dim">turn 2  </span>{" "}
        <span className="c-cyan">create      </span>{" "}<span className="c-fg">kv_secrets_crud.py</span>{"             "}<span className="c-green">ok</span>{"   "}<span className="c-dimmer">2.1 KB</span>
      </div>
      <div>
        {"     "}<span className="c-dimmer">▸</span>{" "}<span className="c-dim">turn 3  </span>{" "}
        <span className="c-cyan">bash        </span>{" "}<span className="c-dim">python -m py_compile kv_secrets_crud.py</span>{"  "}<span className="c-green">ok</span>
      </div>
      <div>
        {"     "}<span className="c-yellow">{spin}</span>{" "}<span className="c-dim">turn 4  </span>{" "}
        <span className="c-yellow">thinking…</span>
      </div>
      <div>{"\n"}</div>

      <div>
        <span className="c-dimmer">·</span>{"  "}<span className="b">graders</span>
        <span className="c-dimmer">   │  </span><span className="c-dimmer">waiting for generator</span>
      </div>
      <div>
        {"     "}<span className="c-dimmer">·</span>{"  "}<span className="c-dimmer">output_files       queued   </span><span className="c-dimmer">checks min_files · min_bytes_per_file</span>
      </div>
      <div>
        {"     "}<span className="c-dimmer">·</span>{"  "}<span className="c-dimmer">prompt_criteria    queued   </span><span className="c-dimmer">6 criteria from prompt file</span>
      </div>
    </div>
  );
}

Object.assign(window, { TerminalB_Plan, TerminalB_Mid, TerminalB_Queue, TerminalB_Final, TerminalB_Verbose });
