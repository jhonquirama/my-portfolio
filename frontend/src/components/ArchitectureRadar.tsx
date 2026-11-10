"use client";

import React, { useState } from "react";
import { useLanguage } from "./LanguageContext";
import { Cpu, Layers, GitFork, ShieldAlert, CheckCircle2, ArrowRight } from "lucide-react";

export default function ArchitectureRadar() {
  const { t } = useLanguage();
  const [activeTab, setActiveTab] = useState(0);

  const pillars = [
    {
      title: t.architecture.pillars[0].title,
      desc: t.architecture.pillars[0].desc,
      icon: <Layers className="w-5 h-5 text-cyan-400" />,
      codeSnippet: `// Domain Layer: Pure Business Logic decoupled from DB
type OrderService interface {
    CreateOrder(ctx context.Context, order domain.Order) error
}

// Driven Adapter: PostgreSQL / DynamoDB implementation
type PostgresOrderRepo struct {
    db *sql.DB
}

func (r *PostgresOrderRepo) Save(ctx context.Context, o domain.Order) error {
    // Isolated persistence without contaminating business core
    return r.db.ExecContext(ctx, "INSERT INTO orders ...", o.ID)
}`,
    },
    {
      title: t.architecture.pillars[1].title,
      desc: t.architecture.pillars[1].desc,
      icon: <GitFork className="w-5 h-5 text-sky-400" />,
      codeSnippet: `// Event-Driven Producer: Kafka decoupled message streaming
producer, _ := kafka.NewProducer(&kafka.ConfigMap{
    "bootstrap.servers": "kafka-cluster:9092",
    "security.protocol": "SASL_SSL",
    "compression.type":  "snappy",
})

producer.Produce(&kafka.Message{
    TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
    Value:          orderPayloadBytes,
}, deliveryChan)`,
    },
    {
      title: t.architecture.pillars[2].title,
      desc: t.architecture.pillars[2].desc,
      icon: <Cpu className="w-5 h-5 text-violet-400" />,
      codeSnippet: `// Saga Orchestrator: Eventual Consistency Management
type SagaStep interface {
    Execute(ctx context.Context) error
    Compensate(ctx context.Context) error // Automatic rollback on failure
}

// Guarantees checkout, inventory & payment sync without 2PC blocking locks`,
    },
    {
      title: t.architecture.pillars[3].title,
      desc: t.architecture.pillars[3].desc,
      icon: <ShieldAlert className="w-5 h-5 text-emerald-400" />,
      codeSnippet: `// Redis Token Bucket Distributed Rate Limiting
func (rl *RateLimiter) Allow(ctx context.Context, clientKey string) (bool, error) {
    // Atomic Redis Lua Script execution
    // Guarantees 0 race conditions across 10,000 req/sec
    return rl.evalLua(ctx, tokenBucketScript, clientKey, capacity, refillRate)
}`,
    },
  ];

  return (
    <section id="architecture" className="py-24 relative z-10 border-t border-slate-800/60">
      <div className="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        {/* Section Header */}
        <div className="mb-16">
          <div className="inline-block text-xs font-mono font-semibold tracking-wider text-cyan-400 uppercase mb-2">
            {t.architecture.tag}
          </div>
          <h2 className="text-3xl sm:text-4xl lg:text-5xl font-extrabold text-white tracking-tight">
            {t.architecture.title}
          </h2>
          <p className="text-base sm:text-lg text-slate-400 mt-3 max-w-2xl">
            {t.architecture.subtitle}
          </p>
        </div>

        {/* Tab Selector & Code Showcase */}
        <div className="grid grid-cols-1 lg:grid-cols-12 gap-8 items-start">
          {/* Left Column: Pillars */}
          <div className="lg:col-span-5 space-y-4">
            {pillars.map((pillar, idx) => (
              <button
                key={idx}
                onClick={() => setActiveTab(idx)}
                className={`w-full text-left p-5 rounded-2xl border transition-all flex items-start gap-4 ${
                  activeTab === idx
                    ? "bg-slate-900 border-cyan-500/60 shadow-xl shadow-cyan-950/30"
                    : "bg-slate-950/40 border-slate-800/80 hover:bg-slate-900/50 hover:border-slate-700"
                }`}
              >
                <div
                  className={`p-2.5 rounded-xl border ${
                    activeTab === idx
                      ? "bg-cyan-950 border-cyan-700/60"
                      : "bg-slate-800/60 border-slate-700/40"
                  }`}
                >
                  {pillar.icon}
                </div>
                <div>
                  <h3
                    className={`text-base font-bold mb-1 ${
                      activeTab === idx ? "text-cyan-300" : "text-white"
                    }`}
                  >
                    {pillar.title}
                  </h3>
                  <p className="text-xs sm:text-sm text-slate-400 leading-relaxed font-normal">
                    {pillar.desc}
                  </p>
                </div>
              </button>
            ))}
          </div>

          {/* Right Column: Interactive Code Sandbox */}
          <div className="lg:col-span-7 rounded-3xl bg-[#090d16] border border-slate-800 p-6 shadow-2xl relative overflow-hidden">
            {/* Terminal Top bar */}
            <div className="flex items-center justify-between border-b border-slate-800/80 pb-4 mb-4">
              <div className="flex items-center space-x-2">
                <div className="w-3 h-3 rounded-full bg-red-500/80" />
                <div className="w-3 h-3 rounded-full bg-yellow-500/80" />
                <div className="w-3 h-3 rounded-full bg-emerald-500/80" />
                <span className="text-xs font-mono text-slate-400 ml-2">
                  architecture_pattern.go
                </span>
              </div>
              <span className="text-[11px] font-mono text-cyan-400 bg-cyan-950/80 border border-cyan-800/50 px-2.5 py-0.5 rounded-full">
                Hexagonal / EDA
              </span>
            </div>

            {/* Code Body */}
            <pre className="font-mono text-xs sm:text-sm text-slate-300 overflow-x-auto p-2 leading-relaxed">
              <code>{pillars[activeTab].codeSnippet}</code>
            </pre>

            <div className="mt-6 pt-4 border-t border-slate-800/80 flex items-center justify-between text-xs font-mono text-slate-400">
              <span className="flex items-center gap-1.5 text-emerald-400">
                <CheckCircle2 className="w-4 h-4" />
                Producción Verificada en Rappi & Boletia
              </span>
              <a
                href="#contact"
                className="text-cyan-400 hover:text-cyan-300 flex items-center gap-1 font-semibold"
              >
                <span>Consultar esta arquitectura</span>
                <ArrowRight className="w-3.5 h-3.5" />
              </a>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
