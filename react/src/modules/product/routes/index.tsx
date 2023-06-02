import React from "react";
import { LoaderFunction, RouteObject} from "react-router-dom";
import {ProductItemQuery} from "../../../__generated__/graphql";
import {PRODUCT_ITEM} from "../../../common/backend/graph/query/products";
import apolloClientContext from "../../../common/contexts/apollo-client-context";
import ProductDetailsPage from "../ProductDetailsPage";
import ProductTablePage from "../ProductTablePage";
import ProductTemplatePage from "../ProductTemplatePage";

const productDetailsLoader: LoaderFunction = async ({params}) => {
    return apolloClientContext.query({
        query: PRODUCT_ITEM,
        variables: {
            id: params.id as string
        }
    });
};

export const productRoute: RouteObject = {
    Component: ProductTemplatePage,
    path: '',
    children: [
        {
            path: '',
            Component: ProductTablePage,
        },
        {
            path: '/products/details/:id',
            element: <ProductDetailsPage/>,
            loader: productDetailsLoader,
            handle: {
                crumb: ({data}: { data: ProductItemQuery } ) => {
                    return <span> {data.productItem?.product.name} </span>;
                }
            }
        }
    ]
};