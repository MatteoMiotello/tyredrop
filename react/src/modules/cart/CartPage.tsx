import {faShoppingCart} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import React from "react";
import {useTranslation} from "react-i18next";
import {useSelector} from "react-redux";
import {Link} from "react-router-dom";
import CartTable from "./components/CartTable";
import CheckoutPanel from "./components/CheckoutPanel";
import cartSelector from "./store/cart-selector";

const CartPage: React.FC = () => {
    const cartItems = useSelector(cartSelector.items);
    const {t} = useTranslation();

    return <main className="w-full m-0 px-4 h-full flex flex-col min-h-screen">
        <div className="lg:grid lg:grid-cols-3">
            <div className="lg:col-span-2 px-4">
                {cartItems.length ? <CartTable cartItems={cartItems}/> :
                    <div className="w-full flex justify-center items-center flex-col mt-24 text-base-300">
                        <FontAwesomeIcon icon={faShoppingCart} fontSize="70"/>
                        <span className="text-xl font-bold mt-4"> {t("cart.empty_cart")} </span>
                        <Link to="/" className="btn btn-primary btn-outline mt-4">
                            {t("cart.go_back_to_home")}
                        </Link>
                    </div>
                }
            </div>
            <div className="p-4 relative">
                <CheckoutPanel/>
            </div>
        </div>
    </main>;
};

export default CartPage;