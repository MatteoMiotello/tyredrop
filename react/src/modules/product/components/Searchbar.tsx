import React, {useContext} from "react";
import {useTranslation} from "react-i18next";


import {Partial} from "../../../common/components-library/Tabs";
import TyreSearchForm, {TyreSearchFormRequest, toSearchDataType} from "./TyreSearchForm";
import ProductSearchContext from "../context/product-search-context";

const Searchbar: React.FC = () => {
    const {t} = useTranslation();
    const { setSearchData } = useContext( ProductSearchContext );

    const onSubmit = ( req: TyreSearchFormRequest, vehicleCode: string ) => {
        const searchData = toSearchDataType( req, vehicleCode );
        setSearchData( searchData );
    };

    const tabParts: Partial[] = [
        {
            title: t( 'searchbar.tyre_tab_title' ),
            content: <TyreSearchForm
                onSubmit={onSubmit}
            />
        },
    ];

    return <div className="bg-primary w-full min-h-64 rounded-box">
        <div className="h-full flex md:flex-row flex-col lg:px-24 px-4 py-5 justify-around">
            <TyreSearchForm
                onSubmit={onSubmit}
            />
        </div>
    </div>;
};

export default Searchbar;