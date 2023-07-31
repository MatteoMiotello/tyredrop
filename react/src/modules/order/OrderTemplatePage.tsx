import React from "react";
import {Outlet} from "react-router-dom";

const OrderTemplatePage: React.FC = () => {
    return <>
        <Outlet></Outlet>
    </>;
};

export default OrderTemplatePage;