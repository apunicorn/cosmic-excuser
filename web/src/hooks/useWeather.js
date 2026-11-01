import { useEffect, useState } from "react";

export default function useWeather(lat, long) {
    const [weather, setWeather] = useState(null);

    useEffect(() => {
       fetch(`/api/weather?lat=${lat}&lon=${long}`)
       .then((resp) => resp.json())
       .then(setWeather)
    }, (err) => console.log(err))
    
    return {weather}
}
