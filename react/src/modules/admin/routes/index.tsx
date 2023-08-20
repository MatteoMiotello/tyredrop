import {Link, LoaderFunction, RouteObject} from "react-router-dom";
import {FETCH_ORDER} from "../../../common/backend/graph/query/order";
import apolloClientContext from "../../../common/contexts/apollo-client-context";
import HomePage from "../pages/HomePage";
import OrderAdminDetailsPage from "../pages/OrderAdminDetailsPage";
import UserPage from "../pages/UserPage";
import {userLoader} from "../../user/routes";
import UserDetailsPage from "../pages/UserDetailsPage";
import {Order, User} from "../../../__generated__/graphql";
import UserAdminTemplatePage from "../pages/UserAdminTemplatePage";
import OrderPage from "../pages/OrderPage";

const orderLoader: LoaderFunction = async ({params}) => {
    return apolloClientContext.query({
        query: FETCH_ORDER,
        variables: {
            orderId: params.id as string
        },
        fetchPolicy: "no-cache"
    }).then(res => res.data);
};

export const adminRoutes: RouteObject = {
    path: '',
    Component: HomePage,
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
                },
                {
                    path: ':id',
                    loader: orderLoader,
                    handle: {
                        crumb: ({order}: { order: Order }) => <span> Ordine: #{order.orderNumber} </span>
                    },
                    Component: OrderAdminDetailsPage
                }
            ]
        }
    ]
};