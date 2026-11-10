export type Language = "es" | "en";

export const translations = {
  es: {
    nav: {
      services: "Servicios",
      experience: "Experiencia",
      architecture: "Arquitectura",
      terminal: "Terminal",
      contact: "Contacto",
      hireMe: "Cotizar Proyecto",
      status: "Disponible para proyectos & consultoría",
    },
    hero: {
      badge: "SISTEMAS DISTRIBUIDOS • GOLANG • ARQUITECTURA CLOUD",
      titleStart: "Construyo software a la medida",
      titleHighlight: "baja, media o alta concurrencia de usuarios.",
      description:
        "Ingeniero de Software Senior con más de 6 años especializándome en arquitecturas backend de alto rendimiento, microservicios en Go, pipelines orientados a eventos con Kafka y plataformas en la nube (AWS/Kubernetes). Transformo ideas de negocio complejas en sistemas sólidos, escalables y ultra rápidos.",
      ctaPrimary: "Cotizar Software a la Medida",
      ctaSecondary: "Ver Casos de Éxito",
      downloadCv: "Descargar CV",
      metrics: [
        { label: "Años de Experiencia", value: "6+" },
        { label: "Usuarios Diarios Atendidos", value: "10M+" },
        { label: "Optimización de Latencia", value: "1,000x" },
        { label: "Éxito en Concurrencia", value: "99.99%" },
      ],
    },
    services: {
      tag: "SOLUCIONES A LA MEDIDA",
      title: "¿Necesitas software empresarial robusto?",
      subtitle:
        "Evita rehacer tu plataforma dos veces. Diseño y desarrollo soluciones diseñadas para crecer desde el primer día.",
      items: [
        {
          id: "custom-dev",
          title: "Desarrollo de Software a la Medida & SaaS",
          description:
            "Creación de aplicaciones web, paneles empresariales y plataformas completas desde cero, optimizadas para conversión, seguridad y velocidad.",
          features: [
            "Arquitectura modular escalable",
            "Interfaces fluidas y reactivas (Next.js/React)",
            "Seguridad de datos de nivel bancario",
            "Automatización de procesos de negocio",
          ],
          badge: "Más Demandado",
        },
        {
          id: "backend-distributed",
          title: "Backend de Alto Rendimiento & Microservicios",
          description:
            "Diseño de APIs de ultra baja latencia en Golang, desacoplamiento de sistemas con Apache Kafka y procesamiento masivo de datos.",
          features: [
            "Arquitectura Hexagonal & DDD",
            "Sistemas orientados a eventos (EDA)",
            "Patrón Saga & Event Sourcing",
            "Concurrencia Go (Goroutines/Channels)",
          ],
          badge: "Especialidad Core",
        },
        {
          id: "db-optimization",
          title: "Optimización de Bases de Datos & Caché",
          description:
            "Solución a cuellos de botella en bases de datos relacionales, NoSQL y sistemas de caché en memoria de alto tráfico.",
          features: [
            "PostgreSQL con PostGIS (Consultas geoespaciales)",
            "Redis avanzado (Token Bucket, Geospatial)",
            "Modelado DynamoDB para latencias de un dígito",
            "Disminución drástica de costos cloud",
          ],
          badge: "Rendimiento",
        },
        {
          id: "cloud-devops",
          title: "Arquitectura Cloud & Resiliencia (AWS/K8s)",
          description:
            "Despliegues seguros, configuración de clústeres en Kubernetes, pipelines de CI/CD automatizados y blindaje contra caídas.",
          features: [
            "Contenerización Docker & Kubernetes",
            "Infraestructura como código (CDK, CloudFormation)",
            "Circuit Breakers & Exponential Backoff",
            "Monitoreo & Observabilidad (OTEL)",
          ],
          badge: "Infraestructura",
        },
      ],
    },
    experience: {
      tag: "TRAYECTORIA COMPROBADA",
      title: "Impacto en plataformas de escala masiva",
      subtitle:
        "He trabajado en el núcleo de empresas líderes optimizando sistemas críticos para millones de transacciones simultáneas.",
      roles: [
        {
          company: "Itti",
          role: "Senior Software Engineer",
          period: "Jun 2025 – Presente",
          location: "Remoto (Colombia)",
          description:
            "Arquitectura de una plataforma resiliente de scraping geoespacial en Go utilizando Arquitectura Hexagonal y patrones avanzados de concurrencia.",
          highlights: [
            "Diseñé un esquema de caché multinivel con Redis Geospatial y PostgreSQL (PostGIS / GIST), mejorando los tiempos de respuesta en 1,000x y reduciendo peticiones redundantes un 75%.",
            "Implementé rate limiting distribuido con el algoritmo Token Bucket en Redis, integrando rotación dinámica de proxies y Circuit Breakers (Backoff exponencial) con >85% de éxito.",
            "Aceleré el ciclo de desarrollo incorporando flujos asistidos por IA (PRDs/TRDs técnicos estructurados, Cursor, Gemini).",
          ],
          tech: ["Go", "Redis Geospatial", "PostgreSQL", "PostGIS", "Hexagonal Arch", "AI Workflows"],
        },
        {
          company: "Boletia México",
          role: "Senior Backend Developer",
          period: "Feb 2022 – May 2025",
          location: "Remoto (México/Colombia)",
          description:
            "Diseño y despliegue del motor de compra y reventa de boletos de alta concurrencia en Go para eventos multitudinarios.",
          highlights: [
            "Desacoplé transacciones usando Apache Kafka, eliminando la degradación de servicio durante picos agresivos de venta.",
            "Implementé gestión de transacciones distribuidas con el Patrón Saga y Event Sourcing, garantizando consistencia eventual entre órdenes, pagos e inventario.",
            "Blindé la infraestructura de Kafka en Kubernetes mediante autenticación SASL/TLS, ACLs estrictas y ajuste de compresión Snappy/Gzip.",
            "Lideré y mentoricé un squad multidisciplinario de 4 desarrolladores, estandarizando principios SOLID y CI/CD.",
          ],
          tech: ["Go", "Apache Kafka", "Kubernetes", "Saga Pattern", "Event Sourcing", "Docker", "AWS"],
        },
        {
          company: "Rappi",
          role: "Backend Developer",
          period: "Jul 2019 – Ene 2022",
          location: "Colombia",
          description:
            "Descomposición y migración del monolito heredado en Node.js hacia microservicios en Go en uno de los mayores unicornios de LatAm.",
          highlights: [
            "Diseñé el motor central de asignación de pedidos ('vasito logic') para repartidores y shoppers, optimizando rutas y algoritmos de distancia.",
            "Estabilicé la disponibilidad del servicio y el throughput para millones de usuarios activos diarios.",
            "Implementé comunicación asíncrona basada en eventos con productores y consumidores de Kafka en Go.",
          ],
          tech: ["Golang", "Node.js", "Kafka", "Microservices Migration", "AWS", "High Concurrency"],
        },
      ],
    },
    architecture: {
      tag: "INGENIERÍA PROFUNDA",
      title: "Filosofía de Arquitectura de Sistemas",
      subtitle:
        "No construyo soluciones temporales. Aplico principios probados para que tu software soporte el crecimiento.",
      pillars: [
        {
          title: "Arquitectura Hexagonal (Ports & Adapters)",
          desc: "Aislamiento total de las reglas de negocio frente a bases de datos, APIs externas o interfaces de usuario, permitiendo evolución sin fricción.",
        },
        {
          title: "Arquitectura Orientada a Eventos (EDA)",
          desc: "Desacoplamiento temporal y espacial mediante Apache Kafka y colas distribuidas para tolerar fallos sin detener el flujo transaccional.",
        },
        {
          title: "Transacciones Distribuidas (Saga Pattern)",
          desc: "Garantía de consistencia eventual a través de múltiples microservicios con compensaciones automáticas ante contingencias.",
        },
        {
          title: "Resiliencia & Rate Limiting",
          desc: "Protección proactiva contra sobrecargas mediante algoritmos Token Bucket, Circuit Breakers y Backoff Exponencial.",
        },
      ],
    },
    terminal: {
      tag: "CONSOLA DE DESARROLLADOR",
      title: "Terminal Interactiva",
      helpText: "Escribe 'help' para ver los comandos disponibles.",
      welcome: "Jhon Quirama CLI v1.0.0 (x86_64-darwin)\nEscribe 'help' o presiona uno de los atajos abajo para explorar.",
    },
    contact: {
      tag: "HABLEMOS DE TU PROYECTO",
      title: "¿Tienes un proyecto en mente?",
      subtitle:
        "Cuéntame qué necesitas construir o resolver. Te responderé en menos de 24 horas con una propuesta técnica clara.",
      form: {
        nameLabel: "Tu Nombre o Empresa",
        namePlaceholder: "Ej. Carlos Gómez / Mi Startup",
        emailLabel: "Correo Electrónico de Contacto",
        emailPlaceholder: "carlos@empresa.com",
        projectTypeLabel: "¿Qué tipo de servicio necesitas?",
        types: [
          { value: "custom_software", label: "Desarrollo de Software a la Medida / MVP" },
          { value: "backend_api", label: "Backend de Alto Rendimiento / APIs en Go" },
          { value: "cloud_arch", label: "Arquitectura Cloud, Kafka o Kubernetes" },
          { value: "consulting", label: "Consultoría Técnica / Auditoría de Sistemas" },
          { value: "full_time", label: "Oportunidad de Empleo Senior / Tech Lead" },
        ],
        budgetLabel: "Presupuesto Estimado (USD)",
        budgets: [
          { value: "under_2k", label: "< $2,000 USD (Diagnóstico / MVP básico)" },
          { value: "2k_5k", label: "$2,000 - $5,000 USD (Módulo / API robusta)" },
          { value: "5k_15k", label: "$5,000 - $15,000 USD (Plataforma completa)" },
          { value: "15k_plus", label: "+ $15,000 USD (Enterprise / Sistema distribuido)" },
          { value: "discuss", label: "Por definir en llamada" },
        ],
        messageLabel: "Cuéntame los detalles del proyecto o necesidad",
        messagePlaceholder:
          "Explica brevemente tu objetivo, desafíos técnicos actuales o plazos previstos...",
        submitBtn: "Enviar Solicitud de Proyecto",
        submitting: "Protegiendo y enviando...",
        successTitle: "¡Mensaje recibido con éxito!",
        successDesc:
          "Gracias por contactarme. He recibido los detalles en mi bandeja de entrada y te responderé en breve.",
        errorTitle: "Hubo un inconveniente",
        errorDesc:
          "Por favor intenta de nuevo o escríbeme directamente a quirama80@gmail.com o por WhatsApp.",
      },
      directInfo: {
        title: "Contacto Directo",
        email: "quirama80@gmail.com",
        phone: "+57 318 226 74 91",
        location: "Colombia (Disponible Remoto Global)",
        scheduleCall: "Agendar una llamada de 20 minutos",
      },
    },
    counter: {
      title: "Métricas del Sitio en Tiempo Real",
      totalViews: "Visualizaciones Totales",
      uniqueVisitors: "Visitantes Únicos",
      status: "Monitoreo Activo",
    },
    footer: {
      rights: "Todos los derechos reservados.",
      builtWith: "Diseñado con Next.js 15, TypeScript, Tailwind CSS, Supabase y Three.js.",
      securityNote: "Blindado con Cloudflare WAF, Row-Level Security y Zero-Trust.",
    },
  },
  en: {
    nav: {
      services: "Services",
      experience: "Experience",
      architecture: "Architecture",
      terminal: "Terminal",
      contact: "Contact",
      hireMe: "Request a Quote",
      status: "Available for projects & consulting",
    },
    hero: {
      badge: "DISTRIBUTED SYSTEMS • GOLANG • CLOUD ARCHITECTURE",
      titleStart: "I build bespoke software",
      titleHighlight: "low, medium and high user attendance.",
      description:
        "Senior Software Engineer with 6+ years specializing in high-performance backend architectures, Go microservices, event-driven pipelines with Apache Kafka, and cloud platforms (AWS/Kubernetes). I turn complex business challenges into rock-solid, ultra-fast distributed systems.",
      ctaPrimary: "Request Custom Software Quote",
      ctaSecondary: "View Case Studies",
      downloadCv: "Download CV",
      metrics: [
        { label: "Years of Experience", value: "6+" },
        { label: "Daily Active Users Handled", value: "10M+" },
        { label: "Query Latency Optimization", value: "1,000x" },
        { label: "Concurrency Success Rate", value: "99.99%" },
      ],
    },
    services: {
      tag: "BESPOKE SOLUTIONS",
      title: "Need enterprise-grade software that scales?",
      subtitle:
        "Avoid rebuilding your system twice. I engineer software designed to scale gracefully from day one.",
      items: [
        {
          id: "custom-dev",
          title: "Bespoke Software Development & SaaS",
          description:
            "End-to-end creation of web applications, custom enterprise dashboards, and MVPs built for conversion, security, and high performance.",
          features: [
            "Modular and scalable architecture",
            "Fluid, interactive UIs (Next.js/React)",
            "Bank-grade data security & encryption",
            "Business process automation",
          ],
          badge: "Most Requested",
        },
        {
          id: "backend-distributed",
          title: "High-Performance Backend & Microservices",
          description:
            "Design ultra low-latency APIs in Golang, decouple transaction systems with Kafka, and process high-volume throughput effortlessly.",
          features: [
            "Hexagonal Architecture & DDD",
            "Event-Driven Architecture (EDA)",
            "Saga Pattern & Event Sourcing",
            "Go Concurrency (Goroutines/Channels)",
          ],
          badge: "Core Specialty",
        },
        {
          id: "db-optimization",
          title: "Database Optimization & Multi-Level Caching",
          description:
            "Eliminate bottlenecks across relational databases, NoSQL, and high-traffic in-memory cache layers.",
          features: [
            "PostgreSQL with PostGIS (Geospatial queries)",
            "Advanced Redis (Token Bucket, Geospatial)",
            "DynamoDB single-digit millisecond modeling",
            "Drastic cloud infrastructure cost cuts",
          ],
          badge: "Performance",
        },
        {
          id: "cloud-devops",
          title: "Cloud Architecture & System Resilience (AWS/K8s)",
          description:
            "Hardened deployments, production Kubernetes cluster setup, automated CI/CD pipelines, and zero-downtime reliability.",
          features: [
            "Docker & Kubernetes containerization",
            "Infrastructure as Code (CDK, CloudFormation)",
            "Circuit Breakers & Exponential Backoff",
            "Observability & Distributed Tracing (OTEL)",
          ],
          badge: "Infrastructure",
        },
      ],
    },
    experience: {
      tag: "PROVEN TRACK RECORD",
      title: "Massive scale engineering impact",
      subtitle:
        "Deep technical experience architecting mission-critical platforms handling millions of concurrent users.",
      roles: [
        {
          company: "Itti",
          role: "Senior Software Engineer",
          period: "Jun 2025 – Present",
          location: "Remote (Colombia)",
          description:
            "Architected a resilient geospatial scraping platform in Go utilizing Hexagonal Architecture and advanced concurrency patterns.",
          highlights: [
            "Engineered a multi-level caching layer with Redis Geospatial and PostgreSQL (PostGIS / GIST indexes), improving query response times by 1,000x and decreasing redundant API requests by 75%.",
            "Designed distributed rate limiting using the Token Bucket algorithm in Redis, integrating dynamic proxy rotation and Circuit Breaker patterns (Exponential Backoff, Jitter) achieving 85%+ success rate.",
            "Accelerated engineering velocity by integrating AI-assisted coding workflows (structured PRDs/TRDs, Cursor, Gemini).",
          ],
          tech: ["Go", "Redis Geospatial", "PostgreSQL", "PostGIS", "Hexagonal Arch", "AI Workflows"],
        },
        {
          company: "Boletia México",
          role: "Senior Backend Developer",
          period: "Feb 2022 – May 2025",
          location: "Remote (Mexico/Colombia)",
          description:
            "Designed and deployed a high-concurrency ticket purchase and resale engine in Go, decoupling transactions with Apache Kafka.",
          highlights: [
            "Handled peak flash sales traffic without system degradation through event-driven Kafka messaging.",
            "Implemented distributed transaction management using the Saga Pattern and Event Sourcing, guaranteeing eventual consistency across order, checkout, and inventory domains.",
            "Hardened Kafka infrastructure on Kubernetes by configuring SASL/TLS authentication, role-based ACLs, and fine-tuning Snappy/Gzip compression and log retention.",
            "Mentored and led a cross-functional squad of 4 developers, standardizing code review practices, CI/CD pipelines, and SOLID architectural principles.",
          ],
          tech: ["Go", "Apache Kafka", "Kubernetes", "Saga Pattern", "Event Sourcing", "Docker", "AWS"],
        },
        {
          company: "Rappi",
          role: "Backend Developer",
          period: "Jul 2019 – Jan 2022",
          location: "Colombia",
          description:
            "Led the decomposition and migration of a legacy Node.js monolith into Golang microservices at one of LatAm's biggest unicorn tech platforms.",
          highlights: [
            "Engineered the core slot allocation engine ('vasito logic') for couriers and shoppers, optimizing route efficiency, distance algorithms, and real-time operational capacity mapping.",
            "Stabilized service availability and throughput for millions of daily active users.",
            "Implemented event-driven communication via Kafka producers and consumers in Go, streamlining asynchronous fulfillment pipelines.",
          ],
          tech: ["Golang", "Node.js", "Kafka", "Microservices Migration", "AWS", "High Concurrency"],
        },
      ],
    },
    architecture: {
      tag: "DEEP ENGINEERING",
      title: "System Architecture Philosophy",
      subtitle:
        "I do not build fragile quick fixes. I apply battle-tested engineering patterns so your product can scale effortlessly.",
      pillars: [
        {
          title: "Hexagonal Architecture (Ports & Adapters)",
          desc: "Complete decoupling of core business logic from database engines, external APIs, and UI layers, allowing clean evolution.",
        },
        {
          title: "Event-Driven Architecture (EDA)",
          desc: "Temporal and spatial decoupling through Apache Kafka to isolate failures and guarantee seamless throughput under load spikes.",
        },
        {
          title: "Distributed Transactions (Saga Pattern)",
          desc: "Ensuring eventual consistency across decoupled microservices with automated compensating actions when errors occur.",
        },
        {
          title: "Resilience & Rate Limiting",
          desc: "Proactive defense against cascading failures using Token Bucket algorithms, Circuit Breakers, and Exponential Backoff.",
        },
      ],
    },
    terminal: {
      tag: "DEVELOPER CONSOLE",
      title: "Interactive Terminal",
      helpText: "Type 'help' to see available commands.",
      welcome: "Jhon Quirama CLI v1.0.0 (x86_64-darwin)\nType 'help' or click any of the shortcut commands below to explore.",
    },
    contact: {
      tag: "LET'S BUILD TOGETHER",
      title: "Have a project in mind?",
      subtitle:
        "Tell me what you need to build or solve. I'll get back to you within 24 hours with a concrete technical plan.",
      form: {
        nameLabel: "Your Name or Company",
        namePlaceholder: "e.g. Sarah Miller / Acme Inc.",
        emailLabel: "Contact Email",
        emailPlaceholder: "sarah@acme.com",
        projectTypeLabel: "What kind of project do you need?",
        types: [
          { value: "custom_software", label: "Bespoke Software Development / MVP" },
          { value: "backend_api", label: "High-Performance Backend / Go APIs" },
          { value: "cloud_arch", label: "Cloud Architecture, Kafka or Kubernetes" },
          { value: "consulting", label: "Technical Advisory / Architecture Audit" },
          { value: "full_time", label: "Senior / Tech Lead Opportunity" },
        ],
        budgetLabel: "Estimated Budget (USD)",
        budgets: [
          { value: "under_2k", label: "< $2,000 USD (Architecture Audit / Small MVP)" },
          { value: "2k_5k", label: "$2,000 - $5,000 USD (Core API / Module)" },
          { value: "5k_15k", label: "$5,000 - $15,000 USD (Complete Platform)" },
          { value: "15k_plus", label: "+ $15,000 USD (Enterprise / Distributed System)" },
          { value: "discuss", label: "To be discussed on call" },
        ],
        messageLabel: "Project Details & Specifics",
        messagePlaceholder:
          "Briefly describe your goals, current technical roadblocks, or expected deadlines...",
        submitBtn: "Submit Project Request",
        submitting: "Securing & submitting...",
        successTitle: "Message successfully received!",
        successDesc:
          "Thank you for reaching out. I've received your request in my inbox and will get back to you shortly.",
        errorTitle: "Something went wrong",
        errorDesc:
          "Please try again or email me directly at quirama80@gmail.com or via WhatsApp.",
      },
      directInfo: {
        title: "Direct Channels",
        email: "quirama80@gmail.com",
        phone: "+57 318 226 74 91",
        location: "Colombia (Available Globally Remote)",
        scheduleCall: "Book a 20-min discovery call",
      },
    },
    counter: {
      title: "Live Visitor & Traffic Metrics",
      totalViews: "Total Page Views",
      uniqueVisitors: "Unique Visitors",
      status: "Live Tracking Active",
    },
    footer: {
      rights: "All rights reserved.",
      builtWith: "Engineered with Next.js 15, TypeScript, Tailwind CSS, Supabase, and Three.js.",
      securityNote: "Hardened with Cloudflare WAF, Row-Level Security, and Zero-Trust architecture.",
    },
  },
};
