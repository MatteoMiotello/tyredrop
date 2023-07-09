import {Link, LoaderFunction, RouteObject} from "react-router-dom";
import { Order} from "../../../__generated__/graphql";
import OrderTemplatePage from "../OrderTemplatePage";
import OrderDetailsPage from "../OrderDetailsPage";
import apolloClientContext from "../../../common/contexts/apollo-client-context";
import {FETCH_ORDER} from "../../../common/backend/graph/query/order";


const orderLoader: LoaderFunction = async ({params}) => {
    return apolloClientContext.query( {
        query: FETCH_ORDER,
        variables: {
            orderId: params.id as string
        }
    } ).then( res  => res.data );
};

export const orderRoutes: RouteObject = {
    path: 'order',
    Component: OrderTemplatePage,
    children: [
        {
            path: 'details/:id',
            loader: orderLoader,
            Component: OrderDetailsPage,
            handle: {
                crumb: ( {order}: {order: Order} ) => <Link to={`/order/details/${order.id}`}> Ordine: #{ order.id } </Link>
            }
        }
    ]
};