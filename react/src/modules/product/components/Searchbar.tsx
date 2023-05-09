import React, {useState} from "react";
import {useTranslation} from "react-i18next";


import Tabs, {Partial} from "../../../common/components-library/Tabs";
import TyreSpecificSearchForm from "./TyreSpecificSearchForm";
import BasicTyreSearchForm from "./BasicTyreSearchForm";

const Searchbar: React.FC = () => {
    const {t} = useTranslation();
    const [searchParams, setSearchParams ] = useState([]);

    const tabParts: Partial[] = [
        {
            title: t( 'searchbar.basic_search_tab_title' ),
            content: <BasicTyreSearchForm/>
        },
        {
            title: t( 'searchbar.tyre_tab_title' ),
            content: <TyreSpecificSearchForm/>
        },
    ];

    return <div className="bg-primary w-full h-64 ">
        <div className="h-full flex md:flex-row flex-col lg:px-24 pt-10 justify-around">
            <Tabs parts={tabParts}></Tabs>
        </div>
    </div>;
};

export default Searchbar;