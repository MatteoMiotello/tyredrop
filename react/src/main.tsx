import React from 'react';
import ReactDOM from 'react-dom/client';
import {Provider} from "react-redux";
import './index.css';
import {RouterProvider, createBrowserRouter} from "react-router-dom";
import App from "./App";
import {LoginPage} from "./modules/auth/LoginPage";
import {I18nextProvider} from "react-i18next";
import i18n from "./common/i18n";
import RegisterPage from "./modules/auth/RegisterPage";
import AuthTemplate from "./modules/auth/AuthTemplate";
import {store} from "./store/store";

const router = createBrowserRouter([
    {
        path: '/',
        element: <App/>,
        children: [
            {
                path: 'auth',
                element: <AuthTemplate/>,
                children: [
                    {
                        path: 'login',
                        element: <LoginPage/>
                    },
                    {
                        path: 'register',
                        element: <RegisterPage/>
                    }
                ]
            },
        ]
    },
    ]);

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>
        <Provider store={store}>
            <I18nextProvider i18n={i18n}>
                <RouterProvider router={router}/>
            </I18nextProvider>
        </Provider>
    </React.StrictMode>
);
