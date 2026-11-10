# Jhon Quirama — Portfolio & Custom Software Engineering

> **Senior Backend Engineer & Distributed Systems Architect**  
> *Golang • Apache Kafka • AWS & Kubernetes • Hexagonal Architecture • Bespoke Software Development*

Este repositorio contiene el código fuente completo del portafolio profesional y plataforma de captación de clientes de software a la medida de **Jhon Quirama**.

---

## 🚀 Arquitectura del Proyecto

El proyecto está diseñado bajo una arquitectura híbrida y modular:

```
my-portfolio/
├── frontend/                        # Aplicación Web Next.js 14/15 (App Router)
│   ├── src/
│   │   ├── app/                     # Rutas, Layout, SEO (Sitemap & Robots), APIs
│   │   │   ├── api/contact/         # API de contacto blindada (Rate limiting, Zod, Resend, Supabase)
│   │   │   ├── api/counter/         # API de métricas y contador de visitas en tiempo real
│   │   │   ├── sitemap.ts           # Generador dinámico de Sitemap XML para Google
│   │   │   └── robots.ts            # Reglas de indexación SEO
│   │   ├── components/              # Componentes React con Tailwind y Framer Motion
│   │   │   ├── BackgroundCanvas.tsx # Fondo 3D interactivo inspirado en Richard Mattka
│   │   │   ├── Navbar.tsx           # Navegación flotante con selector de idioma (ES/EN)
│   │   │   ├── Hero.tsx             # Propuesta de valor dual (Clientes + Tech Leads)
│   │   │   ├── ServicesSection.tsx  # Catálogo de servicios de software a la medida
│   │   │   ├── ExperienceSection.tsx# Casos de éxito (Rappi, Boletia, Itti)
│   │   │   ├── ArchitectureRadar.tsx# Showcase de arquitectura hexagonal y patrones de concurrencia
│   │   │   ├── InteractiveTerminal.tsx# Consola CLI interactiva para visitantes
│   │   │   ├── ContactSection.tsx   # Cotizador y formulario con honeypot anti-spam
│   │   │   ├── LiveCounter.tsx      # Contador de visitas en vivo
│   │   │   └── LanguageContext.tsx  # Motor de internacionalización bilingüe (ES/EN)
│   │   ├── content/
│   │   │   └── translations.ts      # Diccionario completo de textos y datos profesionales
│   │   ├── lib/
│   │   │   ├── supabase.ts          # Cliente y administración Supabase con RLS
│   │   │   ├── rate-limiter.ts      # Limitador de peticiones anti-DDoS / anti-spam
│   │   │   └── utils.ts             # Utilidades de estilo y formateo
│   │   └── public/
│   │       ├── cv-jhon-quirama-es.pdf # Hoja de vida en español
│   │       └── cv-jhon-quirama-en.pdf # Resume en inglés
│   └── supabase/
│       └── schema.sql               # Script SQL listo para ejecutar en Supabase (RLS y RPC)
│
├── cmd/ & internal/ & pkg/          # Microservicio backend en Go (Arquitectura Hexagonal)
└── README.md
```

---

## ⚡ Inicio Rápido en Local

### 1. Requisitos Previos
- **Node.js** (v18+)
- **npm** (v9+)

### 2. Ejecutar la Aplicación Web
```bash
cd frontend
npm install
npm run dev
```
Abre tu navegador en `http://localhost:3000`.

---

## 🛠️ Configuración de Servicios ($0 / mes)

### 1. Base de Datos en Supabase
1. Ingresa a [supabase.com](https://supabase.com) y crea un nuevo proyecto gratuito.
2. Ve al **SQL Editor** y pega el contenido del archivo `frontend/supabase/schema.sql`.
3. Haz clic en **Run**. Esto creará automáticamente:
   - La tabla `contact_leads` protegida con **Row Level Security (RLS)** (nadie externo puede leer tus leads).
   - Las tablas `site_analytics` y `visitor_daily_logs`.
   - La función atómica `record_page_visit` para el contador de visitas diario único.
4. En **Project Settings -> API**, copia tu `URL` y tu `anon public key`.

### 2. Notificaciones por Correo con Resend
1. Regístrate gratis en [resend.com](https://resend.com) (incluye 3,000 correos gratis al mes).
2. Genera una API Key en **API Keys -> Create API Key**.
3. Pega la clave en tu archivo `.env.local` en `RESEND_API_KEY`.

### 3. Variables de Entorno (`frontend/.env.local`)
Copia `.env.example` a `.env.local` y completa tus credenciales:
```env
NEXT_PUBLIC_SITE_URL=https://tu-dominio.com
NEXT_PUBLIC_SUPABASE_URL=https://tu-id.supabase.co
NEXT_PUBLIC_SUPABASE_ANON_KEY=tu-anon-key
SUPABASE_SERVICE_ROLE_KEY=tu-service-role-key
RESEND_API_KEY=re_tu_api_key
CONTACT_RECEIVER_EMAIL=quirama80@gmail.com
```

---

## 🌐 Despliegue en Producción en 2 Minutos (Vercel — $0 / mes)

1. Sube tus cambios a GitHub:
   ```bash
   git add .
   git commit -m "feat: complete modern portfolio & bespoke software studio"
   git push origin main
   ```
2. Entra a [vercel.com](https://vercel.com) e inicia sesión con tu cuenta de GitHub.
3. Haz clic en **"Add New Project"** y selecciona este repositorio.
4. En **Root Directory**, selecciona la carpeta `frontend`.
5. Agrega las variables de entorno (`NEXT_PUBLIC_SUPABASE_URL`, `RESEND_API_KEY`, etc.).
6. Haz clic en **Deploy**. ¡Tu sitio estará en vivo en menos de 60 segundos con HTTPS automático, CDN global y protección DDoS!

---

## 🔒 Blindaje de Seguridad Implementado

- **Cabeceras HTTP de Alta Seguridad:** HSTS, X-Frame-Options: DENY, X-Content-Type-Options, Referrer-Policy y Permissions-Policy configuradas en `next.config.mjs`.
- **Protección Anti-Spam / Bots:** Campo invisibilizado *Honeypot* en el formulario de contacto para desviar bots automáticamente.
- **Rate Limiting:** Limitador de tasa por IP en memoria para prevenir saturación o abusos en `/api/contact` y `/api/counter`.
- **Sanitización Estricta:** Validación con Zod para impedir ataques de Cross-Site Scripting (XSS) e inyecciones de datos maliciosos.
- **Row-Level Security (RLS):** La clave anónima del navegador únicamente tiene permiso de escritura condicional de nuevos leads; ninguna persona externa puede consultar información de tus clientes.

---

## 📈 Buenas Prácticas de SEO Implementadas

- **Renderizado Estático / SSG:** Código HTML servido al instante para los bots de indexación de Google con tiempo de carga LCP < 1.2s.
- **Datos Estructurados Schema.org:** Schemas JSON-LD `Person` y `ProfessionalService` integrados en el `<head>`.
- **Sitemap & Robots Automáticos:** Generados en `/sitemap.xml` y `/robots.txt`.
- **OpenGraph & Twitter Cards:** Metadatos configurados para previsualizaciones impecables al compartir tu URL en LinkedIn, Twitter o WhatsApp.
- **Soporte Multilingüe:** Alternancia instantánea entre Español e Inglés con detección automática y persistencia.
