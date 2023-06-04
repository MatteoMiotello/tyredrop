import {RouteObject} from "react-router-dom";
import i18n from "../../../common/i18n";
import CartPage from "../CartPage";
import CartTemplate from "../CartTemplate";

export const cartRoute: RouteObject = {
    Component: CartTemplate,
    path: 'cart',
    children: [
        {
            Component: CartPage,
            path: '',
            handle: {
                crumb: () => {
                    return <span> {i18n.t("cart.breadcrumb_title")} </span>;
                }
            }
        }
    ]
};