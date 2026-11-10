"use client";

import React, { useState } from "react";
import { useLanguage } from "./LanguageContext";
import { Mail, Phone, MapPin, Send, CheckCircle2, AlertCircle, ShieldCheck, MessageSquare } from "lucide-react";

export default function ContactSection() {
  const { t } = useLanguage();

  const [formData, setFormData] = useState({
    name: "",
    email: "",
    projectType: "custom_software",
    budget: "discuss",
    message: "",
    honeypot: "", // Bot trap
  });

  const [status, setStatus] = useState<"idle" | "loading" | "success" | "error">("idle");
  const [errorMessage, setErrorMessage] = useState("");

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement>
  ) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    // If honeypot is filled, it's a bot. Silently ignore.
    if (formData.honeypot) return;

    setStatus("loading");
    setErrorMessage("");

    try {
      const res = await fetch("/api/contact", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          name: formData.name,
          email: formData.email,
          projectType: formData.projectType,
          budget: formData.budget,
          message: formData.message,
          honeypot: formData.honeypot,
        }),
      });

      const data = await res.json();

      if (!res.ok) {
        throw new Error(data.error || t.contact.form.errorDesc);
      }

      setStatus("success");
      setFormData({
        name: "",
        email: "",
        projectType: "custom_software",
        budget: "discuss",
        message: "",
        honeypot: "",
      });
    } catch (err: unknown) {
      setStatus("error");
      if (err instanceof Error) {
        setErrorMessage(err.message);
      } else {
        setErrorMessage(t.contact.form.errorDesc);
      }
    }
  };

  return (
    <section id="contact" className="py-24 relative z-10 border-t border-slate-800/60">
      <div className="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        {/* Section Header */}
        <div className="mb-16">
          <div className="inline-block text-xs font-mono font-semibold tracking-wider text-cyan-400 uppercase mb-2">
            {t.contact.tag}
          </div>
          <h2 className="text-3xl sm:text-4xl lg:text-5xl font-extrabold text-white tracking-tight">
            {t.contact.title}
          </h2>
          <p className="text-base sm:text-lg text-slate-400 mt-3 max-w-2xl">
            {t.contact.subtitle}
          </p>
        </div>

        <div className="grid grid-cols-1 lg:grid-cols-12 gap-12">
          {/* Left: Contact Info & Direct Channels */}
          <div className="lg:col-span-5 space-y-8">
            <div className="p-8 rounded-3xl bg-slate-900/60 border border-slate-800 backdrop-blur-xl">
              <h3 className="text-xl font-bold text-white mb-6">
                {t.contact.directInfo.title}
              </h3>

              <div className="space-y-6">
                {/* Email */}
                <a
                  href={`mailto:${t.contact.directInfo.email}`}
                  className="flex items-start gap-4 group text-slate-300 hover:text-cyan-400 transition-colors"
                >
                  <div className="p-3 rounded-2xl bg-slate-800/80 border border-slate-700/60 group-hover:border-cyan-500/50">
                    <Mail className="w-5 h-5 text-cyan-400" />
                  </div>
                  <div>
                    <span className="text-xs text-slate-500 uppercase font-mono">Email</span>
                    <p className="font-semibold text-white group-hover:text-cyan-400">
                      {t.contact.directInfo.email}
                    </p>
                  </div>
                </a>

                {/* WhatsApp / Phone */}
                <a
                  href="https://wa.me/573182267491"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="flex items-start gap-4 group text-slate-300 hover:text-emerald-400 transition-colors"
                >
                  <div className="p-3 rounded-2xl bg-slate-800/80 border border-slate-700/60 group-hover:border-emerald-500/50">
                    <Phone className="w-5 h-5 text-emerald-400" />
                  </div>
                  <div>
                    <span className="text-xs text-slate-500 uppercase font-mono">
                      WhatsApp / Móvil
                    </span>
                    <p className="font-semibold text-white group-hover:text-emerald-400">
                      {t.contact.directInfo.phone}
                    </p>
                  </div>
                </a>

                {/* Location */}
                <div className="flex items-start gap-4 text-slate-300">
                  <div className="p-3 rounded-2xl bg-slate-800/80 border border-slate-700/60">
                    <MapPin className="w-5 h-5 text-violet-400" />
                  </div>
                  <div>
                    <span className="text-xs text-slate-500 uppercase font-mono">Modalidad</span>
                    <p className="font-semibold text-white">{t.contact.directInfo.location}</p>
                  </div>
                </div>
              </div>

              {/* Direct Booking Link */}
              <div className="mt-8 pt-6 border-t border-slate-800">
                <a
                  href="https://wa.me/573182267491?text=Hola%20Jhon,%20quiero%20cotizar%20un%20proyecto%20de%20software"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="w-full inline-flex items-center justify-center gap-2 py-3 px-4 rounded-xl text-xs sm:text-sm font-semibold text-emerald-400 bg-emerald-950/40 border border-emerald-800/60 hover:bg-emerald-900/40 transition-colors"
                >
                  <MessageSquare className="w-4 h-4" />
                  <span>Chatear directamente en WhatsApp</span>
                </a>
              </div>
            </div>

            {/* Security Assurance */}
            <div className="p-5 rounded-2xl bg-slate-900/40 border border-slate-800/60 flex items-center gap-3 text-xs text-slate-400">
              <ShieldCheck className="w-5 h-5 text-cyan-400 shrink-0" />
              <span>
                Formulario protegido con encriptación, validación estricta y blindaje anti-spam.
              </span>
            </div>
          </div>

          {/* Right: Lead Capture & Quote Form */}
          <div className="lg:col-span-7">
            <div className="p-8 sm:p-10 rounded-3xl bg-slate-900/80 border border-slate-800 backdrop-blur-xl shadow-2xl relative">
              {/* Bot Honeypot field (hidden from humans) */}
              <input
                type="text"
                name="honeypot"
                value={formData.honeypot}
                onChange={handleChange}
                style={{ display: "none" }}
                tabIndex={-1}
                autoComplete="off"
              />

              {status === "success" ? (
                <div className="py-12 text-center space-y-4">
                  <div className="w-16 h-16 rounded-full bg-emerald-950 border border-emerald-500/50 flex items-center justify-center mx-auto text-emerald-400">
                    <CheckCircle2 className="w-8 h-8" />
                  </div>
                  <h3 className="text-2xl font-bold text-white">
                    {t.contact.form.successTitle}
                  </h3>
                  <p className="text-slate-300 text-sm max-w-md mx-auto">
                    {t.contact.form.successDesc}
                  </p>
                  <button
                    onClick={() => setStatus("idle")}
                    className="mt-4 px-6 py-2.5 rounded-xl text-xs font-mono font-semibold text-cyan-400 bg-cyan-950/60 border border-cyan-800/50 hover:bg-cyan-900/50 transition-colors"
                  >
                    Enviar otra consulta
                  </button>
                </div>
              ) : (
                <form onSubmit={handleSubmit} className="space-y-6">
                  {status === "error" && (
                    <div className="p-4 rounded-xl bg-red-950/60 border border-red-800/60 flex items-center gap-3 text-red-300 text-xs sm:text-sm">
                      <AlertCircle className="w-5 h-5 shrink-0" />
                      <span>{errorMessage || t.contact.form.errorDesc}</span>
                    </div>
                  )}

                  <div className="grid grid-cols-1 sm:grid-cols-2 gap-6">
                    {/* Name */}
                    <div>
                      <label className="block text-xs font-mono text-slate-300 mb-2">
                        {t.contact.form.nameLabel} *
                      </label>
                      <input
                        type="text"
                        name="name"
                        required
                        value={formData.name}
                        onChange={handleChange}
                        placeholder={t.contact.form.namePlaceholder}
                        className="w-full px-4 py-3 rounded-xl bg-slate-950/60 border border-slate-700/80 text-white placeholder:text-slate-600 focus:outline-none focus:border-cyan-400 text-sm transition-colors"
                      />
                    </div>

                    {/* Email */}
                    <div>
                      <label className="block text-xs font-mono text-slate-300 mb-2">
                        {t.contact.form.emailLabel} *
                      </label>
                      <input
                        type="email"
                        name="email"
                        required
                        value={formData.email}
                        onChange={handleChange}
                        placeholder={t.contact.form.emailPlaceholder}
                        className="w-full px-4 py-3 rounded-xl bg-slate-950/60 border border-slate-700/80 text-white placeholder:text-slate-600 focus:outline-none focus:border-cyan-400 text-sm transition-colors"
                      />
                    </div>
                  </div>

                  <div className="grid grid-cols-1 sm:grid-cols-2 gap-6">
                    {/* Project Type */}
                    <div>
                      <label className="block text-xs font-mono text-slate-300 mb-2">
                        {t.contact.form.projectTypeLabel}
                      </label>
                      <select
                        name="projectType"
                        value={formData.projectType}
                        onChange={handleChange}
                        className="w-full px-4 py-3 rounded-xl bg-slate-950 border border-slate-700/80 text-slate-200 focus:outline-none focus:border-cyan-400 text-sm transition-colors"
                      >
                        {t.contact.form.types.map((type) => (
                          <option key={type.value} value={type.value} className="bg-slate-900">
                            {type.label}
                          </option>
                        ))}
                      </select>
                    </div>

                    {/* Budget */}
                    <div>
                      <label className="block text-xs font-mono text-slate-300 mb-2">
                        {t.contact.form.budgetLabel}
                      </label>
                      <select
                        name="budget"
                        value={formData.budget}
                        onChange={handleChange}
                        className="w-full px-4 py-3 rounded-xl bg-slate-950 border border-slate-700/80 text-slate-200 focus:outline-none focus:border-cyan-400 text-sm transition-colors"
                      >
                        {t.contact.form.budgets.map((b) => (
                          <option key={b.value} value={b.value} className="bg-slate-900">
                            {b.label}
                          </option>
                        ))}
                      </select>
                    </div>
                  </div>

                  {/* Message */}
                  <div>
                    <label className="block text-xs font-mono text-slate-300 mb-2">
                      {t.contact.form.messageLabel} *
                    </label>
                    <textarea
                      name="message"
                      required
                      rows={4}
                      value={formData.message}
                      onChange={handleChange}
                      placeholder={t.contact.form.messagePlaceholder}
                      className="w-full px-4 py-3 rounded-xl bg-slate-950/60 border border-slate-700/80 text-white placeholder:text-slate-600 focus:outline-none focus:border-cyan-400 text-sm transition-colors resize-none"
                    />
                  </div>

                  {/* Submit Button */}
                  <button
                    type="submit"
                    disabled={status === "loading"}
                    className="w-full py-4 rounded-xl text-sm font-bold text-black bg-cyan-400 hover:bg-cyan-300 disabled:opacity-50 transition-all flex items-center justify-center gap-2 shadow-lg shadow-cyan-500/20"
                  >
                    {status === "loading" ? (
                      <span>{t.contact.form.submitting}</span>
                    ) : (
                      <>
                        <span>{t.contact.form.submitBtn}</span>
                        <Send className="w-4 h-4" />
                      </>
                    )}
                  </button>
                </form>
              )}
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
