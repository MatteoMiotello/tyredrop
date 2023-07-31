import {faMinus, faPlus} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import React, {useState} from "react";
import {useTranslation} from "react-i18next";
import {useDispatch, useSelector} from "react-redux";
import {ThunkDispatch} from "redux-thunk";
import Button from "../../../common/components-library/Button";
import LoadingSpinner from "../../../common/components-library/LoadingSpinner";
import cartSelector from "../store/cart-selector";
import {editCartItem} from "../store/cart-slice";
import {useToast} from "../../../store/toast";

type CartQuantityButtonsProps = {
    cartId: string
}

const CartQuantityButtons: React.FC<CartQuantityButtonsProps> = (props) => {
    const cart = useSelector(cartSelector.cart(props.cartId));
    const {error} = useToast();
    const dispatch = useDispatch<ThunkDispatch<any, any, any>>();
    const {t} = useTranslation();
    const [loading, setLoading] = useState<boolean>(false);

    const editCart = (quantity: number) => {
        if (!cart) {
            return;
        }

        setLoading(true);
        dispatch(editCartItem({itemId: cart.id, quantity: quantity}))
            .unwrap()
            .catch(() => {
                error(t("cart.item_quantity_edit_error"));
            })
            .finally(() => setLoading(false));
    };

    return <div className="w-full flex justify-between items-center">
        <Button type="primary" onClick={() => {
            editCart( cart?.quantity as number - 1 );
        }}>
            {loading ? <LoadingSpinner/> : <FontAwesomeIcon icon={faMinus}/>}
        </Button>
        <span className="mx-2">
            {cart ? cart.quantity : 0}
        </span>
        <Button type="primary" onClick={() => {
            editCart( cart?.quantity as number + 1 );
        }}>
            {loading ? <LoadingSpinner/> : <FontAwesomeIcon icon={faPlus}/>}
        </Button>
    </div>;
};

export default CartQuantityButtons;