import {RouteObject} from "react-router-dom";
import i18n from "../../../common/i18n";
import CartPage from "../CartPage";

export const cartRoute: RouteObject = {
    Component: CartPage,
    path: 'cart',
    handle: {
        crumb: () => {
            return <span> {i18n.t("cart.breadcrumb_title")} </span>;
        }
    }
};