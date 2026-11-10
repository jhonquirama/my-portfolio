import { LanguageProvider } from "@/components/LanguageContext";
import BackgroundCanvas from "@/components/BackgroundCanvas";
import Navbar from "@/components/Navbar";
import Hero from "@/components/Hero";
import LiveCounter from "@/components/LiveCounter";
import ServicesSection from "@/components/ServicesSection";
import ExperienceSection from "@/components/ExperienceSection";
import ArchitectureRadar from "@/components/ArchitectureRadar";
import InteractiveTerminal from "@/components/InteractiveTerminal";
import ContactSection from "@/components/ContactSection";
import Footer from "@/components/Footer";

export default function Home() {
  return (
    <LanguageProvider>
      <div className="relative min-h-screen bg-[#07090e] text-slate-100 overflow-x-hidden">
        {/* Richard Mattka-inspired dynamic 3D Background */}
        <BackgroundCanvas />

        {/* Floating Top Navigation */}
        <Navbar />

        {/* Main Content Area */}
        <main className="relative z-10 flex flex-col">
          {/* Hero Value Proposition */}
          <Hero />

          {/* Real-time Visitor Counter */}
          <LiveCounter />

          {/* Bespoke Software Services */}
          <ServicesSection />

          {/* Work Experience & Impact */}
          <ExperienceSection />

          {/* Deep Systems Architecture Showcase */}
          <ArchitectureRadar />

          {/* Developer Interactive Console */}
          <InteractiveTerminal />

          {/* Lead Capture & Project Estimator Form */}
          <ContactSection />
        </main>

        {/* Hardened Footer */}
        <Footer />
      </div>
    </LanguageProvider>
  );
}
