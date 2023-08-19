import React, {useContext} from "react";
import {useTranslation} from "react-i18next";
import video from "../../../assets/video-background.mp4";

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

    return <div className="bg-primary w-full min-h-64 rounded-box flex justify-center relative">
        <div className="w-full h-full absolute overflow-hidden rounded-box flex items-center">
        <video loop muted autoPlay className="absolute w-full cover-full">
            <source src={video} type="video/mp4"/>
        </video>
        </div>
        <div className="h-full flex md:flex-row flex-col max-w-3xl px-4 py-5 justify-around z-20 ">
            <TyreSearchForm
                onSubmit={onSubmit}
            />
        </div>
    </div>;
};

export default Searchbar;