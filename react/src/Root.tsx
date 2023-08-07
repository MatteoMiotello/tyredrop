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
import {adminRoutes} from "./modules/admin/routes";
import ContactsPage from "./common/pages/ContactsPage";
import CommonTemplate from "./CommonTemplate";
import GeneralTermsPage from "./common/pages/GeneralTermsPage";
import AboutPage from "./common/pages/AboutPage";
import FaqPage from "./common/pages/FaqPage";
import LegalMentionsPage from "./common/pages/LegalMentionsPage";
import PrivacyPage from "./common/pages/PrivacyPage";

Moment.globalMoment = moment;
Moment.globalLocale = 'it';
Moment.globalTimezone = 'Europe/Rome';
Moment.globalFormat = 'DD/MM/YYYY HH:mm';
Moment.globalLocal = true;
moment.updateLocale( 'it', {
    workingWeekdays: [1, 2, 3, 4, 5]
} );

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
            adminRoutes
        ]
    },
    {
        path: '/not_confirmed',
        element: <NotConfirmedPage/>
    },
    {
        path: '/',
        element: <CommonTemplate/>,
        children: [
            {
                path: 'general-terms',
                handle: {
                    crumb: () => <span> Condizioni generali di vendita </span>
                },
                Component: GeneralTermsPage,
            },
            {
                path: 'contacts',
                handle: {
                    crumb: () => <span> Contatti </span>
                },
                Component: ContactsPage
            },
            {
                path: 'about',
                handle: {
                    crumb: () => <span> A proposito di Tyres in the world </span>
                },
                Component: AboutPage
            },
            {
                path: 'privacy',
                handle: {
                    crumb: () => <span> Dichiarazione sulla privacy </span>
                },
                Component: PrivacyPage
            },
            {
                path: 'faq',
                handle: {
                    crumb: () => <span> FAQ </span>
                },
                Component: FaqPage,
            },
            {
                path: 'legal-mentions',
                handle: {
                    crumb: () => <span> Menzioni legali </span>
                },
                Component: LegalMentionsPage,
            }
        ]
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