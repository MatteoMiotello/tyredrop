import {ApolloProvider} from "@apollo/client";
import React, {useState} from "react";
import {Link, RouterProvider, createBrowserRouter} from "react-router-dom";
import App from "./App";
import client from "./common/contexts/apollo-client-context";
import ToastContext from "./common/contexts/toast-context";
import {authRoutes} from "./modules/auth/routes";
import {billingRoute} from "./modules/billing/routes";
import {cartRoute} from "./modules/cart/routes";
import {productRoute} from "./modules/product/routes";
import {store} from "./store/store";
import i18n from "./common/i18n";
import {I18nextProvider} from "react-i18next";
import {Provider} from "react-redux";
import {CustomToast, ToastConfig} from "./common/components/CustomToast";
import NotConfirmedPage from "./NotConfirmedPage";

const router = createBrowserRouter([
    {
        path: '/',
        element: <App/>,
        handle: {
            crumb: () => <Link className="link" to="/"> Home </Link>
        },
        children: [
            billingRoute,
            productRoute,
            cartRoute
        ]
    },
    {
        path: '/not_confirmed',
        element: <NotConfirmedPage/>
    },
    authRoutes
]);

const Root: React.FC = () => {
    const [toasts, setToasts] = useState<ToastConfig[]>([]);
    const value = {toasts: toasts, setToasts: setToasts};

    return <>
        <ApolloProvider client={client}>
            <Provider store={store}>
                <I18nextProvider i18n={i18n}>
                    <ToastContext.Provider value={value}>
                        <RouterProvider router={router}/>
                        <CustomToast toasts={toasts}/>
                    </ToastContext.Provider>
                </I18nextProvider>
            </Provider>
        </ApolloProvider>
    </>;
};

export default Root;