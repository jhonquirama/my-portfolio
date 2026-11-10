import type { Metadata } from "next";
import "./globals.css";

const siteUrl =
  process.env.NEXT_PUBLIC_SITE_URL && !process.env.NEXT_PUBLIC_SITE_URL.includes("localhost")
    ? process.env.NEXT_PUBLIC_SITE_URL.replace(/\/$/, "")
    : "https://my-portfolio-jhonqs-projects.vercel.app";

export const metadata: Metadata = {
  metadataBase: new URL(siteUrl),
  title: {
    default: "Jhon Quirama | Senior Backend Engineer & Software a la Medida",
    template: "%s | Jhon Quirama",
  },
  description:
    "Senior Backend Engineer con 6+ años de experiencia (Rappi, Boletia, Itti). Experto en Golang, Apache Kafka, AWS, Kubernetes y arquitectura de sistemas distribuidos. Desarrollo de software a la medida de alta disponibilidad.",
  keywords: [
    "Jhon Quirama",
    "software a la medida",
    "desarrollo de software a la medida",
    "senior backend engineer",
    "golang developer",
    "arquitectura hexagonal",
    "apache kafka",
    "microservicios go",
    "kubernetes aws",
    "consultoria de software",
    "freelance software engineer",
    "custom software development",
  ],
  authors: [{ name: "Jhon Quirama", url: siteUrl }],
  creator: "Jhon Quirama",
  publisher: "Jhon Quirama",
  robots: {
    index: true,
    follow: true,
    googleBot: {
      index: true,
      follow: true,
      "max-video-preview": -1,
      "max-image-preview": "large",
      "max-snippet": -1,
    },
  },
  openGraph: {
    type: "website",
    locale: "es_ES",
    alternateLocale: "en_US",
    url: siteUrl,
    title: "Jhon Quirama | Senior Backend Engineer & Software a la Medida",
    description:
      "Construyo sistemas distribuidos y software a la medida de alto rendimiento que no se caen bajo millones de usuarios. Experiencia en Rappi, Boletia e Itti.",
    siteName: "Jhon Quirama Portfolio",
  },
  twitter: {
    card: "summary_large_image",
    title: "Jhon Quirama | Senior Backend Engineer & Software a la Medida",
    description:
      "Senior Backend Engineer con 6+ años en Golang, Apache Kafka, AWS y Kubernetes. Software a la medida y consultoría técnica.",
    creator: "@jhonquirama",
  },
  alternates: {
    canonical: siteUrl,
    languages: {
      "es-ES": siteUrl,
      "en-US": siteUrl,
    },
  },
  verification: {
    google: "v4txeyUSamF8xnXlHtM6wrsYeB1ql5GUmzc2tPHD7w0",
  },
};

const jsonLdPerson = {
  "@context": "https://schema.org",
  "@type": "Person",
  name: "Jhon Quirama",
  jobTitle: "Senior Backend Engineer & Software Architect",
  url: siteUrl,
  email: "quirama80@gmail.com",
  telephone: "+573182267491",
  sameAs: [
    "https://www.linkedin.com/in/jhon-quirama",
    "https://github.com/jhonquirama",
  ],
  knowsAbout: [
    "Golang",
    "Apache Kafka",
    "Microservices",
    "Hexagonal Architecture",
    "Distributed Systems",
    "Amazon Web Services (AWS)",
    "Kubernetes",
    "PostgreSQL",
    "PostGIS",
    "Redis",
    "Next.js",
    "TypeScript",
    "Custom Software Engineering",
  ],
  worksFor: {
    "@type": "Organization",
    name: "Independent Consulting & Engineering",
  },
};

const jsonLdService = {
  "@context": "https://schema.org",
  "@type": "ProfessionalService",
  name: "Jhon Quirama - Servicios de Software a la Medida & Consultoría Backend",
  image: `${siteUrl}/og-image.png`,
  url: siteUrl,
  telephone: "+573182267491",
  priceRange: "$$",
  address: {
    "@type": "PostalAddress",
    addressCountry: "CO",
  },
  description:
    "Desarrollo de software a la medida, arquitectura de sistemas backend de alto tráfico, microservicios en Golang, colas con Kafka y modernización cloud.",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="es" className="dark">
      <head>
        <meta
          name="google-site-verification"
          content="v4txeyUSamF8xnXlHtM6wrsYeB1ql5GUmzc2tPHD7w0"
        />
        <script
          type="application/ld+json"
          dangerouslySetInnerHTML={{ __html: JSON.stringify(jsonLdPerson) }}
        />
        <script
          type="application/ld+json"
          dangerouslySetInnerHTML={{ __html: JSON.stringify(jsonLdService) }}
        />
      </head>
      <body className="bg-[#07090e] text-slate-100 min-h-screen selection:bg-cyan-400 selection:text-black">
        {children}
      </body>
    </html>
  );
}
