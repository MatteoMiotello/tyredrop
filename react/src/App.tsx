import {useState} from 'react'
import './App.css'
import {theme} from "../theme";
import {Outlet} from "react-router-dom";

function App() {
    const [count, setCount] = useState(0)

    return (
        <>
            <Outlet></Outlet>
        </>
    )
}

export default App
