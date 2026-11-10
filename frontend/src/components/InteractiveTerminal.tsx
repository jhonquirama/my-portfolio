"use client";

import React, { useState, useRef, useEffect } from "react";
import { useLanguage } from "./LanguageContext";
import { Terminal as TerminalIcon, CornerDownLeft, Sparkles } from "lucide-react";

interface HistoryEntry {
  command: string;
  output: string | React.ReactNode;
}

export default function InteractiveTerminal() {
  const { t, lang } = useLanguage();
  const [inputVal, setInputVal] = useState("");
  const [history, setHistory] = useState<HistoryEntry[]>([
    {
      command: "init",
      output: t.terminal.welcome,
    },
  ]);

  const bottomRef = useRef<HTMLDivElement | null>(null);

  useEffect(() => {
    bottomRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [history]);

  const handleCommand = (cmd: string) => {
    const trimmed = cmd.trim().toLowerCase();
    if (!trimmed) return;

    if (trimmed === "clear") {
      setHistory([]);
      setInputVal("");
      return;
    }

    let result: React.ReactNode = "";

    switch (trimmed) {
      case "help":
        result = (
          <div className="space-y-1 text-slate-300">
            <p className="text-cyan-400 font-bold">Comandos disponibles:</p>
            <p>• <span className="text-cyan-300">services</span>: Ver servicios de software a la medida</p>
            <p>• <span className="text-cyan-300">experience</span>: Ver trayectoria en Rappi, Boletia e Itti</p>
            <p>• <span className="text-cyan-300">skills</span>: Lista de tecnologías y herramientas</p>
            <p>• <span className="text-cyan-300">contact</span>: Ver teléfono, email y canales directos</p>
            <p>• <span className="text-cyan-300">hire</span>: Ir al formulario de cotización</p>
            <p>• <span className="text-cyan-300">cv</span>: Descargar hoja de vida en PDF</p>
            <p>• <span className="text-cyan-300">clear</span>: Limpiar pantalla</p>
          </div>
        );
        break;

      case "services":
        result = (
          <div className="space-y-1 text-slate-300">
            <p className="text-emerald-400 font-bold">Servicios disponibles:</p>
            <p>1. Desarrollo de Software a la Medida, Paneles y MVPs escalables.</p>
            <p>2. Backend de Alto Rendimiento en Golang y microservicios.</p>
            <p>3. Mensajería desacoplada con Apache Kafka & Patrón Saga.</p>
            <p>4. Optimización de bases de datos (PostgreSQL, PostGIS, Redis).</p>
            <p>5. Despliegues en AWS y orquestación con Kubernetes.</p>
          </div>
        );
        break;

      case "skills":
        result = (
          <div className="space-y-1 text-slate-300">
            <p className="text-yellow-400 font-bold">Technical Stack:</p>
            <p><span className="text-slate-400">Languages:</span> Golang, TypeScript, JavaScript, Node.js</p>
            <p><span className="text-slate-400">Data & Messaging:</span> Apache Kafka, PostgreSQL, PostGIS, Redis Geospatial, DynamoDB, MongoDB</p>
            <p><span className="text-slate-400">Architecture:</span> Hexagonal, Event-Driven, Saga Pattern, Event Sourcing, Microservices</p>
            <p><span className="text-slate-400">Cloud & DevOps:</span> AWS (Lambda, API GW, SQS, S3, CDK), Kubernetes, Docker, CI/CD</p>
            <p><span className="text-slate-400">AI Tooling:</span> Cursor, GitHub Copilot, Gemini, AI-driven PRDs/TRDs</p>
          </div>
        );
        break;

      case "experience":
        result = (
          <div className="space-y-2 text-slate-300">
            <p className="text-cyan-400 font-bold">Historial Profesional:</p>
            <div>
              <p className="font-semibold text-white">● Itti — Senior Software Engineer (Jun 2025 - Presente)</p>
              <p className="text-slate-400 text-xs">Scraping geoespacial con Go, Redis Geo, PostGIS (1,000x más rápido).</p>
            </div>
            <div>
              <p className="font-semibold text-white">● Boletia México — Senior Backend Developer (Feb 2022 - May 2025)</p>
              <p className="text-slate-400 text-xs">Venta de boletos masiva con Go, Kafka, Patrón Saga y Kubernetes hardening.</p>
            </div>
            <div>
              <p className="font-semibold text-white">● Rappi — Backend Developer (Jul 2019 - Jan 2022)</p>
              <p className="text-slate-400 text-xs">Migración de monolito a Go, motor 'vasito logic' de asignación de pedidos.</p>
            </div>
          </div>
        );
        break;

      case "contact":
        result = (
          <div className="space-y-1 text-slate-300">
            <p className="text-cyan-400 font-bold">Canales Directos:</p>
            <p>• Email: <a href="mailto:quirama80@gmail.com" className="text-cyan-300 underline">quirama80@gmail.com</a></p>
            <p>• Celular / WhatsApp: <a href="https://wa.me/573182267491" target="_blank" className="text-emerald-400 underline">+57 318 226 74 91</a></p>
            <p>• LinkedIn: <a href="https://www.linkedin.com/in/jhon-quirama" target="_blank" className="text-sky-400 underline">linkedin.com/in/jhon-quirama</a></p>
            <p>• Ubicación: Colombia (Disponible remoto para todo el mundo)</p>
          </div>
        );
        break;

      case "hire":
        result = <p className="text-cyan-300">Navegando a la sección de cotización...</p>;
        setTimeout(() => {
          document.getElementById("contact")?.scrollIntoView({ behavior: "smooth" });
        }, 500);
        break;

      case "cv":
        result = <p className="text-emerald-300">Descargando hoja de vida...</p>;
        window.open(lang === "es" ? "/cv-jhon-quirama-es.pdf" : "/cv-jhon-quirama-en.pdf", "_blank");
        break;

      default:
        result = (
          <p className="text-red-400">
            Comando no reconocido: &apos;{trimmed}&apos;. Escribe <span className="text-cyan-300 font-bold">&apos;help&apos;</span> para ver los comandos válidos.
          </p>
        );
    }

    setHistory((prev) => [...prev, { command: trimmed, output: result }]);
    setInputVal("");
  };

  const onSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    handleCommand(inputVal);
  };

  const shortcuts = ["help", "services", "skills", "experience", "contact", "hire"];

  return (
    <section id="terminal" className="py-24 relative z-10 border-t border-slate-800/60">
      <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="mb-8">
          <div className="inline-block text-xs font-mono font-semibold tracking-wider text-cyan-400 uppercase mb-2">
            {t.terminal.tag}
          </div>
          <h2 className="text-2xl sm:text-3xl font-extrabold text-white tracking-tight flex items-center gap-2">
            <TerminalIcon className="w-6 h-6 text-cyan-400" />
            <span>{t.terminal.title}</span>
          </h2>
          <p className="text-sm text-slate-400 mt-1">{t.terminal.helpText}</p>
        </div>

        {/* Console Box */}
        <div className="rounded-2xl bg-[#080b12] border border-slate-800 shadow-2xl overflow-hidden font-mono text-xs sm:text-sm">
          {/* Header */}
          <div className="flex items-center justify-between px-4 py-3 bg-[#0d121f] border-b border-slate-800/80">
            <div className="flex items-center space-x-2">
              <div className="w-3 h-3 rounded-full bg-red-500/80" />
              <div className="w-3 h-3 rounded-full bg-yellow-500/80" />
              <div className="w-3 h-3 rounded-full bg-emerald-500/80" />
              <span className="text-slate-400 text-xs ml-2">bash ~ jhon-quirama-cli</span>
            </div>
            <div className="flex items-center gap-1.5 text-[11px] text-cyan-400">
              <Sparkles className="w-3.5 h-3.5" />
              <span>Interactivo</span>
            </div>
          </div>

          {/* Body */}
          <div className="p-4 sm:p-6 min-h-[260px] max-h-[380px] overflow-y-auto space-y-4">
            {history.map((entry, idx) => (
              <div key={idx} className="space-y-1">
                <div className="flex items-center gap-2 text-cyan-400">
                  <span className="text-slate-500">jhon@portfolio:~$</span>
                  <span className="text-white font-semibold">{entry.command}</span>
                </div>
                <div className="text-slate-300 pl-4">{entry.output}</div>
              </div>
            ))}
            <div ref={bottomRef} />
          </div>

          {/* Input Form */}
          <form
            onSubmit={onSubmit}
            className="flex items-center border-t border-slate-800/80 bg-[#0c101b] px-4 py-3"
          >
            <span className="text-cyan-400 mr-2 font-bold">jhon@portfolio:~$</span>
            <input
              type="text"
              value={inputVal}
              onChange={(e) => setInputVal(e.target.value)}
              placeholder="Escribe 'help' o presiona un atajo..."
              className="flex-1 bg-transparent border-none text-white focus:outline-none placeholder:text-slate-600 text-xs sm:text-sm"
            />
            <button
              type="submit"
              className="p-1.5 rounded-lg text-slate-400 hover:text-cyan-400 hover:bg-slate-800 transition-colors"
            >
              <CornerDownLeft className="w-4 h-4" />
            </button>
          </form>

          {/* Quick Shortcuts */}
          <div className="px-4 py-2.5 bg-[#07090e] border-t border-slate-800/40 flex flex-wrap items-center gap-1.5">
            <span className="text-[10px] text-slate-500 uppercase mr-1">Atajos:</span>
            {shortcuts.map((s) => (
              <button
                key={s}
                onClick={() => handleCommand(s)}
                className="px-2 py-0.5 rounded text-[11px] font-mono text-slate-400 hover:text-cyan-300 bg-slate-800/50 hover:bg-slate-800 border border-slate-700/50 transition-colors"
              >
                {s}
              </button>
            ))}
          </div>
        </div>
      </div>
    </section>
  );
}
