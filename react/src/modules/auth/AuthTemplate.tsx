import React from "react";
import {Outlet} from "react-router-dom";
import CustomFooter from "../../common/components/CustomFooter";

const AuthTemplate: React.FC = ( ) => {
    return <>
        <main className="bg-base lg:p-24 p-4">
            <Outlet/>
        </main>
        <CustomFooter/>
    </>;
};

export default AuthTemplate;