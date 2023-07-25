import {ApolloProvider} from "@apollo/client";
import React, {useState} from "react";
import {Link, RouterProvider, createBrowserRouter} from "react-router-dom";
import Admin from "./Admin";
import App from "./App";
import ModalContainer, {ModalData} from "./common/components/ModalContainer";
import client from "./common/contexts/apollo-client-context";
import ModalContext from "./common/contexts/modal-context";
import {authRoutes} from "./modules/auth/routes";
import {billingRoute} from "./modules/billing/routes";
import {cartRoute} from "./modules/cart/routes";
import {productRoute} from "./modules/product/routes";
import userRoutes from "./modules/user/routes";
import {store} from "./store/store";
import i18n from "./common/i18n";
import {I18nextProvider} from "react-i18next";
import {Provider} from "react-redux";
import ToastContainer from "./common/components/ToastContainer";
import NotConfirmedPage from "./NotConfirmedPage";
import {orderRoutes} from "./modules/order/routes";

import moment from 'moment-timezone';
import Moment from "react-moment";
import 'moment/locale/it';

Moment.globalMoment = moment;
Moment.globalLocale = 'it';
Moment.globalTimezone = 'Europe/Rome';
Moment.globalFormat = 'DD/MM/YYYY HH:mm';
Moment.globalLocal = true;

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
    const [modal, setModal] = useState<ModalData | null>(null);

    return <>
        <ApolloProvider client={client}>
            <Provider store={store}>
                <I18nextProvider i18n={i18n}>
                    <ModalContext.Provider value={{modal: modal, setModal: setModal}}>
                            <RouterProvider router={router}/>
                            <ToastContainer/>
                            <ModalContainer modal={modal}/>
                    </ModalContext.Provider>
                </I18nextProvider>
            </Provider>
        </ApolloProvider>
    </>;
};

export default Root;