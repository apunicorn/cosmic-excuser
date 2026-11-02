# Cosmic Excuser

**Did something questionable? Blame the universe!!**

Cosmic Excuser takes whatever you did and writes you a convincing excuse that blames the weather, the planets and your zodiac sign. Mercury was in retrograde, it was 38°C and humid, and you're a Scorpio. None of this was your fault. Feel free to use this excuse for work, your spouse, or anywhere else. 

## How it works

1. You describe what you did and pick your zodiac sign. The browser also shares your location.
2. Go back end:
   - **Your zodiac sign**, checked against a fixed list on the server 
   - **Weather** at your location from [Open-Meteo](https://open-meteo.com/), free with no API key. It is fetched by the server and becomes part of the AI prompt
   - **Planet positions**, including whether Mercury is in retrograde/gatorade/microwave - a classic
3. It sends all of that to an LLM ([Google Gemini](https://ai.google.dev/)) with instructions to write a short, dry excuse 
4. React front end.


## Running it locally

**Prerequisites:** [Go](https://go.dev/dl/) and [Node.js](https://nodejs.org/) (LTS).

Run the server and web parts separately:

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

The server reads a Gemini API key from an environment variable. You can get a free key from [Google AI Studio](https://aistudio.google.com/). The key stays on the server and is not shared/committed

```sh
export AI_API_KEY=your-key-here
```

Run the above before `go run .`. 

## License

Released under the terms of the [LICENSE](LICENSE) file.
