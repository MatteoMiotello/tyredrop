import {faShoppingCart} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import React, {useState} from "react";
import {useTranslation} from "react-i18next";
import {useDispatch} from "react-redux";
import {ThunkDispatch} from "redux-thunk";
import Button from "../../../common/components-library/Button";
import LoadingSpinner from "../../../common/components-library/LoadingSpinner";
import {useToast} from "../../../hooks/useToast";
import {addCartItem} from "../../cart/store/cart-slice";

type AddItemToCartButton = {
    itemId: string
}

const AddItemToCartButton: React.FC<AddItemToCartButton> = ( props ) => {
    const dispatch = useDispatch<ThunkDispatch<any, any, any>>();
    const { setSuccess, setError } = useToast();
    const {t} = useTranslation();
    const [  loading, setLoading ] = useState<boolean>(false);

    return <Button
        className="mx-2 aspect-square"
        type={"primary"}
        onClick={ () => {
            setLoading(true);
            dispatch( addCartItem({itemId: props.itemId} ) )
                .then( () => {
                    setSuccess( t( "cart.item_added_success" ) );
                } )
                .catch( () => {
                    setError( t( "cart.item_add_error" ) );
                } )
                .finally( () =>{
                    setLoading(false);
                } );
        } }
    >
        { loading ? <LoadingSpinner/> : <FontAwesomeIcon icon={faShoppingCart}/>}
    </Button>;
};

export default AddItemToCartButton;