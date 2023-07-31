import {Link, RouteObject} from "react-router-dom";
import UserPage from "../pages/UserPage";
import {userLoader} from "../../user/routes";
import UserDetailsPage from "../pages/UserDetailsPage";
import {User} from "../../../__generated__/graphql";
import UserAdminTemplatePage from "../pages/UserAdminTemplatePage";
import OrderPage from "../pages/OrderPage";



export const adminRoutes: RouteObject = {
    path: '',
    children: [
        {
            path: 'user',
            Component: UserAdminTemplatePage,
            handle: {
                crumb: () => <Link to="/admin/user"> Tutti gli utenti </Link>
            },
            children: [
                {
                    path: '',
                    Component: UserPage,
                },
                {
                    path: ':id',
                    loader: userLoader,
                    handle: {
                        crumb: ({data}: {data: {user: User}}) => <span> {data.user.email} </span>
                    },
                    Component: UserDetailsPage,
                }
            ]
        },
        {
            path: 'order',
            handle: {
                crumb: () => <Link to="/admin/order"> Tutti gli ordini </Link>
            },
            children: [
                {
                    path: '',
                    Component: OrderPage
                }
            ]
        }
    ]
};