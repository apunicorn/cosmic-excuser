import { useState } from 'react'
import './App.css'
import InputComponent from './components/InputComponent.jsx' 
import useLocation from './hooks/useLocation.js'
import useWeather from './hooks/useWeather.js'

function App() { 
  const {coords} = useLocation()
  const {weather} = useWeather(coords)
  console.log(coords, weather);

  const [excuse, setExcuse] = useState('')
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  const getExcuse = async (action, sign) => {
    setLoading(true)
    setError('')
    setExcuse('')
    try {
      const resp = await fetch('/api/excuse', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ action, sign }),
      })

      if (!resp.ok) throw new Error(await resp.text())
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
      <h1>✨✨ Cosmic Excuser ✨✨</h1>
      <InputComponent onSubmit={getExcuse} loading={loading} />
      {error && <p className="error">{error}</p>}
      {excuse && <p className="excuse">{excuse}</p>}
    </div>
  )
}

export default App