# Cosmic Excuser

**Did something questionable? Blame the universe.**

Cosmic Excuser takes whatever you did and writes you a convincing, over-the-top excuse that blames the weather, the planets and your zodiac sign. Mercury was in retrograde, it was 38°C and humid, and you're a Scorpio. None of this was your fault.

> **Status: in active development.** The React front end and Go back end are scaffolded and the weather integration is underway. See the [roadmap](#roadmap) for what works today and what's next.

## How it works

1. You describe what you did and tell the app where you are and your zodiac sign.
2. The Go back end gathers the cosmic evidence, in parallel:
   - **Weather** at your location from [Open-Meteo](https://open-meteo.com/) (free, no API key)
   - **Planet positions**, including whether Mercury is in retrograde
   - **Zodiac traits** for your sign, from a local data file
3. It builds a prompt from all of that and asks an LLM ([Claude](https://www.anthropic.com/claude)) to write the excuse.
4. The React front end displays it.

```
┌──────────────┐   POST /api/excuse    ┌───────────────────────────┐
│  React (Vite)│ ────────────────────► │        Go server          │
│              │ ◄──────────────────── │                           │
└──────────────┘     { excuse }        │  ┌─────────┐ ┌──────────┐ │
                                       │  │ Weather │ │ Planets  │ │  concurrent
                                       │  └────┬────┘ └────┬─────┘ │  requests
                                       │       │  Zodiac   │       │
                                       │       └─────┬─────┘       │
                                       │        ┌────▼────┐        │
                                       │        │  Claude │        │
                                       │        └─────────┘        │
                                       └───────────────────────────┘
```

## Tech stack

| Layer      | Choice                                                         |
| ---------- | -------------------------------------------------------------- |
| Front end  | React, Vite                                                    |
| Back end   | Go (standard library `net/http`)                               |
| Weather    | Open-Meteo forecast and geocoding APIs                         |
| Astrology  | Planet position data plus a local zodiac dataset               |
| AI         | Claude API                                                     |

## Project structure

```
cosmic-excuser/
├── server/        # Go back end (API, weather, zodiac, AI integration)
│   ├── main.go
│   └── weather.go
└── web/           # React front end (Vite)
    └── src/
        ├── App.jsx
        └── components/
```

## Running it locally

**Prerequisites:** [Go](https://go.dev/dl/) and [Node.js](https://nodejs.org/) (LTS).

Run the two halves in separate terminals.

```sh
# Back end
cd server
go run .
```

```sh
# Front end
cd web
npm install
npm run dev
```

The front end is served at http://localhost:5173.

### Configuration

Once the AI step is in, the server will read the Claude API key from an environment variable. It is never sent to or stored in the front end.

```sh
export ANTHROPIC_API_KEY=your-key-here
```

## Roadmap

- [x] React front end scaffolded with Vite
- [x] Go module and server layout
- [x] Excuse input form
- [ ] Open-Meteo weather and city geocoding
- [ ] `POST /api/excuse` endpoint
- [ ] Zodiac sign data and birthday to sign lookup
- [ ] Planet positions and Mercury retrograde detection
- [ ] Claude integration for excuse generation
- [ ] Rate limiting and input validation
- [ ] Styling and polish
- [ ] Deployment

## Design notes

- **Concurrent data fetching.** The weather and planet lookups don't depend on each other, so the server fetches them with goroutines and the request takes as long as the slowest call, not the sum.
- **Graceful degradation.** If one data source is down, the excuse is still generated from what is available.
- **API key stays on the server.** The browser only talks to the Go API, never to the AI provider.
- **Timeouts on all outgoing requests,** so a slow third-party API can't hang the server.

## License

Released under the terms of the [LICENSE](LICENSE) file.
