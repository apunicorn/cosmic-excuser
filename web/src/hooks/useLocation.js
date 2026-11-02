import { useEffect, useState } from "react";

export default function useLocation() {
    const [coords, setCoords] = useState(null);
    const [err, setErr] = useState(null)

    useEffect(() => {
        // get user geolocation - user asked to provide 
        navigator.geolocation.getCurrentPosition(
            (pos) => setCoords({lat: pos.coords.latitude, lon: pos.coords.longitude}),
            (err) => setErr(err.message),
        )
    }, [])
    
    return {coords, err}
}
