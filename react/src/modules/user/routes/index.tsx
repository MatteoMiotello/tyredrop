import {Link, LoaderFunction, RouteObject} from "react-router-dom";
import {FetchUserQuery} from "../../../__generated__/graphql";
import {USER} from "../../../common/backend/graph/query/users";
import apolloClientContext from "../../../common/contexts/apollo-client-context";
import UserAddressPage from "../UserAddressPage";
import UserPage from "../UserPage";
import UserTemplatePage from "../UserTemplatePage";
import UserOrdersPage from "../UserOrdersPage";

const userLoader: LoaderFunction = async ({params}) => {
    return apolloClientContext.query({
        query: USER,
        variables: {
            userId: params.id
        }
    });
};

const userRoutes: RouteObject = {
    path: 'user',
    Component: UserTemplatePage,
    children: [
        {
            path: ':id',
            Component: UserPage,
            loader: userLoader,
            handle: {
                crumb: ({data}: {data: FetchUserQuery }) => {
                    return <Link to={`/user/${data?.user?.id}`}> Utente: {data?.user?.email} </Link>;
                }
            },
        },
        {
            path: ':id/address',
            loader: userLoader,
            Component: UserAddressPage,
            handle: {
                crumb: () => <span> Indirizzi </span>
            }
        },
        {
            path: ':id/orders',
            loader: userLoader,
            Component: UserOrdersPage,
            handle: {
                crumb: () => <span> Ordini </span>
            }
        }
    ]
};

export default userRoutes;