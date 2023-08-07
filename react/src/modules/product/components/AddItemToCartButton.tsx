import {faMinus, faPlus, faShoppingCart} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import React, {useState} from "react";
import {useTranslation} from "react-i18next";
import {useDispatch} from "react-redux";
import {ThunkDispatch} from "redux-thunk";
import LoadingSpinner from "../../../common/components-library/LoadingSpinner";

import {addCartItem} from "../../cart/store/cart-slice";
import {useToast} from "../../../store/toast";
import {Button, Input, Join} from "../../../common/components/shelly-ui";

type AddItemToCartButton = {
    itemId: string,
    quantity: number
}

const AddItemToCartButton: React.FC<AddItemToCartButton> = (props) => {
    const dispatch = useDispatch<ThunkDispatch<any, any, any>>();
    const [quantity, setQuantity] = useState<number>(1);
    const {success, error} = useToast();
    const {t} = useTranslation();
    const [loading, setLoading] = useState<boolean>(false);

    return <div className="flex">
        <Join>
            <Button className="join-item" disabled={quantity == 0} onClick={() => {
                if (quantity <= 0) {
                    return;
                }
                setQuantity(quantity - 1);
            }}><FontAwesomeIcon icon={faMinus}/></Button>
            <Input.FormControl className="join-item">
                <Input
                    min={0}
                    bordered={false}
                    className="!w-14 !outline-none input-ghost text-center"
                    value={String( quantity )}
                    type="number"
                    name="quantity"
                    placeholder="1"
                    onValueChange={ (val) => {
                        setQuantity(Number(val));
                    }}
                    validators={[(value) => {
                        if (Number(value) > props.quantity) {
                            return 'La quantità selezionata non è disponibile';
                        }

                        return null;
                    }]}
                />
            </Input.FormControl>
            <Button className="join-item" onClick={() => {
                setQuantity(quantity + 1);
            }}><FontAwesomeIcon icon={faPlus}/></Button>
        </Join>
        {props.quantity >= 4 &&
            <Button
                className="mx-2 aspect-square"
                buttonType="primary"
                onClick={() => {
                    if (quantity <= 0) {
                        error('Selezionare una quantità valida');
                        return;
                    }

                    if (quantity > props.quantity) {
                        error('La quantità selezionata non è disponibile');
                        return;
                    }

                    setLoading(true);

                    dispatch(addCartItem({itemId: props.itemId, quantity: quantity}))
                        .unwrap()
                        .then(() => {
                            success(t("cart.item_added_success"));
                        })
                        .catch(() => {
                            error(t("cart.item_add_error"));
                        })
                        .finally(() => {
                            setLoading(false);
                        });
                }}
            >
                {loading ? <LoadingSpinner/> : <FontAwesomeIcon icon={faShoppingCart}/>}
            </Button>
        }
    </div>;
};

export default AddItemToCartButton;