import {ApolloProvider} from "@apollo/client";
import React, {useState} from "react";
import {Link, RouterProvider, createBrowserRouter} from "react-router-dom";
import Admin from "./Admin";
import App from "./App";
import ModalContainer, {ModalData} from "./common/components/ModalContainer";
import client from "./common/contexts/apollo-client-context";
import ModalContext from "./common/contexts/modal-context";
import ToastContext from "./common/contexts/toast-context";
import {authRoutes} from "./modules/auth/routes";
import {billingRoute} from "./modules/billing/routes";
import {cartRoute} from "./modules/cart/routes";
import {productRoute} from "./modules/product/routes";
import userRoutes from "./modules/user/routes";
import {store} from "./store/store";
import i18n from "./common/i18n";
import {I18nextProvider} from "react-i18next";
import {Provider} from "react-redux";
import {ToastConfig, ToastContainer} from "./common/components/ToastContainer";
import NotConfirmedPage from "./NotConfirmedPage";
import {orderRoutes} from "./modules/order/routes";

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
            cartRoute,
            userRoutes,
            orderRoutes,
        ]
    },
    {
        path: '/admin',
        element: <Admin/>,
        handle: {
            crumb: () => <Link className="link" to="/admin"> Area privata </Link>
        },
        children: [

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
    const [modal, setModal] = useState<ModalData | null>(null);

    return <>
        <ApolloProvider client={client}>
            <Provider store={store}>
                <I18nextProvider i18n={i18n}>
                    <ModalContext.Provider value={{modal: modal, setModal: setModal}}>
                        <ToastContext.Provider value={{toasts: toasts, setToasts: setToasts}}>
                            <RouterProvider router={router}/>
                            <ToastContainer toasts={toasts}/>
                            <ModalContainer modal={modal}/>
                        </ToastContext.Provider>
                    </ModalContext.Provider>
                </I18nextProvider>
            </Provider>
        </ApolloProvider>
    </>;
};

export default Root;