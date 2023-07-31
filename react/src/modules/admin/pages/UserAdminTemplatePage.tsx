import React from "react";
import {Outlet} from "react-router-dom";

const UserAdminTemplatePage: React.FC = () => {
    return <>
        <Outlet/>
    </>;
};

export default UserAdminTemplatePage;