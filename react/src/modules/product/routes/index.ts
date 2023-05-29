import {RouteObject} from "react-router-dom";
import ProductDetailsPage from "../ProductDetailsPage";
import ProductTablePage from "../ProductTablePage";
import ProductTemplatePage from "../ProductTemplatePage";

export const productRoute: RouteObject = {
    Component: ProductTemplatePage,
    path: '',
    children: [
        {
            path: '',
            Component: ProductTablePage
        },
        {
            path: '/products/details/:id',
            Component: ProductDetailsPage,
        }
    ]
};