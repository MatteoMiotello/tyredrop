import {faShoppingCart} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import React, {useState} from "react";
import {useTranslation} from "react-i18next";
import {useDispatch} from "react-redux";
import {ThunkDispatch} from "redux-thunk";
import Button from "../../../common/components-library/Button";
import Input from "../../../common/components-library/Input";
import LoadingSpinner from "../../../common/components-library/LoadingSpinner";

import {addCartItem} from "../../cart/store/cart-slice";
import {useToast} from "../../../store/toast";

type AddItemToCartButton = {
    itemId: string,
    quantity: number
}

const AddItemToCartButton: React.FC<AddItemToCartButton> = (props) => {
    const dispatch = useDispatch<ThunkDispatch<any, any, any>>();
    const [quantity, setQuantity] = useState(1);
    const {success, error} = useToast();
    const {t} = useTranslation();
    const [loading, setLoading] = useState<boolean>(false);

    return <div className="flex">
        <Input.FormControl>
            <Input.Input
                className="!w-14"
                value={quantity}
                type="number"
                name="quantity"
                placeholder="1"
                onValueChange={setQuantity}
                validators={[(value) => {
                    if (value > props.quantity) {
                        return 'La quantità selezionata non è disponibile';
                    }

                    return null;
                }]}
            />
        </Input.FormControl>
        {props.quantity >= 4 &&
            <Button
                className="mx-2 aspect-square"
                type={"primary"}
                onClick={() => {
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