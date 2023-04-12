import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import './index.css'
import {createBrowserRouter, RouterProvider} from "react-router-dom";
import {LoginPage} from "./modules/login/LoginPage";
import {theme} from "../theme";
import {ConfigProvider} from "antd";

const router = createBrowserRouter([
    {
        path: '/',
        element: <App/>
    },
    {
        path: '/login',
        element: <LoginPage/>
    }
])

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <ConfigProvider theme={theme}>
        <React.StrictMode>
            <RouterProvider router={router}/>
        </React.StrictMode>
    </ConfigProvider>
)
