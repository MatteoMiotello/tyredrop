import React from "react";
import {Outlet} from "react-router-dom";
import CustomFooter from "../../common/components/CustomFooter";

const AuthTemplate: React.FC = ( ) => {
    return <>
        <Outlet/>
        <CustomFooter/>
    </>;
};

export default AuthTemplate;