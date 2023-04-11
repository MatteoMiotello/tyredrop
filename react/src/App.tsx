import {useState} from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import {ConfigProvider} from "antd";
import {theme} from "../theme";

function App() {
    const [count, setCount] = useState(0)

    return (
            <div className="text-default">
                helloooo
            </div>
    )
}

export default App
