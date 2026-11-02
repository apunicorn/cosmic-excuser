import { useState } from 'react'
import './App.css'
import InputComponent from './components/InputComponent.jsx' 
import useLocation from './hooks/useLocation.js'

function App() { 
  const {coords, err: locationError} = useLocation() // browser location
  const [excuse, setExcuse] = useState('')
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  const getExcuse = async (action, sign) => {
    setLoading(true)
    setError('')
    setExcuse('')

    try {
      const resp = await fetch('/api/excuse', { // get excuse from gemini on backend
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ action, sign, lat: coords.lat, lon: coords.lon }), // weather API uses lat/long
      })

      if (!resp.ok) throw new Error(await resp.text()) // show the server's error
        
      const data = await resp.json()
      setExcuse(data.excuse)
    } catch (e) {
      setError(e.message || 'Unable to get excuse.')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="app">
      <h1>✨ Cosmic Excuser ✨</h1>
      <InputComponent onSubmit={getExcuse} loading={loading} />
      {error && <p className="error">{error}</p>}
      {locationError && <p className="error">Couldn't get your location to blame the weather.</p>}
      {excuse && <p className="excuse">{excuse}</p>}
    </div>
  )
}

export default App