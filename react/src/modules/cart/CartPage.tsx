import { faShoppingCart} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import React from "react";
import {useTranslation} from "react-i18next";
import {useDispatch, useSelector} from "react-redux";
import {Link} from "react-router-dom";
import Panel from "../../common/components-library/Panel";
import CartTable from "./components/CartTable";
import CheckoutPanel from "./components/CheckoutPanel";
import cartSelector from "./store/cart-selector";
import Button from "../../common/components-library/Button";
import {ThunkDispatch} from "redux-thunk";
import {emptyCart} from "./store/cart-slice";

const CartPage: React.FC = () => {
    const cartItems = useSelector(cartSelector.items);
    const dispatch = useDispatch<ThunkDispatch<any, any, any>>();
    const {t} = useTranslation();

    return <main className="w-full m-0 p-1 h-full flex flex-col min-h-screen">
        <div className="lg:grid lg:grid-cols-3 gap-4">
            <Panel className="lg:col-span-2 px-4">
                <Panel.Title>
                    Il mio carrello
                    <Button
                        type="ghost"
                        size="sm"
                        onClick={ () => dispatch( emptyCart() ) }
                    >
                        Svuota carrello
                    </Button>
                </Panel.Title>
                {cartItems.length ? <CartTable cartItems={cartItems}/> :
                    <div className="w-full flex justify-center items-center flex-col mt-24 text-base-300">
                        <FontAwesomeIcon icon={faShoppingCart} fontSize="70"/>
                        <span className="text-xl font-bold mt-4"> {t("cart.empty_cart")} </span>
                        <Link to="/" className="btn btn-primary btn-outline mt-4">
                            {t("cart.go_back_to_home")}
                        </Link>
                    </div>
                }
            </Panel>
            <Panel className="p-4 ">
                <CheckoutPanel/>
            </Panel>
        </div>
    </main>;
};

export default CartPage;