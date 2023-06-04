import { faShoppingCart} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import React from "react";
import {useTranslation} from "react-i18next";
import {useSelector} from "react-redux";
import {Link} from "react-router-dom";
import Button from "../../common/components-library/Button";
import {Currency} from "../../common/utilities/currency";
import CartTable from "./components/CartTable";
import UserAddressSelector from "./components/UserAddressSelector";
import cartSelector from "./store/cart-selector";

const CartPage: React.FC = () => {
    const cartItems = useSelector(cartSelector.items);
    const totalPrice = useSelector(cartSelector.amount);

    const {t} = useTranslation();
    
    const getPrice = () => {
        if ( !totalPrice || !totalPrice.currency ) {
            return "-";
        }

        return Currency.defaultFormat(totalPrice.value, totalPrice.currency?.iso_code as string);
    };

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
                <div className="sticky top-20 w-full bg-base-200 rounded-box p-4 flex flex-col">
                    <UserAddressSelector/>
                    <div className="ml-auto mt-10 flex flex-col text-secondary text-sm">
                        Prezzo totale:
                        <span className="text-4xl font-semibold text-primary">
                        {getPrice()}
                    </span>
                    </div>
                    <Button className="ml-auto mt-4" type="secondary" onClick={() => openModal()}>
                        Conferma Ordine
                    </Button>
                </div>
            </div>
        </div>
    </main>;
};

export default CartPage;