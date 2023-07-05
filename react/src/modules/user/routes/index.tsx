import {Link, RouteObject} from "react-router-dom";
import UserAddressPage from "../UserAddressPage";
import UserPage from "../UserPage";
import UserTemplatePage from "../UserTemplatePage";
import UserOrdersPage from "../UserOrdersPage";

const userRoutes: RouteObject = {
    path: 'user',
    Component: UserTemplatePage,
    handle: {
        crumb: () => <Link to="/user"> Utente </Link>
    },
    children: [
        {
            path: '',
            Component: UserPage,
        },
        {
            path: 'address',
            Component: UserAddressPage,
            handle: {
                crumb: () => <span> Indirizzi </span>
            }
        },
        {
            path: 'orders',
            Component: UserOrdersPage,
            handle: {
                crumb: () => <span> Ordini </span>
            }
        }
    ]
};

export default userRoutes;