import {Link, RouteObject} from "react-router-dom";
import UserAddressPage from "../UserAddressPage";
import UserPage from "../UserPage";
import UserTemplatePage from "../UserTemplatePage";

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
        }
    ]
};

export default userRoutes;