import { useEffect, useState } from "react";

export default function useWeather(coords) {
  const [weather, setWeather] = useState(null)
  const [error, setError] = useState(null)

  useEffect(() => {
    if (!coords) return
    fetch(`/api/weather?lat=${coords.lat}&lon=${coords.lon}`)
      .then((r) => (r.ok ? r.json() : Promise.reject(new Error(r.statusText))))
      .then(setWeather)
      .catch((e) => setError(e.message))
  }, [coords])

  return { weather, error }
}
