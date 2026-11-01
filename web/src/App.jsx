import InputComponent from './components/InputComponent.jsx' 
import useLocation from './hooks/useLocation.js'
import useWeather from './hooks/useWeather.js'

function App() { 
  const {coords, err} = useLocation()
  const {weather} = useWeather(coords)
  console.log(coords, weather);

  return (
    <div>
      <InputComponent />
    </div>
  )
}

export default App