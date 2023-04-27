import React, { useState} from "react";
import {RouterProvider, createBrowserRouter} from "react-router-dom";
import App from "./App";
import {authRoutes} from "./modules/auth/routes";
import {store} from "./store/store";
import i18n from "./common/i18n";
import {I18nextProvider} from "react-i18next";
import {Provider} from "react-redux";
import {CustomToast} from "./common/components/CustomToast";


const router = createBrowserRouter([
    {
        path: '/',
        element: <App/>,
    },
    authRoutes
]);

export const ToastContext = React.createContext({
    toast: null,
    setToasts: (): void => {}
});

const Root: React.FC = () => {
    const [toast, setToasts] = useState(null );
    const value = { toast: toast, setToasts };

    return <>
        <Provider store={store}>
            <I18nextProvider i18n={i18n}>
                <ToastContext.Provider value={value}>
                    <RouterProvider router={router}/>
                    <CustomToast toast={toast}/>
                </ToastContext.Provider>
            </I18nextProvider>
        </Provider>
    </>;
};

export default Root;