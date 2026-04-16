import {
  Zap, BarChart3, GitCompare, Languages, Shield, Layers,
  ArrowRight, CheckCircle2, Code2, Bot, ClipboardCheck
} from "lucide-react";
import { Link } from "react-router";
import { motion } from "motion/react";

const services = ["Identity", "Storage", "Key Vault", "Databases", "Messaging", "Compute"];
const languages = ["Go", "Python", ".NET", "Java", "JavaScript", "TypeScript"];

const features = [
  { icon: GitCompare, title: "Side-by-Side Comparison", desc: "Compare multiple AI models and configurations to find the best approach for each evaluation scenario." },
  { icon: BarChart3, title: "Deep Metrics", desc: "Track pass rates, durations, token usage, tool calls, and scoring across every evaluation run." },
  { icon: Shield, title: "Multi-Grader Pipeline", desc: "Combine multiple graders — LLM reviewers, program checks, file validators — to assess agent output comprehensively." },
  { icon: Languages, title: "Polyglot Support", desc: "Evaluate agent output across Go, Python, .NET, Java, JavaScript, TypeScript, and more." },
  { icon: Layers, title: "Flexible Grading", desc: "Mix and match grader types: AI reviews, build verification, behavior analysis, custom program checks." },
  { icon: Zap, title: "Difficulty Levels", desc: "Prompts range from basic to advanced, covering authentication, CRUD operations, error handling, and complex workflows." },
];

const steps = [
  { icon: Code2, title: "Define Prompts", desc: "Write natural language prompts describing the tasks you want AI agents to perform." },
  { icon: Bot, title: "Run Evaluations", desc: "AI agents execute tasks across multiple configurations and models simultaneously." },
  { icon: ClipboardCheck, title: "Review Results", desc: "Get detailed grader reports, session analysis, and insights across all evaluation runs." },
];

export function HomePage() {
  return (
    <div className="bg-[#0a0a0f]" style={{ fontFamily: "'Inter', sans-serif" }}>
      {/* Hero */}
      <section className="relative overflow-hidden px-6 py-24 md:py-36">
        {/* Gradient bg */}
        <div className="absolute inset-0 bg-[radial-gradient(ellipse_at_top,_rgba(16,185,129,0.12)_0%,_transparent_60%)]" />
        <div className="relative mx-auto max-w-4xl text-center">
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.6 }}
          >
            <div className="mb-6 inline-flex items-center gap-2 rounded-full border border-emerald-500/20 bg-emerald-500/10 px-4 py-1.5">
              <Zap className="h-3.5 w-3.5 text-emerald-400" />
              <span className="text-emerald-400" style={{ fontSize: 13 }}>Developer Evaluation Tool</span>
            </div>

            <h1 className="mb-6 text-white" style={{ fontFamily: "'JetBrains Mono', monospace", fontSize: "clamp(2rem, 5vw, 3.5rem)", lineHeight: 1.1 }}>
              Evaluate AI Agent
              <br />
              <span className="text-emerald-400">Output Quality</span>
            </h1>

            <p className="mx-auto mb-10 max-w-2xl text-white/50" style={{ fontSize: 17, lineHeight: 1.7 }}>
              hyoka evaluates how well AI agents perform tasks from natural language prompts —
              across multiple models, configurations, and evaluation criteria.
            </p>

            <div className="flex flex-col items-center justify-center gap-4 sm:flex-row">
              <Link
                to="/dashboard"
                className="inline-flex items-center gap-2 rounded-xl bg-emerald-500 px-6 py-3 text-black no-underline transition-colors hover:bg-emerald-400"
                style={{ fontWeight: 600, fontSize: 15 }}
              >
                View Dashboard <ArrowRight className="h-4 w-4" />
              </Link>
              <Link
                to="/how-it-works"
                className="inline-flex items-center gap-2 rounded-xl border border-white/15 px-6 py-3 text-white/70 no-underline transition-colors hover:bg-white/5"
                style={{ fontSize: 15 }}
              >
                How It Works
              </Link>
            </div>
          </motion.div>
        </div>
      </section>

      {/* Ticker */}
      <section className="border-y border-white/10 bg-white/[0.02] px-6 py-8">
        <div className="mx-auto flex max-w-5xl flex-wrap items-center justify-center gap-3">
          {services.map((s) => (
            <span key={s} className="rounded-lg border border-white/10 bg-white/5 px-3 py-1.5 text-white/50" style={{ fontSize: 13 }}>
              {s}
            </span>
          ))}
          <span className="mx-2 text-white/20">·</span>
          {languages.map((l) => (
            <span key={l} className="rounded-lg border border-emerald-500/15 bg-emerald-500/5 px-3 py-1.5 text-emerald-400/70" style={{ fontSize: 13 }}>
              {l}
            </span>
          ))}
        </div>
      </section>

      {/* Features */}
      <section className="px-6 py-24">
        <div className="mx-auto max-w-6xl">
          <div className="mb-16 text-center">
            <h2 className="mb-3 text-white" style={{ fontFamily: "'JetBrains Mono', monospace" }}>
              Why hyoka?
            </h2>
            <p className="text-white/40" style={{ fontSize: 15 }}>
              Everything you need to evaluate AI agent performance.
            </p>
          </div>

          <div className="grid gap-5 md:grid-cols-2 lg:grid-cols-3">
            {features.map((f, i) => (
              <motion.div
                key={f.title}
                initial={{ opacity: 0, y: 20 }}
                whileInView={{ opacity: 1, y: 0 }}
                viewport={{ once: true }}
                transition={{ delay: i * 0.08 }}
                className="group rounded-2xl border border-white/8 bg-white/[0.03] p-6 transition-colors hover:border-emerald-500/20 hover:bg-emerald-500/[0.03]"
              >
                <div className="mb-4 flex h-10 w-10 items-center justify-center rounded-xl bg-emerald-500/10">
                  <f.icon className="h-5 w-5 text-emerald-400" />
                </div>
                <h3 className="mb-2 text-white" style={{ fontSize: 16 }}>{f.title}</h3>
                <p className="text-white/40" style={{ fontSize: 14, lineHeight: 1.6 }}>{f.desc}</p>
              </motion.div>
            ))}
          </div>
        </div>
      </section>

      {/* How it works preview */}
      <section className="border-t border-white/10 bg-white/[0.01] px-6 py-24">
        <div className="mx-auto max-w-4xl">
          <div className="mb-16 text-center">
            <h2 className="mb-3 text-white" style={{ fontFamily: "'JetBrains Mono', monospace" }}>
              Three Simple Steps
            </h2>
          </div>
          <div className="grid gap-8 md:grid-cols-3">
            {steps.map((s, i) => (
              <div key={s.title} className="text-center">
                <div className="mx-auto mb-4 flex h-14 w-14 items-center justify-center rounded-2xl border border-emerald-500/20 bg-emerald-500/10">
                  <s.icon className="h-6 w-6 text-emerald-400" />
                </div>
                <div className="mb-2 text-emerald-500/60" style={{ fontFamily: "'JetBrains Mono', monospace", fontSize: 12 }}>
                  0{i + 1}
                </div>
                <h3 className="mb-2 text-white" style={{ fontSize: 16 }}>{s.title}</h3>
                <p className="text-white/40" style={{ fontSize: 14, lineHeight: 1.6 }}>{s.desc}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* CTA */}
      <section className="px-6 py-24">
        <div className="mx-auto max-w-3xl rounded-2xl border border-emerald-500/20 bg-gradient-to-br from-emerald-500/10 to-transparent p-12 text-center">
          <h2 className="mb-3 text-white" style={{ fontFamily: "'JetBrains Mono', monospace" }}>
            Ready to evaluate?
          </h2>
          <p className="mb-8 text-white/50" style={{ fontSize: 15 }}>
            Start measuring AI agent performance across your evaluation prompts.
          </p>
          <Link
            to="/dashboard"
            className="inline-flex items-center gap-2 rounded-xl bg-emerald-500 px-6 py-3 text-black no-underline transition-colors hover:bg-emerald-400"
            style={{ fontWeight: 600, fontSize: 15 }}
          >
            Explore Dashboard <ArrowRight className="h-4 w-4" />
          </Link>
        </div>
      </section>
    </div>
  );
}
