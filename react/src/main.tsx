import React, {Suspense} from 'react';
import ReactDOM from 'react-dom/client';
import {Provider} from "react-redux";
import './index.css';
import {RouterProvider, createBrowserRouter} from "react-router-dom";
import App from "./App";
import {I18nextProvider} from "react-i18next";
import i18n from "./common/i18n";
import {authRoutes} from "./modules/auth/routes";
import {store} from "./store/store";
import Spinner from "./common/components/Spinner";

const router = createBrowserRouter([
    {
        path: '/',
        element: <App/>,
    },
    authRoutes
]);

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <Suspense fallback={<Spinner/>}>
        <React.StrictMode>
            <Provider store={store}>
                <I18nextProvider i18n={i18n}>
                    <RouterProvider router={router}/>
                </I18nextProvider>
            </Provider>
        </React.StrictMode>
    </Suspense>
);
