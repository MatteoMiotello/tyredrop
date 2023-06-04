import React from "react";
import {Outlet} from "react-router-dom";

const CartTemplate: React.FC = () => {
    return <>
        <Outlet></Outlet>
    </>;
};

export default CartTemplate;