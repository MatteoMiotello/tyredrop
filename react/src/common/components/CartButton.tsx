import {faCartShopping} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import React from "react";
import {useSelector} from "react-redux";
import {Link} from "react-router-dom";
import cartSelector from "../../modules/cart/store/cart-selector";

const CartButton: React.FC = () => {
    const cartCounts = useSelector( cartSelector.count );

    return <div className="indicator mr-6">
        <span className="indicator-item badge badge-secondary">{cartCounts}</span>
        <Link to="/cart" className="avatar rounded-full border-2 p-3 aspect-square">
            <FontAwesomeIcon icon={faCartShopping} className="text-primary text-md"/>
        </Link>
    </div>;
};

export default CartButton;