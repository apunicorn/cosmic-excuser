# ✨✨✨ Cosmic Excuser ✨✨✨

**Did something questionable? Blame the universe!!**

Cosmic Excuser takes whatever you did and writes you a convincing excuse that blames the weather, the planets and your zodiac sign. Mercury was in retrograde, it was 38°C and humid, and you're a Scorpio. None of this was your fault. Feel free to use this excuse for work, your spouse, or anywhere else. 

<img width="967" height="863" alt="Screenshot 2026-11-02 at 10 23 55 AM" src="https://github.com/user-attachments/assets/a56379ac-8839-4ad7-aadb-3b9316c0dccc" />


## How it works

1. You describe what you did and pick your zodiac sign. The browser also shares your location.
2. Go back end:
   - **Your zodiac sign**, checked against a fixed list on the server 
   - **Weather** at your location from [Open-Meteo](https://open-meteo.com/), free with no API key. It is fetched by the server and becomes part of the AI prompt
3. It sends all of that to an LLM ([Google Gemini](https://ai.google.dev/)) with instructions to write a short, dry excuse 
4. The React front end 
    - form for what you did
    - zodiac sign drop-down
    - browser geolocation
    - excuse renders


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

## Tests

Go unit tests cover the API validation (bad input, fake zodiac signs, out-of-range coordinates), the temperature conversion, and the weather text sent to Gemini. 

```sh
cd server
go test ./... -v
```

## License

Released under the terms of the [LICENSE](LICENSE) file.
