import { useEffect, useState } from "react";

export default function useLocation() {
    const [coords, setCoords] = useState(null);
    const [err, setErr] = useState(null)

    useEffect(() => {
        navigator.geolocation.getCurrentPosition(
            (pos) => setCoords({lat: pos.coords.latitude}, {long: pos.coords.longitude}),
            (err) => setErr(err.message),
        )
    }, [])
    
    return {coords, err}
}
